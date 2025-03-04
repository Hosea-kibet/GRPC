// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.15.8
// source: coffee_shop.proto

package GRPC_EXAMPLE

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

type MenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MenuRequest) Reset() {
	*x = MenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coffee_shop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuRequest) ProtoMessage() {}

func (x *MenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coffee_shop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuRequest.ProtoReflect.Descriptor instead.
func (*MenuRequest) Descriptor() ([]byte, []int) {
	return file_coffee_shop_proto_rawDescGZIP(), []int{0}
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// repeated indicates that we are working with lists
	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coffee_shop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_coffee_shop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_coffee_shop_proto_rawDescGZIP(), []int{1}
}

func (x *Order) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Receipt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Receipt) Reset() {
	*x = Receipt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coffee_shop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Receipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Receipt) ProtoMessage() {}

func (x *Receipt) ProtoReflect() protoreflect.Message {
	mi := &file_coffee_shop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Receipt.ProtoReflect.Descriptor instead.
func (*Receipt) Descriptor() ([]byte, []int) {
	return file_coffee_shop_proto_rawDescGZIP(), []int{2}
}

func (x *Receipt) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type OrderStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *OrderStatusResponse) Reset() {
	*x = OrderStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coffee_shop_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderStatusResponse) ProtoMessage() {}

func (x *OrderStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coffee_shop_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderStatusResponse.ProtoReflect.Descriptor instead.
func (*OrderStatusResponse) Descriptor() ([]byte, []int) {
	return file_coffee_shop_proto_rawDescGZIP(), []int{3}
}

func (x *OrderStatusResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Menu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Menu) Reset() {
	*x = Menu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coffee_shop_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Menu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Menu) ProtoMessage() {}

func (x *Menu) ProtoReflect() protoreflect.Message {
	mi := &file_coffee_shop_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Menu.ProtoReflect.Descriptor instead.
func (*Menu) Descriptor() ([]byte, []int) {
	return file_coffee_shop_proto_rawDescGZIP(), []int{4}
}

func (x *Menu) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coffee_shop_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_coffee_shop_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_coffee_shop_proto_rawDescGZIP(), []int{5}
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_coffee_shop_proto protoreflect.FileDescriptor

var file_coffee_shop_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x22,
	0x0d, 0x0a, 0x0b, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2f,
	0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0x19, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x13, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x2e, 0x0a, 0x04, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x26, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x66,
	0x66, 0x65, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x2a, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x32,
	0xc5, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x12, 0x38,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x66, 0x66,
	0x65, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x4d, 0x65, 0x6e, 0x75, 0x22, 0x00, 0x30, 0x01, 0x12, 0x36, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x66, 0x66,
	0x65, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x22, 0x00,
	0x12, 0x45, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x13, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x68, 0x6f,
	0x70, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x6f, 0x73, 0x65, 0x61, 0x2d, 0x6b, 0x69, 0x62, 0x65,
	0x74, 0x2f, 0x47, 0x52, 0x50, 0x43, 0x5f, 0x45, 0x58, 0x41, 0x4d, 0x50, 0x4c, 0x45, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coffee_shop_proto_rawDescOnce sync.Once
	file_coffee_shop_proto_rawDescData = file_coffee_shop_proto_rawDesc
)

func file_coffee_shop_proto_rawDescGZIP() []byte {
	file_coffee_shop_proto_rawDescOnce.Do(func() {
		file_coffee_shop_proto_rawDescData = protoimpl.X.CompressGZIP(file_coffee_shop_proto_rawDescData)
	})
	return file_coffee_shop_proto_rawDescData
}

var file_coffee_shop_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_coffee_shop_proto_goTypes = []interface{}{
	(*MenuRequest)(nil),         // 0: coffeeshop.MenuRequest
	(*Order)(nil),               // 1: coffeeshop.Order
	(*Receipt)(nil),             // 2: coffeeshop.Receipt
	(*OrderStatusResponse)(nil), // 3: coffeeshop.OrderStatusResponse
	(*Menu)(nil),                // 4: coffeeshop.Menu
	(*Item)(nil),                // 5: coffeeshop.Item
}
var file_coffee_shop_proto_depIdxs = []int32{
	5, // 0: coffeeshop.Order.items:type_name -> coffeeshop.Item
	5, // 1: coffeeshop.Menu.items:type_name -> coffeeshop.Item
	0, // 2: coffeeshop.CoffeeShop.GetMenu:input_type -> coffeeshop.MenuRequest
	1, // 3: coffeeshop.CoffeeShop.PlaceOrder:input_type -> coffeeshop.Order
	2, // 4: coffeeshop.CoffeeShop.OrderStatus:input_type -> coffeeshop.Receipt
	4, // 5: coffeeshop.CoffeeShop.GetMenu:output_type -> coffeeshop.Menu
	2, // 6: coffeeshop.CoffeeShop.PlaceOrder:output_type -> coffeeshop.Receipt
	3, // 7: coffeeshop.CoffeeShop.OrderStatus:output_type -> coffeeshop.OrderStatusResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_coffee_shop_proto_init() }
func file_coffee_shop_proto_init() {
	if File_coffee_shop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coffee_shop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuRequest); i {
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
		file_coffee_shop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_coffee_shop_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Receipt); i {
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
		file_coffee_shop_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderStatusResponse); i {
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
		file_coffee_shop_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Menu); i {
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
		file_coffee_shop_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
			RawDescriptor: file_coffee_shop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coffee_shop_proto_goTypes,
		DependencyIndexes: file_coffee_shop_proto_depIdxs,
		MessageInfos:      file_coffee_shop_proto_msgTypes,
	}.Build()
	File_coffee_shop_proto = out.File
	file_coffee_shop_proto_rawDesc = nil
	file_coffee_shop_proto_goTypes = nil
	file_coffee_shop_proto_depIdxs = nil
}
