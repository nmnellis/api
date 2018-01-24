// Code generated by protoc-gen-go. DO NOT EDIT.
// source: routing/v1alpha2/unsafe_envoy_config.proto

package v1alpha2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf2 "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UnsafeEnvoyConfig_Operation int32

const (
	UnsafeEnvoyConfig_INVALID_OP UnsafeEnvoyConfig_Operation = 0
	// Add a listener or cluster.
	UnsafeEnvoyConfig_ADD UnsafeEnvoyConfig_Operation = 1
	// Replace an existing listener or cluster.
	UnsafeEnvoyConfig_REPLACE UnsafeEnvoyConfig_Operation = 2
	// Update an existing listener or cluster. Any missing fields in the
	// configuration will be replaced with information computed by Istio.
	UnsafeEnvoyConfig_UPDATE UnsafeEnvoyConfig_Operation = 3
)

var UnsafeEnvoyConfig_Operation_name = map[int32]string{
	0: "INVALID_OP",
	1: "ADD",
	2: "REPLACE",
	3: "UPDATE",
}
var UnsafeEnvoyConfig_Operation_value = map[string]int32{
	"INVALID_OP": 0,
	"ADD":        1,
	"REPLACE":    2,
	"UPDATE":     3,
}

func (x UnsafeEnvoyConfig_Operation) String() string {
	return proto.EnumName(UnsafeEnvoyConfig_Operation_name, int32(x))
}
func (UnsafeEnvoyConfig_Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor4, []int{0, 0}
}

// Envoy-specific filters that can be used to enable customized processing
// in addition to the features exposed by the routing rule. This feature
// must be used with care, as incorrect configurations could potentially
// destabilize the entire mesh. This configuration is designed for an
// advanced audience with good familiarity with Envoy configuration, and
// Istio Pilot's internals. Common use cases for using the extensions
// include use of Envoy's Lua filters, in addition to enabling any
// proprietary filters developed inhouse (e.g, protocol decoders,
// geo-location filters, etc.).
//
// NOTE: Do not use this feature to enable Fault, Router (HTTP),
// Mongo/Tcp proxy (TCP) filters.
//
// The following example adds a Lua script that calls out to an external
// service, for all inbound calls to the reviews HTTP service on port 8080.
//
//     apiVersion: config.istio.io/v1alpha2
//     kind: UnsafeEnvoyConfig
//     metadata:
//       name: my-lua-script
//     spec:
//       configs:
//       - labels:
//           name: reviews # selects the pods where this config will be applied
//         operation: update # update existing listener
//         name: http_0.0.0.0:8080
//         envoyConfig:
//           "@type": type.googleapis.com/envoy.api.v2.Listener
//           filterChains:
//           - filters:
//             - name: envoy.lua
//               config:
//                 inline_code: |
//                   function envoy_on_request(request_handle)
//                     -- Make an HTTP call to an upstream host with the following headers, body, and timeout.
//                     local headers, body = request_handle:httpCall(
//                     "lua_cluster",
//                     {
//                      [":method"] = "POST",
//                      [":path"] = "/",
//                      [":authority"] = "lua_cluster"
//                     },
//                     "hello world", 5000)
//                   end
//       - labels:
//           name: reviews # selects the pods where this config will be applied
//         operation: add # add the lua_cluster to existing list of clusters
//         name: lua_cluster
//         envoyConfig:
//           "@type": type.googleapis.com/envoy.api.v2.Cluster
//           type: static
//           hosts:
//           - socketAddress:
//               protocol: tcp
//               address: 1.1.1.1
//               portValue: 5555
//
type UnsafeEnvoyConfig struct {
	Configs []*UnsafeEnvoyConfig_EnvoyConfig `protobuf:"bytes,1,rep,name=configs" json:"configs,omitempty"`
}

func (m *UnsafeEnvoyConfig) Reset()                    { *m = UnsafeEnvoyConfig{} }
func (m *UnsafeEnvoyConfig) String() string            { return proto.CompactTextString(m) }
func (*UnsafeEnvoyConfig) ProtoMessage()               {}
func (*UnsafeEnvoyConfig) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *UnsafeEnvoyConfig) GetConfigs() []*UnsafeEnvoyConfig_EnvoyConfig {
	if m != nil {
		return m.Configs
	}
	return nil
}

// Listener/cluster related configuration.
type UnsafeEnvoyConfig_EnvoyConfig struct {
	// REQUIRED. Add/Replace/Update operation
	Operation UnsafeEnvoyConfig_Operation `protobuf:"varint,1,opt,name=operation,enum=istio.routing.v1alpha2.UnsafeEnvoyConfig_Operation" json:"operation,omitempty"`
	// REQUIRED. Unique name associated with the listener or cluster. Istio
	// generated listeners are named using the format
	// (http|tcp)_(ip)_(port). Wildcards can be used when selecting
	// listeners. For example, to select all listeners on port 8080, use
	// *_*_8080.
	//
	// Istio generated clusters are named using the format
	// out.(serviceName)|(portName/portNumber)|subsetName for outbound
	// clusters (client side) and in.(portName/portNumber) for inbound
	// clusters (server side). Wildcards can be used when selecting
	// clusters. For example, to select all clusters for
	// foo.default.svc.cluster.local service, use
	// *.foo.default.svc.cluster.local|*|*
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Restrict the changes to a subset of pods/VMs in the mesh
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// REQUIRED. Envoy listener/cluster configuration. Refer to
	// https://github.com/envoyproxy/data-plane-api/ for details.
	EnvoyConfig *google_protobuf2.Any `protobuf:"bytes,4,opt,name=envoy_config,json=envoyConfig" json:"envoy_config,omitempty"`
}

func (m *UnsafeEnvoyConfig_EnvoyConfig) Reset()         { *m = UnsafeEnvoyConfig_EnvoyConfig{} }
func (m *UnsafeEnvoyConfig_EnvoyConfig) String() string { return proto.CompactTextString(m) }
func (*UnsafeEnvoyConfig_EnvoyConfig) ProtoMessage()    {}
func (*UnsafeEnvoyConfig_EnvoyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{0, 0}
}

func (m *UnsafeEnvoyConfig_EnvoyConfig) GetOperation() UnsafeEnvoyConfig_Operation {
	if m != nil {
		return m.Operation
	}
	return UnsafeEnvoyConfig_INVALID_OP
}

func (m *UnsafeEnvoyConfig_EnvoyConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UnsafeEnvoyConfig_EnvoyConfig) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *UnsafeEnvoyConfig_EnvoyConfig) GetEnvoyConfig() *google_protobuf2.Any {
	if m != nil {
		return m.EnvoyConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*UnsafeEnvoyConfig)(nil), "istio.routing.v1alpha2.UnsafeEnvoyConfig")
	proto.RegisterType((*UnsafeEnvoyConfig_EnvoyConfig)(nil), "istio.routing.v1alpha2.UnsafeEnvoyConfig.EnvoyConfig")
	proto.RegisterEnum("istio.routing.v1alpha2.UnsafeEnvoyConfig_Operation", UnsafeEnvoyConfig_Operation_name, UnsafeEnvoyConfig_Operation_value)
}

func init() { proto.RegisterFile("routing/v1alpha2/unsafe_envoy_config.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x18, 0xc5, 0x6f, 0x92, 0xde, 0x96, 0x7c, 0xb9, 0x94, 0xdc, 0xa1, 0x48, 0x2c, 0x88, 0xa1, 0xab,
	0xe0, 0x62, 0x82, 0x29, 0xe2, 0x1f, 0x70, 0x11, 0x9b, 0x2c, 0x0a, 0xc5, 0xd6, 0x60, 0x05, 0xdd,
	0x94, 0xa9, 0x4c, 0x63, 0x30, 0xce, 0x84, 0xfc, 0x29, 0xe4, 0xf9, 0x7c, 0x1a, 0xdf, 0x42, 0x3a,
	0x69, 0x6c, 0x50, 0x17, 0xe2, 0xee, 0x0c, 0xcc, 0xf9, 0x7d, 0x73, 0xce, 0x7c, 0x70, 0x94, 0xf2,
	0x22, 0x8f, 0x58, 0x68, 0xaf, 0x8f, 0x49, 0x9c, 0x3c, 0x11, 0xc7, 0x2e, 0x58, 0x46, 0x56, 0x74,
	0x41, 0xd9, 0x9a, 0x97, 0x8b, 0x47, 0xce, 0x56, 0x51, 0x88, 0x93, 0x94, 0xe7, 0x1c, 0xed, 0x45,
	0x59, 0x1e, 0x71, 0xbc, 0x75, 0xe0, 0xda, 0xd1, 0xdf, 0x0f, 0x39, 0x0f, 0x63, 0x6a, 0x8b, 0x5b,
	0xcb, 0x62, 0x65, 0x13, 0x56, 0x56, 0x96, 0xc1, 0x9b, 0x02, 0xff, 0xe7, 0x02, 0xe8, 0x6f, 0x78,
	0x23, 0x81, 0x43, 0x53, 0xe8, 0x54, 0xe0, 0xcc, 0x90, 0x4c, 0xc5, 0xd2, 0x9c, 0x13, 0xfc, 0x3d,
	0x1a, 0x7f, 0xf1, 0xe2, 0x86, 0x0e, 0x6a, 0x4a, 0xff, 0x55, 0x06, 0xad, 0x39, 0xe0, 0x06, 0x54,
	0x9e, 0xd0, 0x94, 0xe4, 0x11, 0x67, 0x86, 0x64, 0x4a, 0x56, 0xd7, 0x19, 0xfe, 0x7c, 0xc4, 0xb4,
	0xb6, 0x06, 0x3b, 0x0a, 0x42, 0xd0, 0x62, 0xe4, 0x85, 0x1a, 0xb2, 0x29, 0x59, 0x6a, 0x20, 0x34,
	0xba, 0x87, 0x76, 0x4c, 0x96, 0x34, 0xce, 0x0c, 0x45, 0xc4, 0x70, 0x7f, 0x15, 0x03, 0x4f, 0x04,
	0xc3, 0x67, 0x79, 0x5a, 0x06, 0x5b, 0x20, 0x3a, 0x85, 0x7f, 0xcd, 0x1f, 0x30, 0x5a, 0xa6, 0x64,
	0x69, 0x4e, 0x0f, 0x57, 0x55, 0xe3, 0xba, 0x6a, 0xec, 0xb2, 0x32, 0xd0, 0xe8, 0x0e, 0xd6, 0x3f,
	0x07, 0xad, 0xc1, 0x43, 0x3a, 0x28, 0xcf, 0xb4, 0x14, 0x1d, 0xa8, 0xc1, 0x46, 0xa2, 0x1e, 0xfc,
	0x5d, 0x93, 0xb8, 0xa8, 0x93, 0x54, 0x87, 0x0b, 0xf9, 0x4c, 0x1a, 0x5c, 0x82, 0xfa, 0x11, 0x1d,
	0x75, 0x01, 0xc6, 0xd7, 0x77, 0xee, 0x64, 0xec, 0x2d, 0xa6, 0x33, 0xfd, 0x0f, 0xea, 0x80, 0xe2,
	0x7a, 0x9e, 0x2e, 0x21, 0x0d, 0x3a, 0x81, 0x3f, 0x9b, 0xb8, 0x23, 0x5f, 0x97, 0x11, 0x40, 0x7b,
	0x3e, 0xf3, 0xdc, 0x5b, 0x5f, 0x57, 0xae, 0x0e, 0x1f, 0x0e, 0xaa, 0xf8, 0x11, 0xb7, 0x49, 0x12,
	0xd9, 0x9f, 0x37, 0x6b, 0xd9, 0x16, 0xaf, 0x1e, 0xbe, 0x07, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xb1,
	0x0d, 0x49, 0x74, 0x02, 0x00, 0x00,
}