// Code generated by protoc-gen-go.
// source: acquisition_task.proto
// DO NOT EDIT!

package ts_mon_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Task struct {
	ServiceName string `protobuf:"bytes,20,opt,name=service_name" json:"service_name,omitempty"`
	JobName     string `protobuf:"bytes,30,opt,name=job_name" json:"job_name,omitempty"`
	DataCenter  string `protobuf:"bytes,40,opt,name=data_center" json:"data_center,omitempty"`
	HostName    string `protobuf:"bytes,50,opt,name=host_name" json:"host_name,omitempty"`
	TaskNum     int32  `protobuf:"varint,60,opt,name=task_num" json:"task_num,omitempty"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
