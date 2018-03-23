// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/buildbucket/client/cmd/buildbucket/proto/project_config.proto

/*
Package buildbucket is a generated protocol buffer package.

It is generated from these files:
	go.chromium.org/luci/buildbucket/client/cmd/buildbucket/proto/project_config.proto

It has these top-level messages:
	Acl
	AclSet
	Swarming
	Bucket
	BuildbucketCfg
*/
package buildbucket

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

type Acl_Role int32

const (
	// Can do read-only operations, such as search for builds.
	Acl_READER Acl_Role = 0
	// Same as READER + can schedule and cancel builds.
	Acl_SCHEDULER Acl_Role = 1
	// Can do all write operations.
	Acl_WRITER Acl_Role = 2
)

var Acl_Role_name = map[int32]string{
	0: "READER",
	1: "SCHEDULER",
	2: "WRITER",
}
var Acl_Role_value = map[string]int32{
	"READER":    0,
	"SCHEDULER": 1,
	"WRITER":    2,
}

func (x Acl_Role) Enum() *Acl_Role {
	p := new(Acl_Role)
	*p = x
	return p
}
func (x Acl_Role) String() string {
	return proto.EnumName(Acl_Role_name, int32(x))
}
func (x *Acl_Role) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Acl_Role_value, data, "Acl_Role")
	if err != nil {
		return err
	}
	*x = Acl_Role(value)
	return nil
}
func (Acl_Role) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// A single access control rule.
type Acl struct {
	// Role denotes a list of actions that an identity can perform.
	Role *Acl_Role `protobuf:"varint,1,opt,name=role,enum=buildbucket.Acl_Role" json:"role,omitempty"`
	// Name of the group defined in the auth service.
	Group *string `protobuf:"bytes,2,opt,name=group" json:"group,omitempty"`
	// An email address or a full identity string "kind:name". See auth service
	// on kinds of identities. Anonymous users are "anonymous:anonymous".
	// Either identity or group must be present, not both.
	Identity         *string `protobuf:"bytes,3,opt,name=identity" json:"identity,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Acl) Reset()                    { *m = Acl{} }
func (m *Acl) String() string            { return proto.CompactTextString(m) }
func (*Acl) ProtoMessage()               {}
func (*Acl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Acl) GetRole() Acl_Role {
	if m != nil && m.Role != nil {
		return *m.Role
	}
	return Acl_READER
}

func (m *Acl) GetGroup() string {
	if m != nil && m.Group != nil {
		return *m.Group
	}
	return ""
}

func (m *Acl) GetIdentity() string {
	if m != nil && m.Identity != nil {
		return *m.Identity
	}
	return ""
}

// A set of Acl messages. Can be referenced in a bucket by name.
type AclSet struct {
	// A name of the ACL set. Required. Must match regex '^[a-z0-9_]+$'.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// List of access control rules.
	// The order does not matter.
	Acls             []*Acl `protobuf:"bytes,2,rep,name=acls" json:"acls,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *AclSet) Reset()                    { *m = AclSet{} }
func (m *AclSet) String() string            { return proto.CompactTextString(m) }
func (*AclSet) ProtoMessage()               {}
func (*AclSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AclSet) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *AclSet) GetAcls() []*Acl {
	if m != nil {
		return m.Acls
	}
	return nil
}

// Configuration of buildbucket-swarming integration for one bucket.
type Swarming struct {
	// Hostname of the swarming instance, e.g. "chromium-swarm.appspot.com".
	Hostname *string `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	// Used to generate a URL for Build, may contain parameters
	// {swarming_hostname}, {task_id}, {bucket} and {builder}. Defaults to:
	// https://{swarming_hostname}/user/task/{task_id}
	UrlFormat *string `protobuf:"bytes,2,opt,name=url_format,json=urlFormat" json:"url_format,omitempty"`
	// Defines default values for builders.
	BuilderDefaults *Swarming_Builder `protobuf:"bytes,3,opt,name=builder_defaults,json=builderDefaults" json:"builder_defaults,omitempty"`
	// Configuration for each builder.
	// Swarming tasks are created only for builds for builders that are not
	// explicitly specified.
	Builders []*Swarming_Builder `protobuf:"bytes,4,rep,name=builders" json:"builders,omitempty"`
	// Percentage of builds that should use a canary swarming task template.
	// A value from 0 to 100.
	TaskTemplateCanaryPercentage *uint32 `protobuf:"varint,5,opt,name=task_template_canary_percentage,json=taskTemplateCanaryPercentage" json:"task_template_canary_percentage,omitempty"`
	XXX_unrecognized             []byte  `json:"-"`
}

func (m *Swarming) Reset()                    { *m = Swarming{} }
func (m *Swarming) String() string            { return proto.CompactTextString(m) }
func (*Swarming) ProtoMessage()               {}
func (*Swarming) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Swarming) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *Swarming) GetUrlFormat() string {
	if m != nil && m.UrlFormat != nil {
		return *m.UrlFormat
	}
	return ""
}

func (m *Swarming) GetBuilderDefaults() *Swarming_Builder {
	if m != nil {
		return m.BuilderDefaults
	}
	return nil
}

func (m *Swarming) GetBuilders() []*Swarming_Builder {
	if m != nil {
		return m.Builders
	}
	return nil
}

func (m *Swarming) GetTaskTemplateCanaryPercentage() uint32 {
	if m != nil && m.TaskTemplateCanaryPercentage != nil {
		return *m.TaskTemplateCanaryPercentage
	}
	return 0
}

type Swarming_Recipe struct {
	// Repository URL of the recipe package.
	Repository *string `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	// Name of the recipe to run.
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// colon-separated build properties to set.
	// A property can be overriden by "properties" build parameter.
	//
	// Use this field for string properties and use properties_j for other
	// types.
	Properties []string `protobuf:"bytes,3,rep,name=properties" json:"properties,omitempty"`
	// Same as properties, but the value must valid JSON. For example
	//   properties_j: "a:1"
	// means property a is a number 1, not string "1".
	//
	// If null, it means no property must be defined. In particular, it removes
	// a default value for the property, if any.
	//
	// Fields properties and properties_j can be used together, but cannot both
	// specify values for same property.
	PropertiesJ      []string `protobuf:"bytes,4,rep,name=properties_j,json=propertiesJ" json:"properties_j,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Swarming_Recipe) Reset()                    { *m = Swarming_Recipe{} }
func (m *Swarming_Recipe) String() string            { return proto.CompactTextString(m) }
func (*Swarming_Recipe) ProtoMessage()               {}
func (*Swarming_Recipe) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *Swarming_Recipe) GetRepository() string {
	if m != nil && m.Repository != nil {
		return *m.Repository
	}
	return ""
}

func (m *Swarming_Recipe) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Swarming_Recipe) GetProperties() []string {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Swarming_Recipe) GetPropertiesJ() []string {
	if m != nil {
		return m.PropertiesJ
	}
	return nil
}

// A builder has a name, a category and specifies what should happen if a
// build is scheduled to that builder.
//
// SECURITY WARNING: if adding more fields to this message, keep in mind that
// a user that has permissions to schedule a build to the bucket, can override
// this config.
type Swarming_Builder struct {
	// Name of the builder. Will be propagated to "builder" build tag and
	// "buildername" recipe property.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Builder category. Will be used for visual grouping, for example in Code Review.
	Category *string `protobuf:"bytes,6,opt,name=category" json:"category,omitempty"`
	// Will be become to swarming task tags.
	// Each tag will end up in "swarming_tag" buildbucket tag, for example
	// "swarming_tag:builder:release"
	SwarmingTags []string `protobuf:"bytes,2,rep,name=swarming_tags,json=swarmingTags" json:"swarming_tags,omitempty"`
	// Colon-delimited key-value pair of task dimensions.
	//
	// If value is not specified ("<key>:"), then it excludes a default value.
	Dimensions []string `protobuf:"bytes,3,rep,name=dimensions" json:"dimensions,omitempty"`
	// CIPD packages to install on the builder.
	CipdPackages []*Swarming_Builder_CipdPackage `protobuf:"bytes,8,rep,name=cipd_packages,json=cipdPackages" json:"cipd_packages,omitempty"`
	// Specifies that a recipe to run.
	Recipe *Swarming_Recipe `protobuf:"bytes,4,opt,name=recipe" json:"recipe,omitempty"`
	// Swarming task priority.
	Priority *uint32 `protobuf:"varint,5,opt,name=priority" json:"priority,omitempty"`
	// Maximum build execution time. Not to be confused with pending time.
	ExecutionTimeoutSecs *uint32 `protobuf:"varint,7,opt,name=execution_timeout_secs,json=executionTimeoutSecs" json:"execution_timeout_secs,omitempty"`
	// Caches that should be present on the bot.
	Caches           []*Swarming_Builder_CacheEntry `protobuf:"bytes,9,rep,name=caches" json:"caches,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *Swarming_Builder) Reset()                    { *m = Swarming_Builder{} }
func (m *Swarming_Builder) String() string            { return proto.CompactTextString(m) }
func (*Swarming_Builder) ProtoMessage()               {}
func (*Swarming_Builder) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 1} }

func (m *Swarming_Builder) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Swarming_Builder) GetCategory() string {
	if m != nil && m.Category != nil {
		return *m.Category
	}
	return ""
}

func (m *Swarming_Builder) GetSwarmingTags() []string {
	if m != nil {
		return m.SwarmingTags
	}
	return nil
}

func (m *Swarming_Builder) GetDimensions() []string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *Swarming_Builder) GetCipdPackages() []*Swarming_Builder_CipdPackage {
	if m != nil {
		return m.CipdPackages
	}
	return nil
}

func (m *Swarming_Builder) GetRecipe() *Swarming_Recipe {
	if m != nil {
		return m.Recipe
	}
	return nil
}

func (m *Swarming_Builder) GetPriority() uint32 {
	if m != nil && m.Priority != nil {
		return *m.Priority
	}
	return 0
}

func (m *Swarming_Builder) GetExecutionTimeoutSecs() uint32 {
	if m != nil && m.ExecutionTimeoutSecs != nil {
		return *m.ExecutionTimeoutSecs
	}
	return 0
}

func (m *Swarming_Builder) GetCaches() []*Swarming_Builder_CacheEntry {
	if m != nil {
		return m.Caches
	}
	return nil
}

type Swarming_Builder_CipdPackage struct {
	// A template of a full CIPD package name, e.g
	// "infra/tools/luci-auth/${platform}". This can be parametrized using
	// ${platform} and ${os_ver} parameters, where ${platform} will be
	// expanded into "<os>-<architecture>" and ${os_ver} will be expanded to
	// OS version name.
	PackageName *string `protobuf:"bytes,1,opt,name=package_name,json=packageName" json:"package_name,omitempty"`
	// Path to dir, relative to the task working dir, where to install the
	// package. The path cannot be empty or start with a slash.
	Path *string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	// Valid package version for all packages matched by package name.
	Version          *string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Swarming_Builder_CipdPackage) Reset()         { *m = Swarming_Builder_CipdPackage{} }
func (m *Swarming_Builder_CipdPackage) String() string { return proto.CompactTextString(m) }
func (*Swarming_Builder_CipdPackage) ProtoMessage()    {}
func (*Swarming_Builder_CipdPackage) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 1, 0}
}

func (m *Swarming_Builder_CipdPackage) GetPackageName() string {
	if m != nil && m.PackageName != nil {
		return *m.PackageName
	}
	return ""
}

func (m *Swarming_Builder_CipdPackage) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *Swarming_Builder_CipdPackage) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

// Describes a named cache that should be present on the bot.
// See also https://github.com/luci/luci-py/blob/3a2941345cf011a96bcd83d76328395323245bb5/appengine/swarming/swarming_rpcs.py#L166
type Swarming_Builder_CacheEntry struct {
	// Unique name of the cache. Required. Length is limited to 4096.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Relative path to the directory that will be linked to the named cache.
	// Required.
	// A path cannot be shared among multiple caches or CIPD installations.
	// A task will fail if a file/dir with the same name already exists.
	Path             *string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Swarming_Builder_CacheEntry) Reset()         { *m = Swarming_Builder_CacheEntry{} }
func (m *Swarming_Builder_CacheEntry) String() string { return proto.CompactTextString(m) }
func (*Swarming_Builder_CacheEntry) ProtoMessage()    {}
func (*Swarming_Builder_CacheEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 1, 1}
}

func (m *Swarming_Builder_CacheEntry) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Swarming_Builder_CacheEntry) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

// Defines one bucket in buildbucket.cfg
type Bucket struct {
	// Name of the bucket. Names are unique within one instance of buildbucket.
	// If another project already uses this name, a config will be rejected.
	// Name reservation is first-come first-serve.
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// List of access control rules for the bucket.
	// The order does not matter.
	Acls []*Acl `protobuf:"bytes,2,rep,name=acls" json:"acls,omitempty"`
	// A list of ACL set names. Each ACL in each referenced ACL set will be
	// included in this bucket.
	// The order does not matter.
	AclSets []string `protobuf:"bytes,4,rep,name=acl_sets,json=aclSets" json:"acl_sets,omitempty"`
	// Buildbucket-swarming integration.
	Swarming         *Swarming `protobuf:"bytes,3,opt,name=swarming" json:"swarming,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Bucket) Reset()                    { *m = Bucket{} }
func (m *Bucket) String() string            { return proto.CompactTextString(m) }
func (*Bucket) ProtoMessage()               {}
func (*Bucket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Bucket) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Bucket) GetAcls() []*Acl {
	if m != nil {
		return m.Acls
	}
	return nil
}

func (m *Bucket) GetAclSets() []string {
	if m != nil {
		return m.AclSets
	}
	return nil
}

func (m *Bucket) GetSwarming() *Swarming {
	if m != nil {
		return m.Swarming
	}
	return nil
}

// Schema of buildbucket.cfg file, a project config.
type BuildbucketCfg struct {
	// All buckets defined for this project.
	Buckets []*Bucket `protobuf:"bytes,1,rep,name=buckets" json:"buckets,omitempty"`
	// A list of ACL sets. Names must be unique.
	AclSets          []*AclSet `protobuf:"bytes,2,rep,name=acl_sets,json=aclSets" json:"acl_sets,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *BuildbucketCfg) Reset()                    { *m = BuildbucketCfg{} }
func (m *BuildbucketCfg) String() string            { return proto.CompactTextString(m) }
func (*BuildbucketCfg) ProtoMessage()               {}
func (*BuildbucketCfg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BuildbucketCfg) GetBuckets() []*Bucket {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func (m *BuildbucketCfg) GetAclSets() []*AclSet {
	if m != nil {
		return m.AclSets
	}
	return nil
}

func init() {
	proto.RegisterType((*Acl)(nil), "buildbucket.Acl")
	proto.RegisterType((*AclSet)(nil), "buildbucket.AclSet")
	proto.RegisterType((*Swarming)(nil), "buildbucket.Swarming")
	proto.RegisterType((*Swarming_Recipe)(nil), "buildbucket.Swarming.Recipe")
	proto.RegisterType((*Swarming_Builder)(nil), "buildbucket.Swarming.Builder")
	proto.RegisterType((*Swarming_Builder_CipdPackage)(nil), "buildbucket.Swarming.Builder.CipdPackage")
	proto.RegisterType((*Swarming_Builder_CacheEntry)(nil), "buildbucket.Swarming.Builder.CacheEntry")
	proto.RegisterType((*Bucket)(nil), "buildbucket.Bucket")
	proto.RegisterType((*BuildbucketCfg)(nil), "buildbucket.BuildbucketCfg")
	proto.RegisterEnum("buildbucket.Acl_Role", Acl_Role_name, Acl_Role_value)
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/buildbucket/client/cmd/buildbucket/proto/project_config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 735 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xdb, 0x38,
	0x10, 0x5e, 0xd9, 0x8a, 0x7f, 0xc6, 0x71, 0xd6, 0xe0, 0x26, 0x0b, 0xad, 0x91, 0x64, 0xbd, 0xde,
	0x3d, 0x38, 0x87, 0xc8, 0xd8, 0x20, 0x97, 0xde, 0x6a, 0x3b, 0x2e, 0xd2, 0xa2, 0x08, 0x02, 0xda,
	0x45, 0x6f, 0x15, 0x18, 0x9a, 0x96, 0x99, 0x48, 0xa2, 0x40, 0x52, 0x6d, 0x7d, 0xea, 0x2b, 0xb4,
	0x2f, 0xd2, 0x97, 0xe8, 0x8b, 0x15, 0xa2, 0x64, 0x59, 0x29, 0x8c, 0x06, 0xe8, 0x45, 0xe0, 0x7c,
	0xfc, 0x66, 0xf8, 0xcd, 0x7c, 0x14, 0x01, 0xfb, 0xc2, 0xa5, 0x2b, 0x29, 0x42, 0x9e, 0x84, 0xae,
	0x90, 0xfe, 0x30, 0x48, 0x28, 0x1f, 0xde, 0x25, 0x3c, 0x58, 0xdc, 0x25, 0xf4, 0x81, 0xe9, 0x21,
	0x0d, 0x38, 0x8b, 0xf4, 0x90, 0x86, 0x8b, 0x47, 0x70, 0x2c, 0x85, 0x16, 0xe9, 0xf7, 0x9e, 0x51,
	0xed, 0x51, 0x11, 0x2d, 0xb9, 0xef, 0x1a, 0x10, 0xb5, 0x4a, 0xbc, 0xfe, 0x17, 0x0b, 0xaa, 0x23,
	0x1a, 0xa0, 0x33, 0xb0, 0xa5, 0x08, 0x98, 0x63, 0xf5, 0xac, 0xc1, 0xc1, 0xc5, 0x91, 0x5b, 0xe2,
	0xb8, 0x23, 0x1a, 0xb8, 0x58, 0x04, 0x0c, 0x1b, 0x0a, 0x3a, 0x84, 0x3d, 0x5f, 0x8a, 0x24, 0x76,
	0x2a, 0x3d, 0x6b, 0xd0, 0xc4, 0x59, 0x80, 0xba, 0xd0, 0xe0, 0x0b, 0x16, 0x69, 0xae, 0xd7, 0x4e,
	0xd5, 0x6c, 0x14, 0x71, 0xff, 0x1c, 0xec, 0x34, 0x1f, 0x01, 0xd4, 0xf0, 0x74, 0x74, 0x35, 0xc5,
	0x9d, 0xdf, 0x50, 0x1b, 0x9a, 0xb3, 0xc9, 0xf5, 0xf4, 0xea, 0xcd, 0xeb, 0x29, 0xee, 0x58, 0xe9,
	0xd6, 0x5b, 0xfc, 0x72, 0x3e, 0xc5, 0x9d, 0x4a, 0x7f, 0x0c, 0xb5, 0x11, 0x0d, 0x66, 0x4c, 0x23,
	0x04, 0x76, 0x44, 0xc2, 0x4c, 0x55, 0x13, 0x9b, 0x35, 0xfa, 0x0f, 0x6c, 0x42, 0x03, 0xe5, 0x54,
	0x7a, 0xd5, 0x41, 0xeb, 0xa2, 0xf3, 0xa3, 0x52, 0x6c, 0x76, 0xfb, 0xdf, 0xea, 0xd0, 0x98, 0x7d,
	0x20, 0x32, 0xe4, 0x91, 0x9f, 0x6a, 0x5b, 0x09, 0xa5, 0x4b, 0xa5, 0x8a, 0x18, 0x9d, 0x00, 0x24,
	0x32, 0xf0, 0x96, 0x42, 0x86, 0x44, 0xe7, 0x2d, 0x35, 0x13, 0x19, 0xbc, 0x30, 0x00, 0xba, 0x86,
	0x8e, 0x39, 0x80, 0x49, 0x6f, 0xc1, 0x96, 0x24, 0x09, 0xb4, 0x32, 0xed, 0xb5, 0x2e, 0x4e, 0x1e,
	0x9d, 0xbc, 0x39, 0xcb, 0x1d, 0x67, 0x6c, 0xfc, 0x7b, 0x9e, 0x76, 0x95, 0x67, 0xa1, 0x67, 0xd0,
	0xc8, 0x21, 0xe5, 0xd8, 0x46, 0xfb, 0x13, 0x15, 0x0a, 0x3a, 0x9a, 0xc2, 0xdf, 0x9a, 0xa8, 0x07,
	0x4f, 0xb3, 0x30, 0x0e, 0x88, 0x66, 0x1e, 0x25, 0x11, 0x91, 0x6b, 0x2f, 0x66, 0x92, 0xb2, 0x48,
	0x13, 0x9f, 0x39, 0x7b, 0x3d, 0x6b, 0xd0, 0xc6, 0xc7, 0x29, 0x6d, 0x9e, 0xb3, 0x26, 0x86, 0x74,
	0x5b, 0x70, 0xba, 0x9f, 0xa0, 0x86, 0x19, 0xe5, 0x31, 0x43, 0xa7, 0x00, 0x92, 0xc5, 0x42, 0x71,
	0x2d, 0xe4, 0x3a, 0x1f, 0x49, 0x09, 0x29, 0xe6, 0x5e, 0x29, 0xcd, 0xfd, 0x14, 0x20, 0x96, 0x22,
	0x66, 0x52, 0x73, 0x96, 0xce, 0xa0, 0x9a, 0xe6, 0x6c, 0x11, 0xf4, 0x0f, 0xec, 0x6f, 0x23, 0xef,
	0xde, 0xf4, 0xd8, 0xc4, 0xad, 0x2d, 0xf6, 0xaa, 0xfb, 0xd5, 0x86, 0x7a, 0xde, 0xdd, 0x4e, 0x6b,
	0xbb, 0xd0, 0xa0, 0x44, 0x33, 0x3f, 0x15, 0x55, 0xcb, 0x7c, 0xda, 0xc4, 0xe8, 0x5f, 0x68, 0xab,
	0x7c, 0x42, 0x9e, 0x26, 0x7e, 0xe6, 0x7f, 0x13, 0xef, 0x6f, 0xc0, 0x39, 0xf1, 0x55, 0xaa, 0x71,
	0xc1, 0x43, 0x16, 0x29, 0x2e, 0xa2, 0x42, 0xe3, 0x16, 0x41, 0x37, 0xd0, 0xa6, 0x3c, 0x5e, 0x78,
	0x31, 0xa1, 0x0f, 0xc4, 0x67, 0xca, 0x69, 0x18, 0x23, 0xce, 0x7e, 0x6a, 0x84, 0x3b, 0xe1, 0xf1,
	0xe2, 0x36, 0xcb, 0xc0, 0xfb, 0x74, 0x1b, 0x28, 0x74, 0x09, 0x35, 0x69, 0x26, 0xea, 0xd8, 0xe6,
	0x4e, 0x1c, 0xef, 0x2e, 0x94, 0x4d, 0x1d, 0xe7, 0xdc, 0xb4, 0xcd, 0x58, 0x72, 0x21, 0xd3, 0x5f,
	0x25, 0xf3, 0xad, 0x88, 0xd1, 0x25, 0xfc, 0xc9, 0x3e, 0x32, 0x9a, 0x68, 0x2e, 0x22, 0x4f, 0xf3,
	0x90, 0x89, 0x44, 0x7b, 0x8a, 0x51, 0xe5, 0xd4, 0x0d, 0xf3, 0xb0, 0xd8, 0x9d, 0x67, 0x9b, 0x33,
	0x46, 0x15, 0x7a, 0x0e, 0x35, 0x4a, 0xe8, 0x8a, 0x29, 0xa7, 0x69, 0x1a, 0x1a, 0x3c, 0xd1, 0x50,
	0xca, 0x9d, 0x46, 0x5a, 0xae, 0x71, 0x9e, 0xd7, 0x7d, 0x07, 0xad, 0x52, 0x9b, 0xc6, 0xcc, 0x6c,
	0xe9, 0x95, 0x5c, 0x6a, 0xe5, 0xd8, 0x4d, 0x6a, 0x16, 0x02, 0x3b, 0x26, 0x7a, 0xb5, 0xb9, 0x23,
	0xe9, 0x1a, 0x39, 0x50, 0x7f, 0xcf, 0x64, 0x3a, 0xeb, 0xfc, 0x0d, 0xd8, 0x84, 0xdd, 0x4b, 0x80,
	0xed, 0xa9, 0x3b, 0xcd, 0xdf, 0x51, 0xaf, 0xff, 0xd9, 0x82, 0xda, 0xd8, 0x34, 0xf1, 0xeb, 0x4f,
	0x01, 0xfa, 0x0b, 0x1a, 0x84, 0x06, 0x9e, 0x62, 0x5a, 0xe5, 0x97, 0xb2, 0x4e, 0xcc, 0xf3, 0xa2,
	0xd0, 0xff, 0xd0, 0xd8, 0xdc, 0x9f, 0xfc, 0xaf, 0x3e, 0xda, 0x39, 0x39, 0x5c, 0xd0, 0xfa, 0x02,
	0x0e, 0xc6, 0x5b, 0xc6, 0x64, 0xe9, 0xa3, 0x73, 0xa8, 0x67, 0x81, 0x72, 0x2c, 0x23, 0xe4, 0x8f,
	0x47, 0x35, 0x32, 0xfd, 0x78, 0xc3, 0x41, 0x6e, 0x49, 0x4e, 0x65, 0x07, 0x3f, 0x7b, 0xfa, 0x0a,
	0x8d, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x57, 0x6c, 0x51, 0x03, 0x06, 0x00, 0x00,
}