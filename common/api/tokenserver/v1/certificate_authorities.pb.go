// Code generated by protoc-gen-go.
// source: certificate_authorities.proto
// DO NOT EDIT!

/*
Package tokenserver is a generated protocol buffer package.

It is generated from these files:
	certificate_authorities.proto
	config.proto
	service_account.proto
	service_accounts.proto
	token_minter.proto
	tokens.proto

It has these top-level messages:
	ImportConfigRequest
	ImportConfigResponse
	FetchCRLRequest
	FetchCRLResponse
	ListCAsResponse
	GetCAStatusRequest
	GetCAStatusResponse
	IsRevokedCertRequest
	IsRevokedCertResponse
	CheckCertificateRequest
	CheckCertificateResponse
	CRLStatus
	TokenServerConfig
	CertificateAuthorityConfig
	DomainConfig
	ServiceAccount
	CreateServiceAccountRequest
	CreateServiceAccountResponse
	MintAccessTokenRequest
	MintAccessTokenResponse
	MintTokenRequest
	TokenRequest
	MintTokenResponse
	OAuth2AccessToken
*/
package tokenserver

import prpccommon "github.com/luci/luci-go/common/prpc"
import prpc "github.com/luci/luci-go/server/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/luci/luci-go/common/proto/google"
import google_protobuf1 "github.com/luci/luci-go/common/proto/google"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// ImportConfigRequest is passed to ImportConfig.
type ImportConfigRequest struct {
	// DevConfig is mapping of {config file name -> config file body}.
	//
	// It is used only on devserver to import some mock config in integration
	// tests. Ignored completely in prod.
	DevConfig map[string]string `protobuf:"bytes,1,rep,name=dev_config,json=devConfig" json:"dev_config,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ImportConfigRequest) Reset()                    { *m = ImportConfigRequest{} }
func (m *ImportConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*ImportConfigRequest) ProtoMessage()               {}
func (*ImportConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ImportConfigRequest) GetDevConfig() map[string]string {
	if m != nil {
		return m.DevConfig
	}
	return nil
}

// ImportConfigResponse is returned by ImportConfig on success.
type ImportConfigResponse struct {
	Revision string `protobuf:"bytes,1,opt,name=revision" json:"revision,omitempty"`
}

func (m *ImportConfigResponse) Reset()                    { *m = ImportConfigResponse{} }
func (m *ImportConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*ImportConfigResponse) ProtoMessage()               {}
func (*ImportConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// FetchCRLRequest identifies a name of CA to fetch CRL for.
type FetchCRLRequest struct {
	Cn    string `protobuf:"bytes,1,opt,name=cn" json:"cn,omitempty"`
	Force bool   `protobuf:"varint,2,opt,name=force" json:"force,omitempty"`
}

func (m *FetchCRLRequest) Reset()                    { *m = FetchCRLRequest{} }
func (m *FetchCRLRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchCRLRequest) ProtoMessage()               {}
func (*FetchCRLRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// FetchCRLResponse is returned by FetchCRL.
type FetchCRLResponse struct {
	CrlStatus *CRLStatus `protobuf:"bytes,1,opt,name=crl_status,json=crlStatus" json:"crl_status,omitempty"`
}

func (m *FetchCRLResponse) Reset()                    { *m = FetchCRLResponse{} }
func (m *FetchCRLResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchCRLResponse) ProtoMessage()               {}
func (*FetchCRLResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FetchCRLResponse) GetCrlStatus() *CRLStatus {
	if m != nil {
		return m.CrlStatus
	}
	return nil
}

// ListCAsResponse is returned by ListCAs.
type ListCAsResponse struct {
	Cn []string `protobuf:"bytes,1,rep,name=cn" json:"cn,omitempty"`
}

func (m *ListCAsResponse) Reset()                    { *m = ListCAsResponse{} }
func (m *ListCAsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListCAsResponse) ProtoMessage()               {}
func (*ListCAsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// GetCAStatusRequest identifies a name of CA to fetch.
type GetCAStatusRequest struct {
	Cn string `protobuf:"bytes,1,opt,name=cn" json:"cn,omitempty"`
}

func (m *GetCAStatusRequest) Reset()                    { *m = GetCAStatusRequest{} }
func (m *GetCAStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*GetCAStatusRequest) ProtoMessage()               {}
func (*GetCAStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// GetCAStatusResponse is returned by GetCAStatus method.
//
// If requested CA doesn't exist, all fields are empty.
type GetCAStatusResponse struct {
	Config     *CertificateAuthorityConfig `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	Cert       string                      `protobuf:"bytes,2,opt,name=cert" json:"cert,omitempty"`
	Removed    bool                        `protobuf:"varint,3,opt,name=removed" json:"removed,omitempty"`
	Ready      bool                        `protobuf:"varint,4,opt,name=ready" json:"ready,omitempty"`
	AddedRev   string                      `protobuf:"bytes,5,opt,name=added_rev,json=addedRev" json:"added_rev,omitempty"`
	UpdatedRev string                      `protobuf:"bytes,6,opt,name=updated_rev,json=updatedRev" json:"updated_rev,omitempty"`
	RemovedRev string                      `protobuf:"bytes,7,opt,name=removed_rev,json=removedRev" json:"removed_rev,omitempty"`
	CrlStatus  *CRLStatus                  `protobuf:"bytes,8,opt,name=crl_status,json=crlStatus" json:"crl_status,omitempty"`
}

func (m *GetCAStatusResponse) Reset()                    { *m = GetCAStatusResponse{} }
func (m *GetCAStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*GetCAStatusResponse) ProtoMessage()               {}
func (*GetCAStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetCAStatusResponse) GetConfig() *CertificateAuthorityConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *GetCAStatusResponse) GetCrlStatus() *CRLStatus {
	if m != nil {
		return m.CrlStatus
	}
	return nil
}

// IsRevokedCertRequest contains a name of the CA and a cert serial number.
type IsRevokedCertRequest struct {
	Ca string `protobuf:"bytes,1,opt,name=ca" json:"ca,omitempty"`
	Sn string `protobuf:"bytes,2,opt,name=sn" json:"sn,omitempty"`
}

func (m *IsRevokedCertRequest) Reset()                    { *m = IsRevokedCertRequest{} }
func (m *IsRevokedCertRequest) String() string            { return proto.CompactTextString(m) }
func (*IsRevokedCertRequest) ProtoMessage()               {}
func (*IsRevokedCertRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// IsRevokedCertResponse is returned by IsRevokedCert
type IsRevokedCertResponse struct {
	Revoked bool `protobuf:"varint,1,opt,name=revoked" json:"revoked,omitempty"`
}

func (m *IsRevokedCertResponse) Reset()                    { *m = IsRevokedCertResponse{} }
func (m *IsRevokedCertResponse) String() string            { return proto.CompactTextString(m) }
func (*IsRevokedCertResponse) ProtoMessage()               {}
func (*IsRevokedCertResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// CheckCertificateRequest contains a pem encoded certificate to check.
type CheckCertificateRequest struct {
	CertPem string `protobuf:"bytes,1,opt,name=cert_pem,json=certPem" json:"cert_pem,omitempty"`
}

func (m *CheckCertificateRequest) Reset()                    { *m = CheckCertificateRequest{} }
func (m *CheckCertificateRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckCertificateRequest) ProtoMessage()               {}
func (*CheckCertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// CheckCertificateResponse is returned by CheckCertificate.
type CheckCertificateResponse struct {
	IsValid       bool   `protobuf:"varint,1,opt,name=is_valid,json=isValid" json:"is_valid,omitempty"`
	InvalidReason string `protobuf:"bytes,2,opt,name=invalid_reason,json=invalidReason" json:"invalid_reason,omitempty"`
}

func (m *CheckCertificateResponse) Reset()                    { *m = CheckCertificateResponse{} }
func (m *CheckCertificateResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckCertificateResponse) ProtoMessage()               {}
func (*CheckCertificateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// CRLStatus describes the latest known state of imported CRL.
type CRLStatus struct {
	LastUpdateTime    *google_protobuf1.Timestamp `protobuf:"bytes,1,opt,name=last_update_time,json=lastUpdateTime" json:"last_update_time,omitempty"`
	LastFetchTime     *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=last_fetch_time,json=lastFetchTime" json:"last_fetch_time,omitempty"`
	LastFetchEtag     string                      `protobuf:"bytes,3,opt,name=last_fetch_etag,json=lastFetchEtag" json:"last_fetch_etag,omitempty"`
	RevokedCertsCount int64                       `protobuf:"varint,4,opt,name=revoked_certs_count,json=revokedCertsCount" json:"revoked_certs_count,omitempty"`
}

func (m *CRLStatus) Reset()                    { *m = CRLStatus{} }
func (m *CRLStatus) String() string            { return proto.CompactTextString(m) }
func (*CRLStatus) ProtoMessage()               {}
func (*CRLStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *CRLStatus) GetLastUpdateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func (m *CRLStatus) GetLastFetchTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.LastFetchTime
	}
	return nil
}

func init() {
	proto.RegisterType((*ImportConfigRequest)(nil), "tokenserver.ImportConfigRequest")
	proto.RegisterType((*ImportConfigResponse)(nil), "tokenserver.ImportConfigResponse")
	proto.RegisterType((*FetchCRLRequest)(nil), "tokenserver.FetchCRLRequest")
	proto.RegisterType((*FetchCRLResponse)(nil), "tokenserver.FetchCRLResponse")
	proto.RegisterType((*ListCAsResponse)(nil), "tokenserver.ListCAsResponse")
	proto.RegisterType((*GetCAStatusRequest)(nil), "tokenserver.GetCAStatusRequest")
	proto.RegisterType((*GetCAStatusResponse)(nil), "tokenserver.GetCAStatusResponse")
	proto.RegisterType((*IsRevokedCertRequest)(nil), "tokenserver.IsRevokedCertRequest")
	proto.RegisterType((*IsRevokedCertResponse)(nil), "tokenserver.IsRevokedCertResponse")
	proto.RegisterType((*CheckCertificateRequest)(nil), "tokenserver.CheckCertificateRequest")
	proto.RegisterType((*CheckCertificateResponse)(nil), "tokenserver.CheckCertificateResponse")
	proto.RegisterType((*CRLStatus)(nil), "tokenserver.CRLStatus")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for CertificateAuthorities service

type CertificateAuthoritiesClient interface {
	// ImportConfig makes the server read its config from luci-config right now.
	//
	// Note that regularly configs are read in background each 5 min. ImportConfig
	// can be used to force config reread immediately. It will block until configs
	// are read.
	ImportConfig(ctx context.Context, in *ImportConfigRequest, opts ...grpc.CallOption) (*ImportConfigResponse, error)
	// FetchCRL makes the server fetch a CRL for some CA.
	FetchCRL(ctx context.Context, in *FetchCRLRequest, opts ...grpc.CallOption) (*FetchCRLResponse, error)
	// ListCAs returns a list of Common Names of registered CAs.
	ListCAs(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ListCAsResponse, error)
	// GetCAStatus returns configuration of some CA defined in the config.
	GetCAStatus(ctx context.Context, in *GetCAStatusRequest, opts ...grpc.CallOption) (*GetCAStatusResponse, error)
	// IsRevokedCert says whether a certificate serial number is in the CRL.
	IsRevokedCert(ctx context.Context, in *IsRevokedCertRequest, opts ...grpc.CallOption) (*IsRevokedCertResponse, error)
	// CheckCertificate says whether a certificate is valid or not.
	CheckCertificate(ctx context.Context, in *CheckCertificateRequest, opts ...grpc.CallOption) (*CheckCertificateResponse, error)
}
type certificateAuthoritiesPRPCClient struct {
	client *prpccommon.Client
}

func NewCertificateAuthoritiesPRPCClient(client *prpccommon.Client) CertificateAuthoritiesClient {
	return &certificateAuthoritiesPRPCClient{client}
}

func (c *certificateAuthoritiesPRPCClient) ImportConfig(ctx context.Context, in *ImportConfigRequest, opts ...grpc.CallOption) (*ImportConfigResponse, error) {
	out := new(ImportConfigResponse)
	err := c.client.Call(ctx, "tokenserver.CertificateAuthorities", "ImportConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesPRPCClient) FetchCRL(ctx context.Context, in *FetchCRLRequest, opts ...grpc.CallOption) (*FetchCRLResponse, error) {
	out := new(FetchCRLResponse)
	err := c.client.Call(ctx, "tokenserver.CertificateAuthorities", "FetchCRL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesPRPCClient) ListCAs(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ListCAsResponse, error) {
	out := new(ListCAsResponse)
	err := c.client.Call(ctx, "tokenserver.CertificateAuthorities", "ListCAs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesPRPCClient) GetCAStatus(ctx context.Context, in *GetCAStatusRequest, opts ...grpc.CallOption) (*GetCAStatusResponse, error) {
	out := new(GetCAStatusResponse)
	err := c.client.Call(ctx, "tokenserver.CertificateAuthorities", "GetCAStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesPRPCClient) IsRevokedCert(ctx context.Context, in *IsRevokedCertRequest, opts ...grpc.CallOption) (*IsRevokedCertResponse, error) {
	out := new(IsRevokedCertResponse)
	err := c.client.Call(ctx, "tokenserver.CertificateAuthorities", "IsRevokedCert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesPRPCClient) CheckCertificate(ctx context.Context, in *CheckCertificateRequest, opts ...grpc.CallOption) (*CheckCertificateResponse, error) {
	out := new(CheckCertificateResponse)
	err := c.client.Call(ctx, "tokenserver.CertificateAuthorities", "CheckCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type certificateAuthoritiesClient struct {
	cc *grpc.ClientConn
}

func NewCertificateAuthoritiesClient(cc *grpc.ClientConn) CertificateAuthoritiesClient {
	return &certificateAuthoritiesClient{cc}
}

func (c *certificateAuthoritiesClient) ImportConfig(ctx context.Context, in *ImportConfigRequest, opts ...grpc.CallOption) (*ImportConfigResponse, error) {
	out := new(ImportConfigResponse)
	err := grpc.Invoke(ctx, "/tokenserver.CertificateAuthorities/ImportConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesClient) FetchCRL(ctx context.Context, in *FetchCRLRequest, opts ...grpc.CallOption) (*FetchCRLResponse, error) {
	out := new(FetchCRLResponse)
	err := grpc.Invoke(ctx, "/tokenserver.CertificateAuthorities/FetchCRL", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesClient) ListCAs(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*ListCAsResponse, error) {
	out := new(ListCAsResponse)
	err := grpc.Invoke(ctx, "/tokenserver.CertificateAuthorities/ListCAs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesClient) GetCAStatus(ctx context.Context, in *GetCAStatusRequest, opts ...grpc.CallOption) (*GetCAStatusResponse, error) {
	out := new(GetCAStatusResponse)
	err := grpc.Invoke(ctx, "/tokenserver.CertificateAuthorities/GetCAStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesClient) IsRevokedCert(ctx context.Context, in *IsRevokedCertRequest, opts ...grpc.CallOption) (*IsRevokedCertResponse, error) {
	out := new(IsRevokedCertResponse)
	err := grpc.Invoke(ctx, "/tokenserver.CertificateAuthorities/IsRevokedCert", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthoritiesClient) CheckCertificate(ctx context.Context, in *CheckCertificateRequest, opts ...grpc.CallOption) (*CheckCertificateResponse, error) {
	out := new(CheckCertificateResponse)
	err := grpc.Invoke(ctx, "/tokenserver.CertificateAuthorities/CheckCertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CertificateAuthorities service

type CertificateAuthoritiesServer interface {
	// ImportConfig makes the server read its config from luci-config right now.
	//
	// Note that regularly configs are read in background each 5 min. ImportConfig
	// can be used to force config reread immediately. It will block until configs
	// are read.
	ImportConfig(context.Context, *ImportConfigRequest) (*ImportConfigResponse, error)
	// FetchCRL makes the server fetch a CRL for some CA.
	FetchCRL(context.Context, *FetchCRLRequest) (*FetchCRLResponse, error)
	// ListCAs returns a list of Common Names of registered CAs.
	ListCAs(context.Context, *google_protobuf.Empty) (*ListCAsResponse, error)
	// GetCAStatus returns configuration of some CA defined in the config.
	GetCAStatus(context.Context, *GetCAStatusRequest) (*GetCAStatusResponse, error)
	// IsRevokedCert says whether a certificate serial number is in the CRL.
	IsRevokedCert(context.Context, *IsRevokedCertRequest) (*IsRevokedCertResponse, error)
	// CheckCertificate says whether a certificate is valid or not.
	CheckCertificate(context.Context, *CheckCertificateRequest) (*CheckCertificateResponse, error)
}

func RegisterCertificateAuthoritiesServer(s prpc.Registrar, srv CertificateAuthoritiesServer) {
	s.RegisterService(&_CertificateAuthorities_serviceDesc, srv)
}

func _CertificateAuthorities_ImportConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ImportConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CertificateAuthoritiesServer).ImportConfig(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _CertificateAuthorities_FetchCRL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(FetchCRLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CertificateAuthoritiesServer).FetchCRL(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _CertificateAuthorities_ListCAs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CertificateAuthoritiesServer).ListCAs(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _CertificateAuthorities_GetCAStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetCAStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CertificateAuthoritiesServer).GetCAStatus(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _CertificateAuthorities_IsRevokedCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(IsRevokedCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CertificateAuthoritiesServer).IsRevokedCert(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _CertificateAuthorities_CheckCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CheckCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CertificateAuthoritiesServer).CheckCertificate(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _CertificateAuthorities_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tokenserver.CertificateAuthorities",
	HandlerType: (*CertificateAuthoritiesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ImportConfig",
			Handler:    _CertificateAuthorities_ImportConfig_Handler,
		},
		{
			MethodName: "FetchCRL",
			Handler:    _CertificateAuthorities_FetchCRL_Handler,
		},
		{
			MethodName: "ListCAs",
			Handler:    _CertificateAuthorities_ListCAs_Handler,
		},
		{
			MethodName: "GetCAStatus",
			Handler:    _CertificateAuthorities_GetCAStatus_Handler,
		},
		{
			MethodName: "IsRevokedCert",
			Handler:    _CertificateAuthorities_IsRevokedCert_Handler,
		},
		{
			MethodName: "CheckCertificate",
			Handler:    _CertificateAuthorities_CheckCertificate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 758 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xdb, 0x4e, 0xdb, 0x40,
	0x10, 0x55, 0x12, 0x20, 0xc9, 0x04, 0x42, 0xba, 0x50, 0x9a, 0x9a, 0x22, 0xc0, 0x82, 0x96, 0x27,
	0xa3, 0xa6, 0x57, 0x55, 0x95, 0x2a, 0x1a, 0x28, 0x42, 0x42, 0x15, 0x32, 0x2d, 0x4f, 0x95, 0x2c,
	0x63, 0x6f, 0x82, 0x45, 0x7c, 0xe9, 0xee, 0xda, 0x52, 0x5e, 0xfb, 0x2b, 0xfd, 0x81, 0xfe, 0x58,
	0xff, 0xa1, 0x7b, 0xb3, 0x89, 0x43, 0x80, 0xbe, 0xed, 0xcc, 0x9c, 0xb9, 0xf8, 0x78, 0xe6, 0xc0,
	0x86, 0x87, 0x09, 0x0b, 0x06, 0x81, 0xe7, 0x32, 0xec, 0xb8, 0x29, 0xbb, 0x8a, 0x49, 0xc0, 0x02,
	0x4c, 0xad, 0x84, 0xc4, 0x2c, 0x46, 0x2d, 0x16, 0x5f, 0xe3, 0x88, 0x62, 0x92, 0x61, 0x62, 0xac,
	0x0f, 0xe3, 0x78, 0x38, 0xc2, 0xfb, 0x32, 0x74, 0x99, 0x0e, 0xf6, 0x71, 0x98, 0xb0, 0xb1, 0x42,
	0x1a, 0x9b, 0xd3, 0x41, 0x16, 0x84, 0x98, 0x32, 0x37, 0x4c, 0x34, 0x60, 0xd1, 0x8b, 0xa3, 0x41,
	0x30, 0x54, 0x96, 0xf9, 0xbb, 0x02, 0x2b, 0x27, 0x61, 0x12, 0x13, 0xd6, 0x97, 0x6e, 0x1b, 0xff,
	0x4c, 0x39, 0x1c, 0x7d, 0x05, 0xf0, 0x71, 0xe6, 0x28, 0x6c, 0xb7, 0xb2, 0x55, 0xdb, 0x6b, 0xf5,
	0xf6, 0xad, 0x89, 0x29, 0xac, 0x19, 0x59, 0xd6, 0x21, 0xce, 0x94, 0xe3, 0x28, 0x62, 0x64, 0x6c,
	0x37, 0xfd, 0xdc, 0x36, 0x3e, 0x42, 0xbb, 0x1c, 0x44, 0x1d, 0xa8, 0x5d, 0xe3, 0x31, 0x2f, 0x5d,
	0xd9, 0x6b, 0xda, 0xe2, 0x89, 0x56, 0x61, 0x3e, 0x73, 0x47, 0x29, 0xee, 0x56, 0xa5, 0x4f, 0x19,
	0x1f, 0xaa, 0xef, 0x2b, 0x66, 0x0f, 0x56, 0xcb, 0xed, 0x68, 0x12, 0xf3, 0x39, 0x90, 0x01, 0x0d,
	0x82, 0xb3, 0x80, 0x06, 0x71, 0xa4, 0x0b, 0x15, 0xb6, 0xf9, 0x0e, 0x96, 0xbf, 0x60, 0xe6, 0x5d,
	0xf5, 0xed, 0xd3, 0xfc, 0xa3, 0xda, 0x50, 0xf5, 0x72, 0x20, 0x7f, 0x89, 0x86, 0x83, 0x98, 0x78,
	0xaa, 0x61, 0xc3, 0x56, 0x86, 0x79, 0x02, 0x9d, 0x9b, 0x44, 0xdd, 0xe8, 0x0d, 0x80, 0x47, 0x46,
	0x0e, 0xe7, 0x91, 0xa5, 0x54, 0x56, 0x68, 0xf5, 0xd6, 0x4a, 0x74, 0x70, 0xf4, 0xb9, 0x8c, 0xda,
	0x4d, 0x8e, 0x54, 0x4f, 0x73, 0x1b, 0x96, 0x4f, 0x03, 0xca, 0xfa, 0x07, 0xb4, 0xa8, 0x94, 0xcf,
	0x50, 0x53, 0x33, 0x98, 0x3b, 0x80, 0x8e, 0x31, 0x47, 0xe8, 0xe4, 0xd9, 0x93, 0x9a, 0x7f, 0xaa,
	0xb0, 0x52, 0x82, 0xe9, 0x6a, 0x9f, 0x60, 0xa1, 0xf8, 0x45, 0x62, 0xa6, 0x17, 0xe5, 0x99, 0x6e,
	0x76, 0xea, 0x40, 0xaf, 0xd4, 0x58, 0x33, 0xa8, 0xd3, 0x10, 0x82, 0x39, 0xb1, 0x79, 0x9a, 0x72,
	0xf9, 0x46, 0x5d, 0xa8, 0x13, 0x1c, 0xc6, 0x19, 0xf6, 0xbb, 0x35, 0x49, 0x4c, 0x6e, 0x0a, 0xc2,
	0x08, 0x76, 0xfd, 0x71, 0x77, 0x4e, 0x11, 0x26, 0x0d, 0xb4, 0x0e, 0x4d, 0xd7, 0xf7, 0xb1, 0xef,
	0x70, 0xee, 0xbb, 0xf3, 0xea, 0x37, 0x48, 0x87, 0x8d, 0x33, 0xb4, 0x09, 0xad, 0x34, 0xf1, 0xf9,
	0x04, 0x2a, 0xbc, 0x20, 0xc3, 0xa0, 0x5d, 0x1a, 0xa0, 0xcb, 0x4b, 0x40, 0x5d, 0x01, 0xb4, 0x4b,
	0x00, 0xca, 0xdc, 0x37, 0xfe, 0x97, 0xfb, 0xb7, 0x7c, 0x67, 0x38, 0x51, 0x19, 0xc7, 0xf9, 0x82,
	0x88, 0x49, 0x6a, 0xdd, 0x82, 0x5a, 0x57, 0xd8, 0x34, 0xd2, 0xdf, 0xcf, 0x5f, 0xe6, 0x4b, 0x78,
	0x3c, 0x95, 0xa7, 0xb9, 0x96, 0xb4, 0x48, 0xb7, 0xcc, 0x96, 0xb4, 0x48, 0xd3, 0x7c, 0x0d, 0x4f,
	0xfa, 0x57, 0xd8, 0xbb, 0x9e, 0xe0, 0x3b, 0xef, 0xf6, 0x14, 0x1a, 0x82, 0x53, 0x27, 0xc1, 0xa1,
	0xee, 0x59, 0x17, 0xf6, 0x19, 0x0e, 0xcd, 0x1f, 0xd0, 0xbd, 0x9d, 0xa5, 0x7b, 0xf1, 0xb4, 0x80,
	0x3a, 0xfc, 0x00, 0x82, 0xa2, 0x59, 0x40, 0x2f, 0x84, 0x89, 0x76, 0xa1, 0x1d, 0x44, 0x32, 0xc2,
	0xf9, 0x72, 0x69, 0x9c, 0xcf, 0xbe, 0xa4, 0xbd, 0xb6, 0x74, 0x9a, 0x7f, 0x2b, 0xd0, 0x2c, 0x78,
	0x41, 0x87, 0xd0, 0x19, 0xb9, 0x94, 0x39, 0x8a, 0x77, 0x47, 0x68, 0x82, 0xde, 0x18, 0xc3, 0x52,
	0x82, 0x61, 0xe5, 0x82, 0x61, 0x7d, 0xcb, 0x05, 0xc3, 0x6e, 0x8b, 0x9c, 0xef, 0x32, 0x45, 0x38,
	0xd1, 0x67, 0x58, 0x96, 0x55, 0x06, 0xe2, 0x3c, 0x54, 0x91, 0xea, 0x83, 0x45, 0x96, 0x44, 0x8a,
	0x3c, 0x28, 0x59, 0xe3, 0x79, 0xa9, 0x06, 0x66, 0xee, 0x50, 0x2e, 0x59, 0x73, 0x02, 0x77, 0xc4,
	0x9d, 0xc8, 0x82, 0x15, 0x4d, 0xaf, 0x23, 0x08, 0xa3, 0x5c, 0x8a, 0xd2, 0x88, 0xc9, 0xc5, 0xab,
	0xd9, 0x8f, 0xc8, 0xcd, 0xff, 0xa1, 0x7d, 0x11, 0xe8, 0xfd, 0x9a, 0x83, 0xb5, 0x19, 0xfb, 0xce,
	0x25, 0x14, 0x9d, 0xc3, 0xe2, 0xa4, 0x7a, 0xa0, 0xad, 0x87, 0x74, 0xcc, 0xd8, 0xbe, 0x07, 0xa1,
	0xff, 0xd0, 0x31, 0x34, 0x72, 0x95, 0x40, 0xcf, 0x4a, 0xf0, 0x29, 0xd5, 0x31, 0x36, 0xee, 0x88,
	0x16, 0x27, 0x5c, 0xd7, 0x1a, 0x81, 0xd6, 0x6e, 0xd1, 0x78, 0x24, 0x94, 0xdd, 0x28, 0xd7, 0x9f,
	0x56, 0x94, 0x33, 0x68, 0x4d, 0x48, 0x03, 0xda, 0x2c, 0x81, 0x6f, 0x6b, 0x8b, 0xb1, 0x75, 0x37,
	0x40, 0x57, 0xbc, 0x80, 0xa5, 0xd2, 0x09, 0xa0, 0x29, 0x3e, 0x66, 0x9c, 0x95, 0x61, 0xde, 0x07,
	0xd1, 0x75, 0x1d, 0xe8, 0x4c, 0x6f, 0x3c, 0xda, 0x29, 0x5f, 0xf2, 0xec, 0x33, 0x32, 0x76, 0x1f,
	0x40, 0xa9, 0x06, 0x97, 0x0b, 0x92, 0xb8, 0x57, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x9a,
	0x1e, 0xe7, 0x4e, 0x07, 0x00, 0x00,
}
