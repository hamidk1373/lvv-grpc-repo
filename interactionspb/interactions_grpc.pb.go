// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: allprotos/interactions.proto

package interactionspb

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

// InteractionsServiceClient is the client API for InteractionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InteractionsServiceClient interface {
	AddSpace(ctx context.Context, in *AddSpaceRequest, opts ...grpc.CallOption) (*AddSpaceResponse, error)
	AddShort(ctx context.Context, in *AddShortRequest, opts ...grpc.CallOption) (*AddShortResponse, error)
	UpdateShortDetails(ctx context.Context, in *UpdateShortDetailsRequest, opts ...grpc.CallOption) (*UpdateShortDetailsResponse, error)
	AddShortComment(ctx context.Context, in *AddShortCommentRequest, opts ...grpc.CallOption) (*AddShortCommentResponse, error)
	EditShortComment(ctx context.Context, in *EditShortCommentRequest, opts ...grpc.CallOption) (*EditShortCommentResponse, error)
	DeleteShortComment(ctx context.Context, in *DeleteShortCommentRequest, opts ...grpc.CallOption) (*DeleteShortCommentResponse, error)
	AddShortCommentReply(ctx context.Context, in *AddShortCommentReplyRequest, opts ...grpc.CallOption) (*AddShortCommentReplyResponse, error)
	EditShortCommentReply(ctx context.Context, in *EditShortCommentReplyRequest, opts ...grpc.CallOption) (*EditShortCommentReplyResponse, error)
	DeleteShortCommentReply(ctx context.Context, in *DeleteShortCommentReplyRequest, opts ...grpc.CallOption) (*DeleteShortCommentReplyResponse, error)
	ListShortComments(ctx context.Context, in *ListShortCommentsRequest, opts ...grpc.CallOption) (*ListShortCommentsResponse, error)
	GetShortsInsightsSummary(ctx context.Context, in *GetShortsInsightsSummaryRequest, opts ...grpc.CallOption) (*GetShortsInsightsSummaryResponse, error)
	GetShortsInsightList(ctx context.Context, in *GetShortsInsightListRequest, opts ...grpc.CallOption) (*GetShortsInsightListResponse, error)
	GetShortFullInsight(ctx context.Context, in *GetShortFullInsightRequest, opts ...grpc.CallOption) (*GetShortFullInsightResponse, error)
	LikeShort(ctx context.Context, in *LikeShortRequest, opts ...grpc.CallOption) (*LikeShortResponse, error)
	UnlikeShort(ctx context.Context, in *UnlikeShortRequest, opts ...grpc.CallOption) (*UnlikeShortResponse, error)
	BookmarkShort(ctx context.Context, in *BookmarkShortRequest, opts ...grpc.CallOption) (*BookmarkShortResponse, error)
	UnbookmarkShort(ctx context.Context, in *UnbookmarkShortRequest, opts ...grpc.CallOption) (*UnbookmarkShortResponse, error)
	InteractWithShortButton(ctx context.Context, in *InteractWithShortButtonRequest, opts ...grpc.CallOption) (*InteractWithShortButtonResponse, error)
	InteractWithShortHotspot(ctx context.Context, in *InteractWithShortHotspotRequest, opts ...grpc.CallOption) (*InteractWithShortHotspotResponse, error)
	InteractWithShortProduct(ctx context.Context, in *InteractWithShortProductRequest, opts ...grpc.CallOption) (*InteractWithShortProductResponse, error)
	ReportShort(ctx context.Context, in *ReportShortRequest, opts ...grpc.CallOption) (*ReportShortResponse, error)
	RecordView(ctx context.Context, in *RecordViewRequest, opts ...grpc.CallOption) (*RecordViewResponse, error)
}

type interactionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInteractionsServiceClient(cc grpc.ClientConnInterface) InteractionsServiceClient {
	return &interactionsServiceClient{cc}
}

func (c *interactionsServiceClient) AddSpace(ctx context.Context, in *AddSpaceRequest, opts ...grpc.CallOption) (*AddSpaceResponse, error) {
	out := new(AddSpaceResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/AddSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionsServiceClient) AddShort(ctx context.Context, in *AddShortRequest, opts ...grpc.CallOption) (*AddShortResponse, error) {
	out := new(AddShortResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/AddShort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionsServiceClient) UpdateShortDetails(ctx context.Context, in *UpdateShortDetailsRequest, opts ...grpc.CallOption) (*UpdateShortDetailsResponse, error) {
	out := new(UpdateShortDetailsResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/UpdateShortDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionsServiceClient) AddShortComment(ctx context.Context, in *AddShortCommentRequest, opts ...grpc.CallOption) (*AddShortCommentResponse, error) {
	out := new(AddShortCommentResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/AddShortComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionsServiceClient) EditShortComment(ctx context.Context, in *EditShortCommentRequest, opts ...grpc.CallOption) (*EditShortCommentResponse, error) {
	out := new(EditShortCommentResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/EditShortComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionsServiceClient) DeleteShortComment(ctx context.Context, in *DeleteShortCommentRequest, opts ...grpc.CallOption) (*DeleteShortCommentResponse, error) {
	out := new(DeleteShortCommentResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/DeleteShortComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionsServiceClient) AddShortCommentReply(ctx context.Context, in *AddShortCommentReplyRequest, opts ...grpc.CallOption) (*AddShortCommentReplyResponse, error) {
	out := new(AddShortCommentReplyResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/AddShortCommentReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionsServiceClient) EditShortCommentReply(ctx context.Context, in *EditShortCommentReplyRequest, opts ...grpc.CallOption) (*EditShortCommentReplyResponse, error) {
	out := new(EditShortCommentReplyResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/EditShortCommentReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionsServiceClient) DeleteShortCommentReply(ctx context.Context, in *DeleteShortCommentReplyRequest, opts ...grpc.CallOption) (*DeleteShortCommentReplyResponse, error) {
	out := new(DeleteShortCommentReplyResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/DeleteShortCommentReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionsServiceClient) ListShortComments(ctx context.Context, in *ListShortCommentsRequest, opts ...grpc.CallOption) (*ListShortCommentsResponse, error) {
	out := new(ListShortCommentsResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/ListShortComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionsServiceClient) GetShortsInsightsSummary(ctx context.Context, in *GetShortsInsightsSummaryRequest, opts ...grpc.CallOption) (*GetShortsInsightsSummaryResponse, error) {
	out := new(GetShortsInsightsSummaryResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/GetShortsInsightsSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionsServiceClient) GetShortsInsightList(ctx context.Context, in *GetShortsInsightListRequest, opts ...grpc.CallOption) (*GetShortsInsightListResponse, error) {
	out := new(GetShortsInsightListResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/GetShortsInsightList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionsServiceClient) GetShortFullInsight(ctx context.Context, in *GetShortFullInsightRequest, opts ...grpc.CallOption) (*GetShortFullInsightResponse, error) {
	out := new(GetShortFullInsightResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/GetShortFullInsight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionsServiceClient) LikeShort(ctx context.Context, in *LikeShortRequest, opts ...grpc.CallOption) (*LikeShortResponse, error) {
	out := new(LikeShortResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/LikeShort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionsServiceClient) UnlikeShort(ctx context.Context, in *UnlikeShortRequest, opts ...grpc.CallOption) (*UnlikeShortResponse, error) {
	out := new(UnlikeShortResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/UnlikeShort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionsServiceClient) BookmarkShort(ctx context.Context, in *BookmarkShortRequest, opts ...grpc.CallOption) (*BookmarkShortResponse, error) {
	out := new(BookmarkShortResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/BookmarkShort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionsServiceClient) UnbookmarkShort(ctx context.Context, in *UnbookmarkShortRequest, opts ...grpc.CallOption) (*UnbookmarkShortResponse, error) {
	out := new(UnbookmarkShortResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/UnbookmarkShort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionsServiceClient) InteractWithShortButton(ctx context.Context, in *InteractWithShortButtonRequest, opts ...grpc.CallOption) (*InteractWithShortButtonResponse, error) {
	out := new(InteractWithShortButtonResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/InteractWithShortButton", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionsServiceClient) InteractWithShortHotspot(ctx context.Context, in *InteractWithShortHotspotRequest, opts ...grpc.CallOption) (*InteractWithShortHotspotResponse, error) {
	out := new(InteractWithShortHotspotResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/InteractWithShortHotspot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionsServiceClient) InteractWithShortProduct(ctx context.Context, in *InteractWithShortProductRequest, opts ...grpc.CallOption) (*InteractWithShortProductResponse, error) {
	out := new(InteractWithShortProductResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/InteractWithShortProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionsServiceClient) ReportShort(ctx context.Context, in *ReportShortRequest, opts ...grpc.CallOption) (*ReportShortResponse, error) {
	out := new(ReportShortResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/ReportShort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionsServiceClient) RecordView(ctx context.Context, in *RecordViewRequest, opts ...grpc.CallOption) (*RecordViewResponse, error) {
	out := new(RecordViewResponse)
	err := c.cc.Invoke(ctx, "/interactions.InteractionsService/RecordView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InteractionsServiceServer is the server API for InteractionsService service.
// All implementations must embed UnimplementedInteractionsServiceServer
// for forward compatibility
type InteractionsServiceServer interface {
	AddSpace(context.Context, *AddSpaceRequest) (*AddSpaceResponse, error)
	AddShort(context.Context, *AddShortRequest) (*AddShortResponse, error)
	UpdateShortDetails(context.Context, *UpdateShortDetailsRequest) (*UpdateShortDetailsResponse, error)
	AddShortComment(context.Context, *AddShortCommentRequest) (*AddShortCommentResponse, error)
	EditShortComment(context.Context, *EditShortCommentRequest) (*EditShortCommentResponse, error)
	DeleteShortComment(context.Context, *DeleteShortCommentRequest) (*DeleteShortCommentResponse, error)
	AddShortCommentReply(context.Context, *AddShortCommentReplyRequest) (*AddShortCommentReplyResponse, error)
	EditShortCommentReply(context.Context, *EditShortCommentReplyRequest) (*EditShortCommentReplyResponse, error)
	DeleteShortCommentReply(context.Context, *DeleteShortCommentReplyRequest) (*DeleteShortCommentReplyResponse, error)
	ListShortComments(context.Context, *ListShortCommentsRequest) (*ListShortCommentsResponse, error)
	GetShortsInsightsSummary(context.Context, *GetShortsInsightsSummaryRequest) (*GetShortsInsightsSummaryResponse, error)
	GetShortsInsightList(context.Context, *GetShortsInsightListRequest) (*GetShortsInsightListResponse, error)
	GetShortFullInsight(context.Context, *GetShortFullInsightRequest) (*GetShortFullInsightResponse, error)
	LikeShort(context.Context, *LikeShortRequest) (*LikeShortResponse, error)
	UnlikeShort(context.Context, *UnlikeShortRequest) (*UnlikeShortResponse, error)
	BookmarkShort(context.Context, *BookmarkShortRequest) (*BookmarkShortResponse, error)
	UnbookmarkShort(context.Context, *UnbookmarkShortRequest) (*UnbookmarkShortResponse, error)
	InteractWithShortButton(context.Context, *InteractWithShortButtonRequest) (*InteractWithShortButtonResponse, error)
	InteractWithShortHotspot(context.Context, *InteractWithShortHotspotRequest) (*InteractWithShortHotspotResponse, error)
	InteractWithShortProduct(context.Context, *InteractWithShortProductRequest) (*InteractWithShortProductResponse, error)
	ReportShort(context.Context, *ReportShortRequest) (*ReportShortResponse, error)
	RecordView(context.Context, *RecordViewRequest) (*RecordViewResponse, error)
	mustEmbedUnimplementedInteractionsServiceServer()
}

// UnimplementedInteractionsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInteractionsServiceServer struct {
}

func (UnimplementedInteractionsServiceServer) AddSpace(context.Context, *AddSpaceRequest) (*AddSpaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSpace not implemented")
}
func (UnimplementedInteractionsServiceServer) AddShort(context.Context, *AddShortRequest) (*AddShortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddShort not implemented")
}
func (UnimplementedInteractionsServiceServer) UpdateShortDetails(context.Context, *UpdateShortDetailsRequest) (*UpdateShortDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateShortDetails not implemented")
}
func (UnimplementedInteractionsServiceServer) AddShortComment(context.Context, *AddShortCommentRequest) (*AddShortCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddShortComment not implemented")
}
func (UnimplementedInteractionsServiceServer) EditShortComment(context.Context, *EditShortCommentRequest) (*EditShortCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditShortComment not implemented")
}
func (UnimplementedInteractionsServiceServer) DeleteShortComment(context.Context, *DeleteShortCommentRequest) (*DeleteShortCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShortComment not implemented")
}
func (UnimplementedInteractionsServiceServer) AddShortCommentReply(context.Context, *AddShortCommentReplyRequest) (*AddShortCommentReplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddShortCommentReply not implemented")
}
func (UnimplementedInteractionsServiceServer) EditShortCommentReply(context.Context, *EditShortCommentReplyRequest) (*EditShortCommentReplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditShortCommentReply not implemented")
}
func (UnimplementedInteractionsServiceServer) DeleteShortCommentReply(context.Context, *DeleteShortCommentReplyRequest) (*DeleteShortCommentReplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShortCommentReply not implemented")
}
func (UnimplementedInteractionsServiceServer) ListShortComments(context.Context, *ListShortCommentsRequest) (*ListShortCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShortComments not implemented")
}
func (UnimplementedInteractionsServiceServer) GetShortsInsightsSummary(context.Context, *GetShortsInsightsSummaryRequest) (*GetShortsInsightsSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShortsInsightsSummary not implemented")
}
func (UnimplementedInteractionsServiceServer) GetShortsInsightList(context.Context, *GetShortsInsightListRequest) (*GetShortsInsightListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShortsInsightList not implemented")
}
func (UnimplementedInteractionsServiceServer) GetShortFullInsight(context.Context, *GetShortFullInsightRequest) (*GetShortFullInsightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShortFullInsight not implemented")
}
func (UnimplementedInteractionsServiceServer) LikeShort(context.Context, *LikeShortRequest) (*LikeShortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeShort not implemented")
}
func (UnimplementedInteractionsServiceServer) UnlikeShort(context.Context, *UnlikeShortRequest) (*UnlikeShortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlikeShort not implemented")
}
func (UnimplementedInteractionsServiceServer) BookmarkShort(context.Context, *BookmarkShortRequest) (*BookmarkShortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookmarkShort not implemented")
}
func (UnimplementedInteractionsServiceServer) UnbookmarkShort(context.Context, *UnbookmarkShortRequest) (*UnbookmarkShortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnbookmarkShort not implemented")
}
func (UnimplementedInteractionsServiceServer) InteractWithShortButton(context.Context, *InteractWithShortButtonRequest) (*InteractWithShortButtonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InteractWithShortButton not implemented")
}
func (UnimplementedInteractionsServiceServer) InteractWithShortHotspot(context.Context, *InteractWithShortHotspotRequest) (*InteractWithShortHotspotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InteractWithShortHotspot not implemented")
}
func (UnimplementedInteractionsServiceServer) InteractWithShortProduct(context.Context, *InteractWithShortProductRequest) (*InteractWithShortProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InteractWithShortProduct not implemented")
}
func (UnimplementedInteractionsServiceServer) ReportShort(context.Context, *ReportShortRequest) (*ReportShortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportShort not implemented")
}
func (UnimplementedInteractionsServiceServer) RecordView(context.Context, *RecordViewRequest) (*RecordViewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordView not implemented")
}
func (UnimplementedInteractionsServiceServer) mustEmbedUnimplementedInteractionsServiceServer() {}

// UnsafeInteractionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InteractionsServiceServer will
// result in compilation errors.
type UnsafeInteractionsServiceServer interface {
	mustEmbedUnimplementedInteractionsServiceServer()
}

func RegisterInteractionsServiceServer(s grpc.ServiceRegistrar, srv InteractionsServiceServer) {
	s.RegisterService(&InteractionsService_ServiceDesc, srv)
}

func _InteractionsService_AddSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSpaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).AddSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/AddSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).AddSpace(ctx, req.(*AddSpaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionsService_AddShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddShortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).AddShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/AddShort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).AddShort(ctx, req.(*AddShortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionsService_UpdateShortDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateShortDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).UpdateShortDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/UpdateShortDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).UpdateShortDetails(ctx, req.(*UpdateShortDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionsService_AddShortComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddShortCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).AddShortComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/AddShortComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).AddShortComment(ctx, req.(*AddShortCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionsService_EditShortComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditShortCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).EditShortComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/EditShortComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).EditShortComment(ctx, req.(*EditShortCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionsService_DeleteShortComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteShortCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).DeleteShortComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/DeleteShortComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).DeleteShortComment(ctx, req.(*DeleteShortCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionsService_AddShortCommentReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddShortCommentReplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).AddShortCommentReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/AddShortCommentReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).AddShortCommentReply(ctx, req.(*AddShortCommentReplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionsService_EditShortCommentReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditShortCommentReplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).EditShortCommentReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/EditShortCommentReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).EditShortCommentReply(ctx, req.(*EditShortCommentReplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionsService_DeleteShortCommentReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteShortCommentReplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).DeleteShortCommentReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/DeleteShortCommentReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).DeleteShortCommentReply(ctx, req.(*DeleteShortCommentReplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionsService_ListShortComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListShortCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).ListShortComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/ListShortComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).ListShortComments(ctx, req.(*ListShortCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionsService_GetShortsInsightsSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShortsInsightsSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).GetShortsInsightsSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/GetShortsInsightsSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).GetShortsInsightsSummary(ctx, req.(*GetShortsInsightsSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionsService_GetShortsInsightList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShortsInsightListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).GetShortsInsightList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/GetShortsInsightList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).GetShortsInsightList(ctx, req.(*GetShortsInsightListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionsService_GetShortFullInsight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShortFullInsightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).GetShortFullInsight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/GetShortFullInsight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).GetShortFullInsight(ctx, req.(*GetShortFullInsightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionsService_LikeShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeShortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).LikeShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/LikeShort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).LikeShort(ctx, req.(*LikeShortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionsService_UnlikeShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlikeShortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).UnlikeShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/UnlikeShort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).UnlikeShort(ctx, req.(*UnlikeShortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionsService_BookmarkShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookmarkShortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).BookmarkShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/BookmarkShort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).BookmarkShort(ctx, req.(*BookmarkShortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionsService_UnbookmarkShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnbookmarkShortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).UnbookmarkShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/UnbookmarkShort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).UnbookmarkShort(ctx, req.(*UnbookmarkShortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionsService_InteractWithShortButton_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InteractWithShortButtonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).InteractWithShortButton(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/InteractWithShortButton",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).InteractWithShortButton(ctx, req.(*InteractWithShortButtonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionsService_InteractWithShortHotspot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InteractWithShortHotspotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).InteractWithShortHotspot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/InteractWithShortHotspot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).InteractWithShortHotspot(ctx, req.(*InteractWithShortHotspotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionsService_InteractWithShortProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InteractWithShortProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).InteractWithShortProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/InteractWithShortProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).InteractWithShortProduct(ctx, req.(*InteractWithShortProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionsService_ReportShort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportShortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).ReportShort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/ReportShort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).ReportShort(ctx, req.(*ReportShortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InteractionsService_RecordView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionsServiceServer).RecordView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interactions.InteractionsService/RecordView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionsServiceServer).RecordView(ctx, req.(*RecordViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InteractionsService_ServiceDesc is the grpc.ServiceDesc for InteractionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InteractionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "interactions.InteractionsService",
	HandlerType: (*InteractionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSpace",
			Handler:    _InteractionsService_AddSpace_Handler,
		},
		{
			MethodName: "AddShort",
			Handler:    _InteractionsService_AddShort_Handler,
		},
		{
			MethodName: "UpdateShortDetails",
			Handler:    _InteractionsService_UpdateShortDetails_Handler,
		},
		{
			MethodName: "AddShortComment",
			Handler:    _InteractionsService_AddShortComment_Handler,
		},
		{
			MethodName: "EditShortComment",
			Handler:    _InteractionsService_EditShortComment_Handler,
		},
		{
			MethodName: "DeleteShortComment",
			Handler:    _InteractionsService_DeleteShortComment_Handler,
		},
		{
			MethodName: "AddShortCommentReply",
			Handler:    _InteractionsService_AddShortCommentReply_Handler,
		},
		{
			MethodName: "EditShortCommentReply",
			Handler:    _InteractionsService_EditShortCommentReply_Handler,
		},
		{
			MethodName: "DeleteShortCommentReply",
			Handler:    _InteractionsService_DeleteShortCommentReply_Handler,
		},
		{
			MethodName: "ListShortComments",
			Handler:    _InteractionsService_ListShortComments_Handler,
		},
		{
			MethodName: "GetShortsInsightsSummary",
			Handler:    _InteractionsService_GetShortsInsightsSummary_Handler,
		},
		{
			MethodName: "GetShortsInsightList",
			Handler:    _InteractionsService_GetShortsInsightList_Handler,
		},
		{
			MethodName: "GetShortFullInsight",
			Handler:    _InteractionsService_GetShortFullInsight_Handler,
		},
		{
			MethodName: "LikeShort",
			Handler:    _InteractionsService_LikeShort_Handler,
		},
		{
			MethodName: "UnlikeShort",
			Handler:    _InteractionsService_UnlikeShort_Handler,
		},
		{
			MethodName: "BookmarkShort",
			Handler:    _InteractionsService_BookmarkShort_Handler,
		},
		{
			MethodName: "UnbookmarkShort",
			Handler:    _InteractionsService_UnbookmarkShort_Handler,
		},
		{
			MethodName: "InteractWithShortButton",
			Handler:    _InteractionsService_InteractWithShortButton_Handler,
		},
		{
			MethodName: "InteractWithShortHotspot",
			Handler:    _InteractionsService_InteractWithShortHotspot_Handler,
		},
		{
			MethodName: "InteractWithShortProduct",
			Handler:    _InteractionsService_InteractWithShortProduct_Handler,
		},
		{
			MethodName: "ReportShort",
			Handler:    _InteractionsService_ReportShort_Handler,
		},
		{
			MethodName: "RecordView",
			Handler:    _InteractionsService_RecordView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "allprotos/interactions.proto",
}
