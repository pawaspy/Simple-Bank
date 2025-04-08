package gapi

import (
	"context"

	"github.com/lib/pq"
	db "github.com/pawaspy/simple_bank/db/sqlc"
	"github.com/pawaspy/simple_bank/pb"
	"github.com/pawaspy/simple_bank/util"
	"github.com/pawaspy/simple_bank/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	violations := validatorACreateUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot hash password: %s", err)
	}

	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
		Email:          req.GetEmail(),
		FullName:       req.GetFullName(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "user already exits: %s", err)
			}
			return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
		}
	}

	rsp := &pb.CreateUserResponse{
		User: convertor(user),
	}
	return rsp, nil
}

func validatorACreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err :=  val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err)) 
	}

	if err :=  val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err)) 
	}

	if err :=  val.ValidateFullName(req.GetFullName()); err != nil {
		violations = append(violations, fieldViolation("full_name", err)) 
	}

	if err :=  val.ValidateUsername(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err)) 
	}
	return
}