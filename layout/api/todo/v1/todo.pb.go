// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: api/todo/v1/todo.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Detail   string                 `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
	Deadline *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Status   int32                  `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_todo_v1_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_api_todo_v1_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_api_todo_v1_todo_proto_rawDescGZIP(), []int{0}
}

func (x *Todo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Todo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Todo) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *Todo) GetDeadline() *timestamppb.Timestamp {
	if x != nil {
		return x.Deadline
	}
	return nil
}

func (x *Todo) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type CreateTodoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *CreateTodoReq) Reset() {
	*x = CreateTodoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_todo_v1_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTodoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoReq) ProtoMessage() {}

func (x *CreateTodoReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_todo_v1_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoReq.ProtoReflect.Descriptor instead.
func (*CreateTodoReq) Descriptor() ([]byte, []int) {
	return file_api_todo_v1_todo_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTodoReq) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type CreateTodoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateTodoReply) Reset() {
	*x = CreateTodoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_todo_v1_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTodoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoReply) ProtoMessage() {}

func (x *CreateTodoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_todo_v1_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoReply.ProtoReflect.Descriptor instead.
func (*CreateTodoReply) Descriptor() ([]byte, []int) {
	return file_api_todo_v1_todo_proto_rawDescGZIP(), []int{2}
}

type UpdateTodoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo                  `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	Mask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=mask,proto3" json:"mask,omitempty"`
}

func (x *UpdateTodoReq) Reset() {
	*x = UpdateTodoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_todo_v1_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTodoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTodoReq) ProtoMessage() {}

func (x *UpdateTodoReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_todo_v1_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTodoReq.ProtoReflect.Descriptor instead.
func (*UpdateTodoReq) Descriptor() ([]byte, []int) {
	return file_api_todo_v1_todo_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateTodoReq) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

func (x *UpdateTodoReq) GetMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.Mask
	}
	return nil
}

type UpdateTodoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *UpdateTodoReply) Reset() {
	*x = UpdateTodoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_todo_v1_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTodoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTodoReply) ProtoMessage() {}

func (x *UpdateTodoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_todo_v1_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTodoReply.ProtoReflect.Descriptor instead.
func (*UpdateTodoReply) Descriptor() ([]byte, []int) {
	return file_api_todo_v1_todo_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTodoReply) GetTodo() *Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type ListTodoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListTodoReq) Reset() {
	*x = ListTodoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_todo_v1_todo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTodoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTodoReq) ProtoMessage() {}

func (x *ListTodoReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_todo_v1_todo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTodoReq.ProtoReflect.Descriptor instead.
func (*ListTodoReq) Descriptor() ([]byte, []int) {
	return file_api_todo_v1_todo_proto_rawDescGZIP(), []int{5}
}

type ListTodoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Todo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ListTodoReply) Reset() {
	*x = ListTodoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_todo_v1_todo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTodoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTodoReply) ProtoMessage() {}

func (x *ListTodoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_todo_v1_todo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTodoReply.ProtoReflect.Descriptor instead.
func (*ListTodoReply) Descriptor() ([]byte, []int) {
	return file_api_todo_v1_todo_proto_rawDescGZIP(), []int{6}
}

func (x *ListTodoReply) GetResults() []*Todo {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_api_todo_v1_todo_proto protoreflect.FileDescriptor

var file_api_todo_v1_todo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f,
	0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x04, 0x54, 0x6f, 0x64,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x36, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64,
	0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x36, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71,
	0x12, 0x25, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64,
	0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x66, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x04, 0x74,
	0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x74, 0x6f,
	0x64, 0x6f, 0x12, 0x2e, 0x0a, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x6d, 0x61,
	0x73, 0x6b, 0x22, 0x38, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x0d, 0x0a, 0x0b,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x22, 0x3c, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f,
	0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x32, 0xe5, 0x01, 0x0a, 0x0b, 0x54, 0x6f,
	0x64, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f,
	0x64, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f,
	0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64,
	0x6f, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x42, 0x0a,
	0x08, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f,
	0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x10, 0x5a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_todo_v1_todo_proto_rawDescOnce sync.Once
	file_api_todo_v1_todo_proto_rawDescData = file_api_todo_v1_todo_proto_rawDesc
)

func file_api_todo_v1_todo_proto_rawDescGZIP() []byte {
	file_api_todo_v1_todo_proto_rawDescOnce.Do(func() {
		file_api_todo_v1_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_todo_v1_todo_proto_rawDescData)
	})
	return file_api_todo_v1_todo_proto_rawDescData
}

var file_api_todo_v1_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_todo_v1_todo_proto_goTypes = []interface{}{
	(*Todo)(nil),                  // 0: api.todo.v1.Todo
	(*CreateTodoReq)(nil),         // 1: api.todo.v1.CreateTodoReq
	(*CreateTodoReply)(nil),       // 2: api.todo.v1.CreateTodoReply
	(*UpdateTodoReq)(nil),         // 3: api.todo.v1.UpdateTodoReq
	(*UpdateTodoReply)(nil),       // 4: api.todo.v1.UpdateTodoReply
	(*ListTodoReq)(nil),           // 5: api.todo.v1.ListTodoReq
	(*ListTodoReply)(nil),         // 6: api.todo.v1.ListTodoReply
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil), // 8: google.protobuf.FieldMask
}
var file_api_todo_v1_todo_proto_depIdxs = []int32{
	7, // 0: api.todo.v1.Todo.deadline:type_name -> google.protobuf.Timestamp
	0, // 1: api.todo.v1.CreateTodoReq.todo:type_name -> api.todo.v1.Todo
	0, // 2: api.todo.v1.UpdateTodoReq.todo:type_name -> api.todo.v1.Todo
	8, // 3: api.todo.v1.UpdateTodoReq.mask:type_name -> google.protobuf.FieldMask
	0, // 4: api.todo.v1.UpdateTodoReply.todo:type_name -> api.todo.v1.Todo
	0, // 5: api.todo.v1.ListTodoReply.results:type_name -> api.todo.v1.Todo
	1, // 6: api.todo.v1.TodoService.CreateTodo:input_type -> api.todo.v1.CreateTodoReq
	3, // 7: api.todo.v1.TodoService.UpdateTodo:input_type -> api.todo.v1.UpdateTodoReq
	5, // 8: api.todo.v1.TodoService.ListTodo:input_type -> api.todo.v1.ListTodoReq
	2, // 9: api.todo.v1.TodoService.CreateTodo:output_type -> api.todo.v1.CreateTodoReply
	4, // 10: api.todo.v1.TodoService.UpdateTodo:output_type -> api.todo.v1.UpdateTodoReply
	6, // 11: api.todo.v1.TodoService.ListTodo:output_type -> api.todo.v1.ListTodoReply
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_todo_v1_todo_proto_init() }
func file_api_todo_v1_todo_proto_init() {
	if File_api_todo_v1_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_todo_v1_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
		file_api_todo_v1_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTodoReq); i {
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
		file_api_todo_v1_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTodoReply); i {
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
		file_api_todo_v1_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTodoReq); i {
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
		file_api_todo_v1_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTodoReply); i {
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
		file_api_todo_v1_todo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTodoReq); i {
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
		file_api_todo_v1_todo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTodoReply); i {
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
			RawDescriptor: file_api_todo_v1_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_todo_v1_todo_proto_goTypes,
		DependencyIndexes: file_api_todo_v1_todo_proto_depIdxs,
		MessageInfos:      file_api_todo_v1_todo_proto_msgTypes,
	}.Build()
	File_api_todo_v1_todo_proto = out.File
	file_api_todo_v1_todo_proto_rawDesc = nil
	file_api_todo_v1_todo_proto_goTypes = nil
	file_api_todo_v1_todo_proto_depIdxs = nil
}