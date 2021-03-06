// Code generated by protoc-gen-go.
// source: peer/chaincode.proto
// DO NOT EDIT!

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Confidentiality Levels
type ConfidentialityLevel int32

const (
	ConfidentialityLevel_PUBLIC       ConfidentialityLevel = 0
	ConfidentialityLevel_CONFIDENTIAL ConfidentialityLevel = 1
)

var ConfidentialityLevel_name = map[int32]string{
	0: "PUBLIC",
	1: "CONFIDENTIAL",
}
var ConfidentialityLevel_value = map[string]int32{
	"PUBLIC":       0,
	"CONFIDENTIAL": 1,
}

func (x ConfidentialityLevel) String() string {
	return proto.EnumName(ConfidentialityLevel_name, int32(x))
}
func (ConfidentialityLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ChaincodeSpec_Type int32

const (
	ChaincodeSpec_UNDEFINED ChaincodeSpec_Type = 0
	ChaincodeSpec_GOLANG    ChaincodeSpec_Type = 1
	ChaincodeSpec_NODE      ChaincodeSpec_Type = 2
	ChaincodeSpec_CAR       ChaincodeSpec_Type = 3
	ChaincodeSpec_JAVA      ChaincodeSpec_Type = 4
)

var ChaincodeSpec_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "GOLANG",
	2: "NODE",
	3: "CAR",
	4: "JAVA",
}
var ChaincodeSpec_Type_value = map[string]int32{
	"UNDEFINED": 0,
	"GOLANG":    1,
	"NODE":      2,
	"CAR":       3,
	"JAVA":      4,
}

func (x ChaincodeSpec_Type) String() string {
	return proto.EnumName(ChaincodeSpec_Type_name, int32(x))
}
func (ChaincodeSpec_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 0} }

type ChaincodeDeploymentSpec_ExecutionEnvironment int32

const (
	ChaincodeDeploymentSpec_DOCKER ChaincodeDeploymentSpec_ExecutionEnvironment = 0
	ChaincodeDeploymentSpec_SYSTEM ChaincodeDeploymentSpec_ExecutionEnvironment = 1
)

var ChaincodeDeploymentSpec_ExecutionEnvironment_name = map[int32]string{
	0: "DOCKER",
	1: "SYSTEM",
}
var ChaincodeDeploymentSpec_ExecutionEnvironment_value = map[string]int32{
	"DOCKER": 0,
	"SYSTEM": 1,
}

func (x ChaincodeDeploymentSpec_ExecutionEnvironment) String() string {
	return proto.EnumName(ChaincodeDeploymentSpec_ExecutionEnvironment_name, int32(x))
}
func (ChaincodeDeploymentSpec_ExecutionEnvironment) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{3, 0}
}

type ChaincodeMessage_Type int32

const (
	ChaincodeMessage_UNDEFINED          ChaincodeMessage_Type = 0
	ChaincodeMessage_REGISTER           ChaincodeMessage_Type = 1
	ChaincodeMessage_REGISTERED         ChaincodeMessage_Type = 2
	ChaincodeMessage_INIT               ChaincodeMessage_Type = 3
	ChaincodeMessage_READY              ChaincodeMessage_Type = 4
	ChaincodeMessage_TRANSACTION        ChaincodeMessage_Type = 5
	ChaincodeMessage_COMPLETED          ChaincodeMessage_Type = 6
	ChaincodeMessage_ERROR              ChaincodeMessage_Type = 7
	ChaincodeMessage_GET_STATE          ChaincodeMessage_Type = 8
	ChaincodeMessage_PUT_STATE          ChaincodeMessage_Type = 9
	ChaincodeMessage_DEL_STATE          ChaincodeMessage_Type = 10
	ChaincodeMessage_INVOKE_CHAINCODE   ChaincodeMessage_Type = 11
	ChaincodeMessage_RESPONSE           ChaincodeMessage_Type = 13
	ChaincodeMessage_GET_STATE_BY_RANGE ChaincodeMessage_Type = 14
	ChaincodeMessage_GET_QUERY_RESULT   ChaincodeMessage_Type = 15
	ChaincodeMessage_QUERY_STATE_NEXT   ChaincodeMessage_Type = 16
	ChaincodeMessage_QUERY_STATE_CLOSE  ChaincodeMessage_Type = 17
	ChaincodeMessage_KEEPALIVE          ChaincodeMessage_Type = 18
)

var ChaincodeMessage_Type_name = map[int32]string{
	0:  "UNDEFINED",
	1:  "REGISTER",
	2:  "REGISTERED",
	3:  "INIT",
	4:  "READY",
	5:  "TRANSACTION",
	6:  "COMPLETED",
	7:  "ERROR",
	8:  "GET_STATE",
	9:  "PUT_STATE",
	10: "DEL_STATE",
	11: "INVOKE_CHAINCODE",
	13: "RESPONSE",
	14: "GET_STATE_BY_RANGE",
	15: "GET_QUERY_RESULT",
	16: "QUERY_STATE_NEXT",
	17: "QUERY_STATE_CLOSE",
	18: "KEEPALIVE",
}
var ChaincodeMessage_Type_value = map[string]int32{
	"UNDEFINED":          0,
	"REGISTER":           1,
	"REGISTERED":         2,
	"INIT":               3,
	"READY":              4,
	"TRANSACTION":        5,
	"COMPLETED":          6,
	"ERROR":              7,
	"GET_STATE":          8,
	"PUT_STATE":          9,
	"DEL_STATE":          10,
	"INVOKE_CHAINCODE":   11,
	"RESPONSE":           13,
	"GET_STATE_BY_RANGE": 14,
	"GET_QUERY_RESULT":   15,
	"QUERY_STATE_NEXT":   16,
	"QUERY_STATE_CLOSE":  17,
	"KEEPALIVE":          18,
}

func (x ChaincodeMessage_Type) String() string {
	return proto.EnumName(ChaincodeMessage_Type_name, int32(x))
}
func (ChaincodeMessage_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{6, 0} }

// ChaincodeID contains the path as specified by the deploy transaction
// that created it as well as the hashCode that is generated by the
// system for the path. From the user level (ie, CLI, REST API and so on)
// deploy transaction is expected to provide the path and other requests
// are expected to provide the hashCode. The other value will be ignored.
// Internally, the structure could contain both values. For instance, the
// hashCode will be set when first generated using the path
type ChaincodeID struct {
	// deploy transaction will use the path
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// all other requests will use the name (really a hashcode) generated by
	// the deploy transaction
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// user friendly version name for the chaincode
	Version string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
}

func (m *ChaincodeID) Reset()                    { *m = ChaincodeID{} }
func (m *ChaincodeID) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeID) ProtoMessage()               {}
func (*ChaincodeID) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// Carries the chaincode function and its arguments.
// UnmarshalJSON in transaction.go converts the string-based REST/JSON input to
// the []byte-based current ChaincodeInput structure.
type ChaincodeInput struct {
	Args [][]byte `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
}

func (m *ChaincodeInput) Reset()                    { *m = ChaincodeInput{} }
func (m *ChaincodeInput) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeInput) ProtoMessage()               {}
func (*ChaincodeInput) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

// Carries the chaincode specification. This is the actual metadata required for
// defining a chaincode.
type ChaincodeSpec struct {
	Type        ChaincodeSpec_Type `protobuf:"varint,1,opt,name=type,enum=protos.ChaincodeSpec_Type" json:"type,omitempty"`
	ChaincodeId *ChaincodeID       `protobuf:"bytes,2,opt,name=chaincode_id,json=chaincodeId" json:"chaincode_id,omitempty"`
	Input       *ChaincodeInput    `protobuf:"bytes,3,opt,name=input" json:"input,omitempty"`
	Timeout     int32              `protobuf:"varint,4,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *ChaincodeSpec) Reset()                    { *m = ChaincodeSpec{} }
func (m *ChaincodeSpec) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeSpec) ProtoMessage()               {}
func (*ChaincodeSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ChaincodeSpec) GetChaincodeId() *ChaincodeID {
	if m != nil {
		return m.ChaincodeId
	}
	return nil
}

func (m *ChaincodeSpec) GetInput() *ChaincodeInput {
	if m != nil {
		return m.Input
	}
	return nil
}

// Specify the deployment of a chaincode.
// TODO: Define `codePackage`.
type ChaincodeDeploymentSpec struct {
	ChaincodeSpec *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincode_spec,json=chaincodeSpec" json:"chaincode_spec,omitempty"`
	// Controls when the chaincode becomes executable.
	EffectiveDate *google_protobuf1.Timestamp                  `protobuf:"bytes,2,opt,name=effective_date,json=effectiveDate" json:"effective_date,omitempty"`
	CodePackage   []byte                                       `protobuf:"bytes,3,opt,name=code_package,json=codePackage,proto3" json:"code_package,omitempty"`
	ExecEnv       ChaincodeDeploymentSpec_ExecutionEnvironment `protobuf:"varint,4,opt,name=exec_env,json=execEnv,enum=protos.ChaincodeDeploymentSpec_ExecutionEnvironment" json:"exec_env,omitempty"`
}

func (m *ChaincodeDeploymentSpec) Reset()                    { *m = ChaincodeDeploymentSpec{} }
func (m *ChaincodeDeploymentSpec) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeDeploymentSpec) ProtoMessage()               {}
func (*ChaincodeDeploymentSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *ChaincodeDeploymentSpec) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

func (m *ChaincodeDeploymentSpec) GetEffectiveDate() *google_protobuf1.Timestamp {
	if m != nil {
		return m.EffectiveDate
	}
	return nil
}

// Carries the chaincode function and its arguments.
type ChaincodeInvocationSpec struct {
	ChaincodeSpec *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincode_spec,json=chaincodeSpec" json:"chaincode_spec,omitempty"`
	// This field can contain a user-specified ID generation algorithm
	// If supplied, this will be used to generate a ID
	// If not supplied (left empty), sha256base64 will be used
	// The algorithm consists of two parts:
	//  1, a hash function
	//  2, a decoding used to decode user (string) input to bytes
	// Currently, SHA256 with BASE64 is supported (e.g. idGenerationAlg='sha256base64')
	IdGenerationAlg string `protobuf:"bytes,2,opt,name=id_generation_alg,json=idGenerationAlg" json:"id_generation_alg,omitempty"`
}

func (m *ChaincodeInvocationSpec) Reset()                    { *m = ChaincodeInvocationSpec{} }
func (m *ChaincodeInvocationSpec) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeInvocationSpec) ProtoMessage()               {}
func (*ChaincodeInvocationSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ChaincodeInvocationSpec) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

// ChaincodeProposalContext contains proposal data that we send to the chaincode
// container shim and allow the chaincode to access through the shim interface.
type ChaincodeProposalContext struct {
	// Creator corresponds to SignatureHeader.Creator
	Creator []byte `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// Transient corresponds to ChaincodeProposalPayload.Transient
	// TODO: The transient field is supposed to carry application-specific
	// data. They might be realted to access-control, encryption and so on.
	// To simply access to this data, replacing bytes with a map
	// is the next step to be carried.
	Transient []byte `protobuf:"bytes,2,opt,name=transient,proto3" json:"transient,omitempty"`
}

func (m *ChaincodeProposalContext) Reset()                    { *m = ChaincodeProposalContext{} }
func (m *ChaincodeProposalContext) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeProposalContext) ProtoMessage()               {}
func (*ChaincodeProposalContext) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type ChaincodeMessage struct {
	Type            ChaincodeMessage_Type       `protobuf:"varint,1,opt,name=type,enum=protos.ChaincodeMessage_Type" json:"type,omitempty"`
	Timestamp       *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Payload         []byte                      `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Txid            string                      `protobuf:"bytes,4,opt,name=txid" json:"txid,omitempty"`
	ProposalContext *ChaincodeProposalContext   `protobuf:"bytes,5,opt,name=proposal_context,json=proposalContext" json:"proposal_context,omitempty"`
	// event emmited by chaincode. Used only with Init or Invoke.
	// This event is then stored (currently)
	// with Block.NonHashData.TransactionResult
	ChaincodeEvent *ChaincodeEvent `protobuf:"bytes,6,opt,name=chaincode_event,json=chaincodeEvent" json:"chaincode_event,omitempty"`
}

func (m *ChaincodeMessage) Reset()                    { *m = ChaincodeMessage{} }
func (m *ChaincodeMessage) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeMessage) ProtoMessage()               {}
func (*ChaincodeMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *ChaincodeMessage) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ChaincodeMessage) GetProposalContext() *ChaincodeProposalContext {
	if m != nil {
		return m.ProposalContext
	}
	return nil
}

func (m *ChaincodeMessage) GetChaincodeEvent() *ChaincodeEvent {
	if m != nil {
		return m.ChaincodeEvent
	}
	return nil
}

type PutStateInfo struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *PutStateInfo) Reset()                    { *m = PutStateInfo{} }
func (m *PutStateInfo) String() string            { return proto.CompactTextString(m) }
func (*PutStateInfo) ProtoMessage()               {}
func (*PutStateInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

type GetStateByRange struct {
	StartKey string `protobuf:"bytes,1,opt,name=startKey" json:"startKey,omitempty"`
	EndKey   string `protobuf:"bytes,2,opt,name=endKey" json:"endKey,omitempty"`
}

func (m *GetStateByRange) Reset()                    { *m = GetStateByRange{} }
func (m *GetStateByRange) String() string            { return proto.CompactTextString(m) }
func (*GetStateByRange) ProtoMessage()               {}
func (*GetStateByRange) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

type GetQueryResult struct {
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
}

func (m *GetQueryResult) Reset()                    { *m = GetQueryResult{} }
func (m *GetQueryResult) String() string            { return proto.CompactTextString(m) }
func (*GetQueryResult) ProtoMessage()               {}
func (*GetQueryResult) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

type QueryStateNext struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *QueryStateNext) Reset()                    { *m = QueryStateNext{} }
func (m *QueryStateNext) String() string            { return proto.CompactTextString(m) }
func (*QueryStateNext) ProtoMessage()               {}
func (*QueryStateNext) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

type QueryStateClose struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *QueryStateClose) Reset()                    { *m = QueryStateClose{} }
func (m *QueryStateClose) String() string            { return proto.CompactTextString(m) }
func (*QueryStateClose) ProtoMessage()               {}
func (*QueryStateClose) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

type QueryStateKeyValue struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *QueryStateKeyValue) Reset()                    { *m = QueryStateKeyValue{} }
func (m *QueryStateKeyValue) String() string            { return proto.CompactTextString(m) }
func (*QueryStateKeyValue) ProtoMessage()               {}
func (*QueryStateKeyValue) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

type QueryStateResponse struct {
	KeysAndValues []*QueryStateKeyValue `protobuf:"bytes,1,rep,name=keys_and_values,json=keysAndValues" json:"keys_and_values,omitempty"`
	HasMore       bool                  `protobuf:"varint,2,opt,name=has_more,json=hasMore" json:"has_more,omitempty"`
	Id            string                `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *QueryStateResponse) Reset()                    { *m = QueryStateResponse{} }
func (m *QueryStateResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryStateResponse) ProtoMessage()               {}
func (*QueryStateResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *QueryStateResponse) GetKeysAndValues() []*QueryStateKeyValue {
	if m != nil {
		return m.KeysAndValues
	}
	return nil
}

func init() {
	proto.RegisterType((*ChaincodeID)(nil), "protos.ChaincodeID")
	proto.RegisterType((*ChaincodeInput)(nil), "protos.ChaincodeInput")
	proto.RegisterType((*ChaincodeSpec)(nil), "protos.ChaincodeSpec")
	proto.RegisterType((*ChaincodeDeploymentSpec)(nil), "protos.ChaincodeDeploymentSpec")
	proto.RegisterType((*ChaincodeInvocationSpec)(nil), "protos.ChaincodeInvocationSpec")
	proto.RegisterType((*ChaincodeProposalContext)(nil), "protos.ChaincodeProposalContext")
	proto.RegisterType((*ChaincodeMessage)(nil), "protos.ChaincodeMessage")
	proto.RegisterType((*PutStateInfo)(nil), "protos.PutStateInfo")
	proto.RegisterType((*GetStateByRange)(nil), "protos.GetStateByRange")
	proto.RegisterType((*GetQueryResult)(nil), "protos.GetQueryResult")
	proto.RegisterType((*QueryStateNext)(nil), "protos.QueryStateNext")
	proto.RegisterType((*QueryStateClose)(nil), "protos.QueryStateClose")
	proto.RegisterType((*QueryStateKeyValue)(nil), "protos.QueryStateKeyValue")
	proto.RegisterType((*QueryStateResponse)(nil), "protos.QueryStateResponse")
	proto.RegisterEnum("protos.ConfidentialityLevel", ConfidentialityLevel_name, ConfidentialityLevel_value)
	proto.RegisterEnum("protos.ChaincodeSpec_Type", ChaincodeSpec_Type_name, ChaincodeSpec_Type_value)
	proto.RegisterEnum("protos.ChaincodeDeploymentSpec_ExecutionEnvironment", ChaincodeDeploymentSpec_ExecutionEnvironment_name, ChaincodeDeploymentSpec_ExecutionEnvironment_value)
	proto.RegisterEnum("protos.ChaincodeMessage_Type", ChaincodeMessage_Type_name, ChaincodeMessage_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for ChaincodeSupport service

type ChaincodeSupportClient interface {
	Register(ctx context.Context, opts ...grpc.CallOption) (ChaincodeSupport_RegisterClient, error)
}

type chaincodeSupportClient struct {
	cc *grpc.ClientConn
}

func NewChaincodeSupportClient(cc *grpc.ClientConn) ChaincodeSupportClient {
	return &chaincodeSupportClient{cc}
}

func (c *chaincodeSupportClient) Register(ctx context.Context, opts ...grpc.CallOption) (ChaincodeSupport_RegisterClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ChaincodeSupport_serviceDesc.Streams[0], c.cc, "/protos.ChaincodeSupport/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &chaincodeSupportRegisterClient{stream}
	return x, nil
}

type ChaincodeSupport_RegisterClient interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ClientStream
}

type chaincodeSupportRegisterClient struct {
	grpc.ClientStream
}

func (x *chaincodeSupportRegisterClient) Send(m *ChaincodeMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chaincodeSupportRegisterClient) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ChaincodeSupport service

type ChaincodeSupportServer interface {
	Register(ChaincodeSupport_RegisterServer) error
}

func RegisterChaincodeSupportServer(s *grpc.Server, srv ChaincodeSupportServer) {
	s.RegisterService(&_ChaincodeSupport_serviceDesc, srv)
}

func _ChaincodeSupport_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChaincodeSupportServer).Register(&chaincodeSupportRegisterServer{stream})
}

type ChaincodeSupport_RegisterServer interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ServerStream
}

type chaincodeSupportRegisterServer struct {
	grpc.ServerStream
}

func (x *chaincodeSupportRegisterServer) Send(m *ChaincodeMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chaincodeSupportRegisterServer) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChaincodeSupport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ChaincodeSupport",
	HandlerType: (*ChaincodeSupportServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _ChaincodeSupport_Register_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("peer/chaincode.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x56, 0xdb, 0x6e, 0xdb, 0x46,
	0x13, 0x0e, 0x25, 0xd9, 0x96, 0x46, 0xb2, 0xb4, 0xd9, 0xdf, 0xf1, 0xaf, 0x18, 0x2d, 0xea, 0x10,
	0x45, 0xe1, 0x06, 0x85, 0xdc, 0xba, 0x41, 0xd0, 0x8b, 0xa0, 0x05, 0x4d, 0x6e, 0x54, 0xd6, 0x32,
	0xa5, 0xac, 0x68, 0x23, 0xee, 0x0d, 0x41, 0x93, 0x63, 0x99, 0x88, 0x4c, 0xb2, 0xe4, 0x4a, 0xb0,
	0xae, 0xf3, 0x3e, 0x7d, 0x84, 0x3e, 0x51, 0xdf, 0xa1, 0xc5, 0x2e, 0x75, 0x72, 0x94, 0x00, 0xb9,
	0xe8, 0x95, 0xf6, 0x9b, 0xf9, 0xe6, 0xb8, 0x33, 0x2b, 0xc2, 0x5e, 0x8a, 0x98, 0x1d, 0x07, 0xb7,
	0x7e, 0x14, 0x07, 0x49, 0x88, 0x9d, 0x34, 0x4b, 0x44, 0x42, 0xb7, 0xd5, 0x4f, 0x7e, 0xf0, 0xf4,
	0xa1, 0x16, 0xa7, 0x18, 0x8b, 0x82, 0x72, 0xf0, 0xd5, 0x28, 0x49, 0x46, 0x63, 0x3c, 0x56, 0xe8,
	0x7a, 0x72, 0x73, 0x2c, 0xa2, 0x3b, 0xcc, 0x85, 0x7f, 0x97, 0x16, 0x04, 0xbd, 0x0f, 0x75, 0x73,
	0x61, 0x68, 0x5b, 0x94, 0x42, 0x25, 0xf5, 0xc5, 0x6d, 0x5b, 0x3b, 0xd4, 0x8e, 0x6a, 0x5c, 0x9d,
	0xa5, 0x2c, 0xf6, 0xef, 0xb0, 0x5d, 0x2a, 0x64, 0xf2, 0x4c, 0xdb, 0xb0, 0x33, 0xc5, 0x2c, 0x8f,
	0x92, 0xb8, 0x5d, 0x56, 0xe2, 0x05, 0xd4, 0xbf, 0x86, 0xe6, 0xca, 0x61, 0x9c, 0x4e, 0x84, 0xb4,
	0xf7, 0xb3, 0x51, 0xde, 0xd6, 0x0e, 0xcb, 0x47, 0x0d, 0xae, 0xce, 0xfa, 0x3f, 0x1a, 0xec, 0x2e,
	0x69, 0xc3, 0x14, 0x03, 0xda, 0x81, 0x8a, 0x98, 0xa5, 0xa8, 0x22, 0x37, 0x4f, 0x0e, 0x8a, 0xf4,
	0xf2, 0xce, 0x03, 0x52, 0xc7, 0x9d, 0xa5, 0xc8, 0x15, 0x8f, 0xbe, 0x84, 0xc6, 0xb2, 0x62, 0x2f,
	0x0a, 0x55, 0x76, 0xf5, 0x93, 0xff, 0x6d, 0xd8, 0xd9, 0x16, 0xaf, 0x2f, 0x89, 0x76, 0x48, 0xbf,
	0x83, 0xad, 0x48, 0xa6, 0xa5, 0xf2, 0xae, 0x9f, 0xec, 0x6f, 0x1a, 0x48, 0x2d, 0x2f, 0x48, 0xb2,
	0x4e, 0xd9, 0xb1, 0x64, 0x22, 0xda, 0x95, 0x43, 0xed, 0x68, 0x8b, 0x2f, 0xa0, 0xfe, 0x33, 0x54,
	0x64, 0x36, 0x74, 0x17, 0x6a, 0x17, 0x8e, 0xc5, 0x5e, 0xdb, 0x0e, 0xb3, 0xc8, 0x23, 0x0a, 0xb0,
	0xdd, 0xed, 0xf7, 0x0c, 0xa7, 0x4b, 0x34, 0x5a, 0x85, 0x8a, 0xd3, 0xb7, 0x18, 0x29, 0xd1, 0x1d,
	0x28, 0x9b, 0x06, 0x27, 0x65, 0x29, 0xfa, 0xcd, 0xb8, 0x34, 0x48, 0x45, 0xff, 0xab, 0x04, 0xff,
	0x5f, 0xc6, 0xb4, 0x30, 0x1d, 0x27, 0xb3, 0x3b, 0x8c, 0x85, 0xea, 0xc5, 0x2b, 0x68, 0xae, 0x6a,
	0xcb, 0x53, 0x0c, 0x54, 0x57, 0xea, 0x27, 0x4f, 0x3e, 0xda, 0x15, 0xbe, 0x1b, 0x3c, 0xe8, 0xa4,
	0x01, 0x4d, 0xbc, 0xb9, 0xc1, 0x40, 0x44, 0x53, 0xf4, 0x42, 0x5f, 0xe0, 0xbc, 0x37, 0x07, 0x9d,
	0x62, 0x18, 0x3a, 0x8b, 0x61, 0xe8, 0xb8, 0x8b, 0x61, 0xe0, 0xbb, 0x4b, 0x0b, 0xcb, 0x17, 0x48,
	0x9f, 0x41, 0x43, 0xc5, 0x4e, 0xfd, 0xe0, 0x9d, 0x3f, 0x42, 0xd5, 0xab, 0x06, 0xaf, 0x4b, 0xd9,
	0xa0, 0x10, 0xd1, 0x3e, 0x54, 0xf1, 0x1e, 0x03, 0x0f, 0xe3, 0xa9, 0x6a, 0x4d, 0xf3, 0xe4, 0xc5,
	0x46, 0x76, 0x0f, 0xcb, 0xea, 0xb0, 0x7b, 0x0c, 0x26, 0x22, 0x4a, 0x62, 0x16, 0x4f, 0xa3, 0x2c,
	0x89, 0xa5, 0x82, 0xef, 0x48, 0x2f, 0x2c, 0x9e, 0xea, 0x1d, 0xd8, 0xfb, 0x18, 0x41, 0x76, 0xd4,
	0xea, 0x9b, 0x67, 0x8c, 0x17, 0xdd, 0x1d, 0x5e, 0x0d, 0x5d, 0x76, 0x4e, 0x34, 0xfd, 0xbd, 0xb6,
	0xd6, 0x40, 0x3b, 0x9e, 0x26, 0x81, 0x2f, 0x4d, 0xff, 0x83, 0x06, 0x3e, 0x87, 0xc7, 0x51, 0xe8,
	0x8d, 0x30, 0xc6, 0x4c, 0xb9, 0xf4, 0xfc, 0xf1, 0x68, 0x3e, 0xfd, 0xad, 0x28, 0xec, 0x2e, 0xe5,
	0xc6, 0x78, 0xa4, 0x73, 0x68, 0x2f, 0x7d, 0x0d, 0xb2, 0x24, 0x4d, 0x72, 0x7f, 0x6c, 0x26, 0xb1,
	0xc0, 0x7b, 0x35, 0x3c, 0x41, 0x86, 0xbe, 0x48, 0x32, 0x15, 0xbe, 0xc1, 0x17, 0x90, 0x7e, 0x01,
	0x35, 0x91, 0xf9, 0x71, 0x1e, 0x61, 0x2c, 0x94, 0xe7, 0x06, 0x5f, 0x09, 0xf4, 0xbf, 0x2b, 0x40,
	0x96, 0x4e, 0xcf, 0x31, 0xcf, 0x65, 0xbf, 0x7f, 0x78, 0xb0, 0x1f, 0x5f, 0x6e, 0x14, 0x32, 0xe7,
	0xad, 0xaf, 0xc8, 0x4f, 0x50, 0x5b, 0xae, 0xfb, 0x67, 0xcc, 0xc0, 0x8a, 0x2c, 0x33, 0x4f, 0xfd,
	0xd9, 0x38, 0xf1, 0xc3, 0xf9, 0xd5, 0x2f, 0xa0, 0x5c, 0x66, 0x71, 0x1f, 0x85, 0xea, 0xca, 0x6b,
	0x5c, 0x9d, 0xe9, 0x19, 0x90, 0x74, 0x5e, 0xba, 0x17, 0x14, 0xb5, 0xb7, 0xb7, 0x54, 0xb8, 0xc3,
	0x8d, 0x34, 0x3f, 0xe8, 0x11, 0x6f, 0xa5, 0x1f, 0x34, 0xed, 0x17, 0x68, 0xad, 0xae, 0x4e, 0x3d,
	0x65, 0xed, 0xed, 0x4f, 0x6c, 0x2a, 0x93, 0x5a, 0xbe, 0xba, 0x69, 0x85, 0xf5, 0x3f, 0x4b, 0x1f,
	0xdf, 0xcc, 0x06, 0x54, 0x39, 0xeb, 0xda, 0x43, 0x97, 0x71, 0xa2, 0xd1, 0x26, 0xc0, 0x02, 0x31,
	0x8b, 0x94, 0xe4, 0x62, 0xda, 0x8e, 0xed, 0x92, 0x32, 0xad, 0xc1, 0x16, 0x67, 0x86, 0x75, 0x45,
	0x2a, 0xb4, 0x05, 0x75, 0x97, 0x1b, 0xce, 0xd0, 0x30, 0x5d, 0xbb, 0xef, 0x90, 0x2d, 0xe9, 0xd2,
	0xec, 0x9f, 0x0f, 0x7a, 0xcc, 0x65, 0x16, 0xd9, 0x96, 0x54, 0xc6, 0x79, 0x9f, 0x93, 0x1d, 0xa9,
	0xe9, 0x32, 0xd7, 0x1b, 0xba, 0x86, 0xcb, 0x48, 0x55, 0xc2, 0xc1, 0xc5, 0x02, 0xd6, 0x24, 0xb4,
	0x58, 0x6f, 0x0e, 0x81, 0xee, 0x01, 0xb1, 0x9d, 0xcb, 0xfe, 0x19, 0xf3, 0xcc, 0x5f, 0x0d, 0xdb,
	0x31, 0xe5, 0x23, 0x51, 0x2f, 0x12, 0x1c, 0x0e, 0xfa, 0xce, 0x90, 0x91, 0x5d, 0xba, 0x0f, 0x74,
	0xe9, 0xd0, 0x3b, 0xbd, 0xf2, 0xb8, 0xe1, 0x74, 0x19, 0x69, 0x4a, 0x5b, 0x29, 0x7f, 0x73, 0xc1,
	0xf8, 0x95, 0xc7, 0xd9, 0xf0, 0xa2, 0xe7, 0x92, 0x96, 0x94, 0x16, 0x92, 0x82, 0xef, 0xb0, 0xb7,
	0x2e, 0x21, 0xf4, 0x09, 0x3c, 0x5e, 0x97, 0x9a, 0xbd, 0xfe, 0x90, 0x91, 0xc7, 0x32, 0x9b, 0x33,
	0xc6, 0x06, 0x46, 0xcf, 0xbe, 0x64, 0x84, 0xea, 0x2f, 0xa1, 0x31, 0x98, 0x88, 0xa1, 0xf0, 0x05,
	0xda, 0xf1, 0x4d, 0x42, 0x09, 0x94, 0xdf, 0xe1, 0x6c, 0xfe, 0x17, 0x20, 0x8f, 0x74, 0x0f, 0xb6,
	0xa6, 0xfe, 0x78, 0x82, 0xf3, 0x51, 0x2d, 0x80, 0xce, 0xa0, 0xd5, 0xc5, 0xc2, 0xee, 0x74, 0xc6,
	0xfd, 0x78, 0x84, 0xf4, 0x00, 0xaa, 0xb9, 0xf0, 0x33, 0x71, 0xb6, 0xb4, 0x5f, 0x62, 0xba, 0x0f,
	0xdb, 0x18, 0x87, 0x52, 0x53, 0xac, 0xd2, 0x1c, 0xe9, 0xdf, 0x40, 0xb3, 0x8b, 0xe2, 0xcd, 0x04,
	0xb3, 0x19, 0xc7, 0x7c, 0x32, 0x16, 0x32, 0xdc, 0x1f, 0x12, 0xce, 0x5d, 0x14, 0x40, 0x3f, 0x84,
	0xa6, 0x22, 0xa9, 0x80, 0x8e, 0x1c, 0x95, 0x26, 0x94, 0xa2, 0x70, 0x4e, 0x2a, 0x45, 0xa1, 0xfe,
	0x0c, 0x5a, 0x2b, 0x86, 0x39, 0x4e, 0x72, 0xdc, 0xa0, 0xbc, 0x02, 0xba, 0xa2, 0x9c, 0xe1, 0xec,
	0x52, 0x56, 0xf2, 0xd9, 0x15, 0xbf, 0xd7, 0xd6, 0xcd, 0x39, 0xe6, 0x69, 0x12, 0xe7, 0x48, 0x4f,
	0xa1, 0xf5, 0x0e, 0x67, 0xb9, 0xe7, 0xc7, 0xa1, 0xa7, 0x88, 0xc5, 0x7f, 0x5d, 0x7d, 0xf5, 0x2f,
	0xb6, 0x19, 0x93, 0xef, 0x4a, 0x13, 0x23, 0x0e, 0x15, 0xca, 0xe9, 0x53, 0xa8, 0xde, 0xfa, 0xb9,
	0x77, 0x97, 0x64, 0x45, 0xcc, 0x2a, 0xdf, 0xb9, 0xf5, 0xf3, 0xf3, 0x24, 0x5b, 0xd4, 0x50, 0x5e,
	0xd4, 0xf0, 0xfc, 0x05, 0xec, 0x99, 0x49, 0x7c, 0x13, 0x85, 0x18, 0x8b, 0xc8, 0x1f, 0x47, 0x62,
	0xd6, 0xc3, 0x29, 0x8e, 0xe5, 0xe3, 0x38, 0xb8, 0x38, 0xed, 0xd9, 0x26, 0x79, 0x44, 0x09, 0x34,
	0xcc, 0xbe, 0xf3, 0xda, 0xb6, 0x98, 0xe3, 0xda, 0x46, 0x8f, 0x68, 0x27, 0x6f, 0xd7, 0xde, 0x94,
	0xe1, 0x24, 0x4d, 0x93, 0x4c, 0x50, 0x0b, 0xaa, 0x1c, 0x47, 0x51, 0x2e, 0x30, 0xa3, 0xed, 0x4f,
	0xbd, 0x28, 0x07, 0x9f, 0xd4, 0xe8, 0x8f, 0x8e, 0xb4, 0xef, 0xb5, 0x53, 0x13, 0xf6, 0x93, 0x6c,
	0xd4, 0xb9, 0x9d, 0xa5, 0x98, 0x8d, 0x31, 0x1c, 0x61, 0x36, 0x37, 0xf8, 0xfd, 0xdb, 0x51, 0x24,
	0x6e, 0x27, 0xd7, 0x9d, 0x20, 0xb9, 0x3b, 0x5e, 0x53, 0x1f, 0xdf, 0xf8, 0xd7, 0x59, 0x14, 0x14,
	0x1f, 0x25, 0xf9, 0xb1, 0xfc, 0x7a, 0xb9, 0x2e, 0xbe, 0x65, 0x7e, 0xfc, 0x37, 0x00, 0x00, 0xff,
	0xff, 0x14, 0x45, 0xe8, 0x6f, 0xea, 0x08, 0x00, 0x00,
}
