package rpc

import (
	"context"
	"github.com/IPAM/kitex_gen/artwork"
	"github.com/IPAM/kitex_gen/artwork/artworkservice"
	"github.com/IPAM/pkg/consts"
	"github.com/IPAM/pkg/errno"
	"github.com/IPAM/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var artworkClient artworkservice.Client

func initArtworkClient() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	c, err := artworkservice.NewClient(
		consts.ArtworkServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	artworkClient = c
}

// GetEthBalance get eth balance
func GetEthBalance(ctx context.Context, req *artwork.GetEthBalanceRequest) (int64, error) {
	resp, err := artworkClient.GetEthBalance(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Balance, nil
}

// GetTokenBalance get token balance
func GetTokenBalance(ctx context.Context, req *artwork.GetOwnerBalanceRequest) (int64, error) {
	resp, err := artworkClient.GetOwnerBalance(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Balance, nil
}

// BuyCollection buy the collection
func BuyCollection(ctx context.Context, req *artwork.BuyCollectionRequest) (int64, error) {
	resp, err := artworkClient.BuyCollection(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Nonce, nil
}
