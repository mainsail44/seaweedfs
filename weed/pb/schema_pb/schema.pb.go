// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: schema.proto

package schema_pb

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

type ScalarType int32

const (
	ScalarType_BOOLEAN ScalarType = 0
	ScalarType_INTEGER ScalarType = 1
	ScalarType_LONG    ScalarType = 3
	ScalarType_FLOAT   ScalarType = 4
	ScalarType_DOUBLE  ScalarType = 5
	ScalarType_BYTES   ScalarType = 6
	ScalarType_STRING  ScalarType = 7
)

// Enum value maps for ScalarType.
var (
	ScalarType_name = map[int32]string{
		0: "BOOLEAN",
		1: "INTEGER",
		3: "LONG",
		4: "FLOAT",
		5: "DOUBLE",
		6: "BYTES",
		7: "STRING",
	}
	ScalarType_value = map[string]int32{
		"BOOLEAN": 0,
		"INTEGER": 1,
		"LONG":    3,
		"FLOAT":   4,
		"DOUBLE":  5,
		"BYTES":   6,
		"STRING":  7,
	}
)

func (x ScalarType) Enum() *ScalarType {
	p := new(ScalarType)
	*p = x
	return p
}

func (x ScalarType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScalarType) Descriptor() protoreflect.EnumDescriptor {
	return file_schema_proto_enumTypes[0].Descriptor()
}

func (ScalarType) Type() protoreflect.EnumType {
	return &file_schema_proto_enumTypes[0]
}

func (x ScalarType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScalarType.Descriptor instead.
func (ScalarType) EnumDescriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{0}
}

type RecordType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []*Field `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *RecordType) Reset() {
	*x = RecordType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordType) ProtoMessage() {}

func (x *RecordType) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordType.ProtoReflect.Descriptor instead.
func (*RecordType) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{0}
}

func (x *RecordType) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type       *Type  `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Index      int32  `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	IsOptional bool   `protobuf:"varint,4,opt,name=is_optional,json=isOptional,proto3" json:"is_optional,omitempty"`
	IsRepeated bool   `protobuf:"varint,5,opt,name=is_repeated,json=isRepeated,proto3" json:"is_repeated,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{1}
}

func (x *Field) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Field) GetType() *Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Field) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Field) GetIsOptional() bool {
	if x != nil {
		return x.IsOptional
	}
	return false
}

func (x *Field) GetIsRepeated() bool {
	if x != nil {
		return x.IsRepeated
	}
	return false
}

type Type struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//
	//	*Type_ScalarType
	//	*Type_RecordType
	//	*Type_MapType
	Kind isType_Kind `protobuf_oneof:"kind"`
}

func (x *Type) Reset() {
	*x = Type{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Type) ProtoMessage() {}

func (x *Type) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Type.ProtoReflect.Descriptor instead.
func (*Type) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{2}
}

func (m *Type) GetKind() isType_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Type) GetScalarType() ScalarType {
	if x, ok := x.GetKind().(*Type_ScalarType); ok {
		return x.ScalarType
	}
	return ScalarType_BOOLEAN
}

func (x *Type) GetRecordType() *RecordType {
	if x, ok := x.GetKind().(*Type_RecordType); ok {
		return x.RecordType
	}
	return nil
}

func (x *Type) GetMapType() *MapType {
	if x, ok := x.GetKind().(*Type_MapType); ok {
		return x.MapType
	}
	return nil
}

type isType_Kind interface {
	isType_Kind()
}

type Type_ScalarType struct {
	ScalarType ScalarType `protobuf:"varint,1,opt,name=scalar_type,json=scalarType,proto3,enum=schema_pb.ScalarType,oneof"`
}

type Type_RecordType struct {
	RecordType *RecordType `protobuf:"bytes,2,opt,name=record_type,json=recordType,proto3,oneof"`
}

type Type_MapType struct {
	MapType *MapType `protobuf:"bytes,3,opt,name=map_type,json=mapType,proto3,oneof"`
}

func (*Type_ScalarType) isType_Kind() {}

func (*Type_RecordType) isType_Kind() {}

func (*Type_MapType) isType_Kind() {}

type MapType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *Type  `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MapType) Reset() {
	*x = MapType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapType) ProtoMessage() {}

func (x *MapType) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapType.ProtoReflect.Descriptor instead.
func (*MapType) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{3}
}

func (x *MapType) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *MapType) GetValue() *Type {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_schema_proto protoreflect.FileDescriptor

var file_schema_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x70, 0x62, 0x22, 0x36, 0x0a, 0x0a, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x5f, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x22, 0x98, 0x01, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x70, 0x62, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x69, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x73, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x69, 0x73, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0xb3, 0x01, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x38, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x6d, 0x61, 0x70,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x5f, 0x70, 0x62, 0x2e, 0x4d, 0x61, 0x70, 0x54, 0x79, 0x70, 0x65, 0x48,
	0x00, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x22, 0x42, 0x0a, 0x07, 0x4d, 0x61, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x70, 0x62, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x5e, 0x0a, 0x0a, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4f, 0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x45, 0x52, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x4c, 0x4f, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x4c, 0x4f, 0x41,
	0x54, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x05, 0x12,
	0x09, 0x0a, 0x05, 0x42, 0x59, 0x54, 0x45, 0x53, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54,
	0x52, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x61, 0x77, 0x65, 0x65, 0x64, 0x66, 0x73, 0x2f, 0x73,
	0x65, 0x61, 0x77, 0x65, 0x65, 0x64, 0x66, 0x73, 0x2f, 0x77, 0x65, 0x65, 0x64, 0x2f, 0x70, 0x62,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_schema_proto_rawDescOnce sync.Once
	file_schema_proto_rawDescData = file_schema_proto_rawDesc
)

func file_schema_proto_rawDescGZIP() []byte {
	file_schema_proto_rawDescOnce.Do(func() {
		file_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_proto_rawDescData)
	})
	return file_schema_proto_rawDescData
}

var file_schema_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_schema_proto_goTypes = []interface{}{
	(ScalarType)(0),    // 0: schema_pb.ScalarType
	(*RecordType)(nil), // 1: schema_pb.RecordType
	(*Field)(nil),      // 2: schema_pb.Field
	(*Type)(nil),       // 3: schema_pb.Type
	(*MapType)(nil),    // 4: schema_pb.MapType
}
var file_schema_proto_depIdxs = []int32{
	2, // 0: schema_pb.RecordType.fields:type_name -> schema_pb.Field
	3, // 1: schema_pb.Field.type:type_name -> schema_pb.Type
	0, // 2: schema_pb.Type.scalar_type:type_name -> schema_pb.ScalarType
	1, // 3: schema_pb.Type.record_type:type_name -> schema_pb.RecordType
	4, // 4: schema_pb.Type.map_type:type_name -> schema_pb.MapType
	3, // 5: schema_pb.MapType.value:type_name -> schema_pb.Type
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_schema_proto_init() }
func file_schema_proto_init() {
	if File_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordType); i {
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
		file_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
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
		file_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Type); i {
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
		file_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapType); i {
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
	file_schema_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Type_ScalarType)(nil),
		(*Type_RecordType)(nil),
		(*Type_MapType)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schema_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_proto_goTypes,
		DependencyIndexes: file_schema_proto_depIdxs,
		EnumInfos:         file_schema_proto_enumTypes,
		MessageInfos:      file_schema_proto_msgTypes,
	}.Build()
	File_schema_proto = out.File
	file_schema_proto_rawDesc = nil
	file_schema_proto_goTypes = nil
	file_schema_proto_depIdxs = nil
}
