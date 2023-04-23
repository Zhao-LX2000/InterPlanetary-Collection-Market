package service

import (
	"context"
	"fmt"
	"github.com/IPAM/kitex_gen/artwork"
	"github.com/IPAM/server/artwork/dal/geth"
	"github.com/cloudwego/kitex/pkg/klog"
	"math/big"
)

type TokenService struct {
	ctx context.Context
}

// NewTokenService get the token Balance
func NewTokenService(ctx context.Context) *TokenService {
	return &TokenService{
		ctx: ctx,
	}
}

// GetTokenBalance get user token balance
func (s *TokenService) GetTokenBalance(req *artwork.GetOwnerBalanceRequest) (int64, error) {

	addr := req.GetAddress()
	//fromAddr := common.HexToAddress(addr)
	//balanceAt, err := geth.Client.BalanceAt(context.Background(), fromAddr, nil)
	//balance, err := geth.GetTokenBalance(addr)
	balance, err := geth.CollectionTokenService.GetTokenBalance(addr)

	if err != nil {
		return 0, err
	}

	fmt.Println("the addr is ", addr)
	fmt.Println("the balance is", balance)

	return balance, nil
}

// TransferTokenTo transfer Token to
func (s *TokenService) TransferTokenTo(req *artwork.BuyCollectionRequest) (uint64, error) {
	fmt.Println("打印req：", req)
	from := req.GetFrom()
	to := req.GetTo()
	password := req.GetPassword()
	filename := req.GetFilename()
	artHash := req.GetArtHash()
	val := req.GetValue()
	fmt.Println("根据 artHash 去数据库查找其价值 ====>", artHash)
	//这里需要请求geth去查找收藏品的价格
	//val := big.NewInt(15)

	//执行交易
	nonce, err := geth.CollectionTokenService.TokenTransfer(from, to, filename, password, big.NewInt(val))

	if err != nil {
		return 0, err
	}

	//最后查一下余额有没有变化 ----需删
	balance, err := geth.CollectionTokenService.GetTokenBalance(from)

	if err != nil {
		return 0, err
	}

	klog.Info("the balance is ", balance)

	return nonce, nil
}
