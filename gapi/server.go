package gapi

import (
	"fmt"

	db "github.com/pawaspy/simple_bank/db/sqlc"
	"github.com/pawaspy/simple_bank/pb"
	"github.com/pawaspy/simple_bank/token"
	"github.com/pawaspy/simple_bank/util"
	"github.com/pawaspy/simple_bank/worker"
)

// Server serves gRPC requests for our banking system
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	taskDistributor worker.TaskDistributor
	tokenMaker token.Maker
}

func NewServer(config util.Config,taskDistributor worker.TaskDistributor, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		taskDistributor: taskDistributor,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
