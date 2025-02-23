// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: Withdraw/withdraw.proto

package withdrawalv1

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

type WithdrawRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	TransactionId  string                 `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`    // UUID из Transaction Service
	CryptoCurrency string                 `protobuf:"bytes,2,opt,name=crypto_currency,json=cryptoCurrency,proto3" json:"crypto_currency,omitempty"` // "BTC", "ETH", "USDT"
	CryptoWallet   string                 `protobuf:"bytes,3,opt,name=crypto_wallet,json=cryptoWallet,proto3" json:"crypto_wallet,omitempty"`       // Адрес кошелька
	CryptoAmount   int64                  `protobuf:"varint,4,opt,name=crypto_amount,json=cryptoAmount,proto3" json:"crypto_amount,omitempty"`      // Сумма в минимальных единицах
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *WithdrawRequest) Reset() {
	*x = WithdrawRequest{}
	mi := &file_Withdraw_withdraw_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WithdrawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawRequest) ProtoMessage() {}

func (x *WithdrawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Withdraw_withdraw_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawRequest.ProtoReflect.Descriptor instead.
func (*WithdrawRequest) Descriptor() ([]byte, []int) {
	return file_Withdraw_withdraw_proto_rawDescGZIP(), []int{0}
}

func (x *WithdrawRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *WithdrawRequest) GetCryptoCurrency() string {
	if x != nil {
		return x.CryptoCurrency
	}
	return ""
}

func (x *WithdrawRequest) GetCryptoWallet() string {
	if x != nil {
		return x.CryptoWallet
	}
	return ""
}

func (x *WithdrawRequest) GetCryptoAmount() int64 {
	if x != nil {
		return x.CryptoAmount
	}
	return 0
}

type WithdrawResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	TxHash        string                 `protobuf:"bytes,2,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"` // Хэш транзакции в блокчейне
	Error         string                 `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`                 // Детали ошибки (если success=false)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WithdrawResponse) Reset() {
	*x = WithdrawResponse{}
	mi := &file_Withdraw_withdraw_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WithdrawResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawResponse) ProtoMessage() {}

func (x *WithdrawResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Withdraw_withdraw_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawResponse.ProtoReflect.Descriptor instead.
func (*WithdrawResponse) Descriptor() ([]byte, []int) {
	return file_Withdraw_withdraw_proto_rawDescGZIP(), []int{1}
}

func (x *WithdrawResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *WithdrawResponse) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *WithdrawResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_Withdraw_withdraw_proto protoreflect.FileDescriptor

var file_Withdraw_withdraw_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x77, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x61, 0x6c, 0x22, 0xab, 0x01, 0x0a, 0x0f, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x27, 0x0a, 0x0f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x5b, 0x0a, 0x10, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x32, 0x53, 0x0a, 0x0a, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x12, 0x45,
	0x0a, 0x08, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x12, 0x1b, 0x2e, 0x77, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x61, 0x6c, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x65, 0x78, 0x42, 0x6f, 0x6b, 0x75, 0x6e, 0x2f, 0x77, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x61, 0x6c, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_Withdraw_withdraw_proto_rawDescOnce sync.Once
	file_Withdraw_withdraw_proto_rawDescData []byte
)

func file_Withdraw_withdraw_proto_rawDescGZIP() []byte {
	file_Withdraw_withdraw_proto_rawDescOnce.Do(func() {
		file_Withdraw_withdraw_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_Withdraw_withdraw_proto_rawDesc), len(file_Withdraw_withdraw_proto_rawDesc)))
	})
	return file_Withdraw_withdraw_proto_rawDescData
}

var file_Withdraw_withdraw_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Withdraw_withdraw_proto_goTypes = []any{
	(*WithdrawRequest)(nil),  // 0: withdrawal.WithdrawRequest
	(*WithdrawResponse)(nil), // 1: withdrawal.WithdrawResponse
}
var file_Withdraw_withdraw_proto_depIdxs = []int32{
	0, // 0: withdrawal.Withdrawal.Withdraw:input_type -> withdrawal.WithdrawRequest
	1, // 1: withdrawal.Withdrawal.Withdraw:output_type -> withdrawal.WithdrawResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Withdraw_withdraw_proto_init() }
func file_Withdraw_withdraw_proto_init() {
	if File_Withdraw_withdraw_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_Withdraw_withdraw_proto_rawDesc), len(file_Withdraw_withdraw_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Withdraw_withdraw_proto_goTypes,
		DependencyIndexes: file_Withdraw_withdraw_proto_depIdxs,
		MessageInfos:      file_Withdraw_withdraw_proto_msgTypes,
	}.Build()
	File_Withdraw_withdraw_proto = out.File
	file_Withdraw_withdraw_proto_goTypes = nil
	file_Withdraw_withdraw_proto_depIdxs = nil
}
