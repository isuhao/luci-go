// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/scheduler/appengine/internal/tq.proto

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	go.chromium.org/luci/scheduler/appengine/internal/tq.proto
	go.chromium.org/luci/scheduler/appengine/internal/triggers.proto

It has these top-level messages:
	ReadProjectConfigTask
	LaunchInvocationTask
	LaunchInvocationsBatchTask
	TriageJobStateTask
	InvocationFinishedTask
	FanOutTriggersTask
	EnqueueTriggersTask
	Trigger
	NoopTriggerData
	GitilesTriggerData
	TriggerList
*/
package internal

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

// ReadProjectConfigTask is used to import jobs of some project.
//
// Queue: "read-project-config".
type ReadProjectConfigTask struct {
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
}

func (m *ReadProjectConfigTask) Reset()                    { *m = ReadProjectConfigTask{} }
func (m *ReadProjectConfigTask) String() string            { return proto.CompactTextString(m) }
func (*ReadProjectConfigTask) ProtoMessage()               {}
func (*ReadProjectConfigTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReadProjectConfigTask) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

// LaunchInvocationTask is used to start running (or retry a lunch of) a single
// invocation.
//
// It is enqueued non-transactionally, but with the deduplication key.
//
// Queue: "launches".
type LaunchInvocationTask struct {
	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	InvId int64  `protobuf:"varint,2,opt,name=inv_id,json=invId" json:"inv_id,omitempty"`
}

func (m *LaunchInvocationTask) Reset()                    { *m = LaunchInvocationTask{} }
func (m *LaunchInvocationTask) String() string            { return proto.CompactTextString(m) }
func (*LaunchInvocationTask) ProtoMessage()               {}
func (*LaunchInvocationTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LaunchInvocationTask) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *LaunchInvocationTask) GetInvId() int64 {
	if m != nil {
		return m.InvId
	}
	return 0
}

// LaunchInvocationsBatchTask is used to kick off several invocations at once.
//
// It is enqueued transactionally. It fans out into many LaunchInvocationTask.
//
// Queue: "batches".
type LaunchInvocationsBatchTask struct {
	Tasks []*LaunchInvocationTask `protobuf:"bytes,1,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *LaunchInvocationsBatchTask) Reset()                    { *m = LaunchInvocationsBatchTask{} }
func (m *LaunchInvocationsBatchTask) String() string            { return proto.CompactTextString(m) }
func (*LaunchInvocationsBatchTask) ProtoMessage()               {}
func (*LaunchInvocationsBatchTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LaunchInvocationsBatchTask) GetTasks() []*LaunchInvocationTask {
	if m != nil {
		return m.Tasks
	}
	return nil
}

// TriageJobStateTask looks at the state of the job and decided what to do next.
//
// Enqueued non-transactionally. It is throttled to run approximately once per
// second. It looks at pending triggers and recently finished invocations and
// launches new invocations (or schedules timers to do it later).
//
// Queue: "triages".
type TriageJobStateTask struct {
	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}

func (m *TriageJobStateTask) Reset()                    { *m = TriageJobStateTask{} }
func (m *TriageJobStateTask) String() string            { return proto.CompactTextString(m) }
func (*TriageJobStateTask) ProtoMessage()               {}
func (*TriageJobStateTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TriageJobStateTask) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

// InvocationFinishedTask is emitted by the invocation when it finishes.
//
// It is enqueued transactionally.
//
// Queue: "completions".
type InvocationFinishedTask struct {
	JobId    string              `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	InvId    int64               `protobuf:"varint,2,opt,name=inv_id,json=invId" json:"inv_id,omitempty"`
	Triggers *FanOutTriggersTask `protobuf:"bytes,3,opt,name=triggers" json:"triggers,omitempty"`
}

func (m *InvocationFinishedTask) Reset()                    { *m = InvocationFinishedTask{} }
func (m *InvocationFinishedTask) String() string            { return proto.CompactTextString(m) }
func (*InvocationFinishedTask) ProtoMessage()               {}
func (*InvocationFinishedTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *InvocationFinishedTask) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *InvocationFinishedTask) GetInvId() int64 {
	if m != nil {
		return m.InvId
	}
	return 0
}

func (m *InvocationFinishedTask) GetTriggers() *FanOutTriggersTask {
	if m != nil {
		return m.Triggers
	}
	return nil
}

// FanOutTriggersTask is a batch task that emits a bunch of triggers.
//
// It is enqueued transactionally. It fans out into many EnqueueTriggersTask,
// one per job ID.
//
// Queue: "triggers".
type FanOutTriggersTask struct {
	JobIds   []string   `protobuf:"bytes,1,rep,name=job_ids,json=jobIds" json:"job_ids,omitempty"`
	Triggers []*Trigger `protobuf:"bytes,2,rep,name=triggers" json:"triggers,omitempty"`
}

func (m *FanOutTriggersTask) Reset()                    { *m = FanOutTriggersTask{} }
func (m *FanOutTriggersTask) String() string            { return proto.CompactTextString(m) }
func (*FanOutTriggersTask) ProtoMessage()               {}
func (*FanOutTriggersTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *FanOutTriggersTask) GetJobIds() []string {
	if m != nil {
		return m.JobIds
	}
	return nil
}

func (m *FanOutTriggersTask) GetTriggers() []*Trigger {
	if m != nil {
		return m.Triggers
	}
	return nil
}

// EnqueueTriggersTask adds given triggers to a job's pending triggers set.
//
// Enqueued non-transactionally.
//
// Queue: "triggers".
type EnqueueTriggersTask struct {
	JobId    string     `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	Triggers []*Trigger `protobuf:"bytes,2,rep,name=triggers" json:"triggers,omitempty"`
}

func (m *EnqueueTriggersTask) Reset()                    { *m = EnqueueTriggersTask{} }
func (m *EnqueueTriggersTask) String() string            { return proto.CompactTextString(m) }
func (*EnqueueTriggersTask) ProtoMessage()               {}
func (*EnqueueTriggersTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *EnqueueTriggersTask) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *EnqueueTriggersTask) GetTriggers() []*Trigger {
	if m != nil {
		return m.Triggers
	}
	return nil
}

func init() {
	proto.RegisterType((*ReadProjectConfigTask)(nil), "internal.tq.ReadProjectConfigTask")
	proto.RegisterType((*LaunchInvocationTask)(nil), "internal.tq.LaunchInvocationTask")
	proto.RegisterType((*LaunchInvocationsBatchTask)(nil), "internal.tq.LaunchInvocationsBatchTask")
	proto.RegisterType((*TriageJobStateTask)(nil), "internal.tq.TriageJobStateTask")
	proto.RegisterType((*InvocationFinishedTask)(nil), "internal.tq.InvocationFinishedTask")
	proto.RegisterType((*FanOutTriggersTask)(nil), "internal.tq.FanOutTriggersTask")
	proto.RegisterType((*EnqueueTriggersTask)(nil), "internal.tq.EnqueueTriggersTask")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/scheduler/appengine/internal/tq.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x5b, 0x4b, 0xeb, 0x40,
	0x14, 0x85, 0x49, 0x4b, 0x7b, 0x4e, 0x77, 0xdf, 0x72, 0x4e, 0xb5, 0x14, 0xc4, 0x98, 0xa7, 0x80,
	0x90, 0x40, 0x85, 0x0a, 0xfa, 0x22, 0x5e, 0x0a, 0x11, 0x41, 0x89, 0xf5, 0x59, 0x26, 0x99, 0x6d,
	0x32, 0xbd, 0xcc, 0xa4, 0x73, 0xe9, 0x1f, 0xf0, 0x8f, 0x4b, 0x93, 0xd6, 0xa6, 0xde, 0xc0, 0xbe,
	0xae, 0xec, 0xbd, 0xbe, 0xb5, 0x57, 0x06, 0xce, 0x52, 0xe1, 0x27, 0x99, 0x14, 0x33, 0x66, 0x66,
	0xbe, 0x90, 0x69, 0x30, 0x35, 0x09, 0x0b, 0x54, 0x92, 0x21, 0x35, 0x53, 0x94, 0x01, 0xc9, 0x73,
	0xe4, 0x29, 0xe3, 0x18, 0x30, 0xae, 0x51, 0x72, 0x32, 0x0d, 0xf4, 0xdc, 0xcf, 0xa5, 0xd0, 0xc2,
	0x6e, 0xaf, 0x25, 0x5f, 0xcf, 0x7b, 0x17, 0x3b, 0x18, 0x49, 0x96, 0xa6, 0x28, 0x55, 0x69, 0xe7,
	0x0e, 0xa0, 0x13, 0x21, 0xa1, 0x0f, 0x52, 0x8c, 0x31, 0xd1, 0x57, 0x82, 0xbf, 0xb0, 0x74, 0x44,
	0xd4, 0xc4, 0x3e, 0x00, 0xc8, 0x4b, 0xf1, 0x99, 0xd1, 0xae, 0xe5, 0x58, 0x5e, 0x2b, 0x6a, 0xad,
	0x94, 0x90, 0xba, 0xd7, 0xf0, 0xff, 0x8e, 0x18, 0x9e, 0x64, 0x21, 0x5f, 0x88, 0x84, 0x68, 0x26,
	0x78, 0xb1, 0xd6, 0x81, 0xe6, 0x58, 0xc4, 0x9b, 0x95, 0xc6, 0x58, 0xc4, 0x21, 0x5d, 0xca, 0x8c,
	0x2f, 0x96, 0x72, 0xcd, 0xb1, 0xbc, 0x7a, 0xd4, 0x60, 0x7c, 0x11, 0x52, 0xf7, 0x09, 0x7a, 0x1f,
	0x5d, 0xd4, 0x25, 0xd1, 0x49, 0x56, 0x78, 0x9d, 0x42, 0x43, 0x13, 0x35, 0x51, 0x5d, 0xcb, 0xa9,
	0x7b, 0xed, 0xfe, 0x91, 0x5f, 0x39, 0xdd, 0xff, 0x8a, 0x1e, 0x95, 0xf3, 0xee, 0x31, 0xd8, 0x23,
	0xc9, 0x48, 0x8a, 0xb7, 0x22, 0x7e, 0xd4, 0x44, 0xe3, 0x0f, 0xd1, 0xdc, 0x57, 0x0b, 0xf6, 0x36,
	0x36, 0x43, 0xc6, 0x99, 0xca, 0x90, 0xfe, 0xfe, 0x18, 0xfb, 0x1c, 0xfe, 0xae, 0xcb, 0xed, 0xd6,
	0x1d, 0xcb, 0x6b, 0xf7, 0x0f, 0xb7, 0x12, 0x0f, 0x09, 0xbf, 0x37, 0x7a, 0xb4, 0x1a, 0x29, 0xf2,
	0xbe, 0x2f, 0xb8, 0x08, 0xf6, 0xe7, 0xef, 0xf6, 0x3e, 0xfc, 0x29, 0x03, 0x94, 0x1d, 0xb4, 0xa2,
	0x66, 0x91, 0x40, 0xd9, 0x83, 0x0a, 0xab, 0x56, 0xb4, 0xd3, 0xab, 0xb0, 0xd6, 0xbf, 0x78, 0xe5,
	0x55, 0xc1, 0x50, 0xf8, 0x77, 0xc3, 0xe7, 0x06, 0x0d, 0x6e, 0x71, 0xbe, 0x39, 0x74, 0x47, 0x4a,
	0xdc, 0x2c, 0xde, 0xd6, 0xc9, 0x5b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45, 0xfa, 0x0e, 0x5a, 0xe8,
	0x02, 0x00, 0x00,
}
