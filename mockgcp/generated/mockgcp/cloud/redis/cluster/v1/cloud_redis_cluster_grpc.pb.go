// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/cloud/redis/cluster/v1/cloud_redis_cluster.proto

package clusterpb

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

// CloudRedisClusterClient is the client API for CloudRedisCluster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudRedisClusterClient interface {
	// Lists all Redis clusters owned by a project in either the specified
	// location (region) or all locations.
	//
	// The location should have the following format:
	//
	// * `projects/{project_id}/locations/{location_id}`
	//
	// If `location_id` is specified as `-` (wildcard), then all regions
	// available to the project are queried, and the results are aggregated.
	ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error)
	// Gets the details of a specific Redis cluster.
	GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*Cluster, error)
	// Updates the metadata and configuration of a specific Redis cluster.
	//
	// Completed longrunning.Operation will contain the new cluster object
	// in the response field. The returned operation is automatically deleted
	// after a few hours, so there is no need to call DeleteOperation.
	UpdateCluster(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Deletes a specific Redis cluster. Cluster stops serving and data is
	// deleted.
	DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Creates a Redis cluster based on the specified properties.
	// The creation is executed asynchronously and callers may check the returned
	// operation to track its progress. Once the operation is completed the Redis
	// cluster will be fully functional. The completed longrunning.Operation will
	// contain the new cluster object in the response field.
	//
	// The returned operation is automatically deleted after a few hours, so there
	// is no need to call DeleteOperation.
	CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error)
	// Gets the details of certificate authority information for Redis cluster.
	GetClusterCertificateAuthority(ctx context.Context, in *GetClusterCertificateAuthorityRequest, opts ...grpc.CallOption) (*CertificateAuthority, error)
}

type cloudRedisClusterClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudRedisClusterClient(cc grpc.ClientConnInterface) CloudRedisClusterClient {
	return &cloudRedisClusterClient{cc}
}

func (c *cloudRedisClusterClient) ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error) {
	out := new(ListClustersResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.redis.cluster.v1.CloudRedisCluster/ListClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudRedisClusterClient) GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.redis.cluster.v1.CloudRedisCluster/GetCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudRedisClusterClient) UpdateCluster(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.redis.cluster.v1.CloudRedisCluster/UpdateCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudRedisClusterClient) DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.redis.cluster.v1.CloudRedisCluster/DeleteCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudRedisClusterClient) CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*longrunningpb.Operation, error) {
	out := new(longrunningpb.Operation)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.redis.cluster.v1.CloudRedisCluster/CreateCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudRedisClusterClient) GetClusterCertificateAuthority(ctx context.Context, in *GetClusterCertificateAuthorityRequest, opts ...grpc.CallOption) (*CertificateAuthority, error) {
	out := new(CertificateAuthority)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.redis.cluster.v1.CloudRedisCluster/GetClusterCertificateAuthority", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudRedisClusterServer is the server API for CloudRedisCluster service.
// All implementations must embed UnimplementedCloudRedisClusterServer
// for forward compatibility
type CloudRedisClusterServer interface {
	// Lists all Redis clusters owned by a project in either the specified
	// location (region) or all locations.
	//
	// The location should have the following format:
	//
	// * `projects/{project_id}/locations/{location_id}`
	//
	// If `location_id` is specified as `-` (wildcard), then all regions
	// available to the project are queried, and the results are aggregated.
	ListClusters(context.Context, *ListClustersRequest) (*ListClustersResponse, error)
	// Gets the details of a specific Redis cluster.
	GetCluster(context.Context, *GetClusterRequest) (*Cluster, error)
	// Updates the metadata and configuration of a specific Redis cluster.
	//
	// Completed longrunning.Operation will contain the new cluster object
	// in the response field. The returned operation is automatically deleted
	// after a few hours, so there is no need to call DeleteOperation.
	UpdateCluster(context.Context, *UpdateClusterRequest) (*longrunningpb.Operation, error)
	// Deletes a specific Redis cluster. Cluster stops serving and data is
	// deleted.
	DeleteCluster(context.Context, *DeleteClusterRequest) (*longrunningpb.Operation, error)
	// Creates a Redis cluster based on the specified properties.
	// The creation is executed asynchronously and callers may check the returned
	// operation to track its progress. Once the operation is completed the Redis
	// cluster will be fully functional. The completed longrunning.Operation will
	// contain the new cluster object in the response field.
	//
	// The returned operation is automatically deleted after a few hours, so there
	// is no need to call DeleteOperation.
	CreateCluster(context.Context, *CreateClusterRequest) (*longrunningpb.Operation, error)
	// Gets the details of certificate authority information for Redis cluster.
	GetClusterCertificateAuthority(context.Context, *GetClusterCertificateAuthorityRequest) (*CertificateAuthority, error)
	mustEmbedUnimplementedCloudRedisClusterServer()
}

// UnimplementedCloudRedisClusterServer must be embedded to have forward compatible implementations.
type UnimplementedCloudRedisClusterServer struct {
}

func (UnimplementedCloudRedisClusterServer) ListClusters(context.Context, *ListClustersRequest) (*ListClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClusters not implemented")
}
func (UnimplementedCloudRedisClusterServer) GetCluster(context.Context, *GetClusterRequest) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCluster not implemented")
}
func (UnimplementedCloudRedisClusterServer) UpdateCluster(context.Context, *UpdateClusterRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCluster not implemented")
}
func (UnimplementedCloudRedisClusterServer) DeleteCluster(context.Context, *DeleteClusterRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCluster not implemented")
}
func (UnimplementedCloudRedisClusterServer) CreateCluster(context.Context, *CreateClusterRequest) (*longrunningpb.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCluster not implemented")
}
func (UnimplementedCloudRedisClusterServer) GetClusterCertificateAuthority(context.Context, *GetClusterCertificateAuthorityRequest) (*CertificateAuthority, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterCertificateAuthority not implemented")
}
func (UnimplementedCloudRedisClusterServer) mustEmbedUnimplementedCloudRedisClusterServer() {}

// UnsafeCloudRedisClusterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudRedisClusterServer will
// result in compilation errors.
type UnsafeCloudRedisClusterServer interface {
	mustEmbedUnimplementedCloudRedisClusterServer()
}

func RegisterCloudRedisClusterServer(s grpc.ServiceRegistrar, srv CloudRedisClusterServer) {
	s.RegisterService(&CloudRedisCluster_ServiceDesc, srv)
}

func _CloudRedisCluster_ListClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudRedisClusterServer).ListClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.redis.cluster.v1.CloudRedisCluster/ListClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudRedisClusterServer).ListClusters(ctx, req.(*ListClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudRedisCluster_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudRedisClusterServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.redis.cluster.v1.CloudRedisCluster/GetCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudRedisClusterServer).GetCluster(ctx, req.(*GetClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudRedisCluster_UpdateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudRedisClusterServer).UpdateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.redis.cluster.v1.CloudRedisCluster/UpdateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudRedisClusterServer).UpdateCluster(ctx, req.(*UpdateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudRedisCluster_DeleteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudRedisClusterServer).DeleteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.redis.cluster.v1.CloudRedisCluster/DeleteCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudRedisClusterServer).DeleteCluster(ctx, req.(*DeleteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudRedisCluster_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudRedisClusterServer).CreateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.redis.cluster.v1.CloudRedisCluster/CreateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudRedisClusterServer).CreateCluster(ctx, req.(*CreateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudRedisCluster_GetClusterCertificateAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterCertificateAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudRedisClusterServer).GetClusterCertificateAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.redis.cluster.v1.CloudRedisCluster/GetClusterCertificateAuthority",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudRedisClusterServer).GetClusterCertificateAuthority(ctx, req.(*GetClusterCertificateAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudRedisCluster_ServiceDesc is the grpc.ServiceDesc for CloudRedisCluster service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudRedisCluster_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.cloud.redis.cluster.v1.CloudRedisCluster",
	HandlerType: (*CloudRedisClusterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListClusters",
			Handler:    _CloudRedisCluster_ListClusters_Handler,
		},
		{
			MethodName: "GetCluster",
			Handler:    _CloudRedisCluster_GetCluster_Handler,
		},
		{
			MethodName: "UpdateCluster",
			Handler:    _CloudRedisCluster_UpdateCluster_Handler,
		},
		{
			MethodName: "DeleteCluster",
			Handler:    _CloudRedisCluster_DeleteCluster_Handler,
		},
		{
			MethodName: "CreateCluster",
			Handler:    _CloudRedisCluster_CreateCluster_Handler,
		},
		{
			MethodName: "GetClusterCertificateAuthority",
			Handler:    _CloudRedisCluster_GetClusterCertificateAuthority_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/cloud/redis/cluster/v1/cloud_redis_cluster.proto",
}
