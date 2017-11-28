// Code generated by protoc-gen-go. DO NOT EDIT.
// source: routing/v1alpha1/egress_rule.proto

package istio_routing_v1alpha1_config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Egress rules describe the properties of a service outside Istio. When transparent proxying
// is used, egress rules signify a white listed set of domains that microserves in the mesh
// are allowed to access. A subset of routing rules and all destination policies can be applied
// on the service targeted by an egress rule. The destination of an egress rule is allowed to
// contain wildcards (e.g., *.foo.com). Currently, only HTTP-based services can be expressed
// through the egress rule. If TLS origination from the sidecar is desired, the protocol
// associated with the service port must be marked as HTTPS, and the service is expected to
// be accessed over HTTP (e.g., http://gmail.com:443). The sidecar will automatically upgrade
// the connection to TLS when initiating a connection with the external service.
//
// For example, the following egress rule describes the set of services hosted under the *.foo.com domain
//
//     kind: EgressRule
//     metadata:
//       name: foo-egress-rule
//     spec:
//       destination:
//         service: *.foo.com
//       ports:
//         - port: 80
//           protocol: http
//         - port: 443
//           protocol: https
//
type EgressRule struct {
	// REQUIRED: Hostname or a wildcard domain name associated with the external service.
	// ONLY the "service" field of destination will be taken into consideration. Name,
	// namespace, domain and labels are ignored. Routing rules and destination policies that
	// refer to these external services must have identical specification for the destination
	// as the corresponding egress rule. Wildcard domain specifications must conform to format
	// allowed by Envoy's Virtual Host specification, such as “*.foo.com” or “*-bar.foo.com”.
	// The character '*' in a domain specification indicates a non-empty string. Hence, a wildcard
	// domain of form “*-bar.foo.com” will match “baz-bar.foo.com” but not “-bar.foo.com”.
	Destination *IstioService `protobuf:"bytes,1,opt,name=destination" json:"destination,omitempty"`
	// REQUIRED: list of ports on which the external service is available.
	Ports []*EgressRule_Port `protobuf:"bytes,2,rep,name=ports" json:"ports,omitempty"`
	// Forward all the external traffic through a dedicated egress proxy. It is used in some scenarios
	// where there is a requirement that all the external traffic goes through special dedicated nodes/pods.
	// These dedicated egress nodes could then be more closely monitored for security vulnerabilities.
	//
	// The default is false, i.e. the sidecar forwards external traffic directly to the external service.
	UseEgressProxy bool `protobuf:"varint,3,opt,name=use_egress_proxy,json=useEgressProxy" json:"use_egress_proxy,omitempty"`
}

func (m *EgressRule) Reset()                    { *m = EgressRule{} }
func (m *EgressRule) String() string            { return proto.CompactTextString(m) }
func (*EgressRule) ProtoMessage()               {}
func (*EgressRule) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *EgressRule) GetDestination() *IstioService {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *EgressRule) GetPorts() []*EgressRule_Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *EgressRule) GetUseEgressProxy() bool {
	if m != nil {
		return m.UseEgressProxy
	}
	return false
}

// Port describes the properties of a specific TCP port of an external service.
type EgressRule_Port struct {
	// A valid non-negative integer port number.
	Port int32 `protobuf:"varint,1,opt,name=port" json:"port,omitempty"`
	// The protocol to communicate with the external services.
	// MUST BE one of HTTP|HTTPS|GRPC|HTTP2.
	Protocol string `protobuf:"bytes,2,opt,name=protocol" json:"protocol,omitempty"`
}

func (m *EgressRule_Port) Reset()                    { *m = EgressRule_Port{} }
func (m *EgressRule_Port) String() string            { return proto.CompactTextString(m) }
func (*EgressRule_Port) ProtoMessage()               {}
func (*EgressRule_Port) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func (m *EgressRule_Port) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *EgressRule_Port) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func init() {
	proto.RegisterType((*EgressRule)(nil), "istio.routing.v1alpha1.config.EgressRule")
	proto.RegisterType((*EgressRule_Port)(nil), "istio.routing.v1alpha1.config.EgressRule.Port")
}

func init() { proto.RegisterFile("routing/v1alpha1/egress_rule.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0xda, 0x4a, 0x9d, 0x80, 0xc8, 0x9e, 0x42, 0x40, 0x88, 0x3d, 0x2d, 0x08, 0x53,
	0x5a, 0xc1, 0x27, 0xd0, 0x83, 0x07, 0xa1, 0xac, 0x0f, 0x50, 0x62, 0x1c, 0xe3, 0x42, 0xc8, 0x84,
	0xd9, 0xdd, 0xa2, 0x67, 0x5f, 0x5c, 0x76, 0x5b, 0xb5, 0x20, 0xe8, 0x6d, 0x76, 0x76, 0xbe, 0x6f,
	0xe6, 0x87, 0x85, 0x70, 0xf0, 0x76, 0xe8, 0x96, 0xbb, 0x55, 0xd3, 0x8f, 0xaf, 0xcd, 0x6a, 0x49,
	0x9d, 0x90, 0x73, 0x5b, 0x09, 0x3d, 0xe1, 0x28, 0xec, 0x59, 0x5d, 0x58, 0xe7, 0x2d, 0xe3, 0x61,
	0x12, 0xbf, 0x26, 0xb1, 0xe5, 0xe1, 0xc5, 0x76, 0xd5, 0xe5, 0x2f, 0x45, 0x6c, 0xd0, 0x91, 0x61,
	0xf1, 0x91, 0x03, 0xdc, 0x25, 0xaf, 0x09, 0x3d, 0xa9, 0x07, 0x28, 0x9e, 0xc9, 0x79, 0x3b, 0x34,
	0xde, 0xf2, 0x50, 0x66, 0x75, 0xa6, 0x8b, 0xf5, 0x15, 0xfe, 0xb9, 0x06, 0xef, 0xe3, 0xef, 0x23,
	0xc9, 0xce, 0xb6, 0x64, 0x8e, 0x79, 0x75, 0x0b, 0xb3, 0x91, 0xc5, 0xbb, 0x32, 0xaf, 0x27, 0xba,
	0x58, 0xe3, 0x3f, 0xa2, 0x9f, 0x43, 0x70, 0xc3, 0xe2, 0xcd, 0x1e, 0x56, 0x1a, 0xce, 0x83, 0xa3,
	0xed, 0x21, 0xfe, 0x28, 0xfc, 0xf6, 0x5e, 0x4e, 0xea, 0x4c, 0xcf, 0xcd, 0x59, 0x70, 0xb4, 0x87,
	0x36, 0xb1, 0x5b, 0xdd, 0xc0, 0x34, 0x82, 0x4a, 0xc1, 0x34, 0xa2, 0xe9, 0xfe, 0x99, 0x49, 0xb5,
	0xaa, 0x60, 0x9e, 0x22, 0xb7, 0xdc, 0x97, 0x79, 0x9d, 0xe9, 0x53, 0xf3, 0xfd, 0x7e, 0x3a, 0x49,
	0xd5, 0xf5, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xb5, 0x60, 0xa9, 0x74, 0x01, 0x00, 0x00,
}
