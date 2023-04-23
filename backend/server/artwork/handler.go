package main

import (
	"context"
	"github.com/IPAM/kitex_gen/artwork"
	"github.com/IPAM/pkg/errno"
	"github.com/IPAM/server/artwork/pack"
	"github.com/IPAM/server/artwork/service"
)

// ArtworkServiceImpl implements the last service interface defined in the IDL.
type ArtworkServiceImpl struct{}

// BuyArtwork implements the ArtworkServiceImpl interface.
func (s *ArtworkServiceImpl) BuyArtwork(ctx context.Context, req *artwork.BuyArtworkRequest) (resp *artwork.BuyArtworkResponse, err error) {
	// TODO: Your code here...
	return
}

// GetOwnerArtwork implements the ArtworkServiceImpl interface.
func (s *ArtworkServiceImpl) GetOwnerArtwork(ctx context.Context, req *artwork.GetOwnerArtworkRequest) (resp *artwork.GetOwnerArtworkResponse, err error) {
	// TODO: Your code here...
	return
}

// GetOwnerBalance implements the ArtworkServiceImpl interface.
func (s *ArtworkServiceImpl) GetOwnerBalance(ctx context.Context, req *artwork.GetOwnerBalanceRequest) (resp *artwork.GetOwnerBalanceResponse, err error) {
	// TODO: Your code here...

	resp = new(artwork.GetOwnerBalanceResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	//balance, err := service.NewGetEthBalanceService(ctx).GetEthBalance(req)
	//进
	balance, err := service.NewTokenService(ctx).GetTokenBalance(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Balance = balance
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetEthBalance implements the ArtworkServiceImpl interface.
func (s *ArtworkServiceImpl) GetEthBalance(ctx context.Context, req *artwork.GetEthBalanceRequest) (resp *artwork.GetEthBalanceResponse, err error) {
	// TODO: Your code here...

	resp = new(artwork.GetEthBalanceResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	balance, err := service.NewGetEthBalanceService(ctx).GetEthBalance(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Balance = balance
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// BuyCollection implements the ArtworkServiceImpl interface.
func (s *ArtworkServiceImpl) BuyCollection(ctx context.Context, req *artwork.BuyCollectionRequest) (resp *artwork.BuyCollectionResponse, err error) {
	// TODO: Your code here...
	resp = new(artwork.BuyCollectionResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	nonce, err := service.NewTokenService(ctx).TransferTokenTo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Nonce = int64(nonce)
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
