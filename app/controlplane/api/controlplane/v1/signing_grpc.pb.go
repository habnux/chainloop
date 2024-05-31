//
// Copyright 2024 The Chainloop Authors.
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
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: controlplane/v1/signing.proto

package v1

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
	SigningService_GenerateSigningCert_FullMethodName = "/controlplane.v1.SigningService/GenerateSigningCert"
)

// SigningServiceClient is the client API for SigningService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SigningServiceClient interface {
	// GenerateSigningCert takes a certificate request and generates a new certificate for attestation signing
	GenerateSigningCert(ctx context.Context, in *GenerateSigningCertRequest, opts ...grpc.CallOption) (*GenerateSigningCertResponse, error)
}

type signingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSigningServiceClient(cc grpc.ClientConnInterface) SigningServiceClient {
	return &signingServiceClient{cc}
}

func (c *signingServiceClient) GenerateSigningCert(ctx context.Context, in *GenerateSigningCertRequest, opts ...grpc.CallOption) (*GenerateSigningCertResponse, error) {
	out := new(GenerateSigningCertResponse)
	err := c.cc.Invoke(ctx, SigningService_GenerateSigningCert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SigningServiceServer is the server API for SigningService service.
// All implementations must embed UnimplementedSigningServiceServer
// for forward compatibility
type SigningServiceServer interface {
	// GenerateSigningCert takes a certificate request and generates a new certificate for attestation signing
	GenerateSigningCert(context.Context, *GenerateSigningCertRequest) (*GenerateSigningCertResponse, error)
	mustEmbedUnimplementedSigningServiceServer()
}

// UnimplementedSigningServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSigningServiceServer struct {
}

func (UnimplementedSigningServiceServer) GenerateSigningCert(context.Context, *GenerateSigningCertRequest) (*GenerateSigningCertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateSigningCert not implemented")
}
func (UnimplementedSigningServiceServer) mustEmbedUnimplementedSigningServiceServer() {}

// UnsafeSigningServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SigningServiceServer will
// result in compilation errors.
type UnsafeSigningServiceServer interface {
	mustEmbedUnimplementedSigningServiceServer()
}

func RegisterSigningServiceServer(s grpc.ServiceRegistrar, srv SigningServiceServer) {
	s.RegisterService(&SigningService_ServiceDesc, srv)
}

func _SigningService_GenerateSigningCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateSigningCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigningServiceServer).GenerateSigningCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SigningService_GenerateSigningCert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigningServiceServer).GenerateSigningCert(ctx, req.(*GenerateSigningCertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SigningService_ServiceDesc is the grpc.ServiceDesc for SigningService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SigningService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "controlplane.v1.SigningService",
	HandlerType: (*SigningServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateSigningCert",
			Handler:    _SigningService_GenerateSigningCert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controlplane/v1/signing.proto",
}
