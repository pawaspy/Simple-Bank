package worker

import (
	"context"

	"github.com/hibiken/asynq"
	db "github.com/pawaspy/simple_bank/db/sqlc"
)

const (
	QueueCritical = "critical"
	QueueDefault = "default"
)

type TaskProcessor interface {
	Start() error
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	sever *asynq.Server
	store db.Store
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store) TaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Queues: map[string]int{
				QueueCritical: 10,
				QueueDefault: 5,
			},
		},
	)

	return &RedisTaskProcessor{
		sever: server,
		store: store,
	}
}

func (processor *RedisTaskProcessor) Start() error{
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskVerifyEmail, processor.ProcessTaskSendVerifyEmail)

	return processor.sever.Start(mux)
}