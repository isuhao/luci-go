// Code generated by protoc-gen-go.
// source: cron.proto
// DO NOT EDIT!

/*
Package messages is a generated protocol buffer package.

It is generated from these files:
	cron.proto

It has these top-level messages:
	Job
	Task
	NoopTask
	UrlFetchTask
	SwarmingTask
	BuildbucketTask
	ProjectConfig
*/
package messages

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Job specifies a single cron job belonging to a project.
type Job struct {
	// Id is a name of the job (unique for the project).
	Id *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Schedule in regular cron expression format.
	Schedule *string `protobuf:"bytes,2,opt,name=schedule" json:"schedule,omitempty"`
	// Disables is true to disable this job.
	Disabled *bool `protobuf:"varint,3,opt,name=disabled" json:"disabled,omitempty"`
	// Task defines what exactly to execute.
	Task             *Task  `protobuf:"bytes,4,opt,name=task" json:"task,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Job) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Job) GetSchedule() string {
	if m != nil && m.Schedule != nil {
		return *m.Schedule
	}
	return ""
}

func (m *Job) GetDisabled() bool {
	if m != nil && m.Disabled != nil {
		return *m.Disabled
	}
	return false
}

func (m *Job) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

// Task defines what exactly to do. One and only one field must be set.
type Task struct {
	// Noop is used for testing. It is "do nothing" task.
	Noop *NoopTask `protobuf:"bytes,1,opt,name=noop" json:"noop,omitempty"`
	// UrlFetch can be used to make a simple HTTP call.
	UrlFetch *UrlFetchTask `protobuf:"bytes,2,opt,name=url_fetch" json:"url_fetch,omitempty"`
	// SwarmingTask can be used to schedule swarming job.
	SwarmingTask *SwarmingTask `protobuf:"bytes,3,opt,name=swarming_task" json:"swarming_task,omitempty"`
	// BuildbucketTask can be used to schedule buildbucket job.
	BuildbucketTask  *BuildbucketTask `protobuf:"bytes,4,opt,name=buildbucket_task" json:"buildbucket_task,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Task) GetNoop() *NoopTask {
	if m != nil {
		return m.Noop
	}
	return nil
}

func (m *Task) GetUrlFetch() *UrlFetchTask {
	if m != nil {
		return m.UrlFetch
	}
	return nil
}

func (m *Task) GetSwarmingTask() *SwarmingTask {
	if m != nil {
		return m.SwarmingTask
	}
	return nil
}

func (m *Task) GetBuildbucketTask() *BuildbucketTask {
	if m != nil {
		return m.BuildbucketTask
	}
	return nil
}

// NoopTask is used for testing. It is "do nothing" task.
type NoopTask struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *NoopTask) Reset()                    { *m = NoopTask{} }
func (m *NoopTask) String() string            { return proto.CompactTextString(m) }
func (*NoopTask) ProtoMessage()               {}
func (*NoopTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// UrlFetchTask specifies parameters for simple HTTP call.
type UrlFetchTask struct {
	// Method is HTTP method to use, such as "GET" or "POST".
	Method *string `protobuf:"bytes,1,opt,name=method,def=GET" json:"method,omitempty"`
	// Url to send the request to.
	Url *string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	// Timeout is how long to wait for request to complete.
	TimeoutSec       *int32 `protobuf:"varint,3,opt,name=timeout_sec,def=60" json:"timeout_sec,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *UrlFetchTask) Reset()                    { *m = UrlFetchTask{} }
func (m *UrlFetchTask) String() string            { return proto.CompactTextString(m) }
func (*UrlFetchTask) ProtoMessage()               {}
func (*UrlFetchTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

const Default_UrlFetchTask_Method string = "GET"
const Default_UrlFetchTask_TimeoutSec int32 = 60

func (m *UrlFetchTask) GetMethod() string {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return Default_UrlFetchTask_Method
}

func (m *UrlFetchTask) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *UrlFetchTask) GetTimeoutSec() int32 {
	if m != nil && m.TimeoutSec != nil {
		return *m.TimeoutSec
	}
	return Default_UrlFetchTask_TimeoutSec
}

// SwarmingTask specifies parameters of Swarming-based cron job.
type SwarmingTask struct {
	// Server is URL of the swarming service to use.
	Server *string `protobuf:"bytes,1,opt,name=server" json:"server,omitempty"`
	// What to run. Only one of 'command' or 'isolated_ref' must be given.
	Command     []string                  `protobuf:"bytes,2,rep,name=command" json:"command,omitempty"`
	IsolatedRef *SwarmingTask_IsolatedRef `protobuf:"bytes,3,opt,name=isolated_ref" json:"isolated_ref,omitempty"`
	// Additional arguments to pass to isolated command.
	ExtraArgs []string `protobuf:"bytes,4,rep,name=extra_args" json:"extra_args,omitempty"`
	// List of "key=value" pairs with additional OS environment variables.
	Env []string `protobuf:"bytes,5,rep,name=env" json:"env,omitempty"`
	// Where to run it. List of "key:value" pairs.
	Dimensions []string `protobuf:"bytes,6,rep,name=dimensions" json:"dimensions,omitempty"`
	// Tags is a list of tags (as "key:value" pairs) to assign to the task.
	Tags []string `protobuf:"bytes,7,rep,name=tags" json:"tags,omitempty"`
	// Priority is task priority (or niceness, lower value - higher priority).
	Priority *int32 `protobuf:"varint,8,opt,name=priority,def=200" json:"priority,omitempty"`
	// Timeouts. All optional. Cron service will set reasonable default values.
	ExecutionTimeoutSecs *int32 `protobuf:"varint,9,opt,name=execution_timeout_secs" json:"execution_timeout_secs,omitempty"`
	GracePeriodSecs      *int32 `protobuf:"varint,10,opt,name=grace_period_secs,def=30" json:"grace_period_secs,omitempty"`
	IoTimeoutSecs        *int32 `protobuf:"varint,11,opt,name=io_timeout_secs" json:"io_timeout_secs,omitempty"`
	XXX_unrecognized     []byte `json:"-"`
}

func (m *SwarmingTask) Reset()                    { *m = SwarmingTask{} }
func (m *SwarmingTask) String() string            { return proto.CompactTextString(m) }
func (*SwarmingTask) ProtoMessage()               {}
func (*SwarmingTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

const Default_SwarmingTask_Priority int32 = 200
const Default_SwarmingTask_GracePeriodSecs int32 = 30

func (m *SwarmingTask) GetServer() string {
	if m != nil && m.Server != nil {
		return *m.Server
	}
	return ""
}

func (m *SwarmingTask) GetCommand() []string {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *SwarmingTask) GetIsolatedRef() *SwarmingTask_IsolatedRef {
	if m != nil {
		return m.IsolatedRef
	}
	return nil
}

func (m *SwarmingTask) GetExtraArgs() []string {
	if m != nil {
		return m.ExtraArgs
	}
	return nil
}

func (m *SwarmingTask) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *SwarmingTask) GetDimensions() []string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *SwarmingTask) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *SwarmingTask) GetPriority() int32 {
	if m != nil && m.Priority != nil {
		return *m.Priority
	}
	return Default_SwarmingTask_Priority
}

func (m *SwarmingTask) GetExecutionTimeoutSecs() int32 {
	if m != nil && m.ExecutionTimeoutSecs != nil {
		return *m.ExecutionTimeoutSecs
	}
	return 0
}

func (m *SwarmingTask) GetGracePeriodSecs() int32 {
	if m != nil && m.GracePeriodSecs != nil {
		return *m.GracePeriodSecs
	}
	return Default_SwarmingTask_GracePeriodSecs
}

func (m *SwarmingTask) GetIoTimeoutSecs() int32 {
	if m != nil && m.IoTimeoutSecs != nil {
		return *m.IoTimeoutSecs
	}
	return 0
}

// IsolatedRef defines a data tree reference, normally a reference to
// an .isolated file
type SwarmingTask_IsolatedRef struct {
	Isolated         *string `protobuf:"bytes,1,opt,name=isolated" json:"isolated,omitempty"`
	IsolatedServer   *string `protobuf:"bytes,2,opt,name=isolated_server" json:"isolated_server,omitempty"`
	Namespace        *string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SwarmingTask_IsolatedRef) Reset()                    { *m = SwarmingTask_IsolatedRef{} }
func (m *SwarmingTask_IsolatedRef) String() string            { return proto.CompactTextString(m) }
func (*SwarmingTask_IsolatedRef) ProtoMessage()               {}
func (*SwarmingTask_IsolatedRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *SwarmingTask_IsolatedRef) GetIsolated() string {
	if m != nil && m.Isolated != nil {
		return *m.Isolated
	}
	return ""
}

func (m *SwarmingTask_IsolatedRef) GetIsolatedServer() string {
	if m != nil && m.IsolatedServer != nil {
		return *m.IsolatedServer
	}
	return ""
}

func (m *SwarmingTask_IsolatedRef) GetNamespace() string {
	if m != nil && m.Namespace != nil {
		return *m.Namespace
	}
	return ""
}

// BuildbucketTask specifies parameters of Buildbucket-based cron job.
type BuildbucketTask struct {
	// Server is URL of the bulildbucket service to use.
	Server *string `protobuf:"bytes,1,opt,name=server" json:"server,omitempty"`
	// Bucket defines what bucket to add the task to.
	Bucket *string `protobuf:"bytes,2,opt,name=bucket" json:"bucket,omitempty"`
	// Builder defines what to run.
	Builder *string `protobuf:"bytes,3,opt,name=builder" json:"builder,omitempty"`
	// Properties is arbitrary "key:value" pairs describing the task.
	Properties []string `protobuf:"bytes,4,rep,name=properties" json:"properties,omitempty"`
	// Tags is a list of tags (as "key:value" pairs) to assign to the task.
	Tags             []string `protobuf:"bytes,5,rep,name=tags" json:"tags,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *BuildbucketTask) Reset()                    { *m = BuildbucketTask{} }
func (m *BuildbucketTask) String() string            { return proto.CompactTextString(m) }
func (*BuildbucketTask) ProtoMessage()               {}
func (*BuildbucketTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *BuildbucketTask) GetServer() string {
	if m != nil && m.Server != nil {
		return *m.Server
	}
	return ""
}

func (m *BuildbucketTask) GetBucket() string {
	if m != nil && m.Bucket != nil {
		return *m.Bucket
	}
	return ""
}

func (m *BuildbucketTask) GetBuilder() string {
	if m != nil && m.Builder != nil {
		return *m.Builder
	}
	return ""
}

func (m *BuildbucketTask) GetProperties() []string {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *BuildbucketTask) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// ProjectConfig defines a schema for cron.cfg files that describe cron jobs
// belonging to some project.
type ProjectConfig struct {
	// Job is a set of jobs defines in the project. It's singular to make
	// text-encoded proto definitions more readable.
	Job              []*Job `protobuf:"bytes,1,rep,name=job" json:"job,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ProjectConfig) Reset()                    { *m = ProjectConfig{} }
func (m *ProjectConfig) String() string            { return proto.CompactTextString(m) }
func (*ProjectConfig) ProtoMessage()               {}
func (*ProjectConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ProjectConfig) GetJob() []*Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func init() {
	proto.RegisterType((*Job)(nil), "messages.Job")
	proto.RegisterType((*Task)(nil), "messages.Task")
	proto.RegisterType((*NoopTask)(nil), "messages.NoopTask")
	proto.RegisterType((*UrlFetchTask)(nil), "messages.UrlFetchTask")
	proto.RegisterType((*SwarmingTask)(nil), "messages.SwarmingTask")
	proto.RegisterType((*SwarmingTask_IsolatedRef)(nil), "messages.SwarmingTask.IsolatedRef")
	proto.RegisterType((*BuildbucketTask)(nil), "messages.BuildbucketTask")
	proto.RegisterType((*ProjectConfig)(nil), "messages.ProjectConfig")
}

var fileDescriptor0 = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x95, 0x38, 0x4d, 0x9d, 0x71, 0xd2, 0xb4, 0x8b, 0x28, 0xa6, 0x02, 0x54, 0xf9, 0x04,
	0x42, 0x44, 0x51, 0x2a, 0x21, 0xd4, 0x23, 0x08, 0x50, 0x41, 0x42, 0x88, 0x96, 0xb3, 0xb5, 0xf1,
	0x4e, 0x9c, 0x6d, 0x6d, 0xaf, 0xb5, 0xbb, 0x2e, 0xe5, 0xc9, 0x78, 0x25, 0x1e, 0x83, 0xf1, 0xc6,
	0x96, 0x93, 0xd2, 0x9b, 0x67, 0xe6, 0x3f, 0x5f, 0xbf, 0xf1, 0x02, 0x24, 0x5a, 0x15, 0xb3, 0x52,
	0x2b, 0xab, 0x98, 0x9f, 0xa3, 0x31, 0x3c, 0x45, 0x13, 0x5d, 0x82, 0xf7, 0x45, 0x2d, 0x19, 0x40,
	0x5f, 0x8a, 0xb0, 0x77, 0xda, 0x7b, 0x39, 0x62, 0x87, 0xe0, 0x9b, 0x64, 0x8d, 0xa2, 0xca, 0x30,
	0xec, 0xb7, 0x1e, 0x21, 0x0d, 0x5f, 0x66, 0x28, 0x42, 0x8f, 0x3c, 0x3e, 0x7b, 0x06, 0x03, 0xcb,
	0xcd, 0x4d, 0x38, 0x20, 0x2b, 0x58, 0x1c, 0xcc, 0xda, 0x7a, 0xb3, 0x2b, 0xf2, 0x46, 0x7f, 0x7a,
	0x30, 0xa8, 0x3f, 0xd8, 0x29, 0x0c, 0x0a, 0xa5, 0x4a, 0x57, 0x38, 0x58, 0xb0, 0x4e, 0xf6, 0x8d,
	0xbc, 0x4e, 0xf1, 0x0a, 0x46, 0x95, 0xce, 0xe2, 0x15, 0xda, 0x64, 0xed, 0xba, 0x05, 0x8b, 0xe3,
	0x4e, 0xf6, 0x53, 0x67, 0x9f, 0xea, 0x88, 0x93, 0xbe, 0x81, 0x89, 0xf9, 0xc5, 0x75, 0x2e, 0x8b,
	0x34, 0x76, 0xcd, 0xbd, 0xfb, 0xf2, 0xcb, 0x26, 0xec, 0xe4, 0x67, 0x70, 0xb8, 0xac, 0x64, 0x26,
	0x96, 0x55, 0x72, 0x83, 0x36, 0xde, 0x1a, 0xf7, 0x69, 0x97, 0xf1, 0xbe, 0x53, 0xb8, 0xc9, 0x01,
	0xfc, 0x76, 0xb4, 0xe8, 0x02, 0xc6, 0x3b, 0xfd, 0x1f, 0xc1, 0x30, 0x47, 0xbb, 0x56, 0x0d, 0xa7,
	0x73, 0xef, 0xf3, 0xc7, 0x2b, 0x16, 0x80, 0x47, 0xf3, 0x37, 0x9c, 0x9e, 0x40, 0x60, 0x65, 0x8e,
	0xaa, 0xb2, 0xb1, 0xc1, 0xc4, 0xcd, 0xb7, 0x77, 0xde, 0x7f, 0x3b, 0x8f, 0xfe, 0xf6, 0x61, 0xbc,
	0x33, 0xdc, 0x01, 0x0c, 0x0d, 0xea, 0x5b, 0xd4, 0x0d, 0xf3, 0x29, 0xec, 0x27, 0x2a, 0xcf, 0x79,
	0x21, 0xa8, 0x94, 0x47, 0x8e, 0x77, 0x30, 0x96, 0x46, 0x65, 0xdc, 0xa2, 0x88, 0x35, 0xae, 0x9a,
	0x5d, 0xa3, 0x87, 0x77, 0x9d, 0x5d, 0x34, 0xd2, 0x1f, 0xb8, 0x62, 0x74, 0x4b, 0xbc, 0xb3, 0x9a,
	0xc7, 0x5c, 0xa7, 0x86, 0x36, 0xae, 0xab, 0xd1, 0x94, 0x58, 0xdc, 0x86, 0x7b, 0xce, 0x20, 0x81,
	0xa0, 0x29, 0x0b, 0x23, 0x55, 0x61, 0xc2, 0xa1, 0xf3, 0x8d, 0xeb, 0x7b, 0x92, 0x7c, 0xdf, 0x59,
	0x8f, 0xc1, 0x2f, 0xb5, 0x54, 0x5a, 0xda, 0xdf, 0xa1, 0xef, 0x96, 0xf0, 0x16, 0xf3, 0x39, 0x7b,
	0x01, 0xc7, 0x78, 0x87, 0x49, 0x65, 0x29, 0x31, 0xde, 0x5a, 0xd4, 0x84, 0xa3, 0x5a, 0xc4, 0x9e,
	0xc3, 0x51, 0xaa, 0x79, 0x82, 0x71, 0x89, 0x94, 0x2d, 0x36, 0x21, 0xd8, 0x40, 0x38, 0x9b, 0x13,
	0x9d, 0xa9, 0x54, 0xbb, 0x79, 0x41, 0x1d, 0x3c, 0xf9, 0x0a, 0xc1, 0xf6, 0x02, 0xf4, 0xb7, 0xb5,
	0xab, 0x37, 0x74, 0xea, 0xcc, 0x16, 0x46, 0x83, 0x6d, 0x03, 0xfc, 0x08, 0x46, 0x05, 0x27, 0x24,
	0x25, 0x75, 0x75, 0x88, 0x46, 0x91, 0x80, 0xe9, 0xbd, 0xa3, 0xfe, 0x07, 0x9b, 0xec, 0x4d, 0xb4,
	0xa9, 0x42, 0xf0, 0xdd, 0x9f, 0x42, 0x02, 0x57, 0xa3, 0x26, 0x44, 0xef, 0x84, 0xb6, 0xb0, 0x12,
	0x5b, 0x84, 0x2d, 0x21, 0xc7, 0x30, 0x7a, 0x0d, 0x93, 0xef, 0x5a, 0x5d, 0x63, 0x62, 0x3f, 0xa8,
	0x62, 0x25, 0x53, 0x76, 0x02, 0xde, 0xb5, 0x5a, 0x52, 0x03, 0x8f, 0xce, 0x34, 0xe9, 0xce, 0x44,
	0x8f, 0xeb, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x44, 0xa9, 0x51, 0x7a, 0x03, 0x00, 0x00,
}
