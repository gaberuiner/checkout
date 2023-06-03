// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: listCart.proto

package listCart

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User uint64 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listCart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_listCart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_listCart_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetUser() uint64 {
	if x != nil {
		return x.User
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items      []*Items `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	TotalPrice uint32   `protobuf:"varint,2,opt,name=totalPrice,proto3" json:"totalPrice,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listCart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_listCart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_listCart_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetItems() []*Items {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Response) GetTotalPrice() uint32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type Items struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku   uint32 `protobuf:"varint,1,opt,name=sku,proto3" json:"sku,omitempty"`
	Count uint32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Price uint32 `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Items) Reset() {
	*x = Items{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listCart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Items) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Items) ProtoMessage() {}

func (x *Items) ProtoReflect() protoreflect.Message {
	mi := &file_listCart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Items.ProtoReflect.Descriptor instead.
func (*Items) Descriptor() ([]byte, []int) {
	return file_listCart_proto_rawDescGZIP(), []int{2}
}

func (x *Items) GetSku() uint32 {
	if x != nil {
		return x.Sku
	}
	return 0
}

func (x *Items) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Items) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Items) GetPrice() uint32 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_listCart_proto protoreflect.FileDescriptor

var file_listCart_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x48, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x59, 0x0a, 0x05, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x73, 0x6b, 0x75, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x32, 0x29, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74,
	0x12, 0x1d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x08, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x1c, 0x5a, 0x1a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x32, 0x35, 0x36, 0x2f, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x6f, 0x75, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_listCart_proto_rawDescOnce sync.Once
	file_listCart_proto_rawDescData = file_listCart_proto_rawDesc
)

func file_listCart_proto_rawDescGZIP() []byte {
	file_listCart_proto_rawDescOnce.Do(func() {
		file_listCart_proto_rawDescData = protoimpl.X.CompressGZIP(file_listCart_proto_rawDescData)
	})
	return file_listCart_proto_rawDescData
}

var file_listCart_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_listCart_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: Request
	(*Response)(nil), // 1: Response
	(*Items)(nil),    // 2: Items
}
var file_listCart_proto_depIdxs = []int32{
	2, // 0: Response.items:type_name -> Items
	0, // 1: ListCart.Create:input_type -> Request
	1, // 2: ListCart.Create:output_type -> Response
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_listCart_proto_init() }
func file_listCart_proto_init() {
	if File_listCart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_listCart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_listCart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_listCart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Items); i {
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
			RawDescriptor: file_listCart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_listCart_proto_goTypes,
		DependencyIndexes: file_listCart_proto_depIdxs,
		MessageInfos:      file_listCart_proto_msgTypes,
	}.Build()
	File_listCart_proto = out.File
	file_listCart_proto_rawDesc = nil
	file_listCart_proto_goTypes = nil
	file_listCart_proto_depIdxs = nil
}
