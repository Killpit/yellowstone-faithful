// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.1
// source: old-faithful.proto

package old_faithful_grpc

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

type VersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VersionRequest) Reset() {
	*x = VersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_old_faithful_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionRequest) ProtoMessage() {}

func (x *VersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_old_faithful_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionRequest.ProtoReflect.Descriptor instead.
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return file_old_faithful_proto_rawDescGZIP(), []int{0}
}

type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_old_faithful_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_old_faithful_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_old_faithful_proto_rawDescGZIP(), []int{1}
}

func (x *VersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type BlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot uint64 `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (x *BlockRequest) Reset() {
	*x = BlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_old_faithful_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockRequest) ProtoMessage() {}

func (x *BlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_old_faithful_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockRequest.ProtoReflect.Descriptor instead.
func (*BlockRequest) Descriptor() ([]byte, []int) {
	return file_old_faithful_proto_rawDescGZIP(), []int{2}
}

func (x *BlockRequest) GetSlot() uint64 {
	if x != nil {
		return x.Slot
	}
	return 0
}

type BlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreviousBlockhash []byte         `protobuf:"bytes,1,opt,name=previous_blockhash,json=previousBlockhash,proto3" json:"previous_blockhash,omitempty"`
	Blockhash         []byte         `protobuf:"bytes,2,opt,name=blockhash,proto3" json:"blockhash,omitempty"`
	ParentSlot        uint64         `protobuf:"varint,3,opt,name=parent_slot,json=parentSlot,proto3" json:"parent_slot,omitempty"`
	Slot              uint64         `protobuf:"varint,4,opt,name=slot,proto3" json:"slot,omitempty"`
	BlockTime         int64          `protobuf:"varint,5,opt,name=block_time,json=blockTime,proto3" json:"block_time,omitempty"`
	BlockHeight       uint64         `protobuf:"varint,6,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Transactions      []*Transaction `protobuf:"bytes,7,rep,name=transactions,proto3" json:"transactions,omitempty"`
	Rewards           []byte         `protobuf:"bytes,8,opt,name=rewards,proto3" json:"rewards,omitempty"` // protobuf or bincode (or empty)
}

func (x *BlockResponse) Reset() {
	*x = BlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_old_faithful_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockResponse) ProtoMessage() {}

func (x *BlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_old_faithful_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockResponse.ProtoReflect.Descriptor instead.
func (*BlockResponse) Descriptor() ([]byte, []int) {
	return file_old_faithful_proto_rawDescGZIP(), []int{3}
}

func (x *BlockResponse) GetPreviousBlockhash() []byte {
	if x != nil {
		return x.PreviousBlockhash
	}
	return nil
}

func (x *BlockResponse) GetBlockhash() []byte {
	if x != nil {
		return x.Blockhash
	}
	return nil
}

func (x *BlockResponse) GetParentSlot() uint64 {
	if x != nil {
		return x.ParentSlot
	}
	return 0
}

func (x *BlockResponse) GetSlot() uint64 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *BlockResponse) GetBlockTime() int64 {
	if x != nil {
		return x.BlockTime
	}
	return 0
}

func (x *BlockResponse) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *BlockResponse) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *BlockResponse) GetRewards() []byte {
	if x != nil {
		return x.Rewards
	}
	return nil
}

type TransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *TransactionRequest) Reset() {
	*x = TransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_old_faithful_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRequest) ProtoMessage() {}

func (x *TransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_old_faithful_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRequest.ProtoReflect.Descriptor instead.
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return file_old_faithful_proto_rawDescGZIP(), []int{4}
}

func (x *TransactionRequest) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type TransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Slot        uint64       `protobuf:"varint,2,opt,name=slot,proto3" json:"slot,omitempty"`
	BlockTime   int64        `protobuf:"varint,3,opt,name=block_time,json=blockTime,proto3" json:"block_time,omitempty"`
	Index       *uint64      `protobuf:"varint,4,opt,name=index,proto3,oneof" json:"index,omitempty"` // position in the block
}

func (x *TransactionResponse) Reset() {
	*x = TransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_old_faithful_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionResponse) ProtoMessage() {}

func (x *TransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_old_faithful_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionResponse.ProtoReflect.Descriptor instead.
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return file_old_faithful_proto_rawDescGZIP(), []int{5}
}

func (x *TransactionResponse) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TransactionResponse) GetSlot() uint64 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *TransactionResponse) GetBlockTime() int64 {
	if x != nil {
		return x.BlockTime
	}
	return 0
}

func (x *TransactionResponse) GetIndex() uint64 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction []byte  `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"` // solana native transaction
	Meta        []byte  `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`               // bincode or protobuf
	Index       *uint64 `protobuf:"varint,4,opt,name=index,proto3,oneof" json:"index,omitempty"`      // position in the block
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_old_faithful_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_old_faithful_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_old_faithful_proto_rawDescGZIP(), []int{6}
}

func (x *Transaction) GetTransaction() []byte {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *Transaction) GetMeta() []byte {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Transaction) GetIndex() uint64 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

var File_old_faithful_proto protoreflect.FileDescriptor

var file_old_faithful_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x6c, 0x64, 0x2d, 0x66, 0x61, 0x69, 0x74, 0x68, 0x66, 0x75, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x4f, 0x6c, 0x64, 0x46, 0x61, 0x69, 0x74, 0x68, 0x66, 0x75,
	0x6c, 0x22, 0x10, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x2b, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x22, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x73, 0x6c, 0x6f, 0x74, 0x22, 0xab, 0x02, 0x0a, 0x0d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x68,
	0x61, 0x73, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x6c,
	0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x53, 0x6c, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x3c, 0x0a, 0x0c, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x4f, 0x6c, 0x64, 0x46, 0x61, 0x69, 0x74, 0x68, 0x66, 0x75, 0x6c, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x73, 0x22, 0x32, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a,
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x4f, 0x6c, 0x64, 0x46, 0x61, 0x69, 0x74, 0x68, 0x66, 0x75,
	0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c,
	0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x22, 0x68, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x88,
	0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x32, 0xee, 0x01, 0x0a,
	0x0b, 0x4f, 0x6c, 0x64, 0x46, 0x61, 0x69, 0x74, 0x68, 0x66, 0x75, 0x6c, 0x12, 0x47, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x4f, 0x6c, 0x64,
	0x46, 0x61, 0x69, 0x74, 0x68, 0x66, 0x75, 0x6c, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x4f, 0x6c, 0x64, 0x46, 0x61, 0x69,
	0x74, 0x68, 0x66, 0x75, 0x6c, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x19, 0x2e, 0x4f, 0x6c, 0x64, 0x46, 0x61, 0x69, 0x74, 0x68, 0x66, 0x75, 0x6c, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x4f,
	0x6c, 0x64, 0x46, 0x61, 0x69, 0x74, 0x68, 0x66, 0x75, 0x6c, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x4f, 0x6c, 0x64,
	0x46, 0x61, 0x69, 0x74, 0x68, 0x66, 0x75, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x4f, 0x6c,
	0x64, 0x46, 0x61, 0x69, 0x74, 0x68, 0x66, 0x75, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4e, 0x5a,
	0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x70, 0x63, 0x70,
	0x6f, 0x6f, 0x6c, 0x2f, 0x79, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x73, 0x74, 0x6f, 0x6e, 0x65, 0x2d,
	0x66, 0x61, 0x69, 0x74, 0x68, 0x66, 0x75, 0x6c, 0x2f, 0x6f, 0x6c, 0x64, 0x2d, 0x66, 0x61, 0x69,
	0x74, 0x68, 0x66, 0x75, 0x6c, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x6f, 0x6c, 0x64, 0x5f,
	0x66, 0x61, 0x69, 0x74, 0x68, 0x66, 0x75, 0x6c, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_old_faithful_proto_rawDescOnce sync.Once
	file_old_faithful_proto_rawDescData = file_old_faithful_proto_rawDesc
)

func file_old_faithful_proto_rawDescGZIP() []byte {
	file_old_faithful_proto_rawDescOnce.Do(func() {
		file_old_faithful_proto_rawDescData = protoimpl.X.CompressGZIP(file_old_faithful_proto_rawDescData)
	})
	return file_old_faithful_proto_rawDescData
}

var file_old_faithful_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_old_faithful_proto_goTypes = []interface{}{
	(*VersionRequest)(nil),      // 0: OldFaithful.VersionRequest
	(*VersionResponse)(nil),     // 1: OldFaithful.VersionResponse
	(*BlockRequest)(nil),        // 2: OldFaithful.BlockRequest
	(*BlockResponse)(nil),       // 3: OldFaithful.BlockResponse
	(*TransactionRequest)(nil),  // 4: OldFaithful.TransactionRequest
	(*TransactionResponse)(nil), // 5: OldFaithful.TransactionResponse
	(*Transaction)(nil),         // 6: OldFaithful.Transaction
}
var file_old_faithful_proto_depIdxs = []int32{
	6, // 0: OldFaithful.BlockResponse.transactions:type_name -> OldFaithful.Transaction
	6, // 1: OldFaithful.TransactionResponse.transaction:type_name -> OldFaithful.Transaction
	0, // 2: OldFaithful.OldFaithful.GetVersion:input_type -> OldFaithful.VersionRequest
	2, // 3: OldFaithful.OldFaithful.GetBlock:input_type -> OldFaithful.BlockRequest
	4, // 4: OldFaithful.OldFaithful.GetTransaction:input_type -> OldFaithful.TransactionRequest
	1, // 5: OldFaithful.OldFaithful.GetVersion:output_type -> OldFaithful.VersionResponse
	3, // 6: OldFaithful.OldFaithful.GetBlock:output_type -> OldFaithful.BlockResponse
	5, // 7: OldFaithful.OldFaithful.GetTransaction:output_type -> OldFaithful.TransactionResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_old_faithful_proto_init() }
func file_old_faithful_proto_init() {
	if File_old_faithful_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_old_faithful_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionRequest); i {
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
		file_old_faithful_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResponse); i {
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
		file_old_faithful_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockRequest); i {
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
		file_old_faithful_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockResponse); i {
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
		file_old_faithful_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionRequest); i {
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
		file_old_faithful_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionResponse); i {
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
		file_old_faithful_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
	file_old_faithful_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_old_faithful_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_old_faithful_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_old_faithful_proto_goTypes,
		DependencyIndexes: file_old_faithful_proto_depIdxs,
		MessageInfos:      file_old_faithful_proto_msgTypes,
	}.Build()
	File_old_faithful_proto = out.File
	file_old_faithful_proto_rawDesc = nil
	file_old_faithful_proto_goTypes = nil
	file_old_faithful_proto_depIdxs = nil
}
