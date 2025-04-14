package main

import (
	"context"
	"database/sql"
	"net"
	"net/http"
	"os"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/pawaspy/simple_bank/api"
	db "github.com/pawaspy/simple_bank/db/sqlc"
	"github.com/pawaspy/simple_bank/gapi"
	"github.com/pawaspy/simple_bank/mail"
	"github.com/pawaspy/simple_bank/pb"
	"github.com/pawaspy/simple_bank/util"
	"github.com/pawaspy/simple_bank/worker"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Info().Msg("Cannot open the config file")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Info().Msg("Cannot connect to db: ")
	}

	store := db.NewStore(conn)

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)
	go runTaskProcessor(config, redisOpt, store)
	go runGatewayServer(config, taskDistributor, store)
	runGrpcServer(config, taskDistributor, store)
}

func runTaskProcessor(config util.Config, redisOpt asynq.RedisClientOpt, store db.Store) {
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store, mailer)
	log.Info().Msg("start task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}
}

func runGrpcServer(config util.Config, taskDistributor worker.TaskDistributor, store db.Store) {
	server, err := gapi.NewServer(config, taskDistributor, store)
	if err != nil {
		log.Info().Msg("cannot create server")
	}

	logger := grpc.UnaryInterceptor(gapi.GrpcLogger)

	grpcServer := grpc.NewServer(logger)
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCAddress)
	if err != nil {
		log.Info().Msg("cannot create listener")
	}
	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Info().Msg("cannot start the server: ")
	}
}

func runGatewayServer(config util.Config, taskDistributor worker.TaskDistributor, store db.Store) {
	server, err := gapi.NewServer(config, taskDistributor, store)
	if err != nil {
		log.Info().Msg("cannot create server: ")
	}

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})
	grpcMux := runtime.NewServeMux(jsonOption)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterSimpleBankHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Info().Msg("cannot register handler server: ")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	fs := http.FileServer(http.Dir("./docs/swagger"))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	listener, err := net.Listen("tcp", config.HTTPAddress)
	if err != nil {
		log.Info().Msg("cannot create listener: ")
	}
	log.Info().Msgf("start HTTP server at %s", listener.Addr().String())

	err = http.Serve(listener, mux)
	if err != nil {
		log.Info().Msg("cannot start the server: ")
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Info().Msg("cannot start server")
	}

	err = server.Start(config.HTTPAddress)
	if err != nil {
		log.Info().Msg("Cannot start the server")
	}
}
