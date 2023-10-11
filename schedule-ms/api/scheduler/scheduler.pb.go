// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: api/scheduler/scheduler.proto

package scheduler

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

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

// Enum value maps for HealthCheckResponse_ServingStatus.
var (
	HealthCheckResponse_ServingStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "SERVING",
		2: "NOT_SERVING",
	}
	HealthCheckResponse_ServingStatus_value = map[string]int32{
		"UNKNOWN":     0,
		"SERVING":     1,
		"NOT_SERVING": 2,
	}
)

func (x HealthCheckResponse_ServingStatus) Enum() *HealthCheckResponse_ServingStatus {
	p := new(HealthCheckResponse_ServingStatus)
	*p = x
	return p
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthCheckResponse_ServingStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_scheduler_scheduler_proto_enumTypes[0].Descriptor()
}

func (HealthCheckResponse_ServingStatus) Type() protoreflect.EnumType {
	return &file_api_scheduler_scheduler_proto_enumTypes[0]
}

func (x HealthCheckResponse_ServingStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthCheckResponse_ServingStatus.Descriptor instead.
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_scheduler_scheduler_proto_rawDescGZIP(), []int{10, 0}
}

type Appointment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApptID       string `protobuf:"bytes,1,opt,name=apptID,proto3" json:"apptID,omitempty"`
	PatientName  string `protobuf:"bytes,2,opt,name=patientName,proto3" json:"patientName,omitempty"`
	StaffID      string `protobuf:"bytes,3,opt,name=staffID,proto3" json:"staffID,omitempty"`
	ApptDateTime string `protobuf:"bytes,4,opt,name=apptDateTime,proto3" json:"apptDateTime,omitempty"`
	ApptType     string `protobuf:"bytes,5,opt,name=apptType,proto3" json:"apptType,omitempty"`
}

func (x *Appointment) Reset() {
	*x = Appointment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_scheduler_scheduler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Appointment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Appointment) ProtoMessage() {}

func (x *Appointment) ProtoReflect() protoreflect.Message {
	mi := &file_api_scheduler_scheduler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Appointment.ProtoReflect.Descriptor instead.
func (*Appointment) Descriptor() ([]byte, []int) {
	return file_api_scheduler_scheduler_proto_rawDescGZIP(), []int{0}
}

func (x *Appointment) GetApptID() string {
	if x != nil {
		return x.ApptID
	}
	return ""
}

func (x *Appointment) GetPatientName() string {
	if x != nil {
		return x.PatientName
	}
	return ""
}

func (x *Appointment) GetStaffID() string {
	if x != nil {
		return x.StaffID
	}
	return ""
}

func (x *Appointment) GetApptDateTime() string {
	if x != nil {
		return x.ApptDateTime
	}
	return ""
}

func (x *Appointment) GetApptType() string {
	if x != nil {
		return x.ApptType
	}
	return ""
}

type CreateAppointment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appointment *Appointment `protobuf:"bytes,1,opt,name=appointment,proto3" json:"appointment,omitempty"`
}

func (x *CreateAppointment) Reset() {
	*x = CreateAppointment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_scheduler_scheduler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAppointment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppointment) ProtoMessage() {}

func (x *CreateAppointment) ProtoReflect() protoreflect.Message {
	mi := &file_api_scheduler_scheduler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppointment.ProtoReflect.Descriptor instead.
func (*CreateAppointment) Descriptor() ([]byte, []int) {
	return file_api_scheduler_scheduler_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAppointment) GetAppointment() *Appointment {
	if x != nil {
		return x.Appointment
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApptID string `protobuf:"bytes,1,opt,name=apptID,proto3" json:"apptID,omitempty"`
	Status uint64 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_scheduler_scheduler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_scheduler_scheduler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_api_scheduler_scheduler_proto_rawDescGZIP(), []int{2}
}

func (x *CreateResponse) GetApptID() string {
	if x != nil {
		return x.ApptID
	}
	return ""
}

func (x *CreateResponse) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetAppointments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAppointments) Reset() {
	*x = GetAppointments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_scheduler_scheduler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppointments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppointments) ProtoMessage() {}

func (x *GetAppointments) ProtoReflect() protoreflect.Message {
	mi := &file_api_scheduler_scheduler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppointments.ProtoReflect.Descriptor instead.
func (*GetAppointments) Descriptor() ([]byte, []int) {
	return file_api_scheduler_scheduler_proto_rawDescGZIP(), []int{3}
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appointments []*Appointment `protobuf:"bytes,1,rep,name=appointments,proto3" json:"appointments,omitempty"`
	Status       uint64         `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Error        string         `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_scheduler_scheduler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_scheduler_scheduler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_api_scheduler_scheduler_proto_rawDescGZIP(), []int{4}
}

func (x *GetResponse) GetAppointments() []*Appointment {
	if x != nil {
		return x.Appointments
	}
	return nil
}

func (x *GetResponse) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type UpdateAppointment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appointment *Appointment `protobuf:"bytes,1,opt,name=appointment,proto3" json:"appointment,omitempty"`
}

func (x *UpdateAppointment) Reset() {
	*x = UpdateAppointment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_scheduler_scheduler_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAppointment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAppointment) ProtoMessage() {}

func (x *UpdateAppointment) ProtoReflect() protoreflect.Message {
	mi := &file_api_scheduler_scheduler_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAppointment.ProtoReflect.Descriptor instead.
func (*UpdateAppointment) Descriptor() ([]byte, []int) {
	return file_api_scheduler_scheduler_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateAppointment) GetAppointment() *Appointment {
	if x != nil {
		return x.Appointment
	}
	return nil
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appointment *Appointment `protobuf:"bytes,1,opt,name=appointment,proto3" json:"appointment,omitempty"`
	Status      uint64       `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Error       string       `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_scheduler_scheduler_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_scheduler_scheduler_proto_msgTypes[6]
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
	return file_api_scheduler_scheduler_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateResponse) GetAppointment() *Appointment {
	if x != nil {
		return x.Appointment
	}
	return nil
}

func (x *UpdateResponse) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UpdateResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type DeleteAppointment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApptID string `protobuf:"bytes,1,opt,name=apptID,proto3" json:"apptID,omitempty"`
}

func (x *DeleteAppointment) Reset() {
	*x = DeleteAppointment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_scheduler_scheduler_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAppointment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAppointment) ProtoMessage() {}

func (x *DeleteAppointment) ProtoReflect() protoreflect.Message {
	mi := &file_api_scheduler_scheduler_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAppointment.ProtoReflect.Descriptor instead.
func (*DeleteAppointment) Descriptor() ([]byte, []int) {
	return file_api_scheduler_scheduler_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteAppointment) GetApptID() string {
	if x != nil {
		return x.ApptID
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApptID string `protobuf:"bytes,1,opt,name=apptID,proto3" json:"apptID,omitempty"`
	Status uint64 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_scheduler_scheduler_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_scheduler_scheduler_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_api_scheduler_scheduler_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteResponse) GetApptID() string {
	if x != nil {
		return x.ApptID
	}
	return ""
}

func (x *DeleteResponse) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DeleteResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_scheduler_scheduler_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_scheduler_scheduler_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_api_scheduler_scheduler_proto_rawDescGZIP(), []int{9}
}

func (x *HealthCheckRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=scheduler.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_scheduler_scheduler_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_scheduler_scheduler_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_api_scheduler_scheduler_proto_rawDescGZIP(), []int{10}
}

func (x *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if x != nil {
		return x.Status
	}
	return HealthCheckResponse_UNKNOWN
}

var File_api_scheduler_scheduler_proto protoreflect.FileDescriptor

var file_api_scheduler_scheduler_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x22, 0xa1, 0x01, 0x0a, 0x0b, 0x41,
	0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70,
	0x70, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x70, 0x74,
	0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x49, 0x44, 0x12, 0x22,
	0x0a, 0x0c, 0x61, 0x70, 0x70, 0x74, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x74, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x70, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4d,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x0b, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x56, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x70, 0x70, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x77, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x4d, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0x78, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x0b, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x2b, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x70, 0x70, 0x74, 0x49, 0x44, 0x22, 0x56, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x70,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x70, 0x74, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x2e, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22,
	0x97, 0x01, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3a, 0x0a,
	0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x32, 0xe7, 0x02, 0x0a, 0x09, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x70, 0x70, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x74, 0x12, 0x1a, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x16, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70,
	0x70, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x1a, 0x19, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x05, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x12, 0x1d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_scheduler_scheduler_proto_rawDescOnce sync.Once
	file_api_scheduler_scheduler_proto_rawDescData = file_api_scheduler_scheduler_proto_rawDesc
)

func file_api_scheduler_scheduler_proto_rawDescGZIP() []byte {
	file_api_scheduler_scheduler_proto_rawDescOnce.Do(func() {
		file_api_scheduler_scheduler_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_scheduler_scheduler_proto_rawDescData)
	})
	return file_api_scheduler_scheduler_proto_rawDescData
}

var file_api_scheduler_scheduler_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_scheduler_scheduler_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_scheduler_scheduler_proto_goTypes = []interface{}{
	(HealthCheckResponse_ServingStatus)(0), // 0: scheduler.HealthCheckResponse.ServingStatus
	(*Appointment)(nil),                    // 1: scheduler.Appointment
	(*CreateAppointment)(nil),              // 2: scheduler.CreateAppointment
	(*CreateResponse)(nil),                 // 3: scheduler.CreateResponse
	(*GetAppointments)(nil),                // 4: scheduler.GetAppointments
	(*GetResponse)(nil),                    // 5: scheduler.GetResponse
	(*UpdateAppointment)(nil),              // 6: scheduler.UpdateAppointment
	(*UpdateResponse)(nil),                 // 7: scheduler.UpdateResponse
	(*DeleteAppointment)(nil),              // 8: scheduler.DeleteAppointment
	(*DeleteResponse)(nil),                 // 9: scheduler.DeleteResponse
	(*HealthCheckRequest)(nil),             // 10: scheduler.HealthCheckRequest
	(*HealthCheckResponse)(nil),            // 11: scheduler.HealthCheckResponse
}
var file_api_scheduler_scheduler_proto_depIdxs = []int32{
	1,  // 0: scheduler.CreateAppointment.appointment:type_name -> scheduler.Appointment
	1,  // 1: scheduler.GetResponse.appointments:type_name -> scheduler.Appointment
	1,  // 2: scheduler.UpdateAppointment.appointment:type_name -> scheduler.Appointment
	1,  // 3: scheduler.UpdateResponse.appointment:type_name -> scheduler.Appointment
	0,  // 4: scheduler.HealthCheckResponse.status:type_name -> scheduler.HealthCheckResponse.ServingStatus
	2,  // 5: scheduler.Scheduler.CreateAppt:input_type -> scheduler.CreateAppointment
	4,  // 6: scheduler.Scheduler.GetAppt:input_type -> scheduler.GetAppointments
	6,  // 7: scheduler.Scheduler.UpdateAppt:input_type -> scheduler.UpdateAppointment
	8,  // 8: scheduler.Scheduler.DeleteAppt:input_type -> scheduler.DeleteAppointment
	10, // 9: scheduler.Scheduler.Check:input_type -> scheduler.HealthCheckRequest
	3,  // 10: scheduler.Scheduler.CreateAppt:output_type -> scheduler.CreateResponse
	5,  // 11: scheduler.Scheduler.GetAppt:output_type -> scheduler.GetResponse
	7,  // 12: scheduler.Scheduler.UpdateAppt:output_type -> scheduler.UpdateResponse
	9,  // 13: scheduler.Scheduler.DeleteAppt:output_type -> scheduler.DeleteResponse
	11, // 14: scheduler.Scheduler.Check:output_type -> scheduler.HealthCheckResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_scheduler_scheduler_proto_init() }
func file_api_scheduler_scheduler_proto_init() {
	if File_api_scheduler_scheduler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_scheduler_scheduler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Appointment); i {
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
		file_api_scheduler_scheduler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAppointment); i {
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
		file_api_scheduler_scheduler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_api_scheduler_scheduler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppointments); i {
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
		file_api_scheduler_scheduler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_api_scheduler_scheduler_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAppointment); i {
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
		file_api_scheduler_scheduler_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_scheduler_scheduler_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAppointment); i {
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
		file_api_scheduler_scheduler_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_api_scheduler_scheduler_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
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
		file_api_scheduler_scheduler_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
			RawDescriptor: file_api_scheduler_scheduler_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_scheduler_scheduler_proto_goTypes,
		DependencyIndexes: file_api_scheduler_scheduler_proto_depIdxs,
		EnumInfos:         file_api_scheduler_scheduler_proto_enumTypes,
		MessageInfos:      file_api_scheduler_scheduler_proto_msgTypes,
	}.Build()
	File_api_scheduler_scheduler_proto = out.File
	file_api_scheduler_scheduler_proto_rawDesc = nil
	file_api_scheduler_scheduler_proto_goTypes = nil
	file_api_scheduler_scheduler_proto_depIdxs = nil
}
