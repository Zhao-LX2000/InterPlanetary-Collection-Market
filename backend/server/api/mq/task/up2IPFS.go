package task

import (
	"encoding/json"
	"github.com/hibiken/asynq"
)

const (
	TypeUpFile2IPFS = "file:up2IPFS"
)

type UpFilePayload struct {
	File    []byte
	Account string
}

func NewUpFileTask(file []byte, account string) (*asynq.Task, error) {
	payload, err := json.Marshal(UpFilePayload{File: file,
		Account: account})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeUpFile2IPFS, payload), nil
}
