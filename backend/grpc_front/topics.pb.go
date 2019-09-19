// Code generated by protoc-gen-go. DO NOT EDIT.
// source: topics.proto

package topics

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Message struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SenderName           string   `protobuf:"bytes,2,opt,name=senderName,proto3" json:"senderName,omitempty"`
	Profile              string   `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	Message              string   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_72af2b0eeeefc419, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetSenderName() string {
	if m != nil {
		return m.SenderName
	}
	return ""
}

func (m *Message) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *Message) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type TopicInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	NumReplies           uint32   `protobuf:"varint,3,opt,name=numReplies,proto3" json:"numReplies,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicInfo) Reset()         { *m = TopicInfo{} }
func (m *TopicInfo) String() string { return proto.CompactTextString(m) }
func (*TopicInfo) ProtoMessage()    {}
func (*TopicInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_72af2b0eeeefc419, []int{1}
}

func (m *TopicInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicInfo.Unmarshal(m, b)
}
func (m *TopicInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicInfo.Marshal(b, m, deterministic)
}
func (m *TopicInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicInfo.Merge(m, src)
}
func (m *TopicInfo) XXX_Size() int {
	return xxx_messageInfo_TopicInfo.Size(m)
}
func (m *TopicInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TopicInfo proto.InternalMessageInfo

func (m *TopicInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TopicInfo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TopicInfo) GetNumReplies() uint32 {
	if m != nil {
		return m.NumReplies
	}
	return 0
}

type TopicRequest struct {
	BoardId              string   `protobuf:"bytes,1,opt,name=boardId,proto3" json:"boardId,omitempty"`
	StartFrom            uint32   `protobuf:"varint,2,opt,name=startFrom,proto3" json:"startFrom,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicRequest) Reset()         { *m = TopicRequest{} }
func (m *TopicRequest) String() string { return proto.CompactTextString(m) }
func (*TopicRequest) ProtoMessage()    {}
func (*TopicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72af2b0eeeefc419, []int{2}
}

func (m *TopicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicRequest.Unmarshal(m, b)
}
func (m *TopicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicRequest.Marshal(b, m, deterministic)
}
func (m *TopicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicRequest.Merge(m, src)
}
func (m *TopicRequest) XXX_Size() int {
	return xxx_messageInfo_TopicRequest.Size(m)
}
func (m *TopicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TopicRequest proto.InternalMessageInfo

func (m *TopicRequest) GetBoardId() string {
	if m != nil {
		return m.BoardId
	}
	return ""
}

func (m *TopicRequest) GetStartFrom() uint32 {
	if m != nil {
		return m.StartFrom
	}
	return 0
}

type MessageRequest struct {
	TopicId              string   `protobuf:"bytes,1,opt,name=topicId,proto3" json:"topicId,omitempty"`
	StartFrom            uint32   `protobuf:"varint,2,opt,name=startFrom,proto3" json:"startFrom,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageRequest) Reset()         { *m = MessageRequest{} }
func (m *MessageRequest) String() string { return proto.CompactTextString(m) }
func (*MessageRequest) ProtoMessage()    {}
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72af2b0eeeefc419, []int{3}
}

func (m *MessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageRequest.Unmarshal(m, b)
}
func (m *MessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageRequest.Marshal(b, m, deterministic)
}
func (m *MessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageRequest.Merge(m, src)
}
func (m *MessageRequest) XXX_Size() int {
	return xxx_messageInfo_MessageRequest.Size(m)
}
func (m *MessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessageRequest proto.InternalMessageInfo

func (m *MessageRequest) GetTopicId() string {
	if m != nil {
		return m.TopicId
	}
	return ""
}

func (m *MessageRequest) GetStartFrom() uint32 {
	if m != nil {
		return m.StartFrom
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "topics.Message")
	proto.RegisterType((*TopicInfo)(nil), "topics.TopicInfo")
	proto.RegisterType((*TopicRequest)(nil), "topics.TopicRequest")
	proto.RegisterType((*MessageRequest)(nil), "topics.MessageRequest")
}

func init() { proto.RegisterFile("topics.proto", fileDescriptor_72af2b0eeeefc419) }

var fileDescriptor_72af2b0eeeefc419 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0x69, 0x75, 0x1b, 0xbd, 0x6c, 0x13, 0xc3, 0x90, 0x30, 0x44, 0xa4, 0x4f, 0x3e, 0x8d,
	0x31, 0xfd, 0x0d, 0xc3, 0x3d, 0x28, 0x18, 0xfd, 0x03, 0xdd, 0x7a, 0x27, 0x81, 0xa6, 0xa9, 0x49,
	0x26, 0xfe, 0x7c, 0xc9, 0x6d, 0x42, 0x5b, 0x85, 0x3d, 0x7e, 0xe7, 0xc0, 0x39, 0x87, 0x7b, 0x61,
	0xea, 0x74, 0x23, 0x0f, 0x76, 0xd5, 0x18, 0xed, 0x34, 0x1b, 0xb7, 0x94, 0x2b, 0x98, 0xbc, 0xa0,
	0xb5, 0xc5, 0x27, 0xb2, 0x39, 0xa4, 0xb2, 0xe4, 0xc9, 0x7d, 0xf2, 0x90, 0x89, 0x54, 0x96, 0xec,
	0x0e, 0xc0, 0x62, 0x5d, 0xa2, 0x79, 0x2d, 0x14, 0xf2, 0x94, 0xf4, 0x9e, 0xc2, 0x38, 0x4c, 0x1a,
	0xa3, 0x8f, 0xb2, 0x42, 0x7e, 0x41, 0x66, 0x44, 0xef, 0xa8, 0x36, 0x94, 0x5f, 0xb6, 0x4e, 0xc0,
	0xfc, 0x0d, 0xb2, 0x0f, 0x5f, 0xbc, 0xab, 0x8f, 0xfa, 0x5f, 0xe1, 0x02, 0x46, 0x4e, 0xba, 0x2a,
	0x76, 0xb5, 0xe0, 0x67, 0xd4, 0x27, 0x25, 0xb0, 0xa9, 0x24, 0x5a, 0x6a, 0x9a, 0x89, 0x9e, 0x92,
	0x6f, 0x61, 0x4a, 0x91, 0x02, 0xbf, 0x4e, 0x68, 0x9d, 0x2f, 0xdf, 0xeb, 0xc2, 0x94, 0xbb, 0x18,
	0x1d, 0x91, 0xdd, 0x42, 0x66, 0x5d, 0x61, 0xdc, 0xd6, 0x68, 0x45, 0x1d, 0x33, 0xd1, 0x09, 0xf9,
	0x33, 0xcc, 0xc3, 0x25, 0x7a, 0x49, 0x74, 0xa5, 0x2e, 0x29, 0xe0, 0xf9, 0xa4, 0xcd, 0x4f, 0x58,
	0xf4, 0x8e, 0xe6, 0x5b, 0x1e, 0x90, 0x6d, 0x60, 0x44, 0xcc, 0x16, 0xab, 0xf0, 0x83, 0xfe, 0xe0,
	0xe5, 0xf5, 0x40, 0xf5, 0x97, 0x59, 0x27, 0xec, 0xa9, 0xfb, 0xcb, 0x4d, 0xf4, 0x87, 0xf3, 0x96,
	0x57, 0x7f, 0xf4, 0x75, 0xb2, 0x1f, 0xd3, 0x73, 0x1f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xee,
	0xba, 0x2e, 0xe2, 0xec, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TopicServiceClient is the client API for TopicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TopicServiceClient interface {
	Topic(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (TopicService_TopicClient, error)
	Message(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (TopicService_MessageClient, error)
}

type topicServiceClient struct {
	cc *grpc.ClientConn
}

func NewTopicServiceClient(cc *grpc.ClientConn) TopicServiceClient {
	return &topicServiceClient{cc}
}

func (c *topicServiceClient) Topic(ctx context.Context, in *TopicRequest, opts ...grpc.CallOption) (TopicService_TopicClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TopicService_serviceDesc.Streams[0], "/topics.TopicService/Topic", opts...)
	if err != nil {
		return nil, err
	}
	x := &topicServiceTopicClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TopicService_TopicClient interface {
	Recv() (*TopicInfo, error)
	grpc.ClientStream
}

type topicServiceTopicClient struct {
	grpc.ClientStream
}

func (x *topicServiceTopicClient) Recv() (*TopicInfo, error) {
	m := new(TopicInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *topicServiceClient) Message(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (TopicService_MessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TopicService_serviceDesc.Streams[1], "/topics.TopicService/Message", opts...)
	if err != nil {
		return nil, err
	}
	x := &topicServiceMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TopicService_MessageClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type topicServiceMessageClient struct {
	grpc.ClientStream
}

func (x *topicServiceMessageClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TopicServiceServer is the server API for TopicService service.
type TopicServiceServer interface {
	Topic(*TopicRequest, TopicService_TopicServer) error
	Message(*MessageRequest, TopicService_MessageServer) error
}

// UnimplementedTopicServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTopicServiceServer struct {
}

func (*UnimplementedTopicServiceServer) Topic(req *TopicRequest, srv TopicService_TopicServer) error {
	return status.Errorf(codes.Unimplemented, "method Topic not implemented")
}
func (*UnimplementedTopicServiceServer) Message(req *MessageRequest, srv TopicService_MessageServer) error {
	return status.Errorf(codes.Unimplemented, "method Message not implemented")
}

func RegisterTopicServiceServer(s *grpc.Server, srv TopicServiceServer) {
	s.RegisterService(&_TopicService_serviceDesc, srv)
}

func _TopicService_Topic_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TopicRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TopicServiceServer).Topic(m, &topicServiceTopicServer{stream})
}

type TopicService_TopicServer interface {
	Send(*TopicInfo) error
	grpc.ServerStream
}

type topicServiceTopicServer struct {
	grpc.ServerStream
}

func (x *topicServiceTopicServer) Send(m *TopicInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _TopicService_Message_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MessageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TopicServiceServer).Message(m, &topicServiceMessageServer{stream})
}

type TopicService_MessageServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type topicServiceMessageServer struct {
	grpc.ServerStream
}

func (x *topicServiceMessageServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

var _TopicService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "topics.TopicService",
	HandlerType: (*TopicServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Topic",
			Handler:       _TopicService_Topic_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Message",
			Handler:       _TopicService_Message_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "topics.proto",
}
