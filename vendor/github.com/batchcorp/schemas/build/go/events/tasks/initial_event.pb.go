// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: initial_event.proto

package tasks

import (
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

// Used for the onboarding wizard flow to signify that an event was collected
// by a collector and written by writer.
//
// Order of operations:
//
//  1. Frontend asks `ui-bff` to initiate initial event flow
//  2. `ui-bff` emits CREATE_TASK event w/ initial event task set to $collection_id
//  3. All collectors consume CREATE_TASK event and begin watching for first event
//  4. All writers consume CREATE_TASK event and begin watching for first event
//  5. One collector receives event for $collection_id and emits UPDATE_TASK
//     specifying that it received an event, updating state to PROCESSING
//  6. One writer receives event on HSB for $collection_id and emits an UPDATE_TASK
//     event, setting task state to "COMPLETED"
//  7. As the task is getting updates, the task system receives all UPDATE_TASK
//     events and emits details via any associated websockets
//
// ~DS 12.02.2020
type InitialEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionId string `protobuf:"bytes,1,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	Service      string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Timestamp    int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *InitialEvent) Reset() {
	*x = InitialEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_initial_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitialEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitialEvent) ProtoMessage() {}

func (x *InitialEvent) ProtoReflect() protoreflect.Message {
	mi := &file_initial_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitialEvent.ProtoReflect.Descriptor instead.
func (*InitialEvent) Descriptor() ([]byte, []int) {
	return file_initial_event_proto_rawDescGZIP(), []int{0}
}

func (x *InitialEvent) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

func (x *InitialEvent) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *InitialEvent) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_initial_event_proto protoreflect.FileDescriptor

var file_initial_event_proto_rawDesc = []byte{
	0x0a, 0x13, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x6b, 0x0a, 0x0c,
	0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x63, 0x6f, 0x72,
	0x70, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f,
	0x67, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_initial_event_proto_rawDescOnce sync.Once
	file_initial_event_proto_rawDescData = file_initial_event_proto_rawDesc
)

func file_initial_event_proto_rawDescGZIP() []byte {
	file_initial_event_proto_rawDescOnce.Do(func() {
		file_initial_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_initial_event_proto_rawDescData)
	})
	return file_initial_event_proto_rawDescData
}

var file_initial_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_initial_event_proto_goTypes = []interface{}{
	(*InitialEvent)(nil), // 0: tasks.InitialEvent
}
var file_initial_event_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_initial_event_proto_init() }
func file_initial_event_proto_init() {
	if File_initial_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_initial_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitialEvent); i {
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
			RawDescriptor: file_initial_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_initial_event_proto_goTypes,
		DependencyIndexes: file_initial_event_proto_depIdxs,
		MessageInfos:      file_initial_event_proto_msgTypes,
	}.Build()
	File_initial_event_proto = out.File
	file_initial_event_proto_rawDesc = nil
	file_initial_event_proto_goTypes = nil
	file_initial_event_proto_depIdxs = nil
}
