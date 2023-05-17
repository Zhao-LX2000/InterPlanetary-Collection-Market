package main

import (
	"context"
	file "github.com/IPAM/kitex_gen/file"
	"github.com/IPAM/pkg/errno"
	"github.com/IPAM/server/file/pack"
	"github.com/IPAM/server/file/service"
	"log"
)

// FileServiceImpl implements the last service interface defined in the IDL.
type FileServiceImpl struct{}

// UploadCollection implements the FileServiceImpl interface.
func (s *FileServiceImpl) UploadCollection(ctx context.Context, req *file.UploadCollectionRequest) (resp *file.UploadCollectionResponse, err error) {
	resp = new(file.UploadCollectionResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewUploadCollectionService(ctx).UploadCollection(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetCollectionListCollection implements the FileServiceImpl interface.
func (s *FileServiceImpl) GetCollectionListCollection(ctx context.Context, req *file.GetCollectionListRequest) (resp *file.GetCollectionListResponse, err error) {
	resp = new(file.GetCollectionListResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	log.Println("3调用")
	list, err := service.NewGetCollectionListService(ctx).GetCollectionList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Collections = list
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
