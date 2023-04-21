// Code generated by protoc-gen-psrpc v0.3.0, DO NOT EDIT.
// source: rpc/signal.proto

package rpc

import (
	"context"

	"github.com/livekit/psrpc"
	"github.com/livekit/psrpc/pkg/client"
	"github.com/livekit/psrpc/pkg/info"
	"github.com/livekit/psrpc/pkg/server"
	"github.com/livekit/psrpc/version"
)

var _ = version.PsrpcVersion_0_3_0

// =======================
// Signal Client Interface
// =======================

type SignalClient[NodeIdTopicType ~string] interface {
	RelaySignal(ctx context.Context, nodeId NodeIdTopicType, opts ...psrpc.RequestOption) (psrpc.ClientStream[*RelaySignalRequest, *RelaySignalResponse], error)
}

// ===========================
// Signal ServerImpl Interface
// ===========================

type SignalServerImpl interface {
	RelaySignal(psrpc.ServerStream[*RelaySignalResponse, *RelaySignalRequest]) error
}

// =======================
// Signal Server Interface
// =======================

type SignalServer[NodeIdTopicType ~string] interface {
	RegisterRelaySignalTopic(nodeId NodeIdTopicType) error
	DeregisterRelaySignalTopic(nodeId NodeIdTopicType)

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// =============
// Signal Client
// =============

type signalClient[NodeIdTopicType ~string] struct {
	client *client.RPCClient
}

// NewSignalClient creates a psrpc client that implements the SignalClient interface.
func NewSignalClient[NodeIdTopicType ~string](clientID string, bus psrpc.MessageBus, opts ...psrpc.ClientOption) (SignalClient[NodeIdTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "Signal",
		ID:   clientID,
	}

	sd.RegisterMethod("RelaySignal", false, false, false)

	rpcClient, err := client.NewRPCClientWithStreams(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &signalClient[NodeIdTopicType]{
		client: rpcClient,
	}, nil
}

func (c *signalClient[NodeIdTopicType]) RelaySignal(ctx context.Context, nodeId NodeIdTopicType, opts ...psrpc.RequestOption) (psrpc.ClientStream[*RelaySignalRequest, *RelaySignalResponse], error) {
	return client.OpenStream[*RelaySignalRequest, *RelaySignalResponse](ctx, c.client, "RelaySignal", []string{string(nodeId)}, opts...)
}

// =============
// Signal Server
// =============

type signalServer[NodeIdTopicType ~string] struct {
	svc SignalServerImpl
	rpc *server.RPCServer
}

// NewSignalServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewSignalServer[NodeIdTopicType ~string](serverID string, svc SignalServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (SignalServer[NodeIdTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "Signal",
		ID:   serverID,
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("RelaySignal", false, false, false)
	return &signalServer[NodeIdTopicType]{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *signalServer[NodeIdTopicType]) RegisterRelaySignalTopic(nodeId NodeIdTopicType) error {
	return server.RegisterStreamHandler(s.rpc, "RelaySignal", []string{string(nodeId)}, s.svc.RelaySignal, nil)
}

func (s *signalServer[NodeIdTopicType]) DeregisterRelaySignalTopic(nodeId NodeIdTopicType) {
	s.rpc.DeregisterHandler("RelaySignal", []string{string(nodeId)})
}

func (s *signalServer[NodeIdTopicType]) Shutdown() {
	s.rpc.Close(false)
}

func (s *signalServer[NodeIdTopicType]) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor3 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcd, 0x4e, 0x32, 0x31,
	0x14, 0x4d, 0x19, 0x02, 0xdf, 0x57, 0x24, 0xc1, 0x12, 0xa1, 0x99, 0x15, 0x99, 0x15, 0x31, 0xa6,
	0x43, 0x20, 0x6e, 0x5c, 0xfa, 0x08, 0x65, 0xa5, 0x1b, 0x02, 0xa5, 0xc1, 0x06, 0x6c, 0x4b, 0x6f,
	0x21, 0xf1, 0x11, 0x5c, 0xfb, 0x26, 0x3e, 0x85, 0x8f, 0x65, 0x66, 0x3a, 0x33, 0x8c, 0x99, 0xe8,
	0x6a, 0xce, 0xdc, 0xf3, 0xd3, 0x73, 0x5b, 0x3c, 0x70, 0x56, 0xa4, 0xa0, 0x76, 0x7a, 0x7d, 0x60,
	0xd6, 0x19, 0x6f, 0x48, 0xe4, 0xac, 0x88, 0xfb, 0xc6, 0x7a, 0x65, 0x34, 0x84, 0x59, 0x3c, 0x3a,
	0xa8, 0xb3, 0xdc, 0x2b, 0xbf, 0x52, 0xda, 0x4b, 0x57, 0x69, 0xe3, 0xeb, 0x72, 0xee, 0xbc, 0x08,
	0xa3, 0xe4, 0x0b, 0x61, 0xc2, 0xe5, 0x61, 0xfd, 0xb6, 0xcc, 0x43, 0xb9, 0x3c, 0x9e, 0x24, 0x78,
	0xf2, 0x80, 0xfb, 0xe0, 0xd7, 0xce, 0xaf, 0x40, 0x02, 0x28, 0xa3, 0x29, 0x9a, 0xa0, 0x69, 0x6f,
	0x7e, 0xc3, 0x8a, 0x04, 0xb6, 0xcc, 0xd8, 0x65, 0x20, 0xf9, 0x15, 0xd4, 0xfe, 0xc8, 0x0c, 0x77,
	0x5d, 0x88, 0xa1, 0xad, 0xdc, 0x35, 0xba, 0xb8, 0xea, 0x87, 0xf0, 0x52, 0x46, 0xe6, 0xf8, 0x5f,
	0x01, 0x81, 0x46, 0x93, 0xe8, 0x0f, 0x4b, 0xa5, 0x23, 0x03, 0x1c, 0x81, 0x3c, 0xd2, 0xf6, 0x04,
	0x4d, 0xdb, 0x3c, 0x83, 0xc9, 0x07, 0xc2, 0xc3, 0x1f, 0xab, 0x80, 0x35, 0x1a, 0x24, 0x59, 0x64,
	0xe9, 0x01, 0x17, 0x6b, 0x8c, 0x1b, 0xe9, 0x81, 0xe6, 0x95, 0x90, 0xdc, 0xe3, 0xff, 0x25, 0x06,
	0xda, 0xca, 0x3b, 0xfd, 0xea, 0xba, 0x28, 0xcb, 0x56, 0x51, 0xd5, 0x6a, 0x2e, 0x70, 0x27, 0xc8,
	0xc9, 0x13, 0xee, 0xd5, 0xea, 0x91, 0x31, 0x73, 0x56, 0xb0, 0xe6, 0xdd, 0xc7, 0xb4, 0x49, 0x84,
	0xf8, 0x64, 0xfc, 0xf9, 0x8e, 0x86, 0x14, 0x25, 0x7d, 0xd2, 0xd5, 0x66, 0x2b, 0x57, 0x6a, 0x9b,
	0x2d, 0x34, 0x43, 0x8f, 0x77, 0xcf, 0xb7, 0x3b, 0xe5, 0x5f, 0x4e, 0x1b, 0x26, 0xcc, 0x6b, 0x5a,
	0xd4, 0xac, 0xbe, 0x76, 0xbf, 0x4b, 0x41, 0xba, 0xb3, 0x12, 0x32, 0x75, 0x56, 0x6c, 0x3a, 0xf9,
	0xd3, 0x2f, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x19, 0x7d, 0x02, 0x15, 0x4d, 0x02, 0x00, 0x00,
}
