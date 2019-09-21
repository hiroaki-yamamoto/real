// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages/messages.proto

package messages

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
	return fileDescriptor_83994550f81e9f35, []int{0}
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
	return fileDescriptor_83994550f81e9f35, []int{1}
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
	proto.RegisterType((*Message)(nil), "messages.Message")
	proto.RegisterType((*MessageRequest)(nil), "messages.MessageRequest")
}

func init() { proto.RegisterFile("messages/messages.proto", fileDescriptor_83994550f81e9f35) }

var fileDescriptor_83994550f81e9f35 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x60,
	0x7c, 0xa5, 0x5c, 0x2e, 0x76, 0x5f, 0x08, 0x5b, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51,
	0x81, 0x51, 0x83, 0x33, 0x88, 0x29, 0x33, 0x45, 0x48, 0x8e, 0x8b, 0xab, 0x38, 0x35, 0x2f, 0x25,
	0xb5, 0xc8, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x09, 0x2c, 0x8e, 0x24, 0x22, 0x24, 0xc1, 0xc5, 0x5e,
	0x50, 0x94, 0x9f, 0x96, 0x99, 0x93, 0x2a, 0xc1, 0x0c, 0x96, 0x84, 0x71, 0x41, 0x32, 0x50, 0x0b,
	0x24, 0x58, 0x20, 0x32, 0x50, 0xae, 0x92, 0x07, 0x17, 0x1f, 0xd4, 0xba, 0xa0, 0xd4, 0xc2, 0xd2,
	0xd4, 0xe2, 0x12, 0x90, 0xda, 0x92, 0xfc, 0x82, 0xcc, 0x64, 0x4f, 0x98, 0xd5, 0x30, 0xae, 0x90,
	0x0c, 0x17, 0x67, 0x71, 0x49, 0x62, 0x51, 0x89, 0x5b, 0x51, 0x7e, 0x2e, 0xd8, 0x7a, 0xde, 0x20,
	0x84, 0x40, 0x12, 0x1b, 0xd8, 0x27, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x37, 0xe6,
	0x25, 0xe4, 0x00, 0x00, 0x00,
}
