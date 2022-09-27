// Code generated with goa v3.8.5, DO NOT EDIT.
//
// character protocol buffer definition
//
// Command:
// $ goa gen characters/design

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: goadesign_goagen_character.proto

package characterpb

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

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_character_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_character_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_character_proto_rawDescGZIP(), []int{0}
}

type StoredCharacterCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field []*StoredCharacter `protobuf:"bytes,1,rep,name=field,proto3" json:"field,omitempty"`
}

func (x *StoredCharacterCollection) Reset() {
	*x = StoredCharacterCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_character_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoredCharacterCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoredCharacterCollection) ProtoMessage() {}

func (x *StoredCharacterCollection) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_character_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoredCharacterCollection.ProtoReflect.Descriptor instead.
func (*StoredCharacterCollection) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_character_proto_rawDescGZIP(), []int{1}
}

func (x *StoredCharacterCollection) GetField() []*StoredCharacter {
	if x != nil {
		return x.Field
	}
	return nil
}

// A StoredCharacter describes a character retrieved by the db.
type StoredCharacter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the unique id of the character.
	Id string `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
	// Name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Health
	Health float64 `protobuf:"fixed64,4,opt,name=health,proto3" json:"health,omitempty"`
	// Experience
	Experience float64 `protobuf:"fixed64,5,opt,name=experience,proto3" json:"experience,omitempty"`
}

func (x *StoredCharacter) Reset() {
	*x = StoredCharacter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_character_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoredCharacter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoredCharacter) ProtoMessage() {}

func (x *StoredCharacter) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_character_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoredCharacter.ProtoReflect.Descriptor instead.
func (*StoredCharacter) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_character_proto_rawDescGZIP(), []int{2}
}

func (x *StoredCharacter) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StoredCharacter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StoredCharacter) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *StoredCharacter) GetHealth() float64 {
	if x != nil {
		return x.Health
	}
	return 0
}

func (x *StoredCharacter) GetExperience() float64 {
	if x != nil {
		return x.Experience
	}
	return 0
}

type ShowNotFoundError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Message of error
	Message_ string `protobuf:"bytes,1,opt,name=message_,json=message,proto3" json:"message_,omitempty"`
	// ID of missing resource
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ShowNotFoundError) Reset() {
	*x = ShowNotFoundError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_character_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowNotFoundError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowNotFoundError) ProtoMessage() {}

func (x *ShowNotFoundError) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_character_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowNotFoundError.ProtoReflect.Descriptor instead.
func (*ShowNotFoundError) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_character_proto_rawDescGZIP(), []int{3}
}

func (x *ShowNotFoundError) GetMessage_() string {
	if x != nil {
		return x.Message_
	}
	return ""
}

func (x *ShowNotFoundError) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ShowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of character to show
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ShowRequest) Reset() {
	*x = ShowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_character_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowRequest) ProtoMessage() {}

func (x *ShowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_character_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowRequest.ProtoReflect.Descriptor instead.
func (*ShowRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_character_proto_rawDescGZIP(), []int{4}
}

func (x *ShowRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ShowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the unique id of the character.
	Id string `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
	// Name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Health
	Health float64 `protobuf:"fixed64,4,opt,name=health,proto3" json:"health,omitempty"`
	// Experience
	Experience float64 `protobuf:"fixed64,5,opt,name=experience,proto3" json:"experience,omitempty"`
}

func (x *ShowResponse) Reset() {
	*x = ShowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_character_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowResponse) ProtoMessage() {}

func (x *ShowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_character_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowResponse.ProtoReflect.Descriptor instead.
func (*ShowResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_character_proto_rawDescGZIP(), []int{5}
}

func (x *ShowResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShowResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ShowResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ShowResponse) GetHealth() float64 {
	if x != nil {
		return x.Health
	}
	return 0
}

func (x *ShowResponse) GetExperience() float64 {
	if x != nil {
		return x.Experience
	}
	return 0
}

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Description
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Health
	Health float64 `protobuf:"fixed64,3,opt,name=health,proto3" json:"health,omitempty"`
	// Experience
	Experience float64 `protobuf:"fixed64,4,opt,name=experience,proto3" json:"experience,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_character_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_character_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_character_proto_rawDescGZIP(), []int{6}
}

func (x *AddRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddRequest) GetHealth() float64 {
	if x != nil {
		return x.Health
	}
	return 0
}

func (x *AddRequest) GetExperience() float64 {
	if x != nil {
		return x.Experience
	}
	return 0
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_character_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_character_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_character_proto_rawDescGZIP(), []int{7}
}

func (x *AddResponse) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type RemoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of character to remove
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveRequest) Reset() {
	*x = RemoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_character_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRequest) ProtoMessage() {}

func (x *RemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_character_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRequest.ProtoReflect.Descriptor instead.
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_character_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RemoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveResponse) Reset() {
	*x = RemoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_character_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveResponse) ProtoMessage() {}

func (x *RemoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_character_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveResponse.ProtoReflect.Descriptor instead.
func (*RemoveResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_character_proto_rawDescGZIP(), []int{9}
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of character to update
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// character to update
	Character *Character2 `protobuf:"bytes,2,opt,name=character,proto3" json:"character,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_character_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_character_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_character_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRequest) GetCharacter() *Character2 {
	if x != nil {
		return x.Character
	}
	return nil
}

// Character
type Character2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Description
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Health
	Health float64 `protobuf:"fixed64,3,opt,name=health,proto3" json:"health,omitempty"`
	// Experience
	Experience float64 `protobuf:"fixed64,4,opt,name=experience,proto3" json:"experience,omitempty"`
}

func (x *Character2) Reset() {
	*x = Character2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_character_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Character2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Character2) ProtoMessage() {}

func (x *Character2) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_character_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Character2.ProtoReflect.Descriptor instead.
func (*Character2) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_character_proto_rawDescGZIP(), []int{11}
}

func (x *Character2) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Character2) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Character2) GetHealth() float64 {
	if x != nil {
		return x.Health
	}
	return 0
}

func (x *Character2) GetExperience() float64 {
	if x != nil {
		return x.Experience
	}
	return 0
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the unique id of the character.
	Id string `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
	// Name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Health
	Health float64 `protobuf:"fixed64,4,opt,name=health,proto3" json:"health,omitempty"`
	// Experience
	Experience float64 `protobuf:"fixed64,5,opt,name=experience,proto3" json:"experience,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_character_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_character_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_character_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateResponse) GetHealth() float64 {
	if x != nil {
		return x.Health
	}
	return 0
}

func (x *UpdateResponse) GetExperience() float64 {
	if x != nil {
		return x.Experience
	}
	return 0
}

var File_goadesign_goagen_character_proto protoreflect.FileDescriptor

var file_goadesign_goagen_character_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x6f, 0x61, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x67, 0x6f, 0x61, 0x67,
	0x65, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x22, 0x0d, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4d, 0x0a, 0x19,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x8f, 0x01, 0x0a, 0x0f,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x1e, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x3e, 0x0a,
	0x11, 0x53, 0x68, 0x6f, 0x77, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1d, 0x0a,
	0x0b, 0x53, 0x68, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8c, 0x01, 0x0a,
	0x0c, 0x53, 0x68, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x7a, 0x0a, 0x0a, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x23, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x1f, 0x0a, 0x0d,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x10, 0x0a,
	0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x54, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x33, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e,
	0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x32, 0x52, 0x09, 0x63, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x22, 0x7a, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63,
	0x65, 0x22, 0x8e, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x32, 0xbe, 0x02, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x12, 0x44, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x64, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x04, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x16,
	0x2e, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x15, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12,
	0x18, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18,
	0x2e, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goadesign_goagen_character_proto_rawDescOnce sync.Once
	file_goadesign_goagen_character_proto_rawDescData = file_goadesign_goagen_character_proto_rawDesc
)

func file_goadesign_goagen_character_proto_rawDescGZIP() []byte {
	file_goadesign_goagen_character_proto_rawDescOnce.Do(func() {
		file_goadesign_goagen_character_proto_rawDescData = protoimpl.X.CompressGZIP(file_goadesign_goagen_character_proto_rawDescData)
	})
	return file_goadesign_goagen_character_proto_rawDescData
}

var file_goadesign_goagen_character_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_goadesign_goagen_character_proto_goTypes = []interface{}{
	(*ListRequest)(nil),               // 0: character.ListRequest
	(*StoredCharacterCollection)(nil), // 1: character.StoredCharacterCollection
	(*StoredCharacter)(nil),           // 2: character.StoredCharacter
	(*ShowNotFoundError)(nil),         // 3: character.ShowNotFoundError
	(*ShowRequest)(nil),               // 4: character.ShowRequest
	(*ShowResponse)(nil),              // 5: character.ShowResponse
	(*AddRequest)(nil),                // 6: character.AddRequest
	(*AddResponse)(nil),               // 7: character.AddResponse
	(*RemoveRequest)(nil),             // 8: character.RemoveRequest
	(*RemoveResponse)(nil),            // 9: character.RemoveResponse
	(*UpdateRequest)(nil),             // 10: character.UpdateRequest
	(*Character2)(nil),                // 11: character.Character2
	(*UpdateResponse)(nil),            // 12: character.UpdateResponse
}
var file_goadesign_goagen_character_proto_depIdxs = []int32{
	2,  // 0: character.StoredCharacterCollection.field:type_name -> character.StoredCharacter
	11, // 1: character.UpdateRequest.character:type_name -> character.Character2
	0,  // 2: character.Character.List:input_type -> character.ListRequest
	4,  // 3: character.Character.Show:input_type -> character.ShowRequest
	6,  // 4: character.Character.Add:input_type -> character.AddRequest
	8,  // 5: character.Character.Remove:input_type -> character.RemoveRequest
	10, // 6: character.Character.Update:input_type -> character.UpdateRequest
	1,  // 7: character.Character.List:output_type -> character.StoredCharacterCollection
	5,  // 8: character.Character.Show:output_type -> character.ShowResponse
	7,  // 9: character.Character.Add:output_type -> character.AddResponse
	9,  // 10: character.Character.Remove:output_type -> character.RemoveResponse
	12, // 11: character.Character.Update:output_type -> character.UpdateResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_goadesign_goagen_character_proto_init() }
func file_goadesign_goagen_character_proto_init() {
	if File_goadesign_goagen_character_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goadesign_goagen_character_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_goadesign_goagen_character_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoredCharacterCollection); i {
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
		file_goadesign_goagen_character_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoredCharacter); i {
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
		file_goadesign_goagen_character_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowNotFoundError); i {
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
		file_goadesign_goagen_character_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowRequest); i {
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
		file_goadesign_goagen_character_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShowResponse); i {
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
		file_goadesign_goagen_character_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_goadesign_goagen_character_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResponse); i {
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
		file_goadesign_goagen_character_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRequest); i {
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
		file_goadesign_goagen_character_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveResponse); i {
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
		file_goadesign_goagen_character_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_goadesign_goagen_character_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Character2); i {
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
		file_goadesign_goagen_character_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
			RawDescriptor: file_goadesign_goagen_character_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goadesign_goagen_character_proto_goTypes,
		DependencyIndexes: file_goadesign_goagen_character_proto_depIdxs,
		MessageInfos:      file_goadesign_goagen_character_proto_msgTypes,
	}.Build()
	File_goadesign_goagen_character_proto = out.File
	file_goadesign_goagen_character_proto_rawDesc = nil
	file_goadesign_goagen_character_proto_goTypes = nil
	file_goadesign_goagen_character_proto_depIdxs = nil
}
