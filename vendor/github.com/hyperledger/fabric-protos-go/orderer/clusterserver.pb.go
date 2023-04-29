// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: orderer/clusterserver.proto

package orderer

import (
	common "github.com/hyperledger/fabric-protos-go/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ClusterNodeServiceStepRequest wraps a message that is sent to a cluster member.
type ClusterNodeServiceStepRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*ClusterNodeServiceStepRequest_NodeConrequest
	//	*ClusterNodeServiceStepRequest_NodeTranrequest
	//	*ClusterNodeServiceStepRequest_NodeAuthrequest
	Payload isClusterNodeServiceStepRequest_Payload `protobuf_oneof:"payload"`
}

func (x *ClusterNodeServiceStepRequest) Reset() {
	*x = ClusterNodeServiceStepRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_clusterserver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterNodeServiceStepRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterNodeServiceStepRequest) ProtoMessage() {}

func (x *ClusterNodeServiceStepRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_clusterserver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterNodeServiceStepRequest.ProtoReflect.Descriptor instead.
func (*ClusterNodeServiceStepRequest) Descriptor() ([]byte, []int) {
	return file_orderer_clusterserver_proto_rawDescGZIP(), []int{0}
}

func (m *ClusterNodeServiceStepRequest) GetPayload() isClusterNodeServiceStepRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *ClusterNodeServiceStepRequest) GetNodeConrequest() *NodeConsensusRequest {
	if x, ok := x.GetPayload().(*ClusterNodeServiceStepRequest_NodeConrequest); ok {
		return x.NodeConrequest
	}
	return nil
}

func (x *ClusterNodeServiceStepRequest) GetNodeTranrequest() *NodeTransactionOrderRequest {
	if x, ok := x.GetPayload().(*ClusterNodeServiceStepRequest_NodeTranrequest); ok {
		return x.NodeTranrequest
	}
	return nil
}

func (x *ClusterNodeServiceStepRequest) GetNodeAuthrequest() *NodeAuthRequest {
	if x, ok := x.GetPayload().(*ClusterNodeServiceStepRequest_NodeAuthrequest); ok {
		return x.NodeAuthrequest
	}
	return nil
}

type isClusterNodeServiceStepRequest_Payload interface {
	isClusterNodeServiceStepRequest_Payload()
}

type ClusterNodeServiceStepRequest_NodeConrequest struct {
	// node_conrequest is a consensus specific message between the cluster memebers.
	NodeConrequest *NodeConsensusRequest `protobuf:"bytes,1,opt,name=node_conrequest,json=nodeConrequest,proto3,oneof"`
}

type ClusterNodeServiceStepRequest_NodeTranrequest struct {
	// node_tranrequest is a relay of a transaction.
	NodeTranrequest *NodeTransactionOrderRequest `protobuf:"bytes,2,opt,name=node_tranrequest,json=nodeTranrequest,proto3,oneof"`
}

type ClusterNodeServiceStepRequest_NodeAuthrequest struct {
	// Auth authentiates the member that initiated the stream
	NodeAuthrequest *NodeAuthRequest `protobuf:"bytes,3,opt,name=node_authrequest,json=nodeAuthrequest,proto3,oneof"`
}

func (*ClusterNodeServiceStepRequest_NodeConrequest) isClusterNodeServiceStepRequest_Payload() {}

func (*ClusterNodeServiceStepRequest_NodeTranrequest) isClusterNodeServiceStepRequest_Payload() {}

func (*ClusterNodeServiceStepRequest_NodeAuthrequest) isClusterNodeServiceStepRequest_Payload() {}

// ClusterNodeServiceStepResponse is a message received from a cluster member.
type ClusterNodeServiceStepResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*ClusterNodeServiceStepResponse_TranorderRes
	Payload isClusterNodeServiceStepResponse_Payload `protobuf_oneof:"payload"`
}

func (x *ClusterNodeServiceStepResponse) Reset() {
	*x = ClusterNodeServiceStepResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_clusterserver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterNodeServiceStepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterNodeServiceStepResponse) ProtoMessage() {}

func (x *ClusterNodeServiceStepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_clusterserver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterNodeServiceStepResponse.ProtoReflect.Descriptor instead.
func (*ClusterNodeServiceStepResponse) Descriptor() ([]byte, []int) {
	return file_orderer_clusterserver_proto_rawDescGZIP(), []int{1}
}

func (m *ClusterNodeServiceStepResponse) GetPayload() isClusterNodeServiceStepResponse_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *ClusterNodeServiceStepResponse) GetTranorderRes() *TransactionOrderResponse {
	if x, ok := x.GetPayload().(*ClusterNodeServiceStepResponse_TranorderRes); ok {
		return x.TranorderRes
	}
	return nil
}

type isClusterNodeServiceStepResponse_Payload interface {
	isClusterNodeServiceStepResponse_Payload()
}

type ClusterNodeServiceStepResponse_TranorderRes struct {
	TranorderRes *TransactionOrderResponse `protobuf:"bytes,1,opt,name=tranorder_res,json=tranorderRes,proto3,oneof"`
}

func (*ClusterNodeServiceStepResponse_TranorderRes) isClusterNodeServiceStepResponse_Payload() {}

// NodeConsensusRequest is a consensus specific message sent to a cluster member.
type NodeConsensusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload  []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Metadata []byte `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *NodeConsensusRequest) Reset() {
	*x = NodeConsensusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_clusterserver_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeConsensusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeConsensusRequest) ProtoMessage() {}

func (x *NodeConsensusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_clusterserver_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeConsensusRequest.ProtoReflect.Descriptor instead.
func (*NodeConsensusRequest) Descriptor() ([]byte, []int) {
	return file_orderer_clusterserver_proto_rawDescGZIP(), []int{2}
}

func (x *NodeConsensusRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *NodeConsensusRequest) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// NodeTransactionOrderRequest wraps a transaction to be sent for ordering.
type NodeTransactionOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// last_validation_seq denotes the last configuration sequence at which the
	// sender validated this message.
	LastValidationSeq uint64 `protobuf:"varint,1,opt,name=last_validation_seq,json=lastValidationSeq,proto3" json:"last_validation_seq,omitempty"`
	// content is the fabric transaction
	// that is forwarded to the cluster member.
	Payload *common.Envelope `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *NodeTransactionOrderRequest) Reset() {
	*x = NodeTransactionOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_clusterserver_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeTransactionOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeTransactionOrderRequest) ProtoMessage() {}

func (x *NodeTransactionOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_clusterserver_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeTransactionOrderRequest.ProtoReflect.Descriptor instead.
func (*NodeTransactionOrderRequest) Descriptor() ([]byte, []int) {
	return file_orderer_clusterserver_proto_rawDescGZIP(), []int{3}
}

func (x *NodeTransactionOrderRequest) GetLastValidationSeq() uint64 {
	if x != nil {
		return x.LastValidationSeq
	}
	return 0
}

func (x *NodeTransactionOrderRequest) GetPayload() *common.Envelope {
	if x != nil {
		return x.Payload
	}
	return nil
}

// TransactionOrderResponse returns a success
// or failure status to the sender.
type TransactionOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	TxId    string `protobuf:"bytes,2,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	// Status code, which may be used to programatically respond to success/failure.
	Status common.Status `protobuf:"varint,3,opt,name=status,proto3,enum=common.Status" json:"status,omitempty"`
	// Info string which may contain additional information about the returned status.
	Info string `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *TransactionOrderResponse) Reset() {
	*x = TransactionOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_clusterserver_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionOrderResponse) ProtoMessage() {}

func (x *TransactionOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_clusterserver_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionOrderResponse.ProtoReflect.Descriptor instead.
func (*TransactionOrderResponse) Descriptor() ([]byte, []int) {
	return file_orderer_clusterserver_proto_rawDescGZIP(), []int{4}
}

func (x *TransactionOrderResponse) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *TransactionOrderResponse) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *TransactionOrderResponse) GetStatus() common.Status {
	if x != nil {
		return x.Status
	}
	return common.Status(0)
}

func (x *TransactionOrderResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

// NodeAuthRequest for authenticate the stream
// between the cluster members
type NodeAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// version represents the fields on which the signature is computed
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// signature is verifiable using the initiator's public key
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	// timestamp indicates the freshness of the request; expected to be within the margin
	// of the responsder's local time
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// from_id is the numerical identifier of the initiator of the connection
	FromId uint64 `protobuf:"varint,4,opt,name=from_id,json=fromId,proto3" json:"from_id,omitempty"`
	// to_id is the numerical identifier of the node that is being connected to
	ToId uint64 `protobuf:"varint,5,opt,name=to_id,json=toId,proto3" json:"to_id,omitempty"`
	// session_binding is verifiable using application level protocol
	SessionBinding []byte `protobuf:"bytes,6,opt,name=session_binding,json=sessionBinding,proto3" json:"session_binding,omitempty"`
	Channel        string `protobuf:"bytes,7,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *NodeAuthRequest) Reset() {
	*x = NodeAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_clusterserver_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeAuthRequest) ProtoMessage() {}

func (x *NodeAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_clusterserver_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeAuthRequest.ProtoReflect.Descriptor instead.
func (*NodeAuthRequest) Descriptor() ([]byte, []int) {
	return file_orderer_clusterserver_proto_rawDescGZIP(), []int{5}
}

func (x *NodeAuthRequest) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *NodeAuthRequest) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *NodeAuthRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *NodeAuthRequest) GetFromId() uint64 {
	if x != nil {
		return x.FromId
	}
	return 0
}

func (x *NodeAuthRequest) GetToId() uint64 {
	if x != nil {
		return x.ToId
	}
	return 0
}

func (x *NodeAuthRequest) GetSessionBinding() []byte {
	if x != nil {
		return x.SessionBinding
	}
	return nil
}

func (x *NodeAuthRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

var File_orderer_clusterserver_proto protoreflect.FileDescriptor

var file_orderer_clusterserver_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x02, 0x0a,
	0x1d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48,
	0x0a, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65,
	0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x6e, 0x6f, 0x64, 0x65, 0x43, 0x6f,
	0x6e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x51, 0x0a, 0x10, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x6e, 0x6f, 0x64, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x10, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x00, 0x52, 0x0f, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x75, 0x74, 0x68, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x75, 0x0a,
	0x1e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x48, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x74, 0x72, 0x61,
	0x6e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x4c, 0x0a, 0x14, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x73, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x79, 0x0a, 0x1b, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2e, 0x0a, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11,
	0x6c, 0x61, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x71, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x65, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x85, 0x01,
	0x0a, 0x18, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xf4, 0x01, 0x0a, 0x0f, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x66, 0x72,
	0x6f, 0x6d, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x6f, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x32, 0x71, 0x0a, 0x12,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5b, 0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x26, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x74, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42,
	0xae, 0x01, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x42, 0x12, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65,
	0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0xca, 0x02, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72,
	0xe2, 0x02, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orderer_clusterserver_proto_rawDescOnce sync.Once
	file_orderer_clusterserver_proto_rawDescData = file_orderer_clusterserver_proto_rawDesc
)

func file_orderer_clusterserver_proto_rawDescGZIP() []byte {
	file_orderer_clusterserver_proto_rawDescOnce.Do(func() {
		file_orderer_clusterserver_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderer_clusterserver_proto_rawDescData)
	})
	return file_orderer_clusterserver_proto_rawDescData
}

var file_orderer_clusterserver_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_orderer_clusterserver_proto_goTypes = []interface{}{
	(*ClusterNodeServiceStepRequest)(nil),  // 0: orderer.ClusterNodeServiceStepRequest
	(*ClusterNodeServiceStepResponse)(nil), // 1: orderer.ClusterNodeServiceStepResponse
	(*NodeConsensusRequest)(nil),           // 2: orderer.NodeConsensusRequest
	(*NodeTransactionOrderRequest)(nil),    // 3: orderer.NodeTransactionOrderRequest
	(*TransactionOrderResponse)(nil),       // 4: orderer.TransactionOrderResponse
	(*NodeAuthRequest)(nil),                // 5: orderer.NodeAuthRequest
	(*common.Envelope)(nil),                // 6: common.Envelope
	(common.Status)(0),                     // 7: common.Status
	(*timestamppb.Timestamp)(nil),          // 8: google.protobuf.Timestamp
}
var file_orderer_clusterserver_proto_depIdxs = []int32{
	2, // 0: orderer.ClusterNodeServiceStepRequest.node_conrequest:type_name -> orderer.NodeConsensusRequest
	3, // 1: orderer.ClusterNodeServiceStepRequest.node_tranrequest:type_name -> orderer.NodeTransactionOrderRequest
	5, // 2: orderer.ClusterNodeServiceStepRequest.node_authrequest:type_name -> orderer.NodeAuthRequest
	4, // 3: orderer.ClusterNodeServiceStepResponse.tranorder_res:type_name -> orderer.TransactionOrderResponse
	6, // 4: orderer.NodeTransactionOrderRequest.payload:type_name -> common.Envelope
	7, // 5: orderer.TransactionOrderResponse.status:type_name -> common.Status
	8, // 6: orderer.NodeAuthRequest.timestamp:type_name -> google.protobuf.Timestamp
	0, // 7: orderer.ClusterNodeService.Step:input_type -> orderer.ClusterNodeServiceStepRequest
	1, // 8: orderer.ClusterNodeService.Step:output_type -> orderer.ClusterNodeServiceStepResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_orderer_clusterserver_proto_init() }
func file_orderer_clusterserver_proto_init() {
	if File_orderer_clusterserver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orderer_clusterserver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterNodeServiceStepRequest); i {
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
		file_orderer_clusterserver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterNodeServiceStepResponse); i {
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
		file_orderer_clusterserver_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeConsensusRequest); i {
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
		file_orderer_clusterserver_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeTransactionOrderRequest); i {
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
		file_orderer_clusterserver_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionOrderResponse); i {
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
		file_orderer_clusterserver_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeAuthRequest); i {
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
	file_orderer_clusterserver_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ClusterNodeServiceStepRequest_NodeConrequest)(nil),
		(*ClusterNodeServiceStepRequest_NodeTranrequest)(nil),
		(*ClusterNodeServiceStepRequest_NodeAuthrequest)(nil),
	}
	file_orderer_clusterserver_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ClusterNodeServiceStepResponse_TranorderRes)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orderer_clusterserver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orderer_clusterserver_proto_goTypes,
		DependencyIndexes: file_orderer_clusterserver_proto_depIdxs,
		MessageInfos:      file_orderer_clusterserver_proto_msgTypes,
	}.Build()
	File_orderer_clusterserver_proto = out.File
	file_orderer_clusterserver_proto_rawDesc = nil
	file_orderer_clusterserver_proto_goTypes = nil
	file_orderer_clusterserver_proto_depIdxs = nil
}
