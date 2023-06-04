package mq

import (
	"github.com/IPAM/pkg/consts"
	"github.com/hibiken/asynq"
)

var Client *asynq.Client

func Init() {

	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     consts.RedisAddr,
		Password: consts.RedisPass,
	})

	Client = client
}
