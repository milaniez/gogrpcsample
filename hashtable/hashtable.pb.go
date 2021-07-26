// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: hashtable/hashtable.proto

package hashtable

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Code_StatusCode int32

const (
	Code_STATUS_CODE_OK        Code_StatusCode = 0
	Code_STATUS_CODE_ER_DUP    Code_StatusCode = 1
	Code_STATUS_CODE_ER_NO_KEY Code_StatusCode = 2
)

// Enum value maps for Code_StatusCode.
var (
	Code_StatusCode_name = map[int32]string{
		0: "STATUS_CODE_OK",
		1: "STATUS_CODE_ER_DUP",
		2: "STATUS_CODE_ER_NO_KEY",
	}
	Code_StatusCode_value = map[string]int32{
		"STATUS_CODE_OK":        0,
		"STATUS_CODE_ER_DUP":    1,
		"STATUS_CODE_ER_NO_KEY": 2,
	}
)

func (x Code_StatusCode) Enum() *Code_StatusCode {
	p := new(Code_StatusCode)
	*p = x
	return p
}

func (x Code_StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code_StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_hashtable_hashtable_proto_enumTypes[0].Descriptor()
}

func (Code_StatusCode) Type() protoreflect.EnumType {
	return &file_hashtable_hashtable_proto_enumTypes[0]
}

func (x Code_StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code_StatusCode.Descriptor instead.
func (Code_StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_hashtable_hashtable_proto_rawDescGZIP(), []int{1, 0}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashtable_hashtable_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_hashtable_hashtable_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_hashtable_hashtable_proto_rawDescGZIP(), []int{0}
}

type Code struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code Code_StatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=hashtable.Code_StatusCode" json:"code,omitempty"`
}

func (x *Code) Reset() {
	*x = Code{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashtable_hashtable_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code) ProtoMessage() {}

func (x *Code) ProtoReflect() protoreflect.Message {
	mi := &file_hashtable_hashtable_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Code.ProtoReflect.Descriptor instead.
func (*Code) Descriptor() ([]byte, []int) {
	return file_hashtable_hashtable_proto_rawDescGZIP(), []int{1}
}

func (x *Code) GetCode() Code_StatusCode {
	if x != nil {
		return x.Code
	}
	return Code_STATUS_CODE_OK
}

type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashtable_hashtable_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_hashtable_hashtable_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_hashtable_hashtable_proto_rawDescGZIP(), []int{2}
}

func (x *Key) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type Val struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code *Code  `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Val  string `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *Val) Reset() {
	*x = Val{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashtable_hashtable_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Val) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Val) ProtoMessage() {}

func (x *Val) ProtoReflect() protoreflect.Message {
	mi := &file_hashtable_hashtable_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Val.ProtoReflect.Descriptor instead.
func (*Val) Descriptor() ([]byte, []int) {
	return file_hashtable_hashtable_proto_rawDescGZIP(), []int{3}
}

func (x *Val) GetCode() *Code {
	if x != nil {
		return x.Code
	}
	return nil
}

func (x *Val) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

type KeyVal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val string `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *KeyVal) Reset() {
	*x = KeyVal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hashtable_hashtable_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyVal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyVal) ProtoMessage() {}

func (x *KeyVal) ProtoReflect() protoreflect.Message {
	mi := &file_hashtable_hashtable_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyVal.ProtoReflect.Descriptor instead.
func (*KeyVal) Descriptor() ([]byte, []int) {
	return file_hashtable_hashtable_proto_rawDescGZIP(), []int{4}
}

func (x *KeyVal) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyVal) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

var File_hashtable_hashtable_proto protoreflect.FileDescriptor

var file_hashtable_hashtable_proto_rawDesc = []byte{
	0x0a, 0x19, 0x68, 0x61, 0x73, 0x68, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x68, 0x61, 0x73, 0x68,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x68, 0x61, 0x73,
	0x68, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x8b, 0x01, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x53, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x5f, 0x44, 0x55, 0x50,
	0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0x02, 0x22, 0x17, 0x0a,
	0x03, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x3c, 0x0a, 0x03, 0x56, 0x61, 0x6c, 0x12, 0x23, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x68, 0x61,
	0x73, 0x68, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x76, 0x61, 0x6c, 0x22, 0x2c, 0x0a, 0x06, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76,
	0x61, 0x6c, 0x32, 0xbc, 0x01, 0x0a, 0x09, 0x48, 0x61, 0x73, 0x68, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x2b, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x11, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x1a, 0x0f, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a,
	0x03, 0x44, 0x65, 0x6c, 0x12, 0x0e, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x4b, 0x65, 0x79, 0x1a, 0x0f, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0e,
	0x2e, 0x68, 0x61, 0x73, 0x68, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x1a, 0x0e,
	0x2e, 0x68, 0x61, 0x73, 0x68, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x22, 0x00,
	0x12, 0x2f, 0x0a, 0x04, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x10, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x22, 0x00, 0x30,
	0x01, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x52, 0x50, 0x43, 0x54, 0x65, 0x73, 0x74, 0x30, 0x30, 0x31,
	0x2f, 0x68, 0x61, 0x73, 0x68, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_hashtable_hashtable_proto_rawDescOnce sync.Once
	file_hashtable_hashtable_proto_rawDescData = file_hashtable_hashtable_proto_rawDesc
)

func file_hashtable_hashtable_proto_rawDescGZIP() []byte {
	file_hashtable_hashtable_proto_rawDescOnce.Do(func() {
		file_hashtable_hashtable_proto_rawDescData = protoimpl.X.CompressGZIP(file_hashtable_hashtable_proto_rawDescData)
	})
	return file_hashtable_hashtable_proto_rawDescData
}

var file_hashtable_hashtable_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_hashtable_hashtable_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_hashtable_hashtable_proto_goTypes = []interface{}{
	(Code_StatusCode)(0), // 0: hashtable.Code.StatusCode
	(*Empty)(nil),        // 1: hashtable.Empty
	(*Code)(nil),         // 2: hashtable.Code
	(*Key)(nil),          // 3: hashtable.Key
	(*Val)(nil),          // 4: hashtable.Val
	(*KeyVal)(nil),       // 5: hashtable.KeyVal
}
var file_hashtable_hashtable_proto_depIdxs = []int32{
	0, // 0: hashtable.Code.code:type_name -> hashtable.Code.StatusCode
	2, // 1: hashtable.Val.code:type_name -> hashtable.Code
	5, // 2: hashtable.HashTable.Add:input_type -> hashtable.KeyVal
	3, // 3: hashtable.HashTable.Del:input_type -> hashtable.Key
	3, // 4: hashtable.HashTable.Get:input_type -> hashtable.Key
	1, // 5: hashtable.HashTable.Dump:input_type -> hashtable.Empty
	2, // 6: hashtable.HashTable.Add:output_type -> hashtable.Code
	2, // 7: hashtable.HashTable.Del:output_type -> hashtable.Code
	4, // 8: hashtable.HashTable.Get:output_type -> hashtable.Val
	5, // 9: hashtable.HashTable.Dump:output_type -> hashtable.KeyVal
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_hashtable_hashtable_proto_init() }
func file_hashtable_hashtable_proto_init() {
	if File_hashtable_hashtable_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hashtable_hashtable_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hashtable_hashtable_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Code); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hashtable_hashtable_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hashtable_hashtable_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Val); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_hashtable_hashtable_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyVal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hashtable_hashtable_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hashtable_hashtable_proto_goTypes,
		DependencyIndexes: file_hashtable_hashtable_proto_depIdxs,
		EnumInfos:         file_hashtable_hashtable_proto_enumTypes,
		MessageInfos:      file_hashtable_hashtable_proto_msgTypes,
	}.Build()
	File_hashtable_hashtable_proto = out.File
	file_hashtable_hashtable_proto_rawDesc = nil
	file_hashtable_hashtable_proto_goTypes = nil
	file_hashtable_hashtable_proto_depIdxs = nil
}
