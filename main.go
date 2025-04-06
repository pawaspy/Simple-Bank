package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/pawaspy/simple_bank/api"
	db "github.com/pawaspy/simple_bank/db/sqlc"
	"github.com/pawaspy/simple_bank/gapi"
	"github.com/pawaspy/simple_bank/pb"
	"github.com/pawaspy/simple_bank/util"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot open the config file")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	runGrpcServer(config, store)
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}
	log.Printf("start gRPC server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start the server.")
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal(err)
	}

	err = server.Start(config.HTTPAddress)
	if err != nil {
		log.Fatal("Cannot start the server")
	}
}
