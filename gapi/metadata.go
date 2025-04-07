package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader  = "user-agents"
	xForwadedForHeader = "x-forwaded-for"
)

type Metadata struct {
	ClientIP  string
	UserAgent string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	md := &Metadata{}

	if mtdt, ok := metadata.FromIncomingContext(ctx); ok {
		if userAgents := mtdt.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			md.UserAgent = userAgents[0]
		}

		if userAgents := mtdt.Get(userAgentHeader); len(userAgents) > 0 {
			md.UserAgent = userAgents[0]
		}

		if clientIPs := mtdt.Get(xForwadedForHeader); len(clientIPs) >0 {
			md.ClientIP = clientIPs[0]	
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		md.ClientIP = p.Addr.String()
	}

	return md
}
