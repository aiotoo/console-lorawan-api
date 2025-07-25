// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.24.4
// source: external/service/mqtt.proto

package service

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MqttAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// clientid
	Clientid string `protobuf:"bytes,1,opt,name=clientid,proto3" json:"clientid,omitempty"`
	// Username
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// Password
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *MqttAuthRequest) Reset() {
	*x = MqttAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_service_mqtt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MqttAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MqttAuthRequest) ProtoMessage() {}

func (x *MqttAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_service_mqtt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MqttAuthRequest.ProtoReflect.Descriptor instead.
func (*MqttAuthRequest) Descriptor() ([]byte, []int) {
	return file_external_service_mqtt_proto_rawDescGZIP(), []int{0}
}

func (x *MqttAuthRequest) GetClientid() string {
	if x != nil {
		return x.Clientid
	}
	return ""
}

func (x *MqttAuthRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *MqttAuthRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type MqttAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Result "allow" | "deny" | "ignore"
	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MqttAuthResponse) Reset() {
	*x = MqttAuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_service_mqtt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MqttAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MqttAuthResponse) ProtoMessage() {}

func (x *MqttAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_service_mqtt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MqttAuthResponse.ProtoReflect.Descriptor instead.
func (*MqttAuthResponse) Descriptor() ([]byte, []int) {
	return file_external_service_mqtt_proto_rawDescGZIP(), []int{1}
}

func (x *MqttAuthResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type MqttSuperuserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// clientid
	Clientid string `protobuf:"bytes,1,opt,name=clientid,proto3" json:"clientid,omitempty"`
	// Username
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// Password
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *MqttSuperuserRequest) Reset() {
	*x = MqttSuperuserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_service_mqtt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MqttSuperuserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MqttSuperuserRequest) ProtoMessage() {}

func (x *MqttSuperuserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_service_mqtt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MqttSuperuserRequest.ProtoReflect.Descriptor instead.
func (*MqttSuperuserRequest) Descriptor() ([]byte, []int) {
	return file_external_service_mqtt_proto_rawDescGZIP(), []int{2}
}

func (x *MqttSuperuserRequest) GetClientid() string {
	if x != nil {
		return x.Clientid
	}
	return ""
}

func (x *MqttSuperuserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *MqttSuperuserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type MqttSuperuserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Result "allow" | "deny" | "ignore"
	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MqttSuperuserResponse) Reset() {
	*x = MqttSuperuserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_service_mqtt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MqttSuperuserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MqttSuperuserResponse) ProtoMessage() {}

func (x *MqttSuperuserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_service_mqtt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MqttSuperuserResponse.ProtoReflect.Descriptor instead.
func (*MqttSuperuserResponse) Descriptor() ([]byte, []int) {
	return file_external_service_mqtt_proto_rawDescGZIP(), []int{3}
}

func (x *MqttSuperuserResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type MqttAclRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// clientid
	Clientid string `protobuf:"bytes,1,opt,name=clientid,proto3" json:"clientid,omitempty"`
	// Username
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// Access
	Access int64 `protobuf:"varint,3,opt,name=access,proto3" json:"access,omitempty"`
	// Topic
	Topic string `protobuf:"bytes,4,opt,name=topic,proto3" json:"topic,omitempty"`
	// ipaddr
	Ipaddr string `protobuf:"bytes,5,opt,name=ipaddr,proto3" json:"ipaddr,omitempty"`
}

func (x *MqttAclRequest) Reset() {
	*x = MqttAclRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_service_mqtt_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MqttAclRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MqttAclRequest) ProtoMessage() {}

func (x *MqttAclRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_service_mqtt_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MqttAclRequest.ProtoReflect.Descriptor instead.
func (*MqttAclRequest) Descriptor() ([]byte, []int) {
	return file_external_service_mqtt_proto_rawDescGZIP(), []int{4}
}

func (x *MqttAclRequest) GetClientid() string {
	if x != nil {
		return x.Clientid
	}
	return ""
}

func (x *MqttAclRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *MqttAclRequest) GetAccess() int64 {
	if x != nil {
		return x.Access
	}
	return 0
}

func (x *MqttAclRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *MqttAclRequest) GetIpaddr() string {
	if x != nil {
		return x.Ipaddr
	}
	return ""
}

type MqttAclResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Result "allow" | "deny" | "ignore"
	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MqttAclResponse) Reset() {
	*x = MqttAclResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_service_mqtt_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MqttAclResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MqttAclResponse) ProtoMessage() {}

func (x *MqttAclResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_service_mqtt_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MqttAclResponse.ProtoReflect.Descriptor instead.
func (*MqttAclResponse) Descriptor() ([]byte, []int) {
	return file_external_service_mqtt_proto_rawDescGZIP(), []int{5}
}

func (x *MqttAclResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_external_service_mqtt_proto protoreflect.FileDescriptor

var file_external_service_mqtt_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x6d, 0x71, 0x74, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x0f, 0x4d, 0x71, 0x74, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2a, 0x0a, 0x10, 0x4d,
	0x71, 0x74, 0x74, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x6a, 0x0a, 0x14, 0x4d, 0x71, 0x74, 0x74, 0x53,
	0x75, 0x70, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x2f, 0x0a, 0x15, 0x4d, 0x71, 0x74, 0x74, 0x53, 0x75, 0x70, 0x65, 0x72,
	0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x8e, 0x01, 0x0a, 0x0e, 0x4d, 0x71, 0x74, 0x74, 0x41, 0x63, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x70, 0x61, 0x64, 0x64, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69,
	0x70, 0x61, 0x64, 0x64, 0x72, 0x22, 0x29, 0x0a, 0x0f, 0x4d, 0x71, 0x74, 0x74, 0x41, 0x63, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x32, 0xa8, 0x02, 0x0a, 0x0e, 0x4d, 0x71, 0x74, 0x74, 0x41, 0x63, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x18, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x71, 0x74, 0x74, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4d, 0x71, 0x74, 0x74, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6d, 0x71, 0x74, 0x74, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x12, 0x6a, 0x0a, 0x09, 0x53,
	0x75, 0x70, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4d, 0x71, 0x74, 0x74, 0x53, 0x75, 0x70, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4d, 0x71, 0x74, 0x74, 0x53, 0x75, 0x70, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a,
	0x01, 0x2a, 0x22, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x71, 0x74, 0x74, 0x2f, 0x73, 0x75,
	0x70, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x03, 0x41, 0x63, 0x6c, 0x12, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x71, 0x74, 0x74, 0x41, 0x63, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4d, 0x71, 0x74, 0x74, 0x41, 0x63, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6d, 0x71, 0x74, 0x74, 0x2f, 0x61, 0x63, 0x6c, 0x42, 0x3e, 0x5a, 0x3c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x6f, 0x74, 0x6f, 0x6f,
	0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2d, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x34, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_external_service_mqtt_proto_rawDescOnce sync.Once
	file_external_service_mqtt_proto_rawDescData = file_external_service_mqtt_proto_rawDesc
)

func file_external_service_mqtt_proto_rawDescGZIP() []byte {
	file_external_service_mqtt_proto_rawDescOnce.Do(func() {
		file_external_service_mqtt_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_service_mqtt_proto_rawDescData)
	})
	return file_external_service_mqtt_proto_rawDescData
}

var file_external_service_mqtt_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_external_service_mqtt_proto_goTypes = []any{
	(*MqttAuthRequest)(nil),       // 0: service.MqttAuthRequest
	(*MqttAuthResponse)(nil),      // 1: service.MqttAuthResponse
	(*MqttSuperuserRequest)(nil),  // 2: service.MqttSuperuserRequest
	(*MqttSuperuserResponse)(nil), // 3: service.MqttSuperuserResponse
	(*MqttAclRequest)(nil),        // 4: service.MqttAclRequest
	(*MqttAclResponse)(nil),       // 5: service.MqttAclResponse
}
var file_external_service_mqtt_proto_depIdxs = []int32{
	0, // 0: service.MqttAclService.Auth:input_type -> service.MqttAuthRequest
	2, // 1: service.MqttAclService.SuperUser:input_type -> service.MqttSuperuserRequest
	4, // 2: service.MqttAclService.Acl:input_type -> service.MqttAclRequest
	1, // 3: service.MqttAclService.Auth:output_type -> service.MqttAuthResponse
	3, // 4: service.MqttAclService.SuperUser:output_type -> service.MqttSuperuserResponse
	5, // 5: service.MqttAclService.Acl:output_type -> service.MqttAclResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_external_service_mqtt_proto_init() }
func file_external_service_mqtt_proto_init() {
	if File_external_service_mqtt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_service_mqtt_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MqttAuthRequest); i {
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
		file_external_service_mqtt_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MqttAuthResponse); i {
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
		file_external_service_mqtt_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MqttSuperuserRequest); i {
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
		file_external_service_mqtt_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MqttSuperuserResponse); i {
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
		file_external_service_mqtt_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MqttAclRequest); i {
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
		file_external_service_mqtt_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*MqttAclResponse); i {
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
			RawDescriptor: file_external_service_mqtt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_external_service_mqtt_proto_goTypes,
		DependencyIndexes: file_external_service_mqtt_proto_depIdxs,
		MessageInfos:      file_external_service_mqtt_proto_msgTypes,
	}.Build()
	File_external_service_mqtt_proto = out.File
	file_external_service_mqtt_proto_rawDesc = nil
	file_external_service_mqtt_proto_goTypes = nil
	file_external_service_mqtt_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MqttAclServiceClient is the client API for MqttAclService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MqttAclServiceClient interface {
	// MQTT Auth.
	Auth(ctx context.Context, in *MqttAuthRequest, opts ...grpc.CallOption) (*MqttAuthResponse, error)
	// MQTT SuperUser.
	SuperUser(ctx context.Context, in *MqttSuperuserRequest, opts ...grpc.CallOption) (*MqttSuperuserResponse, error)
	// MQTT Acl.
	Acl(ctx context.Context, in *MqttAclRequest, opts ...grpc.CallOption) (*MqttAclResponse, error)
}

type mqttAclServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMqttAclServiceClient(cc grpc.ClientConnInterface) MqttAclServiceClient {
	return &mqttAclServiceClient{cc}
}

func (c *mqttAclServiceClient) Auth(ctx context.Context, in *MqttAuthRequest, opts ...grpc.CallOption) (*MqttAuthResponse, error) {
	out := new(MqttAuthResponse)
	err := c.cc.Invoke(ctx, "/service.MqttAclService/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqttAclServiceClient) SuperUser(ctx context.Context, in *MqttSuperuserRequest, opts ...grpc.CallOption) (*MqttSuperuserResponse, error) {
	out := new(MqttSuperuserResponse)
	err := c.cc.Invoke(ctx, "/service.MqttAclService/SuperUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mqttAclServiceClient) Acl(ctx context.Context, in *MqttAclRequest, opts ...grpc.CallOption) (*MqttAclResponse, error) {
	out := new(MqttAclResponse)
	err := c.cc.Invoke(ctx, "/service.MqttAclService/Acl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MqttAclServiceServer is the server API for MqttAclService service.
type MqttAclServiceServer interface {
	// MQTT Auth.
	Auth(context.Context, *MqttAuthRequest) (*MqttAuthResponse, error)
	// MQTT SuperUser.
	SuperUser(context.Context, *MqttSuperuserRequest) (*MqttSuperuserResponse, error)
	// MQTT Acl.
	Acl(context.Context, *MqttAclRequest) (*MqttAclResponse, error)
}

// UnimplementedMqttAclServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMqttAclServiceServer struct {
}

func (*UnimplementedMqttAclServiceServer) Auth(context.Context, *MqttAuthRequest) (*MqttAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (*UnimplementedMqttAclServiceServer) SuperUser(context.Context, *MqttSuperuserRequest) (*MqttSuperuserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuperUser not implemented")
}
func (*UnimplementedMqttAclServiceServer) Acl(context.Context, *MqttAclRequest) (*MqttAclResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Acl not implemented")
}

func RegisterMqttAclServiceServer(s *grpc.Server, srv MqttAclServiceServer) {
	s.RegisterService(&_MqttAclService_serviceDesc, srv)
}

func _MqttAclService_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MqttAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqttAclServiceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.MqttAclService/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqttAclServiceServer).Auth(ctx, req.(*MqttAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqttAclService_SuperUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MqttSuperuserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqttAclServiceServer).SuperUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.MqttAclService/SuperUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqttAclServiceServer).SuperUser(ctx, req.(*MqttSuperuserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MqttAclService_Acl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MqttAclRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MqttAclServiceServer).Acl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.MqttAclService/Acl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MqttAclServiceServer).Acl(ctx, req.(*MqttAclRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MqttAclService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.MqttAclService",
	HandlerType: (*MqttAclServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _MqttAclService_Auth_Handler,
		},
		{
			MethodName: "SuperUser",
			Handler:    _MqttAclService_SuperUser_Handler,
		},
		{
			MethodName: "Acl",
			Handler:    _MqttAclService_Acl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "external/service/mqtt.proto",
}
