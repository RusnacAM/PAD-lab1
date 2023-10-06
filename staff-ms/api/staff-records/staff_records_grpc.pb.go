// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: api/staff-records/staff_records.proto

package staff_records

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	StaffRecords_Create_FullMethodName          = "/StaffRecords.StaffRecords/Create"
	StaffRecords_Get_FullMethodName             = "/StaffRecords.StaffRecords/Get"
	StaffRecords_GetAvailability_FullMethodName = "/StaffRecords.StaffRecords/GetAvailability"
	StaffRecords_Update_FullMethodName          = "/StaffRecords.StaffRecords/Update"
	StaffRecords_Delete_FullMethodName          = "/StaffRecords.StaffRecords/Delete"
	StaffRecords_Check_FullMethodName           = "/StaffRecords.StaffRecords/Check"
)

// StaffRecordsClient is the client API for StaffRecords service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StaffRecordsClient interface {
	// Service Methods
	Create(ctx context.Context, in *CreateStaff, opts ...grpc.CallOption) (*CreateResponse, error)
	Get(ctx context.Context, in *GetStaffRecords, opts ...grpc.CallOption) (*GetResponse, error)
	GetAvailability(ctx context.Context, in *GetStaffAvailability, opts ...grpc.CallOption) (*GetAvailabilityResponse, error)
	Update(ctx context.Context, in *UpdateStaff, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteStaff, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Health Check Methods
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type staffRecordsClient struct {
	cc grpc.ClientConnInterface
}

func NewStaffRecordsClient(cc grpc.ClientConnInterface) StaffRecordsClient {
	return &staffRecordsClient{cc}
}

func (c *staffRecordsClient) Create(ctx context.Context, in *CreateStaff, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, StaffRecords_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffRecordsClient) Get(ctx context.Context, in *GetStaffRecords, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, StaffRecords_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffRecordsClient) GetAvailability(ctx context.Context, in *GetStaffAvailability, opts ...grpc.CallOption) (*GetAvailabilityResponse, error) {
	out := new(GetAvailabilityResponse)
	err := c.cc.Invoke(ctx, StaffRecords_GetAvailability_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffRecordsClient) Update(ctx context.Context, in *UpdateStaff, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, StaffRecords_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffRecordsClient) Delete(ctx context.Context, in *DeleteStaff, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, StaffRecords_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffRecordsClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, StaffRecords_Check_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaffRecordsServer is the server API for StaffRecords service.
// All implementations must embed UnimplementedStaffRecordsServer
// for forward compatibility
type StaffRecordsServer interface {
	// Service Methods
	Create(context.Context, *CreateStaff) (*CreateResponse, error)
	Get(context.Context, *GetStaffRecords) (*GetResponse, error)
	GetAvailability(context.Context, *GetStaffAvailability) (*GetAvailabilityResponse, error)
	Update(context.Context, *UpdateStaff) (*UpdateResponse, error)
	Delete(context.Context, *DeleteStaff) (*DeleteResponse, error)
	// Health Check Methods
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	mustEmbedUnimplementedStaffRecordsServer()
}

// UnimplementedStaffRecordsServer must be embedded to have forward compatible implementations.
type UnimplementedStaffRecordsServer struct {
}

func (UnimplementedStaffRecordsServer) Create(context.Context, *CreateStaff) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedStaffRecordsServer) Get(context.Context, *GetStaffRecords) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStaffRecordsServer) GetAvailability(context.Context, *GetStaffAvailability) (*GetAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailability not implemented")
}
func (UnimplementedStaffRecordsServer) Update(context.Context, *UpdateStaff) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStaffRecordsServer) Delete(context.Context, *DeleteStaff) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStaffRecordsServer) Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedStaffRecordsServer) mustEmbedUnimplementedStaffRecordsServer() {}

// UnsafeStaffRecordsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StaffRecordsServer will
// result in compilation errors.
type UnsafeStaffRecordsServer interface {
	mustEmbedUnimplementedStaffRecordsServer()
}

func RegisterStaffRecordsServer(s grpc.ServiceRegistrar, srv StaffRecordsServer) {
	s.RegisterService(&StaffRecords_ServiceDesc, srv)
}

func _StaffRecords_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStaff)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffRecordsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaffRecords_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffRecordsServer).Create(ctx, req.(*CreateStaff))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffRecords_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStaffRecords)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffRecordsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaffRecords_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffRecordsServer).Get(ctx, req.(*GetStaffRecords))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffRecords_GetAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStaffAvailability)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffRecordsServer).GetAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaffRecords_GetAvailability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffRecordsServer).GetAvailability(ctx, req.(*GetStaffAvailability))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffRecords_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStaff)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffRecordsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaffRecords_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffRecordsServer).Update(ctx, req.(*UpdateStaff))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffRecords_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStaff)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffRecordsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaffRecords_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffRecordsServer).Delete(ctx, req.(*DeleteStaff))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffRecords_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffRecordsServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StaffRecords_Check_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffRecordsServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StaffRecords_ServiceDesc is the grpc.ServiceDesc for StaffRecords service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StaffRecords_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StaffRecords.StaffRecords",
	HandlerType: (*StaffRecordsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _StaffRecords_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _StaffRecords_Get_Handler,
		},
		{
			MethodName: "GetAvailability",
			Handler:    _StaffRecords_GetAvailability_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StaffRecords_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _StaffRecords_Delete_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _StaffRecords_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/staff-records/staff_records.proto",
}
