package service

import (
	"context"
	"fmt"
	"github.com/IPAM/kitex_gen/artwork"
	"github.com/IPAM/server/artwork/dal/geth"
)

type GetEthBalanceService struct {
	ctx context.Context
}

// NewGetEthBalanceService new this service
func NewGetEthBalanceService(ctx context.Context) *GetEthBalanceService {
	return &GetEthBalanceService{
		ctx: ctx,
	}
}

// GetEthBalance get user eth balance
func (s *GetEthBalanceService) GetEthBalance(req *artwork.GetEthBalanceRequest) (int64, error) {

	//defer geth.CloseClient(client)
	//if err != nil {
	//	return 0, err
	//}

	addr := req.GetAddress()

	fmt.Println("the addr is ", addr)

	balanceAt, err := geth.OriginService.GetEthBalance(addr)
	//fromAddr := common.HexToAddress(addr)
	//
	//balanceAt, err := geth.Client.BalanceAt(context.Background(), fromAddr, nil)

	fmt.Println("the balance is", balanceAt)

	return balanceAt, err
}
