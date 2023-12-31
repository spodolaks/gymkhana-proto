// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: gymkhana/rider.proto

package gymkhana

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

type Rider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Surname    string `protobuf:"bytes,5,opt,name=surname,proto3" json:"surname,omitempty"`
	LocationId string `protobuf:"bytes,9,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	Nickname   string `protobuf:"bytes,12,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Email      string `protobuf:"bytes,10,opt,name=email,proto3" json:"email,omitempty"`
	Phone      string `protobuf:"bytes,20,opt,name=phone,proto3" json:"phone,omitempty"`
	Age        int32  `protobuf:"varint,40,opt,name=age,proto3" json:"age,omitempty"`
	Weight     int32  `protobuf:"varint,50,opt,name=weight,proto3" json:"weight,omitempty"`
	Height     int32  `protobuf:"varint,60,opt,name=height,proto3" json:"height,omitempty"`
	Rank       string `protobuf:"bytes,70,opt,name=rank,proto3" json:"rank,omitempty"`
}

func (x *Rider) Reset() {
	*x = Rider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gymkhana_rider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rider) ProtoMessage() {}

func (x *Rider) ProtoReflect() protoreflect.Message {
	mi := &file_gymkhana_rider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rider.ProtoReflect.Descriptor instead.
func (*Rider) Descriptor() ([]byte, []int) {
	return file_gymkhana_rider_proto_rawDescGZIP(), []int{0}
}

func (x *Rider) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Rider) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Rider) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *Rider) GetLocationId() string {
	if x != nil {
		return x.LocationId
	}
	return ""
}

func (x *Rider) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *Rider) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Rider) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Rider) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Rider) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Rider) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Rider) GetRank() string {
	if x != nil {
		return x.Rank
	}
	return ""
}

type DeleteRiderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RiderId string `protobuf:"bytes,1,opt,name=rider_id,json=riderId,proto3" json:"rider_id,omitempty"`
}

func (x *DeleteRiderRequest) Reset() {
	*x = DeleteRiderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gymkhana_rider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRiderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRiderRequest) ProtoMessage() {}

func (x *DeleteRiderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gymkhana_rider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRiderRequest.ProtoReflect.Descriptor instead.
func (*DeleteRiderRequest) Descriptor() ([]byte, []int) {
	return file_gymkhana_rider_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteRiderRequest) GetRiderId() string {
	if x != nil {
		return x.RiderId
	}
	return ""
}

type DeleteRiderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RiderId string `protobuf:"bytes,1,opt,name=rider_id,json=riderId,proto3" json:"rider_id,omitempty"`
}

func (x *DeleteRiderResponse) Reset() {
	*x = DeleteRiderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gymkhana_rider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRiderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRiderResponse) ProtoMessage() {}

func (x *DeleteRiderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gymkhana_rider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRiderResponse.ProtoReflect.Descriptor instead.
func (*DeleteRiderResponse) Descriptor() ([]byte, []int) {
	return file_gymkhana_rider_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteRiderResponse) GetRiderId() string {
	if x != nil {
		return x.RiderId
	}
	return ""
}

type UpdateRiderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RiderId string `protobuf:"bytes,1,opt,name=rider_id,json=riderId,proto3" json:"rider_id,omitempty"`
}

func (x *UpdateRiderResponse) Reset() {
	*x = UpdateRiderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gymkhana_rider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRiderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRiderResponse) ProtoMessage() {}

func (x *UpdateRiderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gymkhana_rider_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRiderResponse.ProtoReflect.Descriptor instead.
func (*UpdateRiderResponse) Descriptor() ([]byte, []int) {
	return file_gymkhana_rider_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRiderResponse) GetRiderId() string {
	if x != nil {
		return x.RiderId
	}
	return ""
}

type RidersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *RidersRequest) Reset() {
	*x = RidersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gymkhana_rider_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RidersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RidersRequest) ProtoMessage() {}

func (x *RidersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gymkhana_rider_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RidersRequest.ProtoReflect.Descriptor instead.
func (*RidersRequest) Descriptor() ([]byte, []int) {
	return file_gymkhana_rider_proto_rawDescGZIP(), []int{4}
}

func (x *RidersRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type RidersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rider []*Rider `protobuf:"bytes,1,rep,name=Rider,proto3" json:"Rider,omitempty"`
}

func (x *RidersResponse) Reset() {
	*x = RidersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gymkhana_rider_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RidersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RidersResponse) ProtoMessage() {}

func (x *RidersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gymkhana_rider_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RidersResponse.ProtoReflect.Descriptor instead.
func (*RidersResponse) Descriptor() ([]byte, []int) {
	return file_gymkhana_rider_proto_rawDescGZIP(), []int{5}
}

func (x *RidersResponse) GetRider() []*Rider {
	if x != nil {
		return x.Rider
	}
	return nil
}

type RiderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RiderId string `protobuf:"bytes,1,opt,name=rider_id,json=riderId,proto3" json:"rider_id,omitempty"`
}

func (x *RiderRequest) Reset() {
	*x = RiderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gymkhana_rider_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RiderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiderRequest) ProtoMessage() {}

func (x *RiderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gymkhana_rider_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiderRequest.ProtoReflect.Descriptor instead.
func (*RiderRequest) Descriptor() ([]byte, []int) {
	return file_gymkhana_rider_proto_rawDescGZIP(), []int{6}
}

func (x *RiderRequest) GetRiderId() string {
	if x != nil {
		return x.RiderId
	}
	return ""
}

type RiderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rider *Rider `protobuf:"bytes,1,opt,name=Rider,proto3" json:"Rider,omitempty"`
}

func (x *RiderResponse) Reset() {
	*x = RiderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gymkhana_rider_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RiderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiderResponse) ProtoMessage() {}

func (x *RiderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gymkhana_rider_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiderResponse.ProtoReflect.Descriptor instead.
func (*RiderResponse) Descriptor() ([]byte, []int) {
	return file_gymkhana_rider_proto_rawDescGZIP(), []int{7}
}

func (x *RiderResponse) GetRider() *Rider {
	if x != nil {
		return x.Rider
	}
	return nil
}

type UpdateRiderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rider *Rider `protobuf:"bytes,1,opt,name=Rider,proto3" json:"Rider,omitempty"`
}

func (x *UpdateRiderRequest) Reset() {
	*x = UpdateRiderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gymkhana_rider_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRiderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRiderRequest) ProtoMessage() {}

func (x *UpdateRiderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gymkhana_rider_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRiderRequest.ProtoReflect.Descriptor instead.
func (*UpdateRiderRequest) Descriptor() ([]byte, []int) {
	return file_gymkhana_rider_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateRiderRequest) GetRider() *Rider {
	if x != nil {
		return x.Rider
	}
	return nil
}

type CreateRiderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rider *Rider `protobuf:"bytes,1,opt,name=Rider,proto3" json:"Rider,omitempty"`
}

func (x *CreateRiderRequest) Reset() {
	*x = CreateRiderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gymkhana_rider_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRiderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRiderRequest) ProtoMessage() {}

func (x *CreateRiderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gymkhana_rider_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRiderRequest.ProtoReflect.Descriptor instead.
func (*CreateRiderRequest) Descriptor() ([]byte, []int) {
	return file_gymkhana_rider_proto_rawDescGZIP(), []int{9}
}

func (x *CreateRiderRequest) GetRider() *Rider {
	if x != nil {
		return x.Rider
	}
	return nil
}

type CreateRiderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RiderId string `protobuf:"bytes,1,opt,name=rider_id,json=riderId,proto3" json:"rider_id,omitempty"`
}

func (x *CreateRiderResponse) Reset() {
	*x = CreateRiderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gymkhana_rider_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRiderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRiderResponse) ProtoMessage() {}

func (x *CreateRiderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gymkhana_rider_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRiderResponse.ProtoReflect.Descriptor instead.
func (*CreateRiderResponse) Descriptor() ([]byte, []int) {
	return file_gymkhana_rider_proto_rawDescGZIP(), []int{10}
}

func (x *CreateRiderResponse) GetRiderId() string {
	if x != nil {
		return x.RiderId
	}
	return ""
}

var File_gymkhana_rider_proto protoreflect.FileDescriptor

var file_gymkhana_rider_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x2f, 0x72, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61,
	0x22, 0x8d, 0x02, 0x0a, 0x05, 0x52, 0x69, 0x64, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67,
	0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x3c,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x61, 0x6e, 0x6b, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b,
	0x22, 0x2f, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x69, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x69, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x69, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x30, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x69, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x69, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x69, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x69, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x69,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x69,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x27, 0x0a, 0x0d, 0x52, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x37,
	0x0a, 0x0e, 0x52, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x05, 0x52, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x2e, 0x52, 0x69, 0x64, 0x65, 0x72,
	0x52, 0x05, 0x52, 0x69, 0x64, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x0c, 0x52, 0x69, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x69, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x69, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x36, 0x0a, 0x0d, 0x52, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x52, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x2e, 0x52, 0x69,
	0x64, 0x65, 0x72, 0x52, 0x05, 0x52, 0x69, 0x64, 0x65, 0x72, 0x22, 0x3b, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x25, 0x0a, 0x05, 0x52, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x2e, 0x52, 0x69, 0x64, 0x65, 0x72,
	0x52, 0x05, 0x52, 0x69, 0x64, 0x65, 0x72, 0x22, 0x3b, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a,
	0x05, 0x52, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67,
	0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x2e, 0x52, 0x69, 0x64, 0x65, 0x72, 0x52, 0x05, 0x52,
	0x69, 0x64, 0x65, 0x72, 0x22, 0x30, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x69,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72,
	0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x32, 0xda, 0x02, 0x0a, 0x0c, 0x52, 0x69, 0x64, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x1c, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x06, 0x52, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x17, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68,
	0x61, 0x6e, 0x61, 0x2e, 0x52, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x2e, 0x52, 0x69, 0x64,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x52,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x2e,
	0x52, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67,
	0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x2e, 0x52, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1c, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e,
	0x61, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x70, 0x6f, 0x64, 0x6f, 0x6c, 0x61, 0x6b, 0x73, 0x2f, 0x67, 0x79, 0x6d, 0x6b,
	0x68, 0x61, 0x6e, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x79, 0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x3b, 0x67, 0x79,
	0x6d, 0x6b, 0x68, 0x61, 0x6e, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gymkhana_rider_proto_rawDescOnce sync.Once
	file_gymkhana_rider_proto_rawDescData = file_gymkhana_rider_proto_rawDesc
)

func file_gymkhana_rider_proto_rawDescGZIP() []byte {
	file_gymkhana_rider_proto_rawDescOnce.Do(func() {
		file_gymkhana_rider_proto_rawDescData = protoimpl.X.CompressGZIP(file_gymkhana_rider_proto_rawDescData)
	})
	return file_gymkhana_rider_proto_rawDescData
}

var file_gymkhana_rider_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_gymkhana_rider_proto_goTypes = []interface{}{
	(*Rider)(nil),               // 0: gymkhana.Rider
	(*DeleteRiderRequest)(nil),  // 1: gymkhana.DeleteRiderRequest
	(*DeleteRiderResponse)(nil), // 2: gymkhana.DeleteRiderResponse
	(*UpdateRiderResponse)(nil), // 3: gymkhana.UpdateRiderResponse
	(*RidersRequest)(nil),       // 4: gymkhana.RidersRequest
	(*RidersResponse)(nil),      // 5: gymkhana.RidersResponse
	(*RiderRequest)(nil),        // 6: gymkhana.RiderRequest
	(*RiderResponse)(nil),       // 7: gymkhana.RiderResponse
	(*UpdateRiderRequest)(nil),  // 8: gymkhana.UpdateRiderRequest
	(*CreateRiderRequest)(nil),  // 9: gymkhana.CreateRiderRequest
	(*CreateRiderResponse)(nil), // 10: gymkhana.CreateRiderResponse
}
var file_gymkhana_rider_proto_depIdxs = []int32{
	0,  // 0: gymkhana.RidersResponse.Rider:type_name -> gymkhana.Rider
	0,  // 1: gymkhana.RiderResponse.Rider:type_name -> gymkhana.Rider
	0,  // 2: gymkhana.UpdateRiderRequest.Rider:type_name -> gymkhana.Rider
	0,  // 3: gymkhana.CreateRiderRequest.Rider:type_name -> gymkhana.Rider
	9,  // 4: gymkhana.RiderService.Create:input_type -> gymkhana.CreateRiderRequest
	4,  // 5: gymkhana.RiderService.Riders:input_type -> gymkhana.RidersRequest
	6,  // 6: gymkhana.RiderService.Rider:input_type -> gymkhana.RiderRequest
	8,  // 7: gymkhana.RiderService.Update:input_type -> gymkhana.UpdateRiderRequest
	1,  // 8: gymkhana.RiderService.Delete:input_type -> gymkhana.DeleteRiderRequest
	10, // 9: gymkhana.RiderService.Create:output_type -> gymkhana.CreateRiderResponse
	5,  // 10: gymkhana.RiderService.Riders:output_type -> gymkhana.RidersResponse
	7,  // 11: gymkhana.RiderService.Rider:output_type -> gymkhana.RiderResponse
	3,  // 12: gymkhana.RiderService.Update:output_type -> gymkhana.UpdateRiderResponse
	2,  // 13: gymkhana.RiderService.Delete:output_type -> gymkhana.DeleteRiderResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_gymkhana_rider_proto_init() }
func file_gymkhana_rider_proto_init() {
	if File_gymkhana_rider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gymkhana_rider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rider); i {
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
		file_gymkhana_rider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRiderRequest); i {
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
		file_gymkhana_rider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRiderResponse); i {
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
		file_gymkhana_rider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRiderResponse); i {
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
		file_gymkhana_rider_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RidersRequest); i {
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
		file_gymkhana_rider_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RidersResponse); i {
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
		file_gymkhana_rider_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RiderRequest); i {
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
		file_gymkhana_rider_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RiderResponse); i {
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
		file_gymkhana_rider_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRiderRequest); i {
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
		file_gymkhana_rider_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRiderRequest); i {
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
		file_gymkhana_rider_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRiderResponse); i {
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
			RawDescriptor: file_gymkhana_rider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gymkhana_rider_proto_goTypes,
		DependencyIndexes: file_gymkhana_rider_proto_depIdxs,
		MessageInfos:      file_gymkhana_rider_proto_msgTypes,
	}.Build()
	File_gymkhana_rider_proto = out.File
	file_gymkhana_rider_proto_rawDesc = nil
	file_gymkhana_rider_proto_goTypes = nil
	file_gymkhana_rider_proto_depIdxs = nil
}
