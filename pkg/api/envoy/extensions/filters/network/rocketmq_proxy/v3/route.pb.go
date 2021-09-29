// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: envoy/extensions/filters/network/rocketmq_proxy/v3/route.proto

package envoy_extensions_filters_network_rocketmq_proxy_v3

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v32 "github.com/datawire/ambassador/v2/pkg/api/envoy/config/core/v3"
	v31 "github.com/datawire/ambassador/v2/pkg/api/envoy/config/route/v3"
	v3 "github.com/datawire/ambassador/v2/pkg/api/envoy/type/matcher/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RouteConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the route configuration.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The list of routes that will be matched, in order, against incoming requests. The first route
	// that matches will be used.
	Routes []*Route `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *RouteConfiguration) Reset() {
	*x = RouteConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteConfiguration) ProtoMessage() {}

func (x *RouteConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteConfiguration.ProtoReflect.Descriptor instead.
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_rawDescGZIP(), []int{0}
}

func (x *RouteConfiguration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RouteConfiguration) GetRoutes() []*Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Route matching parameters.
	Match *RouteMatch `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	// Route request to some upstream cluster.
	Route *RouteAction `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_rawDescGZIP(), []int{1}
}

func (x *Route) GetMatch() *RouteMatch {
	if x != nil {
		return x.Match
	}
	return nil
}

func (x *Route) GetRoute() *RouteAction {
	if x != nil {
		return x.Route
	}
	return nil
}

type RouteMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the topic.
	Topic *v3.StringMatcher `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	// Specifies a set of headers that the route should match on. The router will check the request’s
	// headers against all the specified headers in the route config. A match will happen if all the
	// headers in the route are present in the request with the same values (or based on presence if
	// the value field is not in the config).
	Headers []*v31.HeaderMatcher `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
}

func (x *RouteMatch) Reset() {
	*x = RouteMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteMatch) ProtoMessage() {}

func (x *RouteMatch) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteMatch.ProtoReflect.Descriptor instead.
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_rawDescGZIP(), []int{2}
}

func (x *RouteMatch) GetTopic() *v3.StringMatcher {
	if x != nil {
		return x.Topic
	}
	return nil
}

func (x *RouteMatch) GetHeaders() []*v31.HeaderMatcher {
	if x != nil {
		return x.Headers
	}
	return nil
}

type RouteAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicates the upstream cluster to which the request should be routed.
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// Optional endpoint metadata match criteria used by the subset load balancer.
	MetadataMatch *v32.Metadata `protobuf:"bytes,2,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
}

func (x *RouteAction) Reset() {
	*x = RouteAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteAction) ProtoMessage() {}

func (x *RouteAction) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteAction.ProtoReflect.Descriptor instead.
func (*RouteAction) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_rawDescGZIP(), []int{3}
}

func (x *RouteAction) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *RouteAction) GetMetadataMatch() *v32.Metadata {
	if x != nil {
		return x.MetadataMatch
	}
	return nil
}

var File_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x6d, 0x71, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x6d, 0x71, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x12, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x51, 0x0a, 0x06,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x72, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x6d, 0x71, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76,
	0x33, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22,
	0xc8, 0x01, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x5e, 0x0a, 0x05, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x72, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x6d, 0x71, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x5f, 0x0a, 0x05, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x72, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x6d, 0x71, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x92, 0x01, 0x0a, 0x0a, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x44, 0x0a, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12,
	0x3e, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x22,
	0x77, 0x0a, 0x0b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21,
	0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x45, 0x0a, 0x0e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0d, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x42, 0x58, 0x0a, 0x40, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x72, 0x6f, 0x63, 0x6b, 0x65,
	0x74, 0x6d, 0x71, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x42, 0x0a, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02,
	0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_rawDescData = file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_rawDesc
)

func file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_rawDescData)
	})
	return file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_rawDescData
}

var file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_goTypes = []interface{}{
	(*RouteConfiguration)(nil), // 0: envoy.extensions.filters.network.rocketmq_proxy.v3.RouteConfiguration
	(*Route)(nil),              // 1: envoy.extensions.filters.network.rocketmq_proxy.v3.Route
	(*RouteMatch)(nil),         // 2: envoy.extensions.filters.network.rocketmq_proxy.v3.RouteMatch
	(*RouteAction)(nil),        // 3: envoy.extensions.filters.network.rocketmq_proxy.v3.RouteAction
	(*v3.StringMatcher)(nil),   // 4: envoy.type.matcher.v3.StringMatcher
	(*v31.HeaderMatcher)(nil),  // 5: envoy.config.route.v3.HeaderMatcher
	(*v32.Metadata)(nil),       // 6: envoy.config.core.v3.Metadata
}
var file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.filters.network.rocketmq_proxy.v3.RouteConfiguration.routes:type_name -> envoy.extensions.filters.network.rocketmq_proxy.v3.Route
	2, // 1: envoy.extensions.filters.network.rocketmq_proxy.v3.Route.match:type_name -> envoy.extensions.filters.network.rocketmq_proxy.v3.RouteMatch
	3, // 2: envoy.extensions.filters.network.rocketmq_proxy.v3.Route.route:type_name -> envoy.extensions.filters.network.rocketmq_proxy.v3.RouteAction
	4, // 3: envoy.extensions.filters.network.rocketmq_proxy.v3.RouteMatch.topic:type_name -> envoy.type.matcher.v3.StringMatcher
	5, // 4: envoy.extensions.filters.network.rocketmq_proxy.v3.RouteMatch.headers:type_name -> envoy.config.route.v3.HeaderMatcher
	6, // 5: envoy.extensions.filters.network.rocketmq_proxy.v3.RouteAction.metadata_match:type_name -> envoy.config.core.v3.Metadata
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_init() }
func file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_init() {
	if File_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteConfiguration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteMatch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto = out.File
	file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_rawDesc = nil
	file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_goTypes = nil
	file_envoy_extensions_filters_network_rocketmq_proxy_v3_route_proto_depIdxs = nil
}
