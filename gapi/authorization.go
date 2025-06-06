package gapi

import (
	"context"
	"fmt"
	"strings"

	"github.com/pawaspy/simple_bank/token"
	"google.golang.org/grpc/metadata"
)

const (
	authHeader = "authorization"
	authorizationType   = "bearer"
)

func (server *Server) authorizeUser(ctx context.Context, accessibleRoles []string) (*token.Payload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	values := md.Get(authHeader)
	if len(values) == 0 {
		return nil, fmt.Errorf("missing authorization header")
	}

	authHead := values[0]
	fields := strings.Fields(authHead)
	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationType {
		return nil, fmt.Errorf("unsupported authorization format")
	}

	accessToken := fields[1]
	payload, err := server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("unverified token")
	}

	if !hasPermission(payload.Role,accessibleRoles) {
		return nil, fmt.Errorf("permission denied")
	}
	return payload, nil
}

func hasPermission(userRole string, accessibleRoles []string) bool {
	for _, role := range accessibleRoles{
		if userRole == role {
			return true
		}
	}
	return false
}
