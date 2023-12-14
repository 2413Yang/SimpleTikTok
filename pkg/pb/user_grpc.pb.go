// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.9.0
// source: user.proto

package pb

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
	UserService_GetUserInfo_FullMethodName              = "/UserService/GetUserInfo"
	UserService_GetUserInfoDict_FullMethodName          = "/UserService/GetUserInfoDict"
	UserService_CheckPassWord_FullMethodName            = "/UserService/CheckPassWord"
	UserService_Register_FullMethodName                 = "/UserService/Register"
	UserService_GetUserInfoList_FullMethodName          = "/UserService/GetUserInfoList"
	UserService_CacheChangeUserCount_FullMethodName     = "/UserService/CacheChangeUserCount"
	UserService_CacheGetAuthor_FullMethodName           = "/UserService/CacheGetAuthor"
	UserService_UpdateUserFavoritedCount_FullMethodName = "/UserService/UpdateUserFavoritedCount"
	UserService_UpdateUserFavoriteCount_FullMethodName  = "/UserService/UpdateUserFavoriteCount"
	UserService_UpdateUserFollowCount_FullMethodName    = "/UserService/UpdateUserFollowCount"
	UserService_UpdateUserFollowerCount_FullMethodName  = "/UserService/UpdateUserFollowerCount"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
	GetUserInfoDict(ctx context.Context, in *GetUserInfoDictRequest, opts ...grpc.CallOption) (*GetUserInfoDictResopnse, error)
	CheckPassWord(ctx context.Context, in *CheckPassWordRequest, opts ...grpc.CallOption) (*CheckPassWordResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	GetUserInfoList(ctx context.Context, in *GetUserInfoListRequest, opts ...grpc.CallOption) (*GetUserInfoListResponse, error)
	CacheChangeUserCount(ctx context.Context, in *CacheChangeUserCountReq, opts ...grpc.CallOption) (*CacheChangeUserCountRsp, error)
	CacheGetAuthor(ctx context.Context, in *CacheGetAuthorReq, opts ...grpc.CallOption) (*CacheGetAuthorRsp, error)
	UpdateUserFavoritedCount(ctx context.Context, in *UpdateUserFavoritedCountReq, opts ...grpc.CallOption) (*UpdateUserFavoritedCountRsp, error)
	UpdateUserFavoriteCount(ctx context.Context, in *UpdateUserFavoriteCountReq, opts ...grpc.CallOption) (*UpdateUserFavoriteCountRsp, error)
	UpdateUserFollowCount(ctx context.Context, in *UpdateUserFollowCountReq, opts ...grpc.CallOption) (*UpdateUserFollowCountRsp, error)
	UpdateUserFollowerCount(ctx context.Context, in *UpdateUserFollowerCountReq, opts ...grpc.CallOption) (*UpdateUserFollowerCountRsp, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	out := new(GetUserInfoResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInfoDict(ctx context.Context, in *GetUserInfoDictRequest, opts ...grpc.CallOption) (*GetUserInfoDictResopnse, error) {
	out := new(GetUserInfoDictResopnse)
	err := c.cc.Invoke(ctx, UserService_GetUserInfoDict_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CheckPassWord(ctx context.Context, in *CheckPassWordRequest, opts ...grpc.CallOption) (*CheckPassWordResponse, error) {
	out := new(CheckPassWordResponse)
	err := c.cc.Invoke(ctx, UserService_CheckPassWord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, UserService_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInfoList(ctx context.Context, in *GetUserInfoListRequest, opts ...grpc.CallOption) (*GetUserInfoListResponse, error) {
	out := new(GetUserInfoListResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserInfoList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CacheChangeUserCount(ctx context.Context, in *CacheChangeUserCountReq, opts ...grpc.CallOption) (*CacheChangeUserCountRsp, error) {
	out := new(CacheChangeUserCountRsp)
	err := c.cc.Invoke(ctx, UserService_CacheChangeUserCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CacheGetAuthor(ctx context.Context, in *CacheGetAuthorReq, opts ...grpc.CallOption) (*CacheGetAuthorRsp, error) {
	out := new(CacheGetAuthorRsp)
	err := c.cc.Invoke(ctx, UserService_CacheGetAuthor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserFavoritedCount(ctx context.Context, in *UpdateUserFavoritedCountReq, opts ...grpc.CallOption) (*UpdateUserFavoritedCountRsp, error) {
	out := new(UpdateUserFavoritedCountRsp)
	err := c.cc.Invoke(ctx, UserService_UpdateUserFavoritedCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserFavoriteCount(ctx context.Context, in *UpdateUserFavoriteCountReq, opts ...grpc.CallOption) (*UpdateUserFavoriteCountRsp, error) {
	out := new(UpdateUserFavoriteCountRsp)
	err := c.cc.Invoke(ctx, UserService_UpdateUserFavoriteCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserFollowCount(ctx context.Context, in *UpdateUserFollowCountReq, opts ...grpc.CallOption) (*UpdateUserFollowCountRsp, error) {
	out := new(UpdateUserFollowCountRsp)
	err := c.cc.Invoke(ctx, UserService_UpdateUserFollowCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserFollowerCount(ctx context.Context, in *UpdateUserFollowerCountReq, opts ...grpc.CallOption) (*UpdateUserFollowerCountRsp, error) {
	out := new(UpdateUserFollowerCountRsp)
	err := c.cc.Invoke(ctx, UserService_UpdateUserFollowerCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error)
	GetUserInfoDict(context.Context, *GetUserInfoDictRequest) (*GetUserInfoDictResopnse, error)
	CheckPassWord(context.Context, *CheckPassWordRequest) (*CheckPassWordResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	GetUserInfoList(context.Context, *GetUserInfoListRequest) (*GetUserInfoListResponse, error)
	CacheChangeUserCount(context.Context, *CacheChangeUserCountReq) (*CacheChangeUserCountRsp, error)
	CacheGetAuthor(context.Context, *CacheGetAuthorReq) (*CacheGetAuthorRsp, error)
	UpdateUserFavoritedCount(context.Context, *UpdateUserFavoritedCountReq) (*UpdateUserFavoritedCountRsp, error)
	UpdateUserFavoriteCount(context.Context, *UpdateUserFavoriteCountReq) (*UpdateUserFavoriteCountRsp, error)
	UpdateUserFollowCount(context.Context, *UpdateUserFollowCountReq) (*UpdateUserFollowCountRsp, error)
	UpdateUserFollowerCount(context.Context, *UpdateUserFollowerCountReq) (*UpdateUserFollowerCountRsp, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserServiceServer) GetUserInfoDict(context.Context, *GetUserInfoDictRequest) (*GetUserInfoDictResopnse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoDict not implemented")
}
func (UnimplementedUserServiceServer) CheckPassWord(context.Context, *CheckPassWordRequest) (*CheckPassWordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPassWord not implemented")
}
func (UnimplementedUserServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServiceServer) GetUserInfoList(context.Context, *GetUserInfoListRequest) (*GetUserInfoListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoList not implemented")
}
func (UnimplementedUserServiceServer) CacheChangeUserCount(context.Context, *CacheChangeUserCountReq) (*CacheChangeUserCountRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CacheChangeUserCount not implemented")
}
func (UnimplementedUserServiceServer) CacheGetAuthor(context.Context, *CacheGetAuthorReq) (*CacheGetAuthorRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CacheGetAuthor not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserFavoritedCount(context.Context, *UpdateUserFavoritedCountReq) (*UpdateUserFavoritedCountRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserFavoritedCount not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserFavoriteCount(context.Context, *UpdateUserFavoriteCountReq) (*UpdateUserFavoriteCountRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserFavoriteCount not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserFollowCount(context.Context, *UpdateUserFollowCountReq) (*UpdateUserFollowCountRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserFollowCount not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserFollowerCount(context.Context, *UpdateUserFollowerCountReq) (*UpdateUserFollowerCountRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserFollowerCount not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInfoDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoDictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfoDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserInfoDict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfoDict(ctx, req.(*GetUserInfoDictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CheckPassWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPassWordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CheckPassWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CheckPassWord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CheckPassWord(ctx, req.(*CheckPassWordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInfoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserInfoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfoList(ctx, req.(*GetUserInfoListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CacheChangeUserCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CacheChangeUserCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CacheChangeUserCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CacheChangeUserCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CacheChangeUserCount(ctx, req.(*CacheChangeUserCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CacheGetAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CacheGetAuthorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CacheGetAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CacheGetAuthor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CacheGetAuthor(ctx, req.(*CacheGetAuthorReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserFavoritedCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserFavoritedCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserFavoritedCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUserFavoritedCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserFavoritedCount(ctx, req.(*UpdateUserFavoritedCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserFavoriteCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserFavoriteCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserFavoriteCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUserFavoriteCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserFavoriteCount(ctx, req.(*UpdateUserFavoriteCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserFollowCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserFollowCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserFollowCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUserFollowCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserFollowCount(ctx, req.(*UpdateUserFollowCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserFollowerCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserFollowerCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserFollowerCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUserFollowerCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserFollowerCount(ctx, req.(*UpdateUserFollowerCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _UserService_GetUserInfo_Handler,
		},
		{
			MethodName: "GetUserInfoDict",
			Handler:    _UserService_GetUserInfoDict_Handler,
		},
		{
			MethodName: "CheckPassWord",
			Handler:    _UserService_CheckPassWord_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "GetUserInfoList",
			Handler:    _UserService_GetUserInfoList_Handler,
		},
		{
			MethodName: "CacheChangeUserCount",
			Handler:    _UserService_CacheChangeUserCount_Handler,
		},
		{
			MethodName: "CacheGetAuthor",
			Handler:    _UserService_CacheGetAuthor_Handler,
		},
		{
			MethodName: "UpdateUserFavoritedCount",
			Handler:    _UserService_UpdateUserFavoritedCount_Handler,
		},
		{
			MethodName: "UpdateUserFavoriteCount",
			Handler:    _UserService_UpdateUserFavoriteCount_Handler,
		},
		{
			MethodName: "UpdateUserFollowCount",
			Handler:    _UserService_UpdateUserFollowCount_Handler,
		},
		{
			MethodName: "UpdateUserFollowerCount",
			Handler:    _UserService_UpdateUserFollowerCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
