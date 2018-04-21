// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mailer.proto

/*
Package mailer is a generated protocol buffer package.

It is generated from these files:
	mailer.proto

It has these top-level messages:
	User
	Mail
	SendMailRequest
	SendMailResponse
	ConsumeQueueRequest
	ConsumeQueueResponse
*/
package mailer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Uuid     string `protobuf:"bytes,1,opt,name=Uuid" json:"Uuid,omitempty"`
	Address  string `protobuf:"bytes,2,opt,name=Address" json:"Address,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Language string `protobuf:"bytes,4,opt,name=Language" json:"Language,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

type Mail struct {
	From            *User    `protobuf:"bytes,1,opt,name=From" json:"From,omitempty"`
	To              []*User  `protobuf:"bytes,3,rep,name=To" json:"To,omitempty"`
	Cc              []*User  `protobuf:"bytes,4,rep,name=Cc" json:"Cc,omitempty"`
	DateSent        int64    `protobuf:"varint,5,opt,name=DateSent" json:"DateSent,omitempty"`
	Subject         string   `protobuf:"bytes,6,opt,name=Subject" json:"Subject,omitempty"`
	ContentPlain    string   `protobuf:"bytes,7,opt,name=ContentPlain" json:"ContentPlain,omitempty"`
	ContentHtml     string   `protobuf:"bytes,8,opt,name=ContentHtml" json:"ContentHtml,omitempty"`
	ContentMarkdown string   `protobuf:"bytes,9,opt,name=ContentMarkdown" json:"ContentMarkdown,omitempty"`
	Attachments     []string `protobuf:"bytes,10,rep,name=Attachments" json:"Attachments,omitempty"`
	// Could be used for Re: ... conversations
	ThreadUuid   string            `protobuf:"bytes,11,opt,name=ThreadUuid" json:"ThreadUuid,omitempty"`
	ThreadIndex  string            `protobuf:"bytes,12,opt,name=ThreadIndex" json:"ThreadIndex,omitempty"`
	TemplateId   string            `protobuf:"bytes,13,opt,name=TemplateId" json:"TemplateId,omitempty"`
	TemplateData map[string]string `protobuf:"bytes,14,rep,name=TemplateData" json:"TemplateData,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Retries      int32             `protobuf:"varint,15,opt,name=Retries" json:"Retries,omitempty"`
	SendErrors   []string          `protobuf:"bytes,16,rep,name=sendErrors" json:"sendErrors,omitempty"`
}

func (m *Mail) Reset()                    { *m = Mail{} }
func (m *Mail) String() string            { return proto.CompactTextString(m) }
func (*Mail) ProtoMessage()               {}
func (*Mail) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Mail) GetFrom() *User {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Mail) GetTo() []*User {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Mail) GetCc() []*User {
	if m != nil {
		return m.Cc
	}
	return nil
}

func (m *Mail) GetDateSent() int64 {
	if m != nil {
		return m.DateSent
	}
	return 0
}

func (m *Mail) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Mail) GetContentPlain() string {
	if m != nil {
		return m.ContentPlain
	}
	return ""
}

func (m *Mail) GetContentHtml() string {
	if m != nil {
		return m.ContentHtml
	}
	return ""
}

func (m *Mail) GetContentMarkdown() string {
	if m != nil {
		return m.ContentMarkdown
	}
	return ""
}

func (m *Mail) GetAttachments() []string {
	if m != nil {
		return m.Attachments
	}
	return nil
}

func (m *Mail) GetThreadUuid() string {
	if m != nil {
		return m.ThreadUuid
	}
	return ""
}

func (m *Mail) GetThreadIndex() string {
	if m != nil {
		return m.ThreadIndex
	}
	return ""
}

func (m *Mail) GetTemplateId() string {
	if m != nil {
		return m.TemplateId
	}
	return ""
}

func (m *Mail) GetTemplateData() map[string]string {
	if m != nil {
		return m.TemplateData
	}
	return nil
}

func (m *Mail) GetRetries() int32 {
	if m != nil {
		return m.Retries
	}
	return 0
}

func (m *Mail) GetSendErrors() []string {
	if m != nil {
		return m.SendErrors
	}
	return nil
}

type SendMailRequest struct {
	Mail    *Mail `protobuf:"bytes,1,opt,name=Mail" json:"Mail,omitempty"`
	InQueue bool  `protobuf:"varint,2,opt,name=InQueue" json:"InQueue,omitempty"`
}

func (m *SendMailRequest) Reset()                    { *m = SendMailRequest{} }
func (m *SendMailRequest) String() string            { return proto.CompactTextString(m) }
func (*SendMailRequest) ProtoMessage()               {}
func (*SendMailRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SendMailRequest) GetMail() *Mail {
	if m != nil {
		return m.Mail
	}
	return nil
}

func (m *SendMailRequest) GetInQueue() bool {
	if m != nil {
		return m.InQueue
	}
	return false
}

type SendMailResponse struct {
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
}

func (m *SendMailResponse) Reset()                    { *m = SendMailResponse{} }
func (m *SendMailResponse) String() string            { return proto.CompactTextString(m) }
func (*SendMailResponse) ProtoMessage()               {}
func (*SendMailResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SendMailResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ConsumeQueueRequest struct {
	MaxEmails int64 `protobuf:"varint,1,opt,name=MaxEmails" json:"MaxEmails,omitempty"`
}

func (m *ConsumeQueueRequest) Reset()                    { *m = ConsumeQueueRequest{} }
func (m *ConsumeQueueRequest) String() string            { return proto.CompactTextString(m) }
func (*ConsumeQueueRequest) ProtoMessage()               {}
func (*ConsumeQueueRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ConsumeQueueRequest) GetMaxEmails() int64 {
	if m != nil {
		return m.MaxEmails
	}
	return 0
}

type ConsumeQueueResponse struct {
	Message    string `protobuf:"bytes,1,opt,name=Message" json:"Message,omitempty"`
	EmailsSent int64  `protobuf:"varint,2,opt,name=EmailsSent" json:"EmailsSent,omitempty"`
}

func (m *ConsumeQueueResponse) Reset()                    { *m = ConsumeQueueResponse{} }
func (m *ConsumeQueueResponse) String() string            { return proto.CompactTextString(m) }
func (*ConsumeQueueResponse) ProtoMessage()               {}
func (*ConsumeQueueResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ConsumeQueueResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ConsumeQueueResponse) GetEmailsSent() int64 {
	if m != nil {
		return m.EmailsSent
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "mailer.User")
	proto.RegisterType((*Mail)(nil), "mailer.Mail")
	proto.RegisterType((*SendMailRequest)(nil), "mailer.SendMailRequest")
	proto.RegisterType((*SendMailResponse)(nil), "mailer.SendMailResponse")
	proto.RegisterType((*ConsumeQueueRequest)(nil), "mailer.ConsumeQueueRequest")
	proto.RegisterType((*ConsumeQueueResponse)(nil), "mailer.ConsumeQueueResponse")
}

func init() { proto.RegisterFile("mailer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 557 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb1, 0xdb, 0xa6, 0x93, 0x94, 0x84, 0xa5, 0x12, 0xab, 0x10, 0x55, 0x96, 0x4f, 0x39,
	0xa0, 0x1c, 0xda, 0x0b, 0xe2, 0x52, 0x95, 0x34, 0x88, 0x08, 0x8c, 0x8a, 0x93, 0x3e, 0xc0, 0xd6,
	0x1e, 0xb5, 0xa6, 0xf6, 0x3a, 0xec, 0xae, 0x4b, 0xfb, 0x30, 0x3c, 0x12, 0xef, 0x84, 0x76, 0xd7,
	0x4e, 0x9c, 0x9f, 0xdb, 0x7e, 0xdf, 0xcc, 0x7c, 0x3b, 0x33, 0xfe, 0xd6, 0xd0, 0xcd, 0x59, 0x9a,
	0xa1, 0x18, 0x2f, 0x45, 0xa1, 0x0a, 0x72, 0x68, 0x51, 0x90, 0x80, 0x77, 0x2b, 0x51, 0x10, 0x02,
	0xde, 0x6d, 0x99, 0x26, 0xd4, 0xf1, 0x9d, 0xd1, 0x71, 0x64, 0xce, 0x84, 0xc2, 0xd1, 0x55, 0x92,
	0x08, 0x94, 0x92, 0xb6, 0x0c, 0x5d, 0x43, 0x9d, 0xfd, 0x83, 0xe5, 0x48, 0x5d, 0x9b, 0xad, 0xcf,
	0x64, 0x00, 0xed, 0xef, 0x8c, 0xdf, 0x97, 0xec, 0x1e, 0xa9, 0x67, 0xf8, 0x15, 0x0e, 0xfe, 0x79,
	0xe0, 0x85, 0x2c, 0xcd, 0x88, 0x0f, 0xde, 0x17, 0x51, 0xe4, 0xe6, 0x9a, 0xce, 0x79, 0x77, 0x5c,
	0xf5, 0xa4, 0x5b, 0x88, 0x4c, 0x84, 0x0c, 0xa1, 0xb5, 0x28, 0xa8, 0xeb, 0xbb, 0x3b, 0xf1, 0xd6,
	0xa2, 0xd0, 0xd1, 0x49, 0x4c, 0xbd, 0x7d, 0xd1, 0x49, 0xac, 0x5b, 0xb8, 0x66, 0x0a, 0xe7, 0xc8,
	0x15, 0x3d, 0xf0, 0x9d, 0x91, 0x1b, 0xad, 0xb0, 0x1e, 0x66, 0x5e, 0xde, 0xfd, 0xc2, 0x58, 0xd1,
	0x43, 0x3b, 0x4c, 0x05, 0x49, 0x00, 0xdd, 0x49, 0xc1, 0x15, 0x72, 0x75, 0x93, 0xb1, 0x94, 0xd3,
	0x23, 0x13, 0xde, 0xe0, 0x88, 0x0f, 0x9d, 0x0a, 0x7f, 0x55, 0x79, 0x46, 0xdb, 0x26, 0xa5, 0x49,
	0x91, 0x11, 0xf4, 0x2a, 0x18, 0x32, 0xf1, 0x98, 0x14, 0x7f, 0x38, 0x3d, 0x36, 0x59, 0xdb, 0xb4,
	0xd6, 0xba, 0x52, 0x8a, 0xc5, 0x0f, 0x39, 0x72, 0x25, 0x29, 0xf8, 0xae, 0xd6, 0x6a, 0x50, 0xe4,
	0x0c, 0x60, 0xf1, 0x20, 0x90, 0x25, 0xe6, 0x93, 0x74, 0x8c, 0x4c, 0x83, 0xd1, 0x0a, 0x16, 0xcd,
	0x78, 0x82, 0xcf, 0xb4, 0x6b, 0xbb, 0x69, 0x50, 0x46, 0x01, 0xf3, 0x65, 0xc6, 0x14, 0xce, 0x12,
	0x7a, 0x52, 0x29, 0xac, 0x18, 0xf2, 0x19, 0xba, 0x35, 0xba, 0x66, 0x8a, 0xd1, 0xd7, 0x66, 0xa3,
	0x67, 0xf5, 0x46, 0xf5, 0xb7, 0x1a, 0x37, 0x13, 0xa6, 0x5c, 0x89, 0x97, 0x68, 0xa3, 0x46, 0x6f,
	0x34, 0x42, 0x25, 0x52, 0x94, 0xb4, 0xe7, 0x3b, 0xa3, 0x83, 0xa8, 0x86, 0xfa, 0x76, 0x89, 0x3c,
	0x99, 0x0a, 0x51, 0x08, 0x49, 0xfb, 0x66, 0xc0, 0x06, 0x33, 0xb8, 0x84, 0x37, 0x3b, 0xe2, 0xa4,
	0x0f, 0xee, 0x23, 0xbe, 0x54, 0x06, 0xd4, 0x47, 0x72, 0x0a, 0x07, 0x4f, 0x2c, 0x2b, 0xb1, 0x72,
	0x9f, 0x05, 0x9f, 0x5a, 0x1f, 0x9d, 0x20, 0x84, 0xde, 0x1c, 0x79, 0xa2, 0xdb, 0x8c, 0xf0, 0x77,
	0x89, 0x52, 0x69, 0x67, 0x69, 0xb8, 0xed, 0x2c, 0x93, 0x62, 0xbd, 0x47, 0xe1, 0x68, 0xc6, 0x7f,
	0x96, 0x58, 0x09, 0xb6, 0xa3, 0x1a, 0x06, 0x1f, 0xa0, 0xbf, 0x96, 0x93, 0xcb, 0x82, 0x4b, 0xb4,
	0x7e, 0x89, 0x63, 0x6d, 0x7e, 0xc7, 0x66, 0x57, 0x30, 0xb8, 0x80, 0xb7, 0x93, 0x82, 0xcb, 0x32,
	0x47, 0x53, 0x5d, 0x37, 0x30, 0x84, 0xe3, 0x90, 0x3d, 0x4f, 0xf5, 0xbd, 0xb6, 0xc4, 0x8d, 0xd6,
	0x44, 0x70, 0x03, 0xa7, 0x9b, 0x45, 0xeb, 0x6b, 0x42, 0x94, 0x52, 0x3f, 0x1a, 0x3b, 0x79, 0x0d,
	0xf5, 0x12, 0x6d, 0xad, 0xb1, 0x73, 0xcb, 0x08, 0x36, 0x98, 0xf3, 0xbf, 0x0e, 0x9c, 0x84, 0x66,
	0xc8, 0x39, 0x8a, 0xa7, 0x34, 0x46, 0x72, 0x09, 0xed, 0x7a, 0x0c, 0xf2, 0xae, 0x5e, 0xc0, 0xd6,
	0x9e, 0x06, 0x74, 0x37, 0x60, 0x5b, 0x09, 0x5e, 0x91, 0x6f, 0xe6, 0x25, 0xac, 0x9a, 0x24, 0xef,
	0xeb, 0xdc, 0x3d, 0xf3, 0x0e, 0x86, 0xfb, 0x83, 0xb5, 0xd8, 0xdd, 0xa1, 0xf9, 0xd1, 0x5c, 0xfc,
	0x0f, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xaf, 0xd3, 0xcd, 0x78, 0x04, 0x00, 0x00,
}
