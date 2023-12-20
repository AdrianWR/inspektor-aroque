// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: inspection/v1/inspection.proto

/*
Package inspectionv1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package inspectionv1

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_InspectionService_GetInspection_0(ctx context.Context, marshaler runtime.Marshaler, client InspectionServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetInspectionRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetInspection(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_InspectionService_GetInspection_0(ctx context.Context, marshaler runtime.Marshaler, server InspectionServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetInspectionRequest
	var metadata runtime.ServerMetadata

	msg, err := server.GetInspection(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterInspectionServiceHandlerServer registers the http handlers for service InspectionService to "mux".
// UnaryRPC     :call InspectionServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterInspectionServiceHandlerFromEndpoint instead.
func RegisterInspectionServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server InspectionServiceServer) error {

	mux.Handle("GET", pattern_InspectionService_GetInspection_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/inspection.v1.InspectionService/GetInspection", runtime.WithHTTPPathPattern("/v1/inspect"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_InspectionService_GetInspection_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InspectionService_GetInspection_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterInspectionServiceHandlerFromEndpoint is same as RegisterInspectionServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterInspectionServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterInspectionServiceHandler(ctx, mux, conn)
}

// RegisterInspectionServiceHandler registers the http handlers for service InspectionService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterInspectionServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterInspectionServiceHandlerClient(ctx, mux, NewInspectionServiceClient(conn))
}

// RegisterInspectionServiceHandlerClient registers the http handlers for service InspectionService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "InspectionServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "InspectionServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "InspectionServiceClient" to call the correct interceptors.
func RegisterInspectionServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client InspectionServiceClient) error {

	mux.Handle("GET", pattern_InspectionService_GetInspection_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/inspection.v1.InspectionService/GetInspection", runtime.WithHTTPPathPattern("/v1/inspect"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_InspectionService_GetInspection_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InspectionService_GetInspection_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_InspectionService_GetInspection_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "inspect"}, ""))
)

var (
	forward_InspectionService_GetInspection_0 = runtime.ForwardResponseMessage
)
