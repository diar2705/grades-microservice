// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: grades-microservice.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GradesService_GetCourseGrades_FullMethodName        = "/grades.GradesService/GetCourseGrades"
	GradesService_GetStudentCourseGrades_FullMethodName = "/grades.GradesService/GetStudentCourseGrades"
	GradesService_GetStudentGrades_FullMethodName       = "/grades.GradesService/GetStudentGrades"
	GradesService_AddHomeworkGrade_FullMethodName       = "/grades.GradesService/AddHomeworkGrade"
	GradesService_AddExamGrade_FullMethodName           = "/grades.GradesService/AddExamGrade"
	GradesService_UpdateHomeworkGrade_FullMethodName    = "/grades.GradesService/UpdateHomeworkGrade"
	GradesService_UpdateExamGrade_FullMethodName        = "/grades.GradesService/UpdateExamGrade"
	GradesService_DeleteHomeworkGrade_FullMethodName    = "/grades.GradesService/DeleteHomeworkGrade"
	GradesService_DeleteExamGrade_FullMethodName        = "/grades.GradesService/DeleteExamGrade"
)

// GradesServiceClient is the client API for GradesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// GradesService is a microservice responsible for managing grades.
type GradesServiceClient interface {
	// GetCourseGrades returns all students grades enrolled in a specific course.
	GetCourseGrades(ctx context.Context, in *GetCourseGradesRequest, opts ...grpc.CallOption) (*GetCourseGradesResponse, error)
	// GetStudentCourseGrades returns a specific student grades in specific course.
	GetStudentCourseGrades(ctx context.Context, in *GetStudentCourseGradesRequest, opts ...grpc.CallOption) (*GetStudentCourseGradesResponse, error)
	// GetStudentGrades returns all grades for a student.
	GetStudentGrades(ctx context.Context, in *StudentId, opts ...grpc.CallOption) (*StudentGrades, error)
	// AddHomeworkGrade adds a homework grade for a student in specific course.
	AddHomeworkGrade(ctx context.Context, in *AddHomeworkGradeRequest, opts ...grpc.CallOption) (*Empty, error)
	// AddExamGrade adds an exam grade for a student in specific course.
	AddExamGrade(ctx context.Context, in *AddExamGradeRequest, opts ...grpc.CallOption) (*Empty, error)
	// UpdateHomeworkGrade updates a homework grade for a student in specific course.
	UpdateHomeworkGrade(ctx context.Context, in *UpdateHomeworkGradeRequest, opts ...grpc.CallOption) (*Empty, error)
	// UpdateExamGrade updates an exam grade for a student in specific course.
	UpdateExamGrade(ctx context.Context, in *UpdateExamGradeRequest, opts ...grpc.CallOption) (*Empty, error)
	// DeleteHomeworkGrade deletes a homework grade for a student in specific course.
	DeleteHomeworkGrade(ctx context.Context, in *DeleteHomeworkGradeRequest, opts ...grpc.CallOption) (*Empty, error)
	// DeleteExamGrade deletes an exam grade for a student in specific course.
	DeleteExamGrade(ctx context.Context, in *DeleteExamGradeRequest, opts ...grpc.CallOption) (*Empty, error)
}

type gradesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGradesServiceClient(cc grpc.ClientConnInterface) GradesServiceClient {
	return &gradesServiceClient{cc}
}

func (c *gradesServiceClient) GetCourseGrades(ctx context.Context, in *GetCourseGradesRequest, opts ...grpc.CallOption) (*GetCourseGradesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCourseGradesResponse)
	err := c.cc.Invoke(ctx, GradesService_GetCourseGrades_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gradesServiceClient) GetStudentCourseGrades(ctx context.Context, in *GetStudentCourseGradesRequest, opts ...grpc.CallOption) (*GetStudentCourseGradesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStudentCourseGradesResponse)
	err := c.cc.Invoke(ctx, GradesService_GetStudentCourseGrades_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gradesServiceClient) GetStudentGrades(ctx context.Context, in *StudentId, opts ...grpc.CallOption) (*StudentGrades, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StudentGrades)
	err := c.cc.Invoke(ctx, GradesService_GetStudentGrades_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gradesServiceClient) AddHomeworkGrade(ctx context.Context, in *AddHomeworkGradeRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, GradesService_AddHomeworkGrade_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gradesServiceClient) AddExamGrade(ctx context.Context, in *AddExamGradeRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, GradesService_AddExamGrade_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gradesServiceClient) UpdateHomeworkGrade(ctx context.Context, in *UpdateHomeworkGradeRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, GradesService_UpdateHomeworkGrade_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gradesServiceClient) UpdateExamGrade(ctx context.Context, in *UpdateExamGradeRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, GradesService_UpdateExamGrade_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gradesServiceClient) DeleteHomeworkGrade(ctx context.Context, in *DeleteHomeworkGradeRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, GradesService_DeleteHomeworkGrade_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gradesServiceClient) DeleteExamGrade(ctx context.Context, in *DeleteExamGradeRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, GradesService_DeleteExamGrade_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GradesServiceServer is the server API for GradesService service.
// All implementations must embed UnimplementedGradesServiceServer
// for forward compatibility.
//
// GradesService is a microservice responsible for managing grades.
type GradesServiceServer interface {
	// GetCourseGrades returns all students grades enrolled in a specific course.
	GetCourseGrades(context.Context, *GetCourseGradesRequest) (*GetCourseGradesResponse, error)
	// GetStudentCourseGrades returns a specific student grades in specific course.
	GetStudentCourseGrades(context.Context, *GetStudentCourseGradesRequest) (*GetStudentCourseGradesResponse, error)
	// GetStudentGrades returns all grades for a student.
	GetStudentGrades(context.Context, *StudentId) (*StudentGrades, error)
	// AddHomeworkGrade adds a homework grade for a student in specific course.
	AddHomeworkGrade(context.Context, *AddHomeworkGradeRequest) (*Empty, error)
	// AddExamGrade adds an exam grade for a student in specific course.
	AddExamGrade(context.Context, *AddExamGradeRequest) (*Empty, error)
	// UpdateHomeworkGrade updates a homework grade for a student in specific course.
	UpdateHomeworkGrade(context.Context, *UpdateHomeworkGradeRequest) (*Empty, error)
	// UpdateExamGrade updates an exam grade for a student in specific course.
	UpdateExamGrade(context.Context, *UpdateExamGradeRequest) (*Empty, error)
	// DeleteHomeworkGrade deletes a homework grade for a student in specific course.
	DeleteHomeworkGrade(context.Context, *DeleteHomeworkGradeRequest) (*Empty, error)
	// DeleteExamGrade deletes an exam grade for a student in specific course.
	DeleteExamGrade(context.Context, *DeleteExamGradeRequest) (*Empty, error)
	mustEmbedUnimplementedGradesServiceServer()
}

// UnimplementedGradesServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGradesServiceServer struct{}

func (UnimplementedGradesServiceServer) GetCourseGrades(context.Context, *GetCourseGradesRequest) (*GetCourseGradesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourseGrades not implemented")
}
func (UnimplementedGradesServiceServer) GetStudentCourseGrades(context.Context, *GetStudentCourseGradesRequest) (*GetStudentCourseGradesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentCourseGrades not implemented")
}
func (UnimplementedGradesServiceServer) GetStudentGrades(context.Context, *StudentId) (*StudentGrades, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentGrades not implemented")
}
func (UnimplementedGradesServiceServer) AddHomeworkGrade(context.Context, *AddHomeworkGradeRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddHomeworkGrade not implemented")
}
func (UnimplementedGradesServiceServer) AddExamGrade(context.Context, *AddExamGradeRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddExamGrade not implemented")
}
func (UnimplementedGradesServiceServer) UpdateHomeworkGrade(context.Context, *UpdateHomeworkGradeRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHomeworkGrade not implemented")
}
func (UnimplementedGradesServiceServer) UpdateExamGrade(context.Context, *UpdateExamGradeRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExamGrade not implemented")
}
func (UnimplementedGradesServiceServer) DeleteHomeworkGrade(context.Context, *DeleteHomeworkGradeRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHomeworkGrade not implemented")
}
func (UnimplementedGradesServiceServer) DeleteExamGrade(context.Context, *DeleteExamGradeRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExamGrade not implemented")
}
func (UnimplementedGradesServiceServer) mustEmbedUnimplementedGradesServiceServer() {}
func (UnimplementedGradesServiceServer) testEmbeddedByValue()                       {}

// UnsafeGradesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GradesServiceServer will
// result in compilation errors.
type UnsafeGradesServiceServer interface {
	mustEmbedUnimplementedGradesServiceServer()
}

func RegisterGradesServiceServer(s grpc.ServiceRegistrar, srv GradesServiceServer) {
	// If the following call pancis, it indicates UnimplementedGradesServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GradesService_ServiceDesc, srv)
}

func _GradesService_GetCourseGrades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCourseGradesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GradesServiceServer).GetCourseGrades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GradesService_GetCourseGrades_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GradesServiceServer).GetCourseGrades(ctx, req.(*GetCourseGradesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GradesService_GetStudentCourseGrades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentCourseGradesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GradesServiceServer).GetStudentCourseGrades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GradesService_GetStudentCourseGrades_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GradesServiceServer).GetStudentCourseGrades(ctx, req.(*GetStudentCourseGradesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GradesService_GetStudentGrades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GradesServiceServer).GetStudentGrades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GradesService_GetStudentGrades_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GradesServiceServer).GetStudentGrades(ctx, req.(*StudentId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GradesService_AddHomeworkGrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddHomeworkGradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GradesServiceServer).AddHomeworkGrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GradesService_AddHomeworkGrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GradesServiceServer).AddHomeworkGrade(ctx, req.(*AddHomeworkGradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GradesService_AddExamGrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddExamGradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GradesServiceServer).AddExamGrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GradesService_AddExamGrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GradesServiceServer).AddExamGrade(ctx, req.(*AddExamGradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GradesService_UpdateHomeworkGrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHomeworkGradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GradesServiceServer).UpdateHomeworkGrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GradesService_UpdateHomeworkGrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GradesServiceServer).UpdateHomeworkGrade(ctx, req.(*UpdateHomeworkGradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GradesService_UpdateExamGrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateExamGradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GradesServiceServer).UpdateExamGrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GradesService_UpdateExamGrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GradesServiceServer).UpdateExamGrade(ctx, req.(*UpdateExamGradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GradesService_DeleteHomeworkGrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHomeworkGradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GradesServiceServer).DeleteHomeworkGrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GradesService_DeleteHomeworkGrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GradesServiceServer).DeleteHomeworkGrade(ctx, req.(*DeleteHomeworkGradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GradesService_DeleteExamGrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteExamGradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GradesServiceServer).DeleteExamGrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GradesService_DeleteExamGrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GradesServiceServer).DeleteExamGrade(ctx, req.(*DeleteExamGradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GradesService_ServiceDesc is the grpc.ServiceDesc for GradesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GradesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grades.GradesService",
	HandlerType: (*GradesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCourseGrades",
			Handler:    _GradesService_GetCourseGrades_Handler,
		},
		{
			MethodName: "GetStudentCourseGrades",
			Handler:    _GradesService_GetStudentCourseGrades_Handler,
		},
		{
			MethodName: "GetStudentGrades",
			Handler:    _GradesService_GetStudentGrades_Handler,
		},
		{
			MethodName: "AddHomeworkGrade",
			Handler:    _GradesService_AddHomeworkGrade_Handler,
		},
		{
			MethodName: "AddExamGrade",
			Handler:    _GradesService_AddExamGrade_Handler,
		},
		{
			MethodName: "UpdateHomeworkGrade",
			Handler:    _GradesService_UpdateHomeworkGrade_Handler,
		},
		{
			MethodName: "UpdateExamGrade",
			Handler:    _GradesService_UpdateExamGrade_Handler,
		},
		{
			MethodName: "DeleteHomeworkGrade",
			Handler:    _GradesService_DeleteHomeworkGrade_Handler,
		},
		{
			MethodName: "DeleteExamGrade",
			Handler:    _GradesService_DeleteExamGrade_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grades-microservice.proto",
}
