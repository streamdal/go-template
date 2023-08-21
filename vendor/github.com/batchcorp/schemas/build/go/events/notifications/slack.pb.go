// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: slack.proto

package notifications

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

type Slack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is filled out by backend based on the slack_id; frontend should not
	// use this field.
	XBotToken   string `protobuf:"bytes,1,opt,name=_bot_token,json=BotToken,proto3" json:"_bot_token,omitempty"`
	ChannelName string `protobuf:"bytes,2,opt,name=channel_name,json=channelName,proto3" json:"channel_name,omitempty"`
	// Frontend should fill this out when creating or updating an alert config.
	//
	// Frontend can get this value by first asking the API the available Slack
	// integrations (GET /v1/integrations/slack).
	//
	// This is because frontend does not have access to bot_token - it is
	// encrypted, stored in db and it will be filled out by backend when
	// generating *_ALERT* events)
	SlackId string `protobuf:"bytes,3,opt,name=slack_id,json=slackId,proto3" json:"slack_id,omitempty"`
}

func (x *Slack) Reset() {
	*x = Slack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slack_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Slack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slack) ProtoMessage() {}

func (x *Slack) ProtoReflect() protoreflect.Message {
	mi := &file_slack_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slack.ProtoReflect.Descriptor instead.
func (*Slack) Descriptor() ([]byte, []int) {
	return file_slack_proto_rawDescGZIP(), []int{0}
}

func (x *Slack) GetXBotToken() string {
	if x != nil {
		return x.XBotToken
	}
	return ""
}

func (x *Slack) GetChannelName() string {
	if x != nil {
		return x.ChannelName
	}
	return ""
}

func (x *Slack) GetSlackId() string {
	if x != nil {
		return x.SlackId
	}
	return ""
}

var File_slack_proto protoreflect.FileDescriptor

var file_slack_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x63, 0x0a, 0x05,
	0x53, 0x6c, 0x61, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x0a, 0x5f, 0x62, 0x6f, 0x74, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x6f, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x49,
	0x64, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_slack_proto_rawDescOnce sync.Once
	file_slack_proto_rawDescData = file_slack_proto_rawDesc
)

func file_slack_proto_rawDescGZIP() []byte {
	file_slack_proto_rawDescOnce.Do(func() {
		file_slack_proto_rawDescData = protoimpl.X.CompressGZIP(file_slack_proto_rawDescData)
	})
	return file_slack_proto_rawDescData
}

var file_slack_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_slack_proto_goTypes = []interface{}{
	(*Slack)(nil), // 0: notifications.Slack
}
var file_slack_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_slack_proto_init() }
func file_slack_proto_init() {
	if File_slack_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_slack_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Slack); i {
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
			RawDescriptor: file_slack_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_slack_proto_goTypes,
		DependencyIndexes: file_slack_proto_depIdxs,
		MessageInfos:      file_slack_proto_msgTypes,
	}.Build()
	File_slack_proto = out.File
	file_slack_proto_rawDesc = nil
	file_slack_proto_goTypes = nil
	file_slack_proto_depIdxs = nil
}