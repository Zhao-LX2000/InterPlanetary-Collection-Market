package service

import (
	"context"
	"github.com/IPAM/kitex_gen/file"
	"github.com/IPAM/server/file/dal/db"
	"github.com/IPAM/server/file/pack"
)

type GetCollectionListService struct {
	ctx context.Context
}

func NewGetCollectionListService(ctx context.Context) *GetCollectionListService {
	return &GetCollectionListService{ctx: ctx}
}

func (s *GetCollectionListService) GetCollectionList(req *file.GetCollectionListRequest) ([]*file.Collection, error) {
	listFile, err := db.ListFile(s.ctx)
	if err != nil {
		return nil, err
	}
	return pack.Files(listFile), nil
}
