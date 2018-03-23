// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/common/proto/config/service_config.proto

package config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Depending on repository type, config service imports configuration files
// differently.
type ConfigSetLocation_StorageType int32

const (
	// Will be used if storage_type is not set.
	ConfigSetLocation_UNSET ConfigSetLocation_StorageType = 0
	// Gitiles REST API is used to fetch config files.
	ConfigSetLocation_GITILES ConfigSetLocation_StorageType = 1
)

var ConfigSetLocation_StorageType_name = map[int32]string{
	0: "UNSET",
	1: "GITILES",
}
var ConfigSetLocation_StorageType_value = map[string]int32{
	"UNSET":   0,
	"GITILES": 1,
}

func (x ConfigSetLocation_StorageType) Enum() *ConfigSetLocation_StorageType {
	p := new(ConfigSetLocation_StorageType)
	*p = x
	return p
}
func (x ConfigSetLocation_StorageType) String() string {
	return proto.EnumName(ConfigSetLocation_StorageType_name, int32(x))
}
func (x *ConfigSetLocation_StorageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ConfigSetLocation_StorageType_value, data, "ConfigSetLocation_StorageType")
	if err != nil {
		return err
	}
	*x = ConfigSetLocation_StorageType(value)
	return nil
}
func (ConfigSetLocation_StorageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{0, 0}
}

// Severity of a validation response message. In JSON encoded as a string.
type ValidationResponseMessage_Severity int32

const (
	ValidationResponseMessage_DEBUG    ValidationResponseMessage_Severity = 10
	ValidationResponseMessage_INFO     ValidationResponseMessage_Severity = 20
	ValidationResponseMessage_WARNING  ValidationResponseMessage_Severity = 30
	ValidationResponseMessage_ERROR    ValidationResponseMessage_Severity = 40
	ValidationResponseMessage_CRITICAL ValidationResponseMessage_Severity = 50
)

var ValidationResponseMessage_Severity_name = map[int32]string{
	10: "DEBUG",
	20: "INFO",
	30: "WARNING",
	40: "ERROR",
	50: "CRITICAL",
}
var ValidationResponseMessage_Severity_value = map[string]int32{
	"DEBUG":    10,
	"INFO":     20,
	"WARNING":  30,
	"ERROR":    40,
	"CRITICAL": 50,
}

func (x ValidationResponseMessage_Severity) Enum() *ValidationResponseMessage_Severity {
	p := new(ValidationResponseMessage_Severity)
	*p = x
	return p
}
func (x ValidationResponseMessage_Severity) String() string {
	return proto.EnumName(ValidationResponseMessage_Severity_name, int32(x))
}
func (x *ValidationResponseMessage_Severity) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ValidationResponseMessage_Severity_value, data, "ValidationResponseMessage_Severity")
	if err != nil {
		return err
	}
	*x = ValidationResponseMessage_Severity(value)
	return nil
}
func (ValidationResponseMessage_Severity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{12, 0}
}

// Used to specify project/service configuration location.
type ConfigSetLocation struct {
	// URL of the repository where project-wide configurations are stored.
	Url *string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	// Type of the configuration repository, e.g. GITILES.
	StorageType      *ConfigSetLocation_StorageType `protobuf:"varint,2,opt,name=storage_type,json=storageType,enum=config.ConfigSetLocation_StorageType" json:"storage_type,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *ConfigSetLocation) Reset()                    { *m = ConfigSetLocation{} }
func (m *ConfigSetLocation) String() string            { return proto.CompactTextString(m) }
func (*ConfigSetLocation) ProtoMessage()               {}
func (*ConfigSetLocation) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ConfigSetLocation) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

func (m *ConfigSetLocation) GetStorageType() ConfigSetLocation_StorageType {
	if m != nil && m.StorageType != nil {
		return *m.StorageType
	}
	return ConfigSetLocation_UNSET
}

// A tenant of a service. Defined in projects.cfg.
type Project struct {
	// Globally unique id of the project.
	Id *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Where to import "projects/<id>" config set from.
	ConfigLocation   *ConfigSetLocation `protobuf:"bytes,2,opt,name=config_location,json=configLocation" json:"config_location,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Project) Reset()                    { *m = Project{} }
func (m *Project) String() string            { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()               {}
func (*Project) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Project) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Project) GetConfigLocation() *ConfigSetLocation {
	if m != nil {
		return m.ConfigLocation
	}
	return nil
}

// Schema of projects.cfg file. Represents luci tenants registry.
type ProjectsCfg struct {
	// All projects served by this instance of Luci.
	Projects         []*Project `protobuf:"bytes,1,rep,name=projects" json:"projects,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *ProjectsCfg) Reset()                    { *m = ProjectsCfg{} }
func (m *ProjectsCfg) String() string            { return proto.CompactTextString(m) }
func (*ProjectsCfg) ProtoMessage()               {}
func (*ProjectsCfg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ProjectsCfg) GetProjects() []*Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

// Describes one luci service.
type Service struct {
	// Globally unique id of the service. Required.
	// Used in "services/<service_id>" config set name.
	Id *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Email addresses of responsible and point-of-contacts for the service.
	Owners []string `protobuf:"bytes,2,rep,name=owners" json:"owners,omitempty"`
	// Where to import "services/<id>" config set from. If config_location.url is
	// relative, it is relative to the current configuration file.
	// If not specified, defaults to "../<id>/".
	// Not yet implemented.
	ConfigLocation *ConfigSetLocation `protobuf:"bytes,3,opt,name=config_location,json=configLocation" json:"config_location,omitempty"`
	// An HTTPS endpoint that returns JSON-encoded ServiceDynamicMetadata in body.
	MetadataUrl *string `protobuf:"bytes,4,opt,name=metadata_url,json=metadataUrl" json:"metadata_url,omitempty"`
	// A list of identities that have access to this service's configs.
	// of:
	// * "group:<group>", where group is defined on auth server.
	// * "<email>"
	// * "<identity>"
	//
	// If not specified, only admins and trusted services have access.
	Access           []string `protobuf:"bytes,5,rep,name=access" json:"access,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *Service) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *Service) GetOwners() []string {
	if m != nil {
		return m.Owners
	}
	return nil
}

func (m *Service) GetConfigLocation() *ConfigSetLocation {
	if m != nil {
		return m.ConfigLocation
	}
	return nil
}

func (m *Service) GetMetadataUrl() string {
	if m != nil && m.MetadataUrl != nil {
		return *m.MetadataUrl
	}
	return ""
}

func (m *Service) GetAccess() []string {
	if m != nil {
		return m.Access
	}
	return nil
}

// Machine-generated service metadata, exposed by a service endpoint.
// Typically implemented by config component, embedded in an app:
// see appengine/components/components/config/endpoint.py
//
// If you add a field here, also add it to ServiceDynamicMetadata in endpoint.py
type ServiceDynamicMetadata struct {
	// Format version. Supported versions: 1.0.
	Version *string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	// What configs this service can validate and how to validate them.
	Validation       *Validator `protobuf:"bytes,2,opt,name=validation" json:"validation,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *ServiceDynamicMetadata) Reset()                    { *m = ServiceDynamicMetadata{} }
func (m *ServiceDynamicMetadata) String() string            { return proto.CompactTextString(m) }
func (*ServiceDynamicMetadata) ProtoMessage()               {}
func (*ServiceDynamicMetadata) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ServiceDynamicMetadata) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *ServiceDynamicMetadata) GetValidation() *Validator {
	if m != nil {
		return m.Validation
	}
	return nil
}

// Schema of services.cfg
type ServicesCfg struct {
	// A list of all luci services. Should be sorted by id.
	Services         []*Service `protobuf:"bytes,1,rep,name=services" json:"services,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *ServicesCfg) Reset()                    { *m = ServicesCfg{} }
func (m *ServicesCfg) String() string            { return proto.CompactTextString(m) }
func (*ServicesCfg) ProtoMessage()               {}
func (*ServicesCfg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *ServicesCfg) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

// Schema of acl.cfg file.
type AclCfg struct {
	// Name of the group that has access to all projects/* config sets.
	// Only trusted services should be in this group.
	ProjectAccessGroup *string `protobuf:"bytes,2,opt,name=project_access_group,json=projectAccessGroup" json:"project_access_group,omitempty"`
	// Name of the group that has admin access to the app.
	AdminGroup *string `protobuf:"bytes,3,opt,name=admin_group,json=adminGroup" json:"admin_group,omitempty"`
	// Name of the group that can access configs by hash.
	ConfigGetByHashGroup *string `protobuf:"bytes,4,opt,name=config_get_by_hash_group,json=configGetByHashGroup" json:"config_get_by_hash_group,omitempty"`
	// Name of the group that may call validation API.
	ValidationGroup  *string `protobuf:"bytes,5,opt,name=validation_group,json=validationGroup" json:"validation_group,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AclCfg) Reset()                    { *m = AclCfg{} }
func (m *AclCfg) String() string            { return proto.CompactTextString(m) }
func (*AclCfg) ProtoMessage()               {}
func (*AclCfg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *AclCfg) GetProjectAccessGroup() string {
	if m != nil && m.ProjectAccessGroup != nil {
		return *m.ProjectAccessGroup
	}
	return ""
}

func (m *AclCfg) GetAdminGroup() string {
	if m != nil && m.AdminGroup != nil {
		return *m.AdminGroup
	}
	return ""
}

func (m *AclCfg) GetConfigGetByHashGroup() string {
	if m != nil && m.ConfigGetByHashGroup != nil {
		return *m.ConfigGetByHashGroup
	}
	return ""
}

func (m *AclCfg) GetValidationGroup() string {
	if m != nil && m.ValidationGroup != nil {
		return *m.ValidationGroup
	}
	return ""
}

// Schema for import.cfg. It specified how to import configuration files from
// external sources.
type ImportCfg struct {
	// Configuration of import from Gitiles repositories.
	Gitiles          *ImportCfg_Gitiles `protobuf:"bytes,1,opt,name=gitiles" json:"gitiles,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *ImportCfg) Reset()                    { *m = ImportCfg{} }
func (m *ImportCfg) String() string            { return proto.CompactTextString(m) }
func (*ImportCfg) ProtoMessage()               {}
func (*ImportCfg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *ImportCfg) GetGitiles() *ImportCfg_Gitiles {
	if m != nil {
		return m.Gitiles
	}
	return nil
}

type ImportCfg_Gitiles struct {
	// Request timeout in seconds when requesting commit log.
	FetchLogDeadline *int32 `protobuf:"varint,1,opt,name=fetch_log_deadline,json=fetchLogDeadline" json:"fetch_log_deadline,omitempty"`
	// Request timeout in seconds when requesting directory archive.
	FetchArchiveDeadline *int32 `protobuf:"varint,2,opt,name=fetch_archive_deadline,json=fetchArchiveDeadline" json:"fetch_archive_deadline,omitempty"`
	// Default ref for project configs.
	ProjectConfigDefaultRef *string `protobuf:"bytes,3,opt,name=project_config_default_ref,json=projectConfigDefaultRef" json:"project_config_default_ref,omitempty"`
	// Default directory for project configs.
	ProjectConfigDefaultPath *string `protobuf:"bytes,4,opt,name=project_config_default_path,json=projectConfigDefaultPath" json:"project_config_default_path,omitempty"`
	// Default directory for ref configs.
	RefConfigDefaultPath *string `protobuf:"bytes,5,opt,name=ref_config_default_path,json=refConfigDefaultPath" json:"ref_config_default_path,omitempty"`
	XXX_unrecognized     []byte  `json:"-"`
}

func (m *ImportCfg_Gitiles) Reset()                    { *m = ImportCfg_Gitiles{} }
func (m *ImportCfg_Gitiles) String() string            { return proto.CompactTextString(m) }
func (*ImportCfg_Gitiles) ProtoMessage()               {}
func (*ImportCfg_Gitiles) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7, 0} }

func (m *ImportCfg_Gitiles) GetFetchLogDeadline() int32 {
	if m != nil && m.FetchLogDeadline != nil {
		return *m.FetchLogDeadline
	}
	return 0
}

func (m *ImportCfg_Gitiles) GetFetchArchiveDeadline() int32 {
	if m != nil && m.FetchArchiveDeadline != nil {
		return *m.FetchArchiveDeadline
	}
	return 0
}

func (m *ImportCfg_Gitiles) GetProjectConfigDefaultRef() string {
	if m != nil && m.ProjectConfigDefaultRef != nil {
		return *m.ProjectConfigDefaultRef
	}
	return ""
}

func (m *ImportCfg_Gitiles) GetProjectConfigDefaultPath() string {
	if m != nil && m.ProjectConfigDefaultPath != nil {
		return *m.ProjectConfigDefaultPath
	}
	return ""
}

func (m *ImportCfg_Gitiles) GetRefConfigDefaultPath() string {
	if m != nil && m.RefConfigDefaultPath != nil {
		return *m.RefConfigDefaultPath
	}
	return ""
}

// Schema of schemas.cfg
type SchemasCfg struct {
	// List of known schemas. They are available at /schemas/<name> as a short
	// mutable link.
	Schemas          []*SchemasCfg_Schema `protobuf:"bytes,1,rep,name=schemas" json:"schemas,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *SchemasCfg) Reset()                    { *m = SchemasCfg{} }
func (m *SchemasCfg) String() string            { return proto.CompactTextString(m) }
func (*SchemasCfg) ProtoMessage()               {}
func (*SchemasCfg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *SchemasCfg) GetSchemas() []*SchemasCfg_Schema {
	if m != nil {
		return m.Schemas
	}
	return nil
}

type SchemasCfg_Schema struct {
	// Name of schema.
	// For service configs, "<config_set>:<path>"
	// For project configs, "projects:<path>"
	// For ref configs, "projects/refs:<path>"
	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// URL to the schema definition, e.g. to a .proto file in a repository.
	Url              *string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SchemasCfg_Schema) Reset()                    { *m = SchemasCfg_Schema{} }
func (m *SchemasCfg_Schema) String() string            { return proto.CompactTextString(m) }
func (*SchemasCfg_Schema) ProtoMessage()               {}
func (*SchemasCfg_Schema) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8, 0} }

func (m *SchemasCfg_Schema) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *SchemasCfg_Schema) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

// Defines a pattern of a config identity. Both config_set and path must
// match.
type ConfigPattern struct {
	// A string pattern for config_set.
	ConfigSet *string `protobuf:"bytes,1,opt,name=config_set,json=configSet" json:"config_set,omitempty"`
	// A string pattern for config file path.
	Path             *string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ConfigPattern) Reset()                    { *m = ConfigPattern{} }
func (m *ConfigPattern) String() string            { return proto.CompactTextString(m) }
func (*ConfigPattern) ProtoMessage()               {}
func (*ConfigPattern) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *ConfigPattern) GetConfigSet() string {
	if m != nil && m.ConfigSet != nil {
		return *m.ConfigSet
	}
	return ""
}

func (m *ConfigPattern) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

// Describes what configuration can be validated and how to validate them.
type Validator struct {
	// A list of configuration patterns that this validator can validate.
	Patterns []*ConfigPattern `protobuf:"bytes,1,rep,name=patterns" json:"patterns,omitempty"`
	// URL of a validation endpoint. The config service will send an HTTP POST
	// request to the endpoint, where body is JSON-encoded
	// ValidationRequestMessage. The endpoint is expected to respond with
	// HTTP status 200 and JSON-encoded ValidationResponseMessage.
	Url              *string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Validator) Reset()                    { *m = Validator{} }
func (m *Validator) String() string            { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()               {}
func (*Validator) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *Validator) GetPatterns() []*ConfigPattern {
	if m != nil {
		return m.Patterns
	}
	return nil
}

func (m *Validator) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

// This message is used only in JSON form. It is sent as request body to an
// external validation endpoint in order to validate a config.
type ValidationRequestMessage struct {
	// Config set of the config file to validate.
	ConfigSet *string `protobuf:"bytes,1,opt,name=config_set,json=configSet" json:"config_set,omitempty"`
	// Path of the config file to validate.
	Path *string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	// Base64-encoded contents of the file.
	Content          []byte `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ValidationRequestMessage) Reset()                    { *m = ValidationRequestMessage{} }
func (m *ValidationRequestMessage) String() string            { return proto.CompactTextString(m) }
func (*ValidationRequestMessage) ProtoMessage()               {}
func (*ValidationRequestMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *ValidationRequestMessage) GetConfigSet() string {
	if m != nil && m.ConfigSet != nil {
		return *m.ConfigSet
	}
	return ""
}

func (m *ValidationRequestMessage) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *ValidationRequestMessage) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

// This message is used only in JSON form. It is expected from an external
// validation endpoint that validates a config.
type ValidationResponseMessage struct {
	// Errors, warnings and other information found during validation.
	// If at least one error is found, the config is considered invalid.
	Messages         []*ValidationResponseMessage_Message `protobuf:"bytes,1,rep,name=messages" json:"messages,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *ValidationResponseMessage) Reset()                    { *m = ValidationResponseMessage{} }
func (m *ValidationResponseMessage) String() string            { return proto.CompactTextString(m) }
func (*ValidationResponseMessage) ProtoMessage()               {}
func (*ValidationResponseMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *ValidationResponseMessage) GetMessages() []*ValidationResponseMessage_Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

// A message that explains why a config is valid or not.
type ValidationResponseMessage_Message struct {
	// Textual representation of the message.
	Text *string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	// If an error, a config is considered invalid. Defaults to INFO.
	Severity         *ValidationResponseMessage_Severity `protobuf:"varint,2,opt,name=severity,enum=config.ValidationResponseMessage_Severity" json:"severity,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (m *ValidationResponseMessage_Message) Reset()         { *m = ValidationResponseMessage_Message{} }
func (m *ValidationResponseMessage_Message) String() string { return proto.CompactTextString(m) }
func (*ValidationResponseMessage_Message) ProtoMessage()    {}
func (*ValidationResponseMessage_Message) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{12, 0}
}

func (m *ValidationResponseMessage_Message) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

func (m *ValidationResponseMessage_Message) GetSeverity() ValidationResponseMessage_Severity {
	if m != nil && m.Severity != nil {
		return *m.Severity
	}
	return ValidationResponseMessage_DEBUG
}

func init() {
	proto.RegisterType((*ConfigSetLocation)(nil), "config.ConfigSetLocation")
	proto.RegisterType((*Project)(nil), "config.Project")
	proto.RegisterType((*ProjectsCfg)(nil), "config.ProjectsCfg")
	proto.RegisterType((*Service)(nil), "config.Service")
	proto.RegisterType((*ServiceDynamicMetadata)(nil), "config.ServiceDynamicMetadata")
	proto.RegisterType((*ServicesCfg)(nil), "config.ServicesCfg")
	proto.RegisterType((*AclCfg)(nil), "config.AclCfg")
	proto.RegisterType((*ImportCfg)(nil), "config.ImportCfg")
	proto.RegisterType((*ImportCfg_Gitiles)(nil), "config.ImportCfg.Gitiles")
	proto.RegisterType((*SchemasCfg)(nil), "config.SchemasCfg")
	proto.RegisterType((*SchemasCfg_Schema)(nil), "config.SchemasCfg.Schema")
	proto.RegisterType((*ConfigPattern)(nil), "config.ConfigPattern")
	proto.RegisterType((*Validator)(nil), "config.Validator")
	proto.RegisterType((*ValidationRequestMessage)(nil), "config.ValidationRequestMessage")
	proto.RegisterType((*ValidationResponseMessage)(nil), "config.ValidationResponseMessage")
	proto.RegisterType((*ValidationResponseMessage_Message)(nil), "config.ValidationResponseMessage.Message")
	proto.RegisterEnum("config.ConfigSetLocation_StorageType", ConfigSetLocation_StorageType_name, ConfigSetLocation_StorageType_value)
	proto.RegisterEnum("config.ValidationResponseMessage_Severity", ValidationResponseMessage_Severity_name, ValidationResponseMessage_Severity_value)
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/common/proto/config/service_config.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 876 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdf, 0x6e, 0x1b, 0xc5,
	0x17, 0xfe, 0xd9, 0x4e, 0x62, 0xfb, 0x38, 0xbf, 0x64, 0x3b, 0x32, 0xa9, 0x1b, 0x04, 0x84, 0x95,
	0x2a, 0xb9, 0x80, 0x6c, 0x1a, 0xfe, 0x5c, 0x80, 0x7a, 0xe1, 0x24, 0xae, 0x6b, 0x29, 0x4d, 0xa3,
	0x71, 0x52, 0xae, 0xd0, 0x6a, 0x58, 0x1f, 0xef, 0x2e, 0xda, 0xdd, 0xd9, 0xce, 0x8c, 0x0d, 0x7e,
	0x0b, 0xae, 0xb8, 0xe6, 0x11, 0x78, 0x02, 0x5e, 0x86, 0x17, 0x41, 0x3b, 0x7f, 0xd6, 0x29, 0x4e,
	0x85, 0xe0, 0xca, 0x73, 0xce, 0xf9, 0xbe, 0x73, 0xbe, 0xf9, 0x76, 0x3c, 0x03, 0xcf, 0x22, 0x3e,
	0x08, 0x63, 0xc1, 0xb3, 0x64, 0x99, 0x0d, 0xb8, 0x88, 0x86, 0xe9, 0x32, 0x4c, 0x86, 0x21, 0xcf,
	0x32, 0x9e, 0x0f, 0x0b, 0xc1, 0x15, 0x1f, 0x86, 0x3c, 0x5f, 0x24, 0xd1, 0x50, 0xa2, 0x58, 0x25,
	0x21, 0x06, 0x26, 0x1c, 0xe8, 0x1a, 0xd9, 0x33, 0x91, 0xff, 0x6b, 0x0d, 0x1e, 0x9c, 0xeb, 0xe5,
	0x0c, 0xd5, 0x25, 0x0f, 0x99, 0x4a, 0x78, 0x4e, 0x3c, 0x68, 0x2c, 0x45, 0xda, 0xab, 0x9d, 0xd4,
	0xfa, 0x6d, 0x5a, 0x2e, 0xc9, 0x0b, 0xd8, 0x97, 0x8a, 0x0b, 0x16, 0x61, 0xa0, 0xd6, 0x05, 0xf6,
	0xea, 0x27, 0xb5, 0xfe, 0xc1, 0xe9, 0xe3, 0x81, 0x6d, 0xba, 0xd5, 0x62, 0x30, 0x33, 0xe8, 0x9b,
	0x75, 0x81, 0xb4, 0x23, 0x37, 0x81, 0xff, 0x18, 0x3a, 0x77, 0x6a, 0xa4, 0x0d, 0xbb, 0xb7, 0x57,
	0xb3, 0xf1, 0x8d, 0xf7, 0x3f, 0xd2, 0x81, 0xe6, 0x64, 0x7a, 0x33, 0xbd, 0x1c, 0xcf, 0xbc, 0x9a,
	0xff, 0x3d, 0x34, 0xaf, 0x05, 0xff, 0x11, 0x43, 0x45, 0x0e, 0xa0, 0x9e, 0xcc, 0xad, 0x98, 0x7a,
	0x32, 0x27, 0x67, 0x70, 0x68, 0xc6, 0x06, 0xa9, 0x9d, 0xa6, 0xe5, 0x74, 0x4e, 0x1f, 0xbd, 0x53,
	0x0e, 0x3d, 0x30, 0x15, 0x17, 0xfb, 0xdf, 0x40, 0xc7, 0xb6, 0x97, 0xe7, 0x8b, 0x88, 0x7c, 0x0a,
	0xad, 0xc2, 0x86, 0xbd, 0xda, 0x49, 0xa3, 0xdf, 0x39, 0x3d, 0x74, 0xbd, 0x2c, 0x8c, 0x56, 0x00,
	0xff, 0xf7, 0x1a, 0x34, 0x67, 0xc6, 0xd4, 0x2d, 0x6d, 0x47, 0xb0, 0xc7, 0x7f, 0xca, 0x51, 0xc8,
	0x5e, 0xfd, 0xa4, 0xd1, 0x6f, 0x53, 0x1b, 0xdd, 0xa7, 0xb9, 0xf1, 0x2f, 0x35, 0x93, 0x8f, 0x61,
	0x3f, 0x43, 0xc5, 0xe6, 0x4c, 0xb1, 0xa0, 0xfc, 0x3c, 0x3b, 0x7a, 0x6a, 0xc7, 0xe5, 0x6e, 0x45,
	0x5a, 0x8e, 0x67, 0x61, 0x88, 0x52, 0xf6, 0x76, 0xcd, 0x78, 0x13, 0xf9, 0x08, 0x47, 0x56, 0xf1,
	0xc5, 0x3a, 0x67, 0x59, 0x12, 0xbe, 0xb4, 0x24, 0xd2, 0x83, 0xe6, 0x0a, 0x85, 0x2c, 0x05, 0x99,
	0x5d, 0xb8, 0x90, 0x3c, 0x05, 0x58, 0xb1, 0x34, 0x99, 0xdf, 0x75, 0xf8, 0x81, 0x53, 0xfb, 0xda,
	0x54, 0xb8, 0xa0, 0x77, 0x40, 0xa5, 0xab, 0x76, 0x8c, 0x73, 0xd5, 0x1e, 0xbe, 0x2d, 0x57, 0x2d,
	0x8c, 0x56, 0x00, 0xff, 0x8f, 0x1a, 0xec, 0x8d, 0xc2, 0xb4, 0xe4, 0x7d, 0x0e, 0x5d, 0x6b, 0x76,
	0x60, 0xf4, 0x07, 0x91, 0xe0, 0xcb, 0x42, 0x6b, 0x68, 0x53, 0x62, 0x6b, 0x23, 0x5d, 0x9a, 0x94,
	0x15, 0xf2, 0x11, 0x74, 0xd8, 0x3c, 0x4b, 0x72, 0x0b, 0x6c, 0x68, 0x20, 0xe8, 0x94, 0x01, 0x7c,
	0x0d, 0x3d, 0xeb, 0x7f, 0x84, 0x2a, 0xf8, 0x61, 0x1d, 0xc4, 0x4c, 0xc6, 0x16, 0x6d, 0x7c, 0xec,
	0x9a, 0xfa, 0x04, 0xd5, 0xd9, 0xfa, 0x05, 0x93, 0xb1, 0xe1, 0x3d, 0x01, 0x6f, 0xb3, 0x3f, 0x8b,
	0xdf, 0xd5, 0xf8, 0xc3, 0x4d, 0x5e, 0x43, 0xfd, 0x3f, 0xeb, 0xd0, 0x9e, 0x66, 0x05, 0x17, 0xaa,
	0xdc, 0xc3, 0x17, 0xd0, 0x8c, 0x12, 0x95, 0xa4, 0x7a, 0xeb, 0x6f, 0x7d, 0xe8, 0x0a, 0x33, 0x98,
	0x18, 0x00, 0x75, 0xc8, 0xe3, 0xdf, 0xea, 0xd0, 0xb4, 0x49, 0xf2, 0x19, 0x90, 0x05, 0xaa, 0x30,
	0x0e, 0x52, 0x1e, 0x05, 0x73, 0x64, 0xf3, 0x34, 0xc9, 0x51, 0xf7, 0xda, 0xa5, 0x9e, 0xae, 0x5c,
	0xf2, 0xe8, 0xc2, 0xe6, 0xc9, 0x97, 0x70, 0x64, 0xd0, 0x4c, 0x84, 0x71, 0xb2, 0xc2, 0x0d, 0xa3,
	0xae, 0x19, 0x5d, 0x5d, 0x1d, 0x99, 0x62, 0xc5, 0xfa, 0x16, 0x8e, 0x9d, 0xd1, 0xd6, 0x9d, 0x39,
	0x2e, 0xd8, 0x32, 0x55, 0x81, 0xc0, 0x85, 0x75, 0xf1, 0xa1, 0x45, 0x98, 0xf3, 0x79, 0x61, 0xea,
	0x14, 0x17, 0xe4, 0x19, 0xbc, 0xff, 0x0e, 0x72, 0xc1, 0x54, 0x6c, 0x5d, 0xed, 0xdd, 0xc7, 0xbe,
	0x66, 0x2a, 0x26, 0x5f, 0xc1, 0x43, 0x81, 0x8b, 0x7b, 0xa9, 0xc6, 0xe0, 0xae, 0xc0, 0xc5, 0x16,
	0xcd, 0x7f, 0x03, 0x30, 0x0b, 0x63, 0xcc, 0x98, 0xb4, 0x2e, 0x4b, 0x13, 0xd9, 0x03, 0x56, 0xb9,
	0xbc, 0x01, 0xd9, 0x25, 0x75, 0xc8, 0xe3, 0x01, 0xec, 0x99, 0x14, 0x21, 0xb0, 0x93, 0xb3, 0x0c,
	0xed, 0xc9, 0xd7, 0x6b, 0x77, 0xf7, 0xd5, 0xab, 0xbb, 0xcf, 0x3f, 0x83, 0xff, 0x1b, 0x1d, 0xd7,
	0x4c, 0x29, 0x14, 0x39, 0xf9, 0x00, 0xc0, 0xca, 0x96, 0xa8, 0x2c, 0xb9, 0x1d, 0xba, 0xff, 0x6f,
	0xd9, 0x55, 0x6f, 0xc3, 0xb4, 0xd0, 0x6b, 0xff, 0x1a, 0xda, 0xd5, 0x5f, 0x86, 0x3c, 0x85, 0x56,
	0x61, 0x5a, 0x39, 0xd9, 0xef, 0xbd, 0x7d, 0x0b, 0xd8, 0x41, 0xb4, 0x82, 0xdd, 0xa3, 0x2a, 0x82,
	0xde, 0xeb, 0xea, 0x04, 0x52, 0x7c, 0xb3, 0x44, 0xa9, 0x5e, 0xa2, 0x94, 0x2c, 0xc2, 0xff, 0x20,
	0xb0, 0xbc, 0x07, 0x42, 0x9e, 0x2b, 0xcc, 0x95, 0xfe, 0xee, 0xfb, 0xd4, 0x85, 0xfe, 0x2f, 0x75,
	0x78, 0x74, 0x77, 0x92, 0x2c, 0x78, 0x2e, 0xd1, 0x8d, 0x1a, 0x43, 0x2b, 0x33, 0x4b, 0xb7, 0x97,
	0x27, 0x7f, 0xbb, 0x23, 0xb6, 0x49, 0x03, 0xfb, 0x4b, 0x2b, 0xea, 0x31, 0x42, 0xd3, 0x75, 0x24,
	0xb0, 0xa3, 0xf0, 0x67, 0x27, 0x5b, 0xaf, 0xc9, 0xf3, 0xf2, 0x26, 0x59, 0xa1, 0x48, 0xd4, 0xda,
	0x3e, 0x3d, 0x9f, 0xfc, 0xf3, 0x94, 0x99, 0x65, 0xd0, 0x8a, 0xeb, 0x8f, 0xa1, 0xe5, 0xb2, 0xe5,
	0xcb, 0x73, 0x31, 0x3e, 0xbb, 0x9d, 0x78, 0x40, 0x5a, 0xb0, 0x33, 0xbd, 0x7a, 0xfe, 0xca, 0xeb,
	0x96, 0x6f, 0xd0, 0x77, 0x23, 0x7a, 0x35, 0xbd, 0x9a, 0x78, 0x1f, 0x96, 0x88, 0x31, 0xa5, 0xaf,
	0xa8, 0xd7, 0x27, 0xfb, 0xd0, 0x3a, 0xa7, 0xd3, 0x9b, 0xe9, 0xf9, 0xe8, 0xd2, 0x3b, 0xfd, 0x2b,
	0x00, 0x00, 0xff, 0xff, 0xd8, 0x41, 0xf7, 0x15, 0x7d, 0x07, 0x00, 0x00,
}