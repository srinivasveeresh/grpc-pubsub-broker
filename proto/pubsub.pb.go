// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/pubsub.proto

package pubsub

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_proto_pubsub_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pubsub_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_pubsub_proto_rawDescGZIP(), []int{0}
}

type Message struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Data          string                 `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_proto_pubsub_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pubsub_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_pubsub_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Message) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type TargetedMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	SubscriberId  string                 `protobuf:"bytes,2,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	Data          string                 `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TargetedMessage) Reset() {
	*x = TargetedMessage{}
	mi := &file_proto_pubsub_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TargetedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetedMessage) ProtoMessage() {}

func (x *TargetedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pubsub_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetedMessage.ProtoReflect.Descriptor instead.
func (*TargetedMessage) Descriptor() ([]byte, []int) {
	return file_proto_pubsub_proto_rawDescGZIP(), []int{2}
}

func (x *TargetedMessage) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *TargetedMessage) GetSubscriberId() string {
	if x != nil {
		return x.SubscriberId
	}
	return ""
}

func (x *TargetedMessage) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type TargetedGroupMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	SubscriberIds []string               `protobuf:"bytes,2,rep,name=subscriber_ids,json=subscriberIds,proto3" json:"subscriber_ids,omitempty"`
	Data          string                 `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TargetedGroupMessage) Reset() {
	*x = TargetedGroupMessage{}
	mi := &file_proto_pubsub_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TargetedGroupMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetedGroupMessage) ProtoMessage() {}

func (x *TargetedGroupMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pubsub_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetedGroupMessage.ProtoReflect.Descriptor instead.
func (*TargetedGroupMessage) Descriptor() ([]byte, []int) {
	return file_proto_pubsub_proto_rawDescGZIP(), []int{3}
}

func (x *TargetedGroupMessage) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *TargetedGroupMessage) GetSubscriberIds() []string {
	if x != nil {
		return x.SubscriberIds
	}
	return nil
}

func (x *TargetedGroupMessage) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type SubscribeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	mi := &file_proto_pubsub_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pubsub_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_proto_pubsub_proto_rawDescGZIP(), []int{4}
}

func (x *SubscribeRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

type UnsubscribeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	SubscriberId  string                 `protobuf:"bytes,2,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnsubscribeRequest) Reset() {
	*x = UnsubscribeRequest{}
	mi := &file_proto_pubsub_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnsubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeRequest) ProtoMessage() {}

func (x *UnsubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pubsub_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeRequest.ProtoReflect.Descriptor instead.
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return file_proto_pubsub_proto_rawDescGZIP(), []int{5}
}

func (x *UnsubscribeRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *UnsubscribeRequest) GetSubscriberId() string {
	if x != nil {
		return x.SubscriberId
	}
	return ""
}

type SubscriptionMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SubscriberId  string                 `protobuf:"bytes,1,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	Topic         string                 `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Data          string                 `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscriptionMessage) Reset() {
	*x = SubscriptionMessage{}
	mi := &file_proto_pubsub_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscriptionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionMessage) ProtoMessage() {}

func (x *SubscriptionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pubsub_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionMessage.ProtoReflect.Descriptor instead.
func (*SubscriptionMessage) Descriptor() ([]byte, []int) {
	return file_proto_pubsub_proto_rawDescGZIP(), []int{6}
}

func (x *SubscriptionMessage) GetSubscriberId() string {
	if x != nil {
		return x.SubscriberId
	}
	return ""
}

func (x *SubscriptionMessage) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *SubscriptionMessage) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type SubscriberMetadata struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SubscriberId  string                 `protobuf:"bytes,1,opt,name=subscriber_id,json=subscriberId,proto3" json:"subscriber_id,omitempty"`
	IpAddress     string                 `protobuf:"bytes,2,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Tags          []string               `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscriberMetadata) Reset() {
	*x = SubscriberMetadata{}
	mi := &file_proto_pubsub_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscriberMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriberMetadata) ProtoMessage() {}

func (x *SubscriberMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pubsub_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriberMetadata.ProtoReflect.Descriptor instead.
func (*SubscriberMetadata) Descriptor() ([]byte, []int) {
	return file_proto_pubsub_proto_rawDescGZIP(), []int{7}
}

func (x *SubscriberMetadata) GetSubscriberId() string {
	if x != nil {
		return x.SubscriberId
	}
	return ""
}

func (x *SubscriberMetadata) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *SubscriberMetadata) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type SubscriberList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Subscribers   []*SubscriberMetadata  `protobuf:"bytes,1,rep,name=subscribers,proto3" json:"subscribers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscriberList) Reset() {
	*x = SubscriberList{}
	mi := &file_proto_pubsub_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscriberList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriberList) ProtoMessage() {}

func (x *SubscriberList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pubsub_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriberList.ProtoReflect.Descriptor instead.
func (*SubscriberList) Descriptor() ([]byte, []int) {
	return file_proto_pubsub_proto_rawDescGZIP(), []int{8}
}

func (x *SubscriberList) GetSubscribers() []*SubscriberMetadata {
	if x != nil {
		return x.Subscribers
	}
	return nil
}

type TopicList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topics        []string               `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TopicList) Reset() {
	*x = TopicList{}
	mi := &file_proto_pubsub_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TopicList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicList) ProtoMessage() {}

func (x *TopicList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pubsub_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicList.ProtoReflect.Descriptor instead.
func (*TopicList) Descriptor() ([]byte, []int) {
	return file_proto_pubsub_proto_rawDescGZIP(), []int{9}
}

func (x *TopicList) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

type Ack struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Ack) Reset() {
	*x = Ack{}
	mi := &file_proto_pubsub_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pubsub_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_proto_pubsub_proto_rawDescGZIP(), []int{10}
}

func (x *Ack) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_pubsub_proto protoreflect.FileDescriptor

const file_proto_pubsub_proto_rawDesc = "" +
	"\n" +
	"\x12proto/pubsub.proto\x12\x06pubsub\"\a\n" +
	"\x05Empty\"3\n" +
	"\aMessage\x12\x14\n" +
	"\x05topic\x18\x01 \x01(\tR\x05topic\x12\x12\n" +
	"\x04data\x18\x02 \x01(\tR\x04data\"`\n" +
	"\x0fTargetedMessage\x12\x14\n" +
	"\x05topic\x18\x01 \x01(\tR\x05topic\x12#\n" +
	"\rsubscriber_id\x18\x02 \x01(\tR\fsubscriberId\x12\x12\n" +
	"\x04data\x18\x03 \x01(\tR\x04data\"g\n" +
	"\x14TargetedGroupMessage\x12\x14\n" +
	"\x05topic\x18\x01 \x01(\tR\x05topic\x12%\n" +
	"\x0esubscriber_ids\x18\x02 \x03(\tR\rsubscriberIds\x12\x12\n" +
	"\x04data\x18\x03 \x01(\tR\x04data\"(\n" +
	"\x10SubscribeRequest\x12\x14\n" +
	"\x05topic\x18\x01 \x01(\tR\x05topic\"O\n" +
	"\x12UnsubscribeRequest\x12\x14\n" +
	"\x05topic\x18\x01 \x01(\tR\x05topic\x12#\n" +
	"\rsubscriber_id\x18\x02 \x01(\tR\fsubscriberId\"d\n" +
	"\x13SubscriptionMessage\x12#\n" +
	"\rsubscriber_id\x18\x01 \x01(\tR\fsubscriberId\x12\x14\n" +
	"\x05topic\x18\x02 \x01(\tR\x05topic\x12\x12\n" +
	"\x04data\x18\x03 \x01(\tR\x04data\"l\n" +
	"\x12SubscriberMetadata\x12#\n" +
	"\rsubscriber_id\x18\x01 \x01(\tR\fsubscriberId\x12\x1d\n" +
	"\n" +
	"ip_address\x18\x02 \x01(\tR\tipAddress\x12\x12\n" +
	"\x04tags\x18\x03 \x03(\tR\x04tags\"N\n" +
	"\x0eSubscriberList\x12<\n" +
	"\vsubscribers\x18\x01 \x03(\v2\x1a.pubsub.SubscriberMetadataR\vsubscribers\"#\n" +
	"\tTopicList\x12\x16\n" +
	"\x06topics\x18\x01 \x03(\tR\x06topics\"\x1d\n" +
	"\x03Ack\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status2\xcd\x03\n" +
	"\x06PubSub\x12'\n" +
	"\aPublish\x12\x0f.pubsub.Message\x1a\v.pubsub.Ack\x12;\n" +
	"\x13PublishToSubscriber\x12\x17.pubsub.TargetedMessage\x1a\v.pubsub.Ack\x12A\n" +
	"\x14PublishToSubscribers\x12\x1c.pubsub.TargetedGroupMessage\x1a\v.pubsub.Ack\x12D\n" +
	"\tSubscribe\x12\x18.pubsub.SubscribeRequest\x1a\x1b.pubsub.SubscriptionMessage0\x01\x12'\n" +
	"\aForward\x12\x0f.pubsub.Message\x1a\v.pubsub.Ack\x12C\n" +
	"\x0fListSubscribers\x12\x18.pubsub.SubscribeRequest\x1a\x16.pubsub.SubscriberList\x12.\n" +
	"\n" +
	"ListTopics\x12\r.pubsub.Empty\x1a\x11.pubsub.TopicList\x126\n" +
	"\vUnsubscribe\x12\x1a.pubsub.UnsubscribeRequest\x1a\v.pubsub.AckB@Z>github.com/srinivasveeresh/openvpn-manager/grpc-pubsub-broker;pubsubb\x06proto3"

var (
	file_proto_pubsub_proto_rawDescOnce sync.Once
	file_proto_pubsub_proto_rawDescData []byte
)

func file_proto_pubsub_proto_rawDescGZIP() []byte {
	file_proto_pubsub_proto_rawDescOnce.Do(func() {
		file_proto_pubsub_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_pubsub_proto_rawDesc), len(file_proto_pubsub_proto_rawDesc)))
	})
	return file_proto_pubsub_proto_rawDescData
}

var file_proto_pubsub_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_pubsub_proto_goTypes = []any{
	(*Empty)(nil),                // 0: pubsub.Empty
	(*Message)(nil),              // 1: pubsub.Message
	(*TargetedMessage)(nil),      // 2: pubsub.TargetedMessage
	(*TargetedGroupMessage)(nil), // 3: pubsub.TargetedGroupMessage
	(*SubscribeRequest)(nil),     // 4: pubsub.SubscribeRequest
	(*UnsubscribeRequest)(nil),   // 5: pubsub.UnsubscribeRequest
	(*SubscriptionMessage)(nil),  // 6: pubsub.SubscriptionMessage
	(*SubscriberMetadata)(nil),   // 7: pubsub.SubscriberMetadata
	(*SubscriberList)(nil),       // 8: pubsub.SubscriberList
	(*TopicList)(nil),            // 9: pubsub.TopicList
	(*Ack)(nil),                  // 10: pubsub.Ack
}
var file_proto_pubsub_proto_depIdxs = []int32{
	7,  // 0: pubsub.SubscriberList.subscribers:type_name -> pubsub.SubscriberMetadata
	1,  // 1: pubsub.PubSub.Publish:input_type -> pubsub.Message
	2,  // 2: pubsub.PubSub.PublishToSubscriber:input_type -> pubsub.TargetedMessage
	3,  // 3: pubsub.PubSub.PublishToSubscribers:input_type -> pubsub.TargetedGroupMessage
	4,  // 4: pubsub.PubSub.Subscribe:input_type -> pubsub.SubscribeRequest
	1,  // 5: pubsub.PubSub.Forward:input_type -> pubsub.Message
	4,  // 6: pubsub.PubSub.ListSubscribers:input_type -> pubsub.SubscribeRequest
	0,  // 7: pubsub.PubSub.ListTopics:input_type -> pubsub.Empty
	5,  // 8: pubsub.PubSub.Unsubscribe:input_type -> pubsub.UnsubscribeRequest
	10, // 9: pubsub.PubSub.Publish:output_type -> pubsub.Ack
	10, // 10: pubsub.PubSub.PublishToSubscriber:output_type -> pubsub.Ack
	10, // 11: pubsub.PubSub.PublishToSubscribers:output_type -> pubsub.Ack
	6,  // 12: pubsub.PubSub.Subscribe:output_type -> pubsub.SubscriptionMessage
	10, // 13: pubsub.PubSub.Forward:output_type -> pubsub.Ack
	8,  // 14: pubsub.PubSub.ListSubscribers:output_type -> pubsub.SubscriberList
	9,  // 15: pubsub.PubSub.ListTopics:output_type -> pubsub.TopicList
	10, // 16: pubsub.PubSub.Unsubscribe:output_type -> pubsub.Ack
	9,  // [9:17] is the sub-list for method output_type
	1,  // [1:9] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_proto_pubsub_proto_init() }
func file_proto_pubsub_proto_init() {
	if File_proto_pubsub_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_pubsub_proto_rawDesc), len(file_proto_pubsub_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_pubsub_proto_goTypes,
		DependencyIndexes: file_proto_pubsub_proto_depIdxs,
		MessageInfos:      file_proto_pubsub_proto_msgTypes,
	}.Build()
	File_proto_pubsub_proto = out.File
	file_proto_pubsub_proto_goTypes = nil
	file_proto_pubsub_proto_depIdxs = nil
}
