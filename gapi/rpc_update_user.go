package gapi

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/pawaspy/simple_bank/db/sqlc"
	"github.com/pawaspy/simple_bank/pb"
	"github.com/pawaspy/simple_bank/util"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	violation := validateUpdateUserRequest(req)
	if violation != nil {
		return nil, invalidArgumentError(violation)
	}

	arg := db.UpdateUserParams{
		Username: req.GetUsername(),
		FullName: sql.NullString{
			String: req.GetFullName(),
			Valid:  req.FullName != nil,
		},
		Email: sql.NullString{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
	}

	if req.Password != nil {
		hashedPass, err := util.HashPassword(req.GetPassword())
		if err != nil {
			return nil, fmt.Errorf("cannot hash password")
		}
		arg.HashedPassword = sql.NullString{
			String: hashedPass,
			Valid: true,
		}
	}
}

func validateUpdateUserRequest(req *pb.UpdateUserRequest) (violation []*errdetails.BadRequest_FieldViolation) {
	return
}
