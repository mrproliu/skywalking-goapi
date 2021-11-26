// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: envoy/service/metrics/v2/metrics_service.proto

package v2

import (
	core "github.com/mrproliu/skywalking-goapi/proto/envoy/api/v2/core"
	client "github.com/mrproliu/skywalking-goapi/proto/io/prometheus/client"
	_ "github.com/mrproliu/skywalking-goapi/proto/udpa/annotations"
	_ "github.com/mrproliu/skywalking-goapi/proto/validate"
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

type StreamMetricsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamMetricsResponse) Reset() {
	*x = StreamMetricsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_metrics_v2_metrics_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamMetricsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamMetricsResponse) ProtoMessage() {}

func (x *StreamMetricsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_metrics_v2_metrics_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamMetricsResponse.ProtoReflect.Descriptor instead.
func (*StreamMetricsResponse) Descriptor() ([]byte, []int) {
	return file_envoy_service_metrics_v2_metrics_service_proto_rawDescGZIP(), []int{0}
}

type StreamMetricsMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier data effectively is a structured metadata. As a performance optimization this will
	// only be sent in the first message on the stream.
	Identifier *StreamMetricsMessage_Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// A list of metric entries
	EnvoyMetrics []*client.MetricFamily `protobuf:"bytes,2,rep,name=envoy_metrics,json=envoyMetrics,proto3" json:"envoy_metrics,omitempty"`
}

func (x *StreamMetricsMessage) Reset() {
	*x = StreamMetricsMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_metrics_v2_metrics_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamMetricsMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamMetricsMessage) ProtoMessage() {}

func (x *StreamMetricsMessage) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_metrics_v2_metrics_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamMetricsMessage.ProtoReflect.Descriptor instead.
func (*StreamMetricsMessage) Descriptor() ([]byte, []int) {
	return file_envoy_service_metrics_v2_metrics_service_proto_rawDescGZIP(), []int{1}
}

func (x *StreamMetricsMessage) GetIdentifier() *StreamMetricsMessage_Identifier {
	if x != nil {
		return x.Identifier
	}
	return nil
}

func (x *StreamMetricsMessage) GetEnvoyMetrics() []*client.MetricFamily {
	if x != nil {
		return x.EnvoyMetrics
	}
	return nil
}

type StreamMetricsMessage_Identifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The node sending metrics over the stream.
	Node *core.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
}

func (x *StreamMetricsMessage_Identifier) Reset() {
	*x = StreamMetricsMessage_Identifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_metrics_v2_metrics_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamMetricsMessage_Identifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamMetricsMessage_Identifier) ProtoMessage() {}

func (x *StreamMetricsMessage_Identifier) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_metrics_v2_metrics_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamMetricsMessage_Identifier.ProtoReflect.Descriptor instead.
func (*StreamMetricsMessage_Identifier) Descriptor() ([]byte, []int) {
	return file_envoy_service_metrics_v2_metrics_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *StreamMetricsMessage_Identifier) GetNode() *core.Node {
	if x != nil {
		return x.Node
	}
	return nil
}

var File_envoy_service_metrics_v2_metrics_service_proto protoreflect.FileDescriptor

var file_envoy_service_metrics_v2_metrics_service_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x69, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x15, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xff, 0x01,
	0x0a, 0x14, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x59, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x47, 0x0a, 0x0d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x52, 0x0c, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x1a, 0x43, 0x0a, 0x0a, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x32,
	0x86, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x74, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x12, 0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x2f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x4a, 0x0a, 0x26, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e,
	0x76, 0x32, 0x42, 0x13, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x88, 0x01, 0x01, 0xba, 0x80, 0xc8, 0xd1,
	0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_service_metrics_v2_metrics_service_proto_rawDescOnce sync.Once
	file_envoy_service_metrics_v2_metrics_service_proto_rawDescData = file_envoy_service_metrics_v2_metrics_service_proto_rawDesc
)

func file_envoy_service_metrics_v2_metrics_service_proto_rawDescGZIP() []byte {
	file_envoy_service_metrics_v2_metrics_service_proto_rawDescOnce.Do(func() {
		file_envoy_service_metrics_v2_metrics_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_service_metrics_v2_metrics_service_proto_rawDescData)
	})
	return file_envoy_service_metrics_v2_metrics_service_proto_rawDescData
}

var file_envoy_service_metrics_v2_metrics_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_service_metrics_v2_metrics_service_proto_goTypes = []interface{}{
	(*StreamMetricsResponse)(nil),           // 0: envoy.service.metrics.v2.StreamMetricsResponse
	(*StreamMetricsMessage)(nil),            // 1: envoy.service.metrics.v2.StreamMetricsMessage
	(*StreamMetricsMessage_Identifier)(nil), // 2: envoy.service.metrics.v2.StreamMetricsMessage.Identifier
	(*client.MetricFamily)(nil),             // 3: io.prometheus.client.MetricFamily
	(*core.Node)(nil),                       // 4: envoy.api.v2.core.Node
}
var file_envoy_service_metrics_v2_metrics_service_proto_depIdxs = []int32{
	2, // 0: envoy.service.metrics.v2.StreamMetricsMessage.identifier:type_name -> envoy.service.metrics.v2.StreamMetricsMessage.Identifier
	3, // 1: envoy.service.metrics.v2.StreamMetricsMessage.envoy_metrics:type_name -> io.prometheus.client.MetricFamily
	4, // 2: envoy.service.metrics.v2.StreamMetricsMessage.Identifier.node:type_name -> envoy.api.v2.core.Node
	1, // 3: envoy.service.metrics.v2.MetricsService.StreamMetrics:input_type -> envoy.service.metrics.v2.StreamMetricsMessage
	0, // 4: envoy.service.metrics.v2.MetricsService.StreamMetrics:output_type -> envoy.service.metrics.v2.StreamMetricsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_service_metrics_v2_metrics_service_proto_init() }
func file_envoy_service_metrics_v2_metrics_service_proto_init() {
	if File_envoy_service_metrics_v2_metrics_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_service_metrics_v2_metrics_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamMetricsResponse); i {
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
		file_envoy_service_metrics_v2_metrics_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamMetricsMessage); i {
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
		file_envoy_service_metrics_v2_metrics_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamMetricsMessage_Identifier); i {
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
			RawDescriptor: file_envoy_service_metrics_v2_metrics_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_envoy_service_metrics_v2_metrics_service_proto_goTypes,
		DependencyIndexes: file_envoy_service_metrics_v2_metrics_service_proto_depIdxs,
		MessageInfos:      file_envoy_service_metrics_v2_metrics_service_proto_msgTypes,
	}.Build()
	File_envoy_service_metrics_v2_metrics_service_proto = out.File
	file_envoy_service_metrics_v2_metrics_service_proto_rawDesc = nil
	file_envoy_service_metrics_v2_metrics_service_proto_goTypes = nil
	file_envoy_service_metrics_v2_metrics_service_proto_depIdxs = nil
}
