// Copyright Istio Authors
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

package sink

import (
	"context"
	"io"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	mcp "istio.io/api/mcp/v1alpha1"
	"istio.io/libistio/pkg/mcp/rate"
)

// RateLimiter is partially representing standard lib's rate limiter
type RateLimiter interface {
	Wait(ctx context.Context) (err error)
}

// AuthChecker is used to check the transport auth info that is associated with each stream. If the function
// returns nil, then the connection will be allowed. If the function returns an error, then it will be
// percolated up to the gRPC stack.
//
// Note that it is possible that this method can be called with nil authInfo. This can happen either if there
// is no peer info, or if the underlying gRPC stream is insecure. The implementations should be resilient in
// this case and apply appropriate policy.
type AuthChecker interface {
	Check(authInfo credentials.AuthInfo) error
}

// Server implements the server for the MCP sink service. The server is the sink and receives configuration
// from the client.
type Server struct {
	authCheck            AuthChecker
	newConnectionLimiter RateLimiter
	sink                 *Sink
}

var _ mcp.ResourceSinkServer = &Server{}

// ServerOptions contains source server specific options
type ServerOptions struct {
	AuthChecker AuthChecker
	RateLimiter rate.Limit
}

// NewServer creates a new instance of a MCP sink server.
func NewServer(sinkOptions *Options, serverOptions *ServerOptions) *Server {
	s := &Server{
		sink:                 New(sinkOptions),
		authCheck:            serverOptions.AuthChecker,
		newConnectionLimiter: serverOptions.RateLimiter,
	}
	return s
}

// EstablishResourceStream implements the ResourceSinkServer interface.
func (s *Server) EstablishResourceStream(stream mcp.ResourceSink_EstablishResourceStreamServer) error {
	// TODO support receiving configuration from multiple sources?
	// TODO MVP - limit to one connection at a time?

	if err := s.newConnectionLimiter.Wait(stream.Context()); err != nil {
		return err
	}
	var authInfo credentials.AuthInfo
	if peerInfo, ok := peer.FromContext(stream.Context()); ok {
		authInfo = peerInfo.AuthInfo
	} else {
		scope.Warnf("No peer info found on the incoming stream.")
	}

	if err := s.authCheck.Check(authInfo); err != nil {
		return status.Errorf(codes.Unauthenticated, "Authentication failure: %v", err)
	}

	err := s.sink.ProcessStream(stream)
	code := status.Code(err)
	if code == codes.OK || code == codes.Canceled || err == io.EOF {
		return nil
	}
	return err
}
