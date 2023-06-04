package mqHandler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/IPAM/pkg/consts"
	"github.com/IPAM/server/api/mq/task"
	"github.com/hibiken/asynq"
	shell "github.com/ipfs/go-ipfs-api"
	"time"
)

func HandleUpFile2IPFSTask(ctx context.Context, t *asynq.Task) error {
	var p task.UpFilePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	fmt.Println("bro i got to sleep for a while")

	time.Sleep(time.Second * 5)

	account := p.Account
	fmt.Println("the accountName is", account)

	file := p.File

	// 连接到本地 IPFS 节点
	sh := shell.NewShell(consts.IPFSAddr)

	// 将文件上传到 IPFS，并获得哈希值
	hash, err := sh.Add(bytes.NewReader(file))
	if err != nil {
		fmt.Println("send file got wrong")
		return err
	}

	// 可选：将哈希值保存到数据库或返回给客户端
	fmt.Println("File uploaded successfully. IPFS hash: ", hash)

	return nil
}
