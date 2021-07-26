package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"
	"sync"

	hspb "github.com/milaniez/gogrpcsample/hashtable"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
)

type ServiceMethodType int

const (
	ServiceMethodTypeAdd ServiceMethodType = iota
	ServiceMethodTypeDel
	ServiceMethodTypeGet
	ServiceMethodTypeDump
	ServiceMethodTypeDumpSz
)

type ControlMsg struct {
	typeSM    ServiceMethodType
	valSM     interface{}
	replyChan interface{}
}

type ClientBundle struct {
	keyToVal map[string]string
	reqChan  chan ControlMsg
}

func (cb *ClientBundle) HandleClientBundle() {
	for {
		msg, recvd := <-cb.reqChan
		if !recvd {
			return
		}
		switch msg.typeSM {
		case ServiceMethodTypeAdd:
			req, ok := msg.valSM.(*hspb.KeyVal)
			if !ok {
				log.Fatal(errors.New("logic issue"))
			}
			ch, ok := msg.replyChan.(chan hspb.Code)
			if !ok {
				log.Fatal(errors.New("logic issue"))
			}
			cb.ClientBundleAdd(req, ch)
		case ServiceMethodTypeGet:
			req, ok := msg.valSM.(*hspb.Key)
			if !ok {
				log.Fatal(errors.New("logic issue"))
			}
			ch, ok := msg.replyChan.(chan hspb.Val)
			if !ok {
				log.Fatal(errors.New("logic issue"))
			}
			cb.ClientBundleGet(req, ch)
		case ServiceMethodTypeDel:
			req, ok := msg.valSM.(*hspb.Key)
			if !ok {
				log.Fatal(errors.New("logic issue"))
			}
			ch, ok := msg.replyChan.(chan hspb.Code)
			if !ok {
				log.Fatal(errors.New("logic issue"))
			}
			cb.ClientBundleDel(req, ch)
		case ServiceMethodTypeDump:
			ch, ok := msg.replyChan.(chan hspb.KeyVal)
			if !ok {
				log.Fatal(errors.New("logic issue"))
			}
			cb.ClientBundleGetDump(ch)
		case ServiceMethodTypeDumpSz:
			ch, ok := msg.replyChan.(chan int)
			if !ok {
				log.Fatal(errors.New("logic issue"))
			}
			cb.ClientBundleGetDumpSz(ch)
		}
	}
}

func (cb *ClientBundle) ClientBundleAdd(req *hspb.KeyVal, ch chan hspb.Code) {
	if _, ok := cb.keyToVal[req.Key]; ok {
		ch <- hspb.Code{
			Code: hspb.Code_STATUS_CODE_ER_DUP,
		}
		close(ch)
		return
	}
	cb.keyToVal[req.Key] = req.Val
	ch <- hspb.Code{
		Code: hspb.Code_STATUS_CODE_OK,
	}
	close(ch)
}

func (cb *ClientBundle) ClientBundleGet(req *hspb.Key, ch chan hspb.Val) {
	if val, ok := cb.keyToVal[req.Key]; !ok {
		ch <- hspb.Val{Val: "", Code: &hspb.Code{Code: hspb.Code_STATUS_CODE_ER_NO_KEY}}
	} else {
		ch <- hspb.Val{Val: val, Code: &hspb.Code{Code: hspb.Code_STATUS_CODE_OK}}
	}
	close(ch)
}

func (cb *ClientBundle) ClientBundleDel(req *hspb.Key, ch chan hspb.Code) {
	if _, ok := cb.keyToVal[req.Key]; !ok {
		ch <- hspb.Code{Code: hspb.Code_STATUS_CODE_ER_NO_KEY}
	} else {
		delete(cb.keyToVal, req.Key)
		ch <- hspb.Code{Code: hspb.Code_STATUS_CODE_OK}
	}
	close(ch)
}

func (cb *ClientBundle) ClientBundleGetDump(ch chan hspb.KeyVal) {
	for key, val := range cb.keyToVal {
		ch <- hspb.KeyVal{Key: key, Val: val}
	}
	close(ch)
}

func (cb *ClientBundle) ClientBundleGetDumpSz(ch chan int) {
	ch <- len(cb.keyToVal)
	close(ch)
}

func NewClientBundle() *ClientBundle {
	ret := &ClientBundle{
		keyToVal: make(map[string]string),
		reqChan:  make(chan ControlMsg),
	}
	go ret.HandleClientBundle()
	return ret
}

type HashTableServerImpl struct {
	hspb.UnimplementedHashTableServer `json:"hspb_unimplemented_hash_table_server"`
	ClientBundleMap                   sync.Map `json:"maps,omitempty"`
}

func NewHashTableServerImpl() *HashTableServerImpl {
	return &HashTableServerImpl{}
}

func (s *HashTableServerImpl) GetClientBundle(ctx context.Context) (*ClientBundle, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, errors.New("issue getting client peer")
	}
	tlsInfo := p.AuthInfo.(credentials.TLSInfo)
	serial := fmt.Sprintf("%v", *tlsInfo.State.VerifiedChains[0][0].SerialNumber)
	if res, ok := s.ClientBundleMap.Load(serial); ok {
		if ret, ok := res.(*ClientBundle); ok {
			return ret, nil
		} else {
			return nil, errors.New("logic error")
		}
	}
	ret := NewClientBundle()
	s.ClientBundleMap.Store(serial, ret)
	return ret, nil
}

func (s *HashTableServerImpl) Add(ctx context.Context, keyVal *hspb.KeyVal) (*hspb.Code, error) {
	cb, err := s.GetClientBundle(ctx)
	if err != nil {
		return nil, err
	}
	ch := make(chan hspb.Code, 1)
	cb.reqChan <- ControlMsg{typeSM: ServiceMethodTypeAdd, valSM: keyVal, replyChan: ch}
	if reply, ok := <-ch; ok {
		return &reply, nil
	} else {
		return nil, errors.New("issue getting reply")
	}
}

func (s *HashTableServerImpl) Del(ctx context.Context, key *hspb.Key) (*hspb.Code, error) {
	cb, err := s.GetClientBundle(ctx)
	if err != nil {
		return nil, err
	}
	ch := make(chan hspb.Code, 1)
	cb.reqChan <- ControlMsg{typeSM: ServiceMethodTypeDel, valSM: key, replyChan: ch}
	if reply, ok := <-ch; ok {
		return &reply, nil
	} else {
		return nil, errors.New("issue getting reply")
	}
}

func (s *HashTableServerImpl) Get(ctx context.Context, key *hspb.Key) (*hspb.Val, error) {
	cb, err := s.GetClientBundle(ctx)
	if err != nil {
		return nil, err
	}
	ch := make(chan hspb.Val, 1)
	cb.reqChan <- ControlMsg{typeSM: ServiceMethodTypeDel, valSM: key, replyChan: ch}
	if reply, ok := <-ch; ok {
		return &reply, nil
	} else {
		return nil, errors.New("issue getting reply")
	}
}

func (s *HashTableServerImpl) Dump(_ *hspb.Empty, stream hspb.HashTable_DumpServer) error {
	cb, err := s.GetClientBundle(stream.Context())
	if err != nil {
		return err
	}
	chSz := make(chan int, 1)
	cb.reqChan <- ControlMsg{typeSM: ServiceMethodTypeDumpSz, replyChan: chSz}
	ch := make(chan hspb.KeyVal, <-chSz)
	cb.reqChan <- ControlMsg{typeSM: ServiceMethodTypeDump, replyChan: ch}
	//goland:noinspection GoVetCopyLock
	for keyVal := range ch {
		if err := stream.Send(&keyVal); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	var (
		certFile = flag.String("cert", "/home/mehdi/mzzcerts/localcerts/server.crt", "Server certificate public key file")
		keyFile  = flag.String("key", "/home/mehdi/mzzcerts/localcerts/server.key", "Server certificate secret key file")
		caDir    = flag.String("cadir", "/home/mehdi/mzzcerts/cacerts/", "Certificate Authority directory")
		addr     = flag.String("address", "localhost:8443", "Server address")
	)
	flag.Parse()
	if !strings.HasSuffix(*caDir, "/") {
		*caDir += "/"
	}
	caCertPool := x509.NewCertPool()
	caCertPoolMemCnt := 0
	caDirFiles, err := ioutil.ReadDir(*caDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range caDirFiles {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".pem") {
			continue
		}
		caFilePath := *caDir + entry.Name()
		caFile, err := ioutil.ReadFile(caFilePath)
		if err != nil {
			log.Fatal(err)
		}
		if !caCertPool.AppendCertsFromPEM(caFile) {
			log.Printf("Warning: Certificate %v not added.\n", caFilePath)
		} else {
			caCertPoolMemCnt++
		}
	}
	if caCertPoolMemCnt == 0 {
		log.Fatal(errors.New("no CA certificates"))
	}
	serverCert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatalf("failed to load certificate pair: %v: ", err)
	}
	listener, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
		VerifyPeerCertificate: func(_ [][]byte, verifiedChains [][]*x509.Certificate) error {
			if len(verifiedChains) != 1 {
				return errors.New("unhandled case")
			}
			if len(verifiedChains[0]) != 2 {
				return errors.New("invalid certificate chain length")
			}
			return nil
		},
		Certificates: []tls.Certificate{serverCert},
	}
	serverOpt := grpc.Creds(credentials.NewTLS(tlsConfig))
	grpcServer := grpc.NewServer(serverOpt)
	hspb.RegisterHashTableServer(grpcServer, NewHashTableServerImpl())
	reflection.Register(grpcServer)
	log.Fatal(grpcServer.Serve(listener))
}
