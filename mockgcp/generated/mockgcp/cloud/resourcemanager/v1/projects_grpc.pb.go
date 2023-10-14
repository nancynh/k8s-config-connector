// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/cloud/resourcemanager/v1/projects.proto

package resourcemanagerpb

import (
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProjectsClient is the client API for Projects service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectsClient interface {
	// Retrieves the project identified by the specified `name` (for example,
	// `projects/415104041262`).
	//
	// The caller must have `resourcemanager.projects.get` permission
	// for this project.
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error)
	// Request that a new project be created. The result is an `Operation` which
	// can be used to track the creation process. This process usually takes a few
	// seconds, but can sometimes take much longer. The tracking `Operation` is
	// automatically deleted after a few hours, so there is no need to call
	// `DeleteOperation`.
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Marks the project identified by the specified
	// `name` (for example, `projects/415104041262`) for deletion.
	//
	// This method will only affect the project if it has a lifecycle state of
	// [ACTIVE][mockgcp.cloud.resourcemanager.v3.Project.State.ACTIVE].
	//
	// This method changes the Project's lifecycle state from
	// [ACTIVE][mockgcp.cloud.resourcemanager.v3.Project.State.ACTIVE]
	// to
	// [DELETE_REQUESTED][mockgcp.cloud.resourcemanager.v3.Project.State.DELETE_REQUESTED].
	// The deletion starts at an unspecified time,
	// at which point the Project is no longer accessible.
	//
	// Until the deletion completes, you can check the lifecycle state
	// checked by retrieving the project with [GetProject]
	// [mockgcp.cloud.resourcemanager.v3.Projects.GetProject],
	// and the project remains visible to [ListProjects]
	// [mockgcp.cloud.resourcemanager.v3.Projects.ListProjects].
	// However, you cannot update the project.
	//
	// After the deletion completes, the project is not retrievable by
	// the  [GetProject]
	// [mockgcp.cloud.resourcemanager.v3.Projects.GetProject],
	// [ListProjects]
	// [mockgcp.cloud.resourcemanager.v3.Projects.ListProjects], and
	// [SearchProjects][mockgcp.cloud.resourcemanager.v3.Projects.SearchProjects]
	// methods.
	//
	// This method behaves idempotently, such that deleting a `DELETE_REQUESTED`
	// project will not cause an error, but also won't do anything.
	//
	// The caller must have `resourcemanager.projects.delete` permissions for this
	// project.
	DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
}

type projectsClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectsClient(cc grpc.ClientConnInterface) ProjectsClient {
	return &projectsClient{cc}
}

func (c *projectsClient) GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.resourcemanager.v1.Projects/GetProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.resourcemanager.v1.Projects/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.resourcemanager.v1.Projects/DeleteProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectsServer is the server API for Projects service.
// All implementations must embed UnimplementedProjectsServer
// for forward compatibility
type ProjectsServer interface {
	// Retrieves the project identified by the specified `name` (for example,
	// `projects/415104041262`).
	//
	// The caller must have `resourcemanager.projects.get` permission
	// for this project.
	GetProject(context.Context, *GetProjectRequest) (*Project, error)
	// Request that a new project be created. The result is an `Operation` which
	// can be used to track the creation process. This process usually takes a few
	// seconds, but can sometimes take much longer. The tracking `Operation` is
	// automatically deleted after a few hours, so there is no need to call
	// `DeleteOperation`.
	CreateProject(context.Context, *CreateProjectRequest) (*longrunningpb.Operation, error)
	// Marks the project identified by the specified
	// `name` (for example, `projects/415104041262`) for deletion.
	//
	// This method will only affect the project if it has a lifecycle state of
	// [ACTIVE][mockgcp.cloud.resourcemanager.v3.Project.State.ACTIVE].
	//
	// This method changes the Project's lifecycle state from
	// [ACTIVE][mockgcp.cloud.resourcemanager.v3.Project.State.ACTIVE]
	// to
	// [DELETE_REQUESTED][mockgcp.cloud.resourcemanager.v3.Project.State.DELETE_REQUESTED].
	// The deletion starts at an unspecified time,
	// at which point the Project is no longer accessible.
	//
	// Until the deletion completes, you can check the lifecycle state
	// checked by retrieving the project with [GetProject]
	// [mockgcp.cloud.resourcemanager.v3.Projects.GetProject],
	// and the project remains visible to [ListProjects]
	// [mockgcp.cloud.resourcemanager.v3.Projects.ListProjects].
	// However, you cannot update the project.
	//
	// After the deletion completes, the project is not retrievable by
	// the  [GetProject]
	// [mockgcp.cloud.resourcemanager.v3.Projects.GetProject],
	// [ListProjects]
	// [mockgcp.cloud.resourcemanager.v3.Projects.ListProjects], and
	// [SearchProjects][mockgcp.cloud.resourcemanager.v3.Projects.SearchProjects]
	// methods.
	//
	// This method behaves idempotently, such that deleting a `DELETE_REQUESTED`
	// project will not cause an error, but also won't do anything.
	//
	// The caller must have `resourcemanager.projects.delete` permissions for this
	// project.
	DeleteProject(context.Context, *DeleteProjectRequest) (*longrunningpb.Operation, error)
	mustEmbedUnimplementedProjectsServer()
}

// UnimplementedProjectsServer must be embedded to have forward compatible implementations.
type UnimplementedProjectsServer struct {
}

func (UnimplementedProjectsServer) GetProject(context.Context, *GetProjectRequest) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedProjectsServer) CreateProject(context.Context, *CreateProjectRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedProjectsServer) DeleteProject(context.Context, *DeleteProjectRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}
func (UnimplementedProjectsServer) mustEmbedUnimplementedProjectsServer() {}

// UnsafeProjectsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectsServer will
// result in compilation errors.
type UnsafeProjectsServer interface {
	mustEmbedUnimplementedProjectsServer()
}

func RegisterProjectsServer(s grpc.ServiceRegistrar, srv ProjectsServer) {
	s.RegisterService(&Projects_ServiceDesc, srv)
}

func _Projects_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.resourcemanager.v1.Projects/GetProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).GetProject(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.resourcemanager.v1.Projects/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.resourcemanager.v1.Projects/DeleteProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).DeleteProject(ctx, req.(*DeleteProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Projects_ServiceDesc is the grpc.ServiceDesc for Projects service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Projects_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.cloud.resourcemanager.v1.Projects",
	HandlerType: (*ProjectsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProject",
			Handler:    _Projects_GetProject_Handler,
		},
		{
			MethodName: "CreateProject",
			Handler:    _Projects_CreateProject_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _Projects_DeleteProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/cloud/resourcemanager/v1/projects.proto",
}
