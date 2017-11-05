// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/config/notify.proto

/*
Package config is a generated protocol buffer package.

It is generated from these files:
	api/config/notify.proto
	api/config/settings.proto

It has these top-level messages:
	ProjectConfig
	Notifier
	Notification
	Builder
	Settings
*/
package config

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

// ProjectConfig is a luci-notify configuration for a particular project.
type ProjectConfig struct {
	// Notifiers is a list of Notifiers which watch builders and send
	// notifications for this project.
	Notifiers []*Notifier `protobuf:"bytes,1,rep,name=notifiers" json:"notifiers,omitempty"`
}

func (m *ProjectConfig) Reset()                    { *m = ProjectConfig{} }
func (m *ProjectConfig) String() string            { return proto.CompactTextString(m) }
func (*ProjectConfig) ProtoMessage()               {}
func (*ProjectConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProjectConfig) GetNotifiers() []*Notifier {
	if m != nil {
		return m.Notifiers
	}
	return nil
}

// Notifier contains a set of notification configurations (which specify
// triggers to send notifications on) and a set of builders that will be
// watched for these triggers.
type Notifier struct {
	// Name is an identifier for the notifier which must be unique within a
	// project.
	//
	// Name must additionally match ^[a-z\-]+$, meaning it must only
	// use an alphabet of lowercase characters and hyphens.
	//
	// Required.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Notifications is a list of notification configurations.
	Notifications []*Notification `protobuf:"bytes,2,rep,name=notifications" json:"notifications,omitempty"`
	// Builders is a list of buildbucket builders this Notifier should watch.
	Builders []*Builder `protobuf:"bytes,3,rep,name=builders" json:"builders,omitempty"`
}

func (m *Notifier) Reset()                    { *m = Notifier{} }
func (m *Notifier) String() string            { return proto.CompactTextString(m) }
func (*Notifier) ProtoMessage()               {}
func (*Notifier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Notifier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Notifier) GetNotifications() []*Notification {
	if m != nil {
		return m.Notifications
	}
	return nil
}

func (m *Notifier) GetBuilders() []*Builder {
	if m != nil {
		return m.Builders
	}
	return nil
}

// Notification specifies the triggers to watch for and send
// notifications on. It also specifies email recipients.
type Notification struct {
	// OnSuccess specifies a trigger to notify recipients on each build success.
	//
	// Optional.
	OnSuccess bool `protobuf:"varint,1,opt,name=on_success,json=onSuccess" json:"on_success,omitempty"`
	// OnFailure specifies a trigger to notify recipients on each build failure.
	//
	// Optional.
	OnFailure bool `protobuf:"varint,2,opt,name=on_failure,json=onFailure" json:"on_failure,omitempty"`
	// OnChange specifies a trigger to notify recipients if the builder's
	// previous build had a different result than the most recent build.
	//
	// Optional.
	OnChange bool `protobuf:"varint,3,opt,name=on_change,json=onChange" json:"on_change,omitempty"`
	// Email is the set of email addresses to notify.
	//
	// Optional.
	Email *Notification_Email `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
}

func (m *Notification) Reset()                    { *m = Notification{} }
func (m *Notification) String() string            { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()               {}
func (*Notification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Notification) GetOnSuccess() bool {
	if m != nil {
		return m.OnSuccess
	}
	return false
}

func (m *Notification) GetOnFailure() bool {
	if m != nil {
		return m.OnFailure
	}
	return false
}

func (m *Notification) GetOnChange() bool {
	if m != nil {
		return m.OnChange
	}
	return false
}

func (m *Notification) GetEmail() *Notification_Email {
	if m != nil {
		return m.Email
	}
	return nil
}

// EmailConfig is a message representing a set of mail recipients (email
// addresses).
type Notification_Email struct {
	// Recipients is a list of email addresses to notify.
	Recipients []string `protobuf:"bytes,1,rep,name=recipients" json:"recipients,omitempty"`
}

func (m *Notification_Email) Reset()                    { *m = Notification_Email{} }
func (m *Notification_Email) String() string            { return proto.CompactTextString(m) }
func (*Notification_Email) ProtoMessage()               {}
func (*Notification_Email) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *Notification_Email) GetRecipients() []string {
	if m != nil {
		return m.Recipients
	}
	return nil
}

// Builder references a buildbucket builder in the current project.
type Builder struct {
	// Bucket is the buildbucket bucket that the builder is a part of.
	//
	// Required.
	Bucket string `protobuf:"bytes,1,opt,name=bucket" json:"bucket,omitempty"`
	// Name is the name of the buildbucket builder.
	//
	// Required.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Builder) Reset()                    { *m = Builder{} }
func (m *Builder) String() string            { return proto.CompactTextString(m) }
func (*Builder) ProtoMessage()               {}
func (*Builder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Builder) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *Builder) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ProjectConfig)(nil), "config.ProjectConfig")
	proto.RegisterType((*Notifier)(nil), "config.Notifier")
	proto.RegisterType((*Notification)(nil), "config.Notification")
	proto.RegisterType((*Notification_Email)(nil), "config.Notification.Email")
	proto.RegisterType((*Builder)(nil), "config.Builder")
}

func init() { proto.RegisterFile("api/config/notify.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xdf, 0x4a, 0xc3, 0x30,
	0x14, 0xc6, 0x49, 0xbb, 0xd5, 0xf6, 0xcc, 0xa1, 0x04, 0xd1, 0x30, 0x51, 0x4a, 0x6f, 0x2c, 0x08,
	0x9d, 0x4c, 0xbc, 0xf1, 0x46, 0x70, 0xe8, 0xa5, 0x48, 0x7c, 0x80, 0x91, 0xc6, 0x74, 0x46, 0xbb,
	0xa4, 0xb4, 0xe9, 0x85, 0x4f, 0xe0, 0x8b, 0xf9, 0x60, 0xb2, 0xa4, 0xdd, 0x1f, 0xf0, 0x2e, 0xe7,
	0xfb, 0x7e, 0xe7, 0x70, 0xce, 0x17, 0x38, 0x63, 0x95, 0x9c, 0x72, 0xad, 0x0a, 0xb9, 0x9c, 0x2a,
	0x6d, 0x64, 0xf1, 0x9d, 0x55, 0xb5, 0x36, 0x1a, 0x07, 0x4e, 0x4c, 0x1e, 0x60, 0xfc, 0x5a, 0xeb,
	0x4f, 0xc1, 0xcd, 0xdc, 0x0a, 0x38, 0x83, 0xc8, 0x82, 0x52, 0xd4, 0x0d, 0x41, 0xb1, 0x9f, 0x8e,
	0x66, 0xc7, 0x99, 0x83, 0xb3, 0x97, 0xce, 0xa0, 0x5b, 0x24, 0xf9, 0x41, 0x10, 0xf6, 0x3a, 0xc6,
	0x30, 0x50, 0x6c, 0x25, 0x08, 0x8a, 0x51, 0x1a, 0x51, 0xfb, 0xc6, 0xf7, 0x30, 0x76, 0x34, 0x67,
	0x46, 0x6a, 0xd5, 0x10, 0xcf, 0x0e, 0x3d, 0xd9, 0x1f, 0xea, 0x4c, 0xba, 0x8f, 0xe2, 0x6b, 0x08,
	0xf3, 0x56, 0x96, 0xef, 0xeb, 0x5d, 0x7c, 0xdb, 0x76, 0xd4, 0xb7, 0x3d, 0x3a, 0x9d, 0x6e, 0x80,
	0xe4, 0x17, 0xc1, 0xe1, 0xee, 0x30, 0x7c, 0x01, 0xa0, 0xd5, 0xa2, 0x69, 0x39, 0x17, 0x4d, 0x63,
	0x77, 0x0a, 0x69, 0xa4, 0xd5, 0x9b, 0x13, 0x3a, 0xbb, 0x60, 0xb2, 0x6c, 0x6b, 0x41, 0xbc, 0xde,
	0x7e, 0x76, 0x02, 0x3e, 0x87, 0x48, 0xab, 0x05, 0xff, 0x60, 0x6a, 0x29, 0x88, 0x6f, 0xdd, 0x50,
	0xab, 0xb9, 0xad, 0xf1, 0x0d, 0x0c, 0xc5, 0x8a, 0xc9, 0x92, 0x0c, 0x62, 0x94, 0x8e, 0x66, 0x93,
	0xff, 0x8e, 0xc9, 0x9e, 0xd6, 0x04, 0x75, 0xe0, 0xe4, 0x0a, 0x86, 0xb6, 0xc6, 0x97, 0x00, 0xb5,
	0xe0, 0xb2, 0x92, 0x42, 0x19, 0x97, 0x70, 0x44, 0x77, 0x94, 0xe4, 0x0e, 0x0e, 0xba, 0xdb, 0xf0,
	0x29, 0x04, 0x79, 0xcb, 0xbf, 0x84, 0xe9, 0x02, 0xed, 0xaa, 0x4d, 0xcc, 0xde, 0x36, 0xe6, 0x3c,
	0xb0, 0xff, 0x7a, 0xfb, 0x17, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xdc, 0xfb, 0xf9, 0xf2, 0x01, 0x00,
	0x00,
}
