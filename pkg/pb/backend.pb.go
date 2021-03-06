// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/backend.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FunctionConfig_Type int32

const (
	FunctionConfig_GRPC FunctionConfig_Type = 0
	FunctionConfig_REST FunctionConfig_Type = 1
)

var FunctionConfig_Type_name = map[int32]string{
	0: "GRPC",
	1: "REST",
}

var FunctionConfig_Type_value = map[string]int32{
	"GRPC": 0,
	"REST": 1,
}

func (x FunctionConfig_Type) String() string {
	return proto.EnumName(FunctionConfig_Type_name, int32(x))
}

func (FunctionConfig_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{0, 0}
}

type AssignmentFailure_Cause int32

const (
	AssignmentFailure_UNKNOWN          AssignmentFailure_Cause = 0
	AssignmentFailure_TICKET_NOT_FOUND AssignmentFailure_Cause = 1
)

var AssignmentFailure_Cause_name = map[int32]string{
	0: "UNKNOWN",
	1: "TICKET_NOT_FOUND",
}

var AssignmentFailure_Cause_value = map[string]int32{
	"UNKNOWN":          0,
	"TICKET_NOT_FOUND": 1,
}

func (x AssignmentFailure_Cause) String() string {
	return proto.EnumName(AssignmentFailure_Cause_name, int32(x))
}

func (AssignmentFailure_Cause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{8, 0}
}

// FunctionConfig specifies a MMF address and client type for Backend to establish connections with the MMF
type FunctionConfig struct {
	Host                 string              `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 int32               `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Type                 FunctionConfig_Type `protobuf:"varint,3,opt,name=type,proto3,enum=openmatch.FunctionConfig_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *FunctionConfig) Reset()         { *m = FunctionConfig{} }
func (m *FunctionConfig) String() string { return proto.CompactTextString(m) }
func (*FunctionConfig) ProtoMessage()    {}
func (*FunctionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{0}
}

func (m *FunctionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FunctionConfig.Unmarshal(m, b)
}
func (m *FunctionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FunctionConfig.Marshal(b, m, deterministic)
}
func (m *FunctionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionConfig.Merge(m, src)
}
func (m *FunctionConfig) XXX_Size() int {
	return xxx_messageInfo_FunctionConfig.Size(m)
}
func (m *FunctionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionConfig proto.InternalMessageInfo

func (m *FunctionConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *FunctionConfig) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *FunctionConfig) GetType() FunctionConfig_Type {
	if m != nil {
		return m.Type
	}
	return FunctionConfig_GRPC
}

type FetchMatchesRequest struct {
	// A configuration for the MatchFunction server of this FetchMatches call.
	Config *FunctionConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	// A MatchProfile that will be sent to the MatchFunction server of this FetchMatches call.
	Profile              *MatchProfile `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FetchMatchesRequest) Reset()         { *m = FetchMatchesRequest{} }
func (m *FetchMatchesRequest) String() string { return proto.CompactTextString(m) }
func (*FetchMatchesRequest) ProtoMessage()    {}
func (*FetchMatchesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{1}
}

func (m *FetchMatchesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchMatchesRequest.Unmarshal(m, b)
}
func (m *FetchMatchesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchMatchesRequest.Marshal(b, m, deterministic)
}
func (m *FetchMatchesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchMatchesRequest.Merge(m, src)
}
func (m *FetchMatchesRequest) XXX_Size() int {
	return xxx_messageInfo_FetchMatchesRequest.Size(m)
}
func (m *FetchMatchesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchMatchesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchMatchesRequest proto.InternalMessageInfo

func (m *FetchMatchesRequest) GetConfig() *FunctionConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *FetchMatchesRequest) GetProfile() *MatchProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type FetchMatchesResponse struct {
	// A Match generated by the user-defined MMF with the specified MatchProfiles.
	// A valid Match response will contain at least one ticket.
	Match                *Match   `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchMatchesResponse) Reset()         { *m = FetchMatchesResponse{} }
func (m *FetchMatchesResponse) String() string { return proto.CompactTextString(m) }
func (*FetchMatchesResponse) ProtoMessage()    {}
func (*FetchMatchesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{2}
}

func (m *FetchMatchesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchMatchesResponse.Unmarshal(m, b)
}
func (m *FetchMatchesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchMatchesResponse.Marshal(b, m, deterministic)
}
func (m *FetchMatchesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchMatchesResponse.Merge(m, src)
}
func (m *FetchMatchesResponse) XXX_Size() int {
	return xxx_messageInfo_FetchMatchesResponse.Size(m)
}
func (m *FetchMatchesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchMatchesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchMatchesResponse proto.InternalMessageInfo

func (m *FetchMatchesResponse) GetMatch() *Match {
	if m != nil {
		return m.Match
	}
	return nil
}

type ReleaseTicketsRequest struct {
	// TicketIds is a list of string representing Open Match generated Ids to be re-enabled for MMF querying
	// because they are no longer awaiting assignment from a previous match result
	TicketIds            []string `protobuf:"bytes,1,rep,name=ticket_ids,json=ticketIds,proto3" json:"ticket_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseTicketsRequest) Reset()         { *m = ReleaseTicketsRequest{} }
func (m *ReleaseTicketsRequest) String() string { return proto.CompactTextString(m) }
func (*ReleaseTicketsRequest) ProtoMessage()    {}
func (*ReleaseTicketsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{3}
}

func (m *ReleaseTicketsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseTicketsRequest.Unmarshal(m, b)
}
func (m *ReleaseTicketsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseTicketsRequest.Marshal(b, m, deterministic)
}
func (m *ReleaseTicketsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseTicketsRequest.Merge(m, src)
}
func (m *ReleaseTicketsRequest) XXX_Size() int {
	return xxx_messageInfo_ReleaseTicketsRequest.Size(m)
}
func (m *ReleaseTicketsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseTicketsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseTicketsRequest proto.InternalMessageInfo

func (m *ReleaseTicketsRequest) GetTicketIds() []string {
	if m != nil {
		return m.TicketIds
	}
	return nil
}

type ReleaseTicketsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseTicketsResponse) Reset()         { *m = ReleaseTicketsResponse{} }
func (m *ReleaseTicketsResponse) String() string { return proto.CompactTextString(m) }
func (*ReleaseTicketsResponse) ProtoMessage()    {}
func (*ReleaseTicketsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{4}
}

func (m *ReleaseTicketsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseTicketsResponse.Unmarshal(m, b)
}
func (m *ReleaseTicketsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseTicketsResponse.Marshal(b, m, deterministic)
}
func (m *ReleaseTicketsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseTicketsResponse.Merge(m, src)
}
func (m *ReleaseTicketsResponse) XXX_Size() int {
	return xxx_messageInfo_ReleaseTicketsResponse.Size(m)
}
func (m *ReleaseTicketsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseTicketsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseTicketsResponse proto.InternalMessageInfo

type ReleaseAllTicketsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseAllTicketsRequest) Reset()         { *m = ReleaseAllTicketsRequest{} }
func (m *ReleaseAllTicketsRequest) String() string { return proto.CompactTextString(m) }
func (*ReleaseAllTicketsRequest) ProtoMessage()    {}
func (*ReleaseAllTicketsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{5}
}

func (m *ReleaseAllTicketsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseAllTicketsRequest.Unmarshal(m, b)
}
func (m *ReleaseAllTicketsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseAllTicketsRequest.Marshal(b, m, deterministic)
}
func (m *ReleaseAllTicketsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseAllTicketsRequest.Merge(m, src)
}
func (m *ReleaseAllTicketsRequest) XXX_Size() int {
	return xxx_messageInfo_ReleaseAllTicketsRequest.Size(m)
}
func (m *ReleaseAllTicketsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseAllTicketsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseAllTicketsRequest proto.InternalMessageInfo

type ReleaseAllTicketsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseAllTicketsResponse) Reset()         { *m = ReleaseAllTicketsResponse{} }
func (m *ReleaseAllTicketsResponse) String() string { return proto.CompactTextString(m) }
func (*ReleaseAllTicketsResponse) ProtoMessage()    {}
func (*ReleaseAllTicketsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{6}
}

func (m *ReleaseAllTicketsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseAllTicketsResponse.Unmarshal(m, b)
}
func (m *ReleaseAllTicketsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseAllTicketsResponse.Marshal(b, m, deterministic)
}
func (m *ReleaseAllTicketsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseAllTicketsResponse.Merge(m, src)
}
func (m *ReleaseAllTicketsResponse) XXX_Size() int {
	return xxx_messageInfo_ReleaseAllTicketsResponse.Size(m)
}
func (m *ReleaseAllTicketsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseAllTicketsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseAllTicketsResponse proto.InternalMessageInfo

// AssignmentGroup contains an Assignment and the Tickets to which it should be applied.
type AssignmentGroup struct {
	// TicketIds is a list of strings representing Open Match generated Ids which apply to an Assignment.
	TicketIds []string `protobuf:"bytes,1,rep,name=ticket_ids,json=ticketIds,proto3" json:"ticket_ids,omitempty"`
	// An Assignment specifies game connection related information to be associated with the TicketIds.
	Assignment           *Assignment `protobuf:"bytes,2,opt,name=assignment,proto3" json:"assignment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AssignmentGroup) Reset()         { *m = AssignmentGroup{} }
func (m *AssignmentGroup) String() string { return proto.CompactTextString(m) }
func (*AssignmentGroup) ProtoMessage()    {}
func (*AssignmentGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{7}
}

func (m *AssignmentGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignmentGroup.Unmarshal(m, b)
}
func (m *AssignmentGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignmentGroup.Marshal(b, m, deterministic)
}
func (m *AssignmentGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignmentGroup.Merge(m, src)
}
func (m *AssignmentGroup) XXX_Size() int {
	return xxx_messageInfo_AssignmentGroup.Size(m)
}
func (m *AssignmentGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignmentGroup.DiscardUnknown(m)
}

var xxx_messageInfo_AssignmentGroup proto.InternalMessageInfo

func (m *AssignmentGroup) GetTicketIds() []string {
	if m != nil {
		return m.TicketIds
	}
	return nil
}

func (m *AssignmentGroup) GetAssignment() *Assignment {
	if m != nil {
		return m.Assignment
	}
	return nil
}

// AssignmentFailure contains the id of the Ticket that failed the Assignment and the failure status.
type AssignmentFailure struct {
	TicketId             string                  `protobuf:"bytes,1,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
	Cause                AssignmentFailure_Cause `protobuf:"varint,2,opt,name=cause,proto3,enum=openmatch.AssignmentFailure_Cause" json:"cause,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *AssignmentFailure) Reset()         { *m = AssignmentFailure{} }
func (m *AssignmentFailure) String() string { return proto.CompactTextString(m) }
func (*AssignmentFailure) ProtoMessage()    {}
func (*AssignmentFailure) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{8}
}

func (m *AssignmentFailure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignmentFailure.Unmarshal(m, b)
}
func (m *AssignmentFailure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignmentFailure.Marshal(b, m, deterministic)
}
func (m *AssignmentFailure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignmentFailure.Merge(m, src)
}
func (m *AssignmentFailure) XXX_Size() int {
	return xxx_messageInfo_AssignmentFailure.Size(m)
}
func (m *AssignmentFailure) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignmentFailure.DiscardUnknown(m)
}

var xxx_messageInfo_AssignmentFailure proto.InternalMessageInfo

func (m *AssignmentFailure) GetTicketId() string {
	if m != nil {
		return m.TicketId
	}
	return ""
}

func (m *AssignmentFailure) GetCause() AssignmentFailure_Cause {
	if m != nil {
		return m.Cause
	}
	return AssignmentFailure_UNKNOWN
}

type AssignTicketsRequest struct {
	// Assignments is a list of assignment groups that contain assignment and the Tickets to which they should be applied.
	Assignments          []*AssignmentGroup `protobuf:"bytes,1,rep,name=assignments,proto3" json:"assignments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AssignTicketsRequest) Reset()         { *m = AssignTicketsRequest{} }
func (m *AssignTicketsRequest) String() string { return proto.CompactTextString(m) }
func (*AssignTicketsRequest) ProtoMessage()    {}
func (*AssignTicketsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{9}
}

func (m *AssignTicketsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignTicketsRequest.Unmarshal(m, b)
}
func (m *AssignTicketsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignTicketsRequest.Marshal(b, m, deterministic)
}
func (m *AssignTicketsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignTicketsRequest.Merge(m, src)
}
func (m *AssignTicketsRequest) XXX_Size() int {
	return xxx_messageInfo_AssignTicketsRequest.Size(m)
}
func (m *AssignTicketsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignTicketsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AssignTicketsRequest proto.InternalMessageInfo

func (m *AssignTicketsRequest) GetAssignments() []*AssignmentGroup {
	if m != nil {
		return m.Assignments
	}
	return nil
}

type AssignTicketsResponse struct {
	// Failures is a list of all the Tickets that failed assignment along with the cause of failure.
	Failures             []*AssignmentFailure `protobuf:"bytes,1,rep,name=failures,proto3" json:"failures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AssignTicketsResponse) Reset()         { *m = AssignTicketsResponse{} }
func (m *AssignTicketsResponse) String() string { return proto.CompactTextString(m) }
func (*AssignTicketsResponse) ProtoMessage()    {}
func (*AssignTicketsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{10}
}

func (m *AssignTicketsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignTicketsResponse.Unmarshal(m, b)
}
func (m *AssignTicketsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignTicketsResponse.Marshal(b, m, deterministic)
}
func (m *AssignTicketsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignTicketsResponse.Merge(m, src)
}
func (m *AssignTicketsResponse) XXX_Size() int {
	return xxx_messageInfo_AssignTicketsResponse.Size(m)
}
func (m *AssignTicketsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignTicketsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AssignTicketsResponse proto.InternalMessageInfo

func (m *AssignTicketsResponse) GetFailures() []*AssignmentFailure {
	if m != nil {
		return m.Failures
	}
	return nil
}

func init() {
	proto.RegisterEnum("openmatch.FunctionConfig_Type", FunctionConfig_Type_name, FunctionConfig_Type_value)
	proto.RegisterEnum("openmatch.AssignmentFailure_Cause", AssignmentFailure_Cause_name, AssignmentFailure_Cause_value)
	proto.RegisterType((*FunctionConfig)(nil), "openmatch.FunctionConfig")
	proto.RegisterType((*FetchMatchesRequest)(nil), "openmatch.FetchMatchesRequest")
	proto.RegisterType((*FetchMatchesResponse)(nil), "openmatch.FetchMatchesResponse")
	proto.RegisterType((*ReleaseTicketsRequest)(nil), "openmatch.ReleaseTicketsRequest")
	proto.RegisterType((*ReleaseTicketsResponse)(nil), "openmatch.ReleaseTicketsResponse")
	proto.RegisterType((*ReleaseAllTicketsRequest)(nil), "openmatch.ReleaseAllTicketsRequest")
	proto.RegisterType((*ReleaseAllTicketsResponse)(nil), "openmatch.ReleaseAllTicketsResponse")
	proto.RegisterType((*AssignmentGroup)(nil), "openmatch.AssignmentGroup")
	proto.RegisterType((*AssignmentFailure)(nil), "openmatch.AssignmentFailure")
	proto.RegisterType((*AssignTicketsRequest)(nil), "openmatch.AssignTicketsRequest")
	proto.RegisterType((*AssignTicketsResponse)(nil), "openmatch.AssignTicketsResponse")
}

func init() { proto.RegisterFile("api/backend.proto", fileDescriptor_8dab762378f455cd) }

var fileDescriptor_8dab762378f455cd = []byte{
	// 929 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xee, 0x3a, 0xce, 0x8f, 0x4f, 0xc0, 0x38, 0x43, 0x52, 0xdc, 0x6d, 0xa1, 0x9b, 0x29, 0x2d,
	0x91, 0xa9, 0xbd, 0xce, 0x12, 0x50, 0x65, 0x7e, 0xd4, 0xd4, 0x4d, 0xaa, 0xa8, 0xc5, 0x29, 0x1b,
	0x17, 0x24, 0x6e, 0xa2, 0xf5, 0xfa, 0x78, 0xbd, 0x64, 0xbd, 0xb3, 0xec, 0xcc, 0xa6, 0x54, 0x48,
	0x08, 0x21, 0x2e, 0x10, 0x57, 0x08, 0x24, 0x2e, 0x78, 0x04, 0x24, 0x2e, 0x78, 0x16, 0x6e, 0x78,
	0x00, 0x1e, 0x04, 0xed, 0xcc, 0xfa, 0xdf, 0x49, 0xaf, 0xbc, 0x33, 0xe7, 0x3b, 0xdf, 0xf7, 0xcd,
	0x99, 0x73, 0x3c, 0xb0, 0xe1, 0x44, 0xbe, 0xd9, 0x71, 0xdc, 0x33, 0x0c, 0xbb, 0xb5, 0x28, 0x66,
	0x82, 0x91, 0x02, 0x8b, 0x30, 0x1c, 0x38, 0xc2, 0xed, 0xeb, 0x24, 0x8d, 0x0e, 0x90, 0x73, 0xc7,
	0x43, 0xae, 0xc2, 0xfa, 0x0d, 0x8f, 0x31, 0x2f, 0x40, 0x33, 0x0d, 0x39, 0x61, 0xc8, 0x84, 0x23,
	0x7c, 0x16, 0x0e, 0xa3, 0x77, 0xe5, 0x8f, 0x5b, 0xf5, 0x30, 0xac, 0xf2, 0xe7, 0x8e, 0xe7, 0x61,
	0x6c, 0xb2, 0x48, 0x22, 0xe6, 0xd1, 0xf4, 0x27, 0x0d, 0x8a, 0x87, 0x49, 0xe8, 0xa6, 0x7b, 0x4d,
	0x16, 0xf6, 0x7c, 0x8f, 0x10, 0xc8, 0xf7, 0x19, 0x17, 0x65, 0xcd, 0xd0, 0x76, 0x0a, 0xb6, 0xfc,
	0x4e, 0xf7, 0x22, 0x16, 0x8b, 0x72, 0xce, 0xd0, 0x76, 0x96, 0x6d, 0xf9, 0x4d, 0x2c, 0xc8, 0x8b,
	0x17, 0x11, 0x96, 0x97, 0x0c, 0x6d, 0xa7, 0x68, 0xbd, 0x55, 0x1b, 0x99, 0xae, 0x4d, 0x13, 0xd6,
	0xda, 0x2f, 0x22, 0xb4, 0x25, 0x96, 0xea, 0x90, 0x4f, 0x57, 0x64, 0x0d, 0xf2, 0x8f, 0xec, 0xa7,
	0xcd, 0xd2, 0x95, 0xf4, 0xcb, 0x3e, 0x38, 0x69, 0x97, 0x34, 0xfa, 0x2d, 0xbc, 0x7e, 0x88, 0xc2,
	0xed, 0x7f, 0x9a, 0x72, 0x20, 0xb7, 0xf1, 0xeb, 0x04, 0xb9, 0x20, 0xbb, 0xb0, 0xe2, 0x4a, 0x1e,
	0x69, 0x68, 0xdd, 0xba, 0x76, 0xa1, 0x90, 0x9d, 0x01, 0xc9, 0x2e, 0xac, 0x46, 0x31, 0xeb, 0xf9,
	0x01, 0x4a, 0xc3, 0xeb, 0xd6, 0x1b, 0x13, 0x39, 0x92, 0xfe, 0xa9, 0x0a, 0xdb, 0x43, 0x1c, 0xfd,
	0x04, 0x36, 0xa7, 0xc5, 0x79, 0xc4, 0x42, 0x8e, 0xe4, 0x0e, 0x2c, 0xcb, 0xb4, 0x4c, 0xbc, 0x34,
	0x4b, 0x64, 0xab, 0x30, 0xfd, 0x00, 0xb6, 0x6c, 0x0c, 0xd0, 0xe1, 0xd8, 0xf6, 0xdd, 0x33, 0x14,
	0x23, 0xfb, 0x6f, 0x02, 0x08, 0xb9, 0x73, 0xea, 0x77, 0x79, 0x59, 0x33, 0x96, 0x76, 0x0a, 0x76,
	0x41, 0xed, 0x1c, 0x75, 0x39, 0x2d, 0xc3, 0xd5, 0xd9, 0x3c, 0xa5, 0x4c, 0x75, 0x28, 0x67, 0x91,
	0xfd, 0x20, 0x98, 0x26, 0xa5, 0xd7, 0xe1, 0xda, 0x82, 0x58, 0x96, 0xe8, 0xc1, 0x6b, 0xfb, 0x9c,
	0xfb, 0x5e, 0x38, 0xc0, 0x50, 0x3c, 0x8a, 0x59, 0x12, 0xbd, 0xc4, 0x04, 0x79, 0x1f, 0xc0, 0x19,
	0x65, 0x64, 0x25, 0xdb, 0x9a, 0x38, 0xe9, 0x98, 0xce, 0x9e, 0x00, 0xd2, 0xdf, 0x35, 0xd8, 0x18,
	0x87, 0x0e, 0x1d, 0x3f, 0x48, 0x62, 0x24, 0xd7, 0xa1, 0x30, 0xd2, 0xca, 0x7a, 0x68, 0x6d, 0x28,
	0x45, 0xee, 0xc1, 0xb2, 0xeb, 0x24, 0x5c, 0xdd, 0x4b, 0xd1, 0xa2, 0x0b, 0x45, 0x32, 0xa6, 0x5a,
	0x33, 0x45, 0xda, 0x2a, 0x81, 0x56, 0x60, 0x59, 0xae, 0xc9, 0x3a, 0xac, 0x3e, 0x6b, 0x3d, 0x6e,
	0x1d, 0x7f, 0xd1, 0x2a, 0x5d, 0x21, 0x9b, 0x50, 0x6a, 0x1f, 0x35, 0x1f, 0x1f, 0xb4, 0x4f, 0x5b,
	0xc7, 0xed, 0xd3, 0xc3, 0xe3, 0x67, 0xad, 0x87, 0x25, 0x8d, 0xb6, 0x61, 0x53, 0xb1, 0xcd, 0xdc,
	0xc5, 0x47, 0xb0, 0x3e, 0xb6, 0xaf, 0xea, 0xb0, 0x6e, 0xe9, 0x0b, 0x3d, 0xc8, 0xba, 0xd9, 0x93,
	0x70, 0xfa, 0x19, 0x6c, 0xcd, 0xb0, 0x66, 0x3d, 0x72, 0x0f, 0xd6, 0x7a, 0xca, 0xf2, 0x90, 0xf3,
	0xc6, 0x65, 0xe7, 0xb2, 0x47, 0x68, 0xeb, 0xaf, 0x3c, 0x14, 0x1f, 0xa8, 0xd1, 0x3f, 0xc1, 0xf8,
	0xdc, 0x77, 0x91, 0x7c, 0x07, 0xaf, 0x4c, 0x36, 0x22, 0x99, 0x9a, 0xab, 0xf9, 0xf1, 0xd0, 0x6f,
	0x5e, 0x18, 0xcf, 0xda, 0xe1, 0xdd, 0x1f, 0xfe, 0xf9, 0xef, 0xb7, 0xdc, 0x6d, 0x6a, 0x98, 0xe7,
	0xbb, 0xc3, 0xff, 0x19, 0xae, 0xc4, 0xcc, 0x81, 0xc2, 0x36, 0x7a, 0x69, 0x62, 0x43, 0xab, 0xd4,
	0x35, 0xf2, 0xbd, 0x06, 0xaf, 0x4e, 0x1d, 0x93, 0xdc, 0x9c, 0x3b, 0xcc, 0x74, 0x59, 0x75, 0xe3,
	0x62, 0x40, 0xe6, 0xe1, 0xae, 0xf4, 0x70, 0x87, 0x6e, 0x2f, 0xf0, 0xa0, 0x7a, 0x83, 0x37, 0x54,
	0xa9, 0x1b, 0x5a, 0x85, 0xfc, 0xa8, 0x41, 0x71, 0x7a, 0x28, 0xc8, 0xa4, 0xc4, 0xc2, 0x39, 0xd3,
	0xb7, 0x2f, 0x41, 0x64, 0x2e, 0xaa, 0xd2, 0xc5, 0x3b, 0x94, 0x5e, 0xe2, 0x22, 0x56, 0xa9, 0xa9,
	0x8d, 0x5f, 0x34, 0xd8, 0x98, 0x9b, 0x32, 0x72, 0x6b, 0x5e, 0x67, 0x6e, 0x3e, 0xf5, 0xb7, 0x2f,
	0x07, 0x65, 0x7e, 0xea, 0xd2, 0x4f, 0x85, 0xde, 0x7e, 0xb9, 0x1f, 0x27, 0x08, 0x1a, 0x5a, 0xe5,
	0xc1, 0xcf, 0x4b, 0xbf, 0xee, 0xff, 0x9b, 0x23, 0x7f, 0x6b, 0xb0, 0x9a, 0xb5, 0x0d, 0x3d, 0x02,
	0x38, 0x8e, 0x30, 0x34, 0xe4, 0xb5, 0x93, 0xab, 0x7d, 0x21, 0x22, 0xde, 0x30, 0xcd, 0x54, 0xbf,
	0xaa, 0x0c, 0x74, 0xf1, 0x5c, 0xbf, 0x35, 0x5e, 0x57, 0xbb, 0x3e, 0x77, 0x13, 0xce, 0xef, 0xab,
	0x57, 0xc4, 0x4b, 0x1b, 0x9d, 0xd7, 0x5c, 0x36, 0xa8, 0x7c, 0x0e, 0x64, 0x3f, 0x72, 0xdc, 0x3e,
	0x1a, 0x56, 0xad, 0x6e, 0x3c, 0xf1, 0x5d, 0x4c, 0xbb, 0xfb, 0xfe, 0x90, 0xd2, 0xf3, 0x45, 0x3f,
	0xe9, 0xa4, 0x48, 0x53, 0xa5, 0xf6, 0x58, 0xec, 0x39, 0x03, 0xe4, 0x13, 0x62, 0x66, 0x27, 0x60,
	0x1d, 0x73, 0xe0, 0x70, 0x81, 0xb1, 0xf9, 0xe4, 0xa8, 0x79, 0xd0, 0x3a, 0x39, 0xb0, 0x96, 0x76,
	0x6b, 0xf5, 0x4a, 0x4e, 0xcb, 0x59, 0x25, 0x27, 0x8a, 0x02, 0xdf, 0x95, 0x0f, 0x90, 0xf9, 0x15,
	0x67, 0x61, 0x63, 0x6e, 0xc7, 0xfe, 0x10, 0x96, 0xf6, 0xea, 0x7b, 0x64, 0x0f, 0x2a, 0x36, 0x8a,
	0x24, 0x0e, 0xb1, 0x6b, 0x3c, 0xef, 0x63, 0x68, 0x88, 0x3e, 0x1a, 0x31, 0x72, 0x96, 0xc4, 0x2e,
	0x1a, 0x5d, 0x86, 0xdc, 0x08, 0x99, 0x30, 0xf0, 0x1b, 0x9f, 0x8b, 0x1a, 0x59, 0x81, 0xfc, 0x1f,
	0x39, 0x6d, 0x35, 0xfe, 0x18, 0xca, 0xe3, 0x62, 0x18, 0x0f, 0x99, 0x9b, 0xa4, 0x73, 0x27, 0xd9,
	0xc9, 0xf6, 0xe2, 0xd2, 0x98, 0xdc, 0x17, 0x68, 0x76, 0x99, 0xcb, 0xcd, 0x2f, 0x8d, 0x99, 0xd0,
	0xc4, 0xb9, 0xa2, 0x33, 0xcf, 0x8c, 0x3a, 0x7f, 0xe6, 0x0a, 0x29, 0xbf, 0xa4, 0xef, 0xac, 0xc8,
	0x17, 0xf4, 0xbd, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xea, 0xe8, 0xfb, 0x0c, 0xc1, 0x07, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BackendServiceClient is the client API for BackendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BackendServiceClient interface {
	// FetchMatches triggers a MatchFunction with the specified MatchProfile and
	// returns a set of matches generated by the Match Making Function, and
	// accepted by the evaluator.
	// Tickets in matches returned by FetchMatches are moved from active to
	// pending, and will not be returned by query.
	FetchMatches(ctx context.Context, in *FetchMatchesRequest, opts ...grpc.CallOption) (BackendService_FetchMatchesClient, error)
	// AssignTickets overwrites the Assignment field of the input TicketIds.
	AssignTickets(ctx context.Context, in *AssignTicketsRequest, opts ...grpc.CallOption) (*AssignTicketsResponse, error)
	// ReleaseTickets moves tickets from the pending state, to the active state.
	// This enables them to be returned by query, and find different matches.
	//
	// BETA FEATURE WARNING:  This call and the associated Request and Response
	// messages are not finalized and still subject to possible change or removal.
	ReleaseTickets(ctx context.Context, in *ReleaseTicketsRequest, opts ...grpc.CallOption) (*ReleaseTicketsResponse, error)
	// ReleaseAllTickets moves all tickets from the pending state, to the active
	// state. This enables them to be returned by query, and find different
	// matches.
	//
	// BETA FEATURE WARNING:  This call and the associated Request and Response
	// messages are not finalized and still subject to possible change or removal.
	ReleaseAllTickets(ctx context.Context, in *ReleaseAllTicketsRequest, opts ...grpc.CallOption) (*ReleaseAllTicketsResponse, error)
}

type backendServiceClient struct {
	cc *grpc.ClientConn
}

func NewBackendServiceClient(cc *grpc.ClientConn) BackendServiceClient {
	return &backendServiceClient{cc}
}

func (c *backendServiceClient) FetchMatches(ctx context.Context, in *FetchMatchesRequest, opts ...grpc.CallOption) (BackendService_FetchMatchesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BackendService_serviceDesc.Streams[0], "/openmatch.BackendService/FetchMatches", opts...)
	if err != nil {
		return nil, err
	}
	x := &backendServiceFetchMatchesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BackendService_FetchMatchesClient interface {
	Recv() (*FetchMatchesResponse, error)
	grpc.ClientStream
}

type backendServiceFetchMatchesClient struct {
	grpc.ClientStream
}

func (x *backendServiceFetchMatchesClient) Recv() (*FetchMatchesResponse, error) {
	m := new(FetchMatchesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *backendServiceClient) AssignTickets(ctx context.Context, in *AssignTicketsRequest, opts ...grpc.CallOption) (*AssignTicketsResponse, error) {
	out := new(AssignTicketsResponse)
	err := c.cc.Invoke(ctx, "/openmatch.BackendService/AssignTickets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendServiceClient) ReleaseTickets(ctx context.Context, in *ReleaseTicketsRequest, opts ...grpc.CallOption) (*ReleaseTicketsResponse, error) {
	out := new(ReleaseTicketsResponse)
	err := c.cc.Invoke(ctx, "/openmatch.BackendService/ReleaseTickets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendServiceClient) ReleaseAllTickets(ctx context.Context, in *ReleaseAllTicketsRequest, opts ...grpc.CallOption) (*ReleaseAllTicketsResponse, error) {
	out := new(ReleaseAllTicketsResponse)
	err := c.cc.Invoke(ctx, "/openmatch.BackendService/ReleaseAllTickets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackendServiceServer is the server API for BackendService service.
type BackendServiceServer interface {
	// FetchMatches triggers a MatchFunction with the specified MatchProfile and
	// returns a set of matches generated by the Match Making Function, and
	// accepted by the evaluator.
	// Tickets in matches returned by FetchMatches are moved from active to
	// pending, and will not be returned by query.
	FetchMatches(*FetchMatchesRequest, BackendService_FetchMatchesServer) error
	// AssignTickets overwrites the Assignment field of the input TicketIds.
	AssignTickets(context.Context, *AssignTicketsRequest) (*AssignTicketsResponse, error)
	// ReleaseTickets moves tickets from the pending state, to the active state.
	// This enables them to be returned by query, and find different matches.
	//
	// BETA FEATURE WARNING:  This call and the associated Request and Response
	// messages are not finalized and still subject to possible change or removal.
	ReleaseTickets(context.Context, *ReleaseTicketsRequest) (*ReleaseTicketsResponse, error)
	// ReleaseAllTickets moves all tickets from the pending state, to the active
	// state. This enables them to be returned by query, and find different
	// matches.
	//
	// BETA FEATURE WARNING:  This call and the associated Request and Response
	// messages are not finalized and still subject to possible change or removal.
	ReleaseAllTickets(context.Context, *ReleaseAllTicketsRequest) (*ReleaseAllTicketsResponse, error)
}

// UnimplementedBackendServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBackendServiceServer struct {
}

func (*UnimplementedBackendServiceServer) FetchMatches(req *FetchMatchesRequest, srv BackendService_FetchMatchesServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchMatches not implemented")
}
func (*UnimplementedBackendServiceServer) AssignTickets(ctx context.Context, req *AssignTicketsRequest) (*AssignTicketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignTickets not implemented")
}
func (*UnimplementedBackendServiceServer) ReleaseTickets(ctx context.Context, req *ReleaseTicketsRequest) (*ReleaseTicketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseTickets not implemented")
}
func (*UnimplementedBackendServiceServer) ReleaseAllTickets(ctx context.Context, req *ReleaseAllTicketsRequest) (*ReleaseAllTicketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseAllTickets not implemented")
}

func RegisterBackendServiceServer(s *grpc.Server, srv BackendServiceServer) {
	s.RegisterService(&_BackendService_serviceDesc, srv)
}

func _BackendService_FetchMatches_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FetchMatchesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BackendServiceServer).FetchMatches(m, &backendServiceFetchMatchesServer{stream})
}

type BackendService_FetchMatchesServer interface {
	Send(*FetchMatchesResponse) error
	grpc.ServerStream
}

type backendServiceFetchMatchesServer struct {
	grpc.ServerStream
}

func (x *backendServiceFetchMatchesServer) Send(m *FetchMatchesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BackendService_AssignTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServiceServer).AssignTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openmatch.BackendService/AssignTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServiceServer).AssignTickets(ctx, req.(*AssignTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackendService_ReleaseTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServiceServer).ReleaseTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openmatch.BackendService/ReleaseTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServiceServer).ReleaseTickets(ctx, req.(*ReleaseTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackendService_ReleaseAllTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseAllTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServiceServer).ReleaseAllTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openmatch.BackendService/ReleaseAllTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServiceServer).ReleaseAllTickets(ctx, req.(*ReleaseAllTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BackendService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openmatch.BackendService",
	HandlerType: (*BackendServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssignTickets",
			Handler:    _BackendService_AssignTickets_Handler,
		},
		{
			MethodName: "ReleaseTickets",
			Handler:    _BackendService_ReleaseTickets_Handler,
		},
		{
			MethodName: "ReleaseAllTickets",
			Handler:    _BackendService_ReleaseAllTickets_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchMatches",
			Handler:       _BackendService_FetchMatches_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/backend.proto",
}
