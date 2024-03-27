// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.1
// source: goods.proto

package goods

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

type GoodsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    int64  `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Name  string `protobuf:"bytes,20,opt,name=Name,proto3" json:"Name,omitempty"`
	Price string `protobuf:"bytes,30,opt,name=Price,proto3" json:"Price,omitempty"`
	Stock int64  `protobuf:"varint,40,opt,name=Stock,proto3" json:"Stock,omitempty"`
}

func (x *GoodsInfo) Reset() {
	*x = GoodsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodsInfo) ProtoMessage() {}

func (x *GoodsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodsInfo.ProtoReflect.Descriptor instead.
func (*GoodsInfo) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{0}
}

func (x *GoodsInfo) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *GoodsInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GoodsInfo) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *GoodsInfo) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type GetGoodsByIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDs []int64 `protobuf:"varint,10,rep,packed,name=IDs,proto3" json:"IDs,omitempty"`
}

func (x *GetGoodsByIdsRequest) Reset() {
	*x = GetGoodsByIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsByIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsByIdsRequest) ProtoMessage() {}

func (x *GetGoodsByIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsByIdsRequest.ProtoReflect.Descriptor instead.
func (*GetGoodsByIdsRequest) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{1}
}

func (x *GetGoodsByIdsRequest) GetIDs() []int64 {
	if x != nil {
		return x.IDs
	}
	return nil
}

type GetGoodsByIdsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*GoodsInfo `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *GetGoodsByIdsResponse) Reset() {
	*x = GetGoodsByIdsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsByIdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsByIdsResponse) ProtoMessage() {}

func (x *GetGoodsByIdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsByIdsResponse.ProtoReflect.Descriptor instead.
func (*GetGoodsByIdsResponse) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{2}
}

func (x *GetGoodsByIdsResponse) GetInfos() []*GoodsInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

type UpdateStockReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID  int64 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Num int64 `protobuf:"varint,20,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (x *UpdateStockReq) Reset() {
	*x = UpdateStockReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStockReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStockReq) ProtoMessage() {}

func (x *UpdateStockReq) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStockReq.ProtoReflect.Descriptor instead.
func (*UpdateStockReq) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateStockReq) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateStockReq) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type UpdateStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodsInfos []*UpdateStockReq `protobuf:"bytes,10,rep,name=GoodsInfos,proto3" json:"GoodsInfos,omitempty"`
}

func (x *UpdateStockRequest) Reset() {
	*x = UpdateStockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStockRequest) ProtoMessage() {}

func (x *UpdateStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStockRequest.ProtoReflect.Descriptor instead.
func (*UpdateStockRequest) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateStockRequest) GetGoodsInfos() []*UpdateStockReq {
	if x != nil {
		return x.GoodsInfos
	}
	return nil
}

type UpdateStockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*GoodsInfo `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *UpdateStockResponse) Reset() {
	*x = UpdateStockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStockResponse) ProtoMessage() {}

func (x *UpdateStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goods_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStockResponse.ProtoReflect.Descriptor instead.
func (*UpdateStockResponse) Descriptor() ([]byte, []int) {
	return file_goods_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateStockResponse) GetInfos() []*GoodsInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

var File_goods_proto protoreflect.FileDescriptor

var file_goods_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x22, 0x5b, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x18, 0x28, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x22, 0x28, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x79, 0x49,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x44, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x49, 0x44, 0x73, 0x22, 0x3f, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x32, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x10,
	0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x4e, 0x75, 0x6d,
	0x22, 0x4b, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0a, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x52, 0x0a, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x3d, 0x0a,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x32, 0x99, 0x01, 0x0a,
	0x05, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x4a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x12, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goods_proto_rawDescOnce sync.Once
	file_goods_proto_rawDescData = file_goods_proto_rawDesc
)

func file_goods_proto_rawDescGZIP() []byte {
	file_goods_proto_rawDescOnce.Do(func() {
		file_goods_proto_rawDescData = protoimpl.X.CompressGZIP(file_goods_proto_rawDescData)
	})
	return file_goods_proto_rawDescData
}

var file_goods_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_goods_proto_goTypes = []interface{}{
	(*GoodsInfo)(nil),             // 0: goods.GoodsInfo
	(*GetGoodsByIdsRequest)(nil),  // 1: goods.GetGoodsByIdsRequest
	(*GetGoodsByIdsResponse)(nil), // 2: goods.GetGoodsByIdsResponse
	(*UpdateStockReq)(nil),        // 3: goods.UpdateStockReq
	(*UpdateStockRequest)(nil),    // 4: goods.UpdateStockRequest
	(*UpdateStockResponse)(nil),   // 5: goods.UpdateStockResponse
}
var file_goods_proto_depIdxs = []int32{
	0, // 0: goods.GetGoodsByIdsResponse.Infos:type_name -> goods.GoodsInfo
	3, // 1: goods.UpdateStockRequest.GoodsInfos:type_name -> goods.UpdateStockReq
	0, // 2: goods.UpdateStockResponse.Infos:type_name -> goods.GoodsInfo
	1, // 3: goods.Goods.GetGoodsByIds:input_type -> goods.GetGoodsByIdsRequest
	4, // 4: goods.Goods.UpdateStock:input_type -> goods.UpdateStockRequest
	2, // 5: goods.Goods.GetGoodsByIds:output_type -> goods.GetGoodsByIdsResponse
	5, // 6: goods.Goods.UpdateStock:output_type -> goods.UpdateStockResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_goods_proto_init() }
func file_goods_proto_init() {
	if File_goods_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goods_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodsInfo); i {
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
		file_goods_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsByIdsRequest); i {
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
		file_goods_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsByIdsResponse); i {
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
		file_goods_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStockReq); i {
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
		file_goods_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStockRequest); i {
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
		file_goods_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStockResponse); i {
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
			RawDescriptor: file_goods_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goods_proto_goTypes,
		DependencyIndexes: file_goods_proto_depIdxs,
		MessageInfos:      file_goods_proto_msgTypes,
	}.Build()
	File_goods_proto = out.File
	file_goods_proto_rawDesc = nil
	file_goods_proto_goTypes = nil
	file_goods_proto_depIdxs = nil
}
