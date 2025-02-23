// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: Convert/convert.proto

package conversionv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConvertRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	TransactionId    string                 `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`             // UUID из Transaction Service
	AmountRubKopecks int64                  `protobuf:"varint,2,opt,name=amount_rub_kopecks,json=amountRubKopecks,proto3" json:"amount_rub_kopecks,omitempty"` // Сумма в копейках (100 RUB = 10000)
	CryptoCurrency   string                 `protobuf:"bytes,3,opt,name=crypto_currency,json=cryptoCurrency,proto3" json:"crypto_currency,omitempty"`          // "BTC", "ETH", "USDT"
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *ConvertRequest) Reset() {
	*x = ConvertRequest{}
	mi := &file_Convert_convert_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConvertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertRequest) ProtoMessage() {}

func (x *ConvertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Convert_convert_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertRequest.ProtoReflect.Descriptor instead.
func (*ConvertRequest) Descriptor() ([]byte, []int) {
	return file_Convert_convert_proto_rawDescGZIP(), []int{0}
}

func (x *ConvertRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *ConvertRequest) GetAmountRubKopecks() int64 {
	if x != nil {
		return x.AmountRubKopecks
	}
	return 0
}

func (x *ConvertRequest) GetCryptoCurrency() string {
	if x != nil {
		return x.CryptoCurrency
	}
	return ""
}

type ConvertResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	CryptoAmount  int64                  `protobuf:"varint,2,opt,name=crypto_amount,json=cryptoAmount,proto3" json:"crypto_amount,omitempty"` // Сумма в минимальных единицах
	Error         string                 `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`                                    // Детали ошибки (если success=false)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConvertResponse) Reset() {
	*x = ConvertResponse{}
	mi := &file_Convert_convert_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConvertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertResponse) ProtoMessage() {}

func (x *ConvertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Convert_convert_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertResponse.ProtoReflect.Descriptor instead.
func (*ConvertResponse) Descriptor() ([]byte, []int) {
	return file_Convert_convert_proto_rawDescGZIP(), []int{1}
}

func (x *ConvertResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ConvertResponse) GetCryptoAmount() int64 {
	if x != nil {
		return x.CryptoAmount
	}
	return 0
}

func (x *ConvertResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type RateRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	CryptoCurrency string                 `protobuf:"bytes,1,opt,name=crypto_currency,json=cryptoCurrency,proto3" json:"crypto_currency,omitempty"` // "BTC", "ETH", "USDT"
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *RateRequest) Reset() {
	*x = RateRequest{}
	mi := &file_Convert_convert_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateRequest) ProtoMessage() {}

func (x *RateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Convert_convert_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateRequest.ProtoReflect.Descriptor instead.
func (*RateRequest) Descriptor() ([]byte, []int) {
	return file_Convert_convert_proto_rawDescGZIP(), []int{2}
}

func (x *RateRequest) GetCryptoCurrency() string {
	if x != nil {
		return x.CryptoCurrency
	}
	return ""
}

type RateResponse struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	RateKopecksPerUnit int64                  `protobuf:"varint,1,opt,name=rate_kopecks_per_unit,json=rateKopecksPerUnit,proto3" json:"rate_kopecks_per_unit,omitempty"` // Стоимость 1 единицы крипты в копейках
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *RateResponse) Reset() {
	*x = RateResponse{}
	mi := &file_Convert_convert_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateResponse) ProtoMessage() {}

func (x *RateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Convert_convert_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateResponse.ProtoReflect.Descriptor instead.
func (*RateResponse) Descriptor() ([]byte, []int) {
	return file_Convert_convert_proto_rawDescGZIP(), []int{3}
}

func (x *RateResponse) GetRateKopecksPerUnit() int64 {
	if x != nil {
		return x.RateKopecksPerUnit
	}
	return 0
}

var File_Convert_convert_proto protoreflect.FileDescriptor

var file_Convert_convert_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x8e, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x12, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x75, 0x62, 0x5f, 0x6b, 0x6f, 0x70, 0x65,
	0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x75, 0x62, 0x4b, 0x6f, 0x70, 0x65, 0x63, 0x6b, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x22, 0x66, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x36, 0x0a, 0x0b,
	0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x22, 0x41, 0x0a, 0x0c, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x15, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x6f, 0x70,
	0x65, 0x63, 0x6b, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x12, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x6f, 0x70, 0x65, 0x63, 0x6b, 0x73,
	0x50, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x32, 0x8e, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x65, 0x78, 0x42, 0x6f, 0x6b, 0x75, 0x6e, 0x2f,
	0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_Convert_convert_proto_rawDescOnce sync.Once
	file_Convert_convert_proto_rawDescData []byte
)

func file_Convert_convert_proto_rawDescGZIP() []byte {
	file_Convert_convert_proto_rawDescOnce.Do(func() {
		file_Convert_convert_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_Convert_convert_proto_rawDesc), len(file_Convert_convert_proto_rawDesc)))
	})
	return file_Convert_convert_proto_rawDescData
}

var file_Convert_convert_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Convert_convert_proto_goTypes = []any{
	(*ConvertRequest)(nil),  // 0: conversion.ConvertRequest
	(*ConvertResponse)(nil), // 1: conversion.ConvertResponse
	(*RateRequest)(nil),     // 2: conversion.RateRequest
	(*RateResponse)(nil),    // 3: conversion.RateResponse
}
var file_Convert_convert_proto_depIdxs = []int32{
	0, // 0: conversion.Conversion.Convert:input_type -> conversion.ConvertRequest
	2, // 1: conversion.Conversion.GetRate:input_type -> conversion.RateRequest
	1, // 2: conversion.Conversion.Convert:output_type -> conversion.ConvertResponse
	3, // 3: conversion.Conversion.GetRate:output_type -> conversion.RateResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Convert_convert_proto_init() }
func file_Convert_convert_proto_init() {
	if File_Convert_convert_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_Convert_convert_proto_rawDesc), len(file_Convert_convert_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Convert_convert_proto_goTypes,
		DependencyIndexes: file_Convert_convert_proto_depIdxs,
		MessageInfos:      file_Convert_convert_proto_msgTypes,
	}.Build()
	File_Convert_convert_proto = out.File
	file_Convert_convert_proto_goTypes = nil
	file_Convert_convert_proto_depIdxs = nil
}
