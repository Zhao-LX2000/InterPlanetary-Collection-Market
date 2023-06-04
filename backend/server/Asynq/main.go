package main

import (
	"github.com/IPAM/pkg/consts"
	"github.com/IPAM/server/api/mq/mqHandler"
	"github.com/IPAM/server/api/mq/task"
	"github.com/hibiken/asynq"
	"log"
)

func main() {

	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     consts.RedisAddr,
			Password: consts.RedisPass,
		},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeUpFile2IPFS, mqHandler.HandleUpFile2IPFSTask)
	// ...register other handlers...

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}

}
