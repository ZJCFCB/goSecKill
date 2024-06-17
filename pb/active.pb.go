// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v5.26.1
// source: active.proto

package pb

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

type Activity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityId   int64   `protobuf:"varint,1,opt,name=ActivityId,proto3" json:"ActivityId,omitempty"`    // 活动Id
	ActivityName string  `protobuf:"bytes,2,opt,name=ActivityName,proto3" json:"ActivityName,omitempty"` // 活动名称
	ProductId    int64   `protobuf:"varint,3,opt,name=ProductId,proto3" json:"ProductId,omitempty"`      // 商品Id
	StartTime    int64   `protobuf:"varint,4,opt,name=StartTime,proto3" json:"StartTime,omitempty"`      // 开始时间
	EndTime      int64   `protobuf:"varint,5,opt,name=EndTime,proto3" json:"EndTime,omitempty"`          // 结束时间
	Total        int64   `protobuf:"varint,6,opt,name=Total,proto3" json:"Total,omitempty"`              // 商品总数
	Status       int64   `protobuf:"varint,7,opt,name=Status,proto3" json:"Status,omitempty"`            // 状态
	StartTimeStr string  `protobuf:"bytes,8,opt,name=StartTimeStr,proto3" json:"StartTimeStr,omitempty"` //
	EndTimeStr   string  `protobuf:"bytes,9,opt,name=EndTimeStr,proto3" json:"EndTimeStr,omitempty"`
	StatusStr    string  `protobuf:"bytes,10,opt,name=StatusStr,proto3" json:"StatusStr,omitempty"`
	Speed        int64   `protobuf:"varint,11,opt,name=Speed,proto3" json:"Speed,omitempty"`
	BuyLimit     int64   `protobuf:"varint,12,opt,name=BuyLimit,proto3" json:"BuyLimit,omitempty"`
	BuyRate      float64 `protobuf:"fixed64,13,opt,name=BuyRate,proto3" json:"BuyRate,omitempty"`
}

func (x *Activity) Reset() {
	*x = Activity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_active_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity) ProtoMessage() {}

func (x *Activity) ProtoReflect() protoreflect.Message {
	mi := &file_active_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity.ProtoReflect.Descriptor instead.
func (*Activity) Descriptor() ([]byte, []int) {
	return file_active_proto_rawDescGZIP(), []int{0}
}

func (x *Activity) GetActivityId() int64 {
	if x != nil {
		return x.ActivityId
	}
	return 0
}

func (x *Activity) GetActivityName() string {
	if x != nil {
		return x.ActivityName
	}
	return ""
}

func (x *Activity) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *Activity) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *Activity) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *Activity) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Activity) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Activity) GetStartTimeStr() string {
	if x != nil {
		return x.StartTimeStr
	}
	return ""
}

func (x *Activity) GetEndTimeStr() string {
	if x != nil {
		return x.EndTimeStr
	}
	return ""
}

func (x *Activity) GetStatusStr() string {
	if x != nil {
		return x.StatusStr
	}
	return ""
}

func (x *Activity) GetSpeed() int64 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *Activity) GetBuyLimit() int64 {
	if x != nil {
		return x.BuyLimit
	}
	return 0
}

func (x *Activity) GetBuyRate() float64 {
	if x != nil {
		return x.BuyRate
	}
	return 0
}

type SecProductInfoConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId         int64   `protobuf:"varint,1,opt,name=ProductId,proto3" json:"ProductId,omitempty"`                 // 商品id
	StartTime         int64   `protobuf:"varint,2,opt,name=StartTime,proto3" json:"StartTime,omitempty"`                 // 开始时间
	EndTime           int64   `protobuf:"varint,3,opt,name=EndTime,proto3" json:"EndTime,omitempty"`                     // 结束时间
	Status            int64   `protobuf:"varint,4,opt,name=Status,proto3" json:"Status,omitempty"`                       // 状态
	Total             int64   `protobuf:"varint,5,opt,name=Total,proto3" json:"Total,omitempty"`                         // 商品总数
	Left              int64   `protobuf:"varint,6,opt,name=Left,proto3" json:"Left,omitempty"`                           // 剩余商品数量
	OnePersonBuyLimit int64   `protobuf:"varint,7,opt,name=OnePersonBuyLimit,proto3" json:"OnePersonBuyLimit,omitempty"` // 一个人购买限制
	BuyRate           float64 `protobuf:"fixed64,8,opt,name=BuyRate,proto3" json:"BuyRate,omitempty"`                    // 买中几率
	SoldMaxLimit      int64   `protobuf:"varint,9,opt,name=SoldMaxLimit,proto3" json:"SoldMaxLimit,omitempty"`           // 每秒最多能卖多少个
}

func (x *SecProductInfoConf) Reset() {
	*x = SecProductInfoConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_active_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecProductInfoConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecProductInfoConf) ProtoMessage() {}

func (x *SecProductInfoConf) ProtoReflect() protoreflect.Message {
	mi := &file_active_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecProductInfoConf.ProtoReflect.Descriptor instead.
func (*SecProductInfoConf) Descriptor() ([]byte, []int) {
	return file_active_proto_rawDescGZIP(), []int{1}
}

func (x *SecProductInfoConf) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *SecProductInfoConf) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *SecProductInfoConf) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *SecProductInfoConf) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SecProductInfoConf) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SecProductInfoConf) GetLeft() int64 {
	if x != nil {
		return x.Left
	}
	return 0
}

func (x *SecProductInfoConf) GetOnePersonBuyLimit() int64 {
	if x != nil {
		return x.OnePersonBuyLimit
	}
	return 0
}

func (x *SecProductInfoConf) GetBuyRate() float64 {
	if x != nil {
		return x.BuyRate
	}
	return 0
}

func (x *SecProductInfoConf) GetSoldMaxLimit() int64 {
	if x != nil {
		return x.SoldMaxLimit
	}
	return 0
}

var File_active_proto protoreflect.FileDescriptor

var file_active_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x73, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x03,
	0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x45, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x53,
	0x74, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x53, 0x74, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x53, 0x74, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x45, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x53, 0x74, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x53, 0x74, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x53, 0x74, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x70, 0x65, 0x65, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x75,
	0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x42, 0x75,
	0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x75, 0x79, 0x52, 0x61, 0x74,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x42, 0x75, 0x79, 0x52, 0x61, 0x74, 0x65,
	0x22, 0x98, 0x02, 0x0a, 0x12, 0x53, 0x65, 0x63, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x4c,
	0x65, 0x66, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x4c, 0x65, 0x66, 0x74, 0x12,
	0x2c, 0x0a, 0x11, 0x4f, 0x6e, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42, 0x75, 0x79, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x4f, 0x6e, 0x65, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42, 0x75, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x42, 0x75, 0x79, 0x52, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07,
	0x42, 0x75, 0x79, 0x52, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x6f, 0x6c, 0x64, 0x4d,
	0x61, 0x78, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x53,
	0x6f, 0x6c, 0x64, 0x4d, 0x61, 0x78, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x32, 0x39, 0x0a, 0x0f, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26,
	0x0a, 0x07, 0x73, 0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x12, 0x0b, 0x2e, 0x53, 0x65, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x53, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_active_proto_rawDescOnce sync.Once
	file_active_proto_rawDescData = file_active_proto_rawDesc
)

func file_active_proto_rawDescGZIP() []byte {
	file_active_proto_rawDescOnce.Do(func() {
		file_active_proto_rawDescData = protoimpl.X.CompressGZIP(file_active_proto_rawDescData)
	})
	return file_active_proto_rawDescData
}

var file_active_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_active_proto_goTypes = []interface{}{
	(*Activity)(nil),           // 0: Activity
	(*SecProductInfoConf)(nil), // 1: SecProductInfoConf
	(*SecRequest)(nil),         // 2: SecRequest
	(*SecResponse)(nil),        // 3: SecResponse
}
var file_active_proto_depIdxs = []int32{
	2, // 0: ActivityService.secKill:input_type -> SecRequest
	3, // 1: ActivityService.secKill:output_type -> SecResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_active_proto_init() }
func file_active_proto_init() {
	if File_active_proto != nil {
		return
	}
	file_seckill_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_active_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activity); i {
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
		file_active_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecProductInfoConf); i {
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
			RawDescriptor: file_active_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_active_proto_goTypes,
		DependencyIndexes: file_active_proto_depIdxs,
		MessageInfos:      file_active_proto_msgTypes,
	}.Build()
	File_active_proto = out.File
	file_active_proto_rawDesc = nil
	file_active_proto_goTypes = nil
	file_active_proto_depIdxs = nil
}