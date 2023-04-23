package sol

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/IPAM/pkg/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

//var client *ethclient.Client
//
//var tokenClient *CollectionToken

// OriginService 提供对以太坊原生的服务
type OriginService struct {
	client *ethclient.Client
}

// CollectionTokenService 提供对收藏品合约的服务
type CollectionTokenService struct {
	client *OriginService

	tokenClient *CollectionToken
}

func NewOriginService() *OriginService {

	client, err := ethclient.Dial(consts.GethURL)
	if err != nil {
		log.Fatal(err)
	}

	return &OriginService{client: client}
}

func (server *OriginService) NewCollectionTokenService() *CollectionTokenService {

	tokenClient, err := NewCollectionToken(common.HexToAddress(consts.CollectionTokenSOlAddress), server.client)
	if err != nil {
		log.Fatal(err)
	}

	return &CollectionTokenService{
		tokenClient: tokenClient,
		client:      server,
	}
}

func (server *OriginService) GetEthBalance(address string) (int64, error) {
	fromAddr := common.HexToAddress(address)
	balanceAt, err := server.client.BalanceAt(context.Background(), fromAddr, nil)
	if err != nil {
		return 0, err
	}

	return balanceAt.Int64(), nil
}

func (server *CollectionTokenService) GetTokenBalance(address string) (int64, error) {

	fmt.Println("addddd" + address)
	balance, err := server.tokenClient.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))

	if err != nil {
		return 0, err
	}

	return balance.Int64(), err
}

func (server *CollectionTokenService) TokenTransfer(from, to, filename, password string, value *big.Int) (uint64, error) {

	auth, err := server.client.PackTransactOpts(from, filename, password)
	if err != nil {
		klog.Info("the err happen", err)
		return 0, err
	}

	transfer, err := server.tokenClient.Transfer(auth, common.HexToAddress(to), value)

	if err != nil {
		klog.Info("the err happen", err)
		return 0, err
	}

	fmt.Println("the tx to is ", transfer.To())
	fmt.Println("the tx value is ", transfer.Value())
	fmt.Println("the tx gas is ", transfer.Gas())

	nonce := transfer.Nonce()

	return nonce, err

}

// MineTokenForAccount 给新账户添加一百个收藏品币
func (server *CollectionTokenService) MineTokenForAccount(toAddr string) error {

	auth, err := server.client.PackTransactOpts(consts.MinerAddress, consts.MinerName, consts.MinerPassword)
	if err != nil {
		klog.Info("the err happen", err)
		return err
	}

	auth.Value = big.NewInt(0) // in wei

	mint, err := server.tokenClient.Mint(auth, common.HexToAddress(toAddr), big.NewInt(100))

	if err != nil {
		klog.Info("the err happen", err)
		return err
	}
	fmt.Println("the tx nonce is ", mint.Nonce())

	return err
}

// MineEthForAccount 给新账号添加1个以太币
func (server *OriginService) MineEthForAccount(toAddr string) error {
	auth, err := server.PackTransactOpts(consts.MinerAddress, consts.MinerName, consts.MinerPassword)
	if err != nil {
		klog.Info("the err happen", err)
		return err
	}
	auth.Value = big.NewInt(1000000000000000000) // in wei

	to := common.HexToAddress(toAddr)

	//var rawTx *types.Transaction
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    auth.Nonce.Uint64(),
		To:       &to,
		Value:    big.NewInt(1000000000000000000),
		Gas:      auth.GasLimit,
		GasPrice: auth.GasPrice,
		Data:     common.FromHex("damn this is 4 u"),
	})

	klog.Info("the chainId is ", tx.ChainId())

	signer, err := auth.Signer(auth.From, tx)

	if err != nil {
		klog.Info("the err happen", err)
	}

	err = server.client.SendTransaction(context.Background(), signer)

	return err
}

func (server *OriginService) PackTransactOpts(from, filename, password string) (*bind.TransactOpts, error) {

	privateKey, err := GetPrivateKey(from, filename, password)
	if err != nil {
		klog.Info("the err happen", err)
		return nil, err
	}

	nonce, err := server.client.PendingNonceAt(context.Background(), common.HexToAddress(from))
	if err != nil {
		klog.Info("the err happen", err)
		return nil, err
	}

	gasPrice, err := server.client.SuggestGasPrice(context.Background())
	if err != nil {
		klog.Info("the err happen", err)
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(consts.ChainID))
	if err != nil {
		klog.Info("the err happen", err)
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(300000) // in units
	//这里设置了gasPrice还是没设置没区别
	auth.GasPrice = gasPrice
	auth.Value = big.NewInt(0) // in wei

	return auth, err
}

func GetPrivateKey(from, filename, password string) (*ecdsa.PrivateKey, error) {

	//keystore存储位置
	ks := keystore.NewKeyStore("/mnt/f/key/tmp/"+filename, keystore.StandardScryptN, keystore.StandardScryptP)
	account := accounts.Account{Address: common.HexToAddress(from)}

	//拿到密钥那些
	keyJSON, err := ks.Export(account, password, password)
	if err != nil {
		klog.Info("the err happen", err)
		return nil, err
	}

	//同上
	key, err := keystore.DecryptKey(keyJSON, password)
	if err != nil {
		klog.Info("the err happen", err)
		return nil, err

	}

	privateKey := key.PrivateKey

	return privateKey, err
}
