package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IPAM/kitex_gen/file"
	"github.com/IPAM/server/file/dal/db"
	"github.com/jlaffaye/ftp"
	"io"
	"net/http"
	"time"
)

type UploadCollectionService struct {
	ctx context.Context
}

func NewUploadCollectionService(ctx context.Context) *UploadCollectionService {
	return &UploadCollectionService{ctx: ctx}
}

func (s *UploadCollectionService) UploadCollection(req *file.UploadCollectionRequest) error {
	// 连接FTP服务器
	ftpClient, err := ftp.Dial("43.136.22.7:21")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer ftpClient.Quit()
	// 登录FTP服务器
	err = ftpClient.Login("ipcmgo", "ipcmgo")
	if err != nil {
		fmt.Println(err)
		return err
	}
	get, err := http.Get(req.Url)
	defer get.Body.Close()
	// 上传文件
	all, err := io.ReadAll(get.Body)
	timestamp := time.Now().Unix()
	t := time.Unix(timestamp, 0)
	tstr := t.Format("2006-01-02 15:04:05")
	nginxfilename := tstr + "-" + req.Filename
	err = ftpClient.Stor(nginxfilename, bytes.NewReader(all))
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = db.AddFile(s.ctx, []*db.IPFSFile{{
		FileName:        req.Filename,
		AuthorName:      req.Authorname,
		Owners:          req.Owners,
		WorkName:        req.Workname,
		WorkDescription: req.Workdescription,
		Hash:            req.Hash,
		Url:             "http://43.136.22.7/images/" + nginxfilename,
		Price:           int(req.Price),
	}})
	if err != nil {
		return err
	}
	return nil
}
