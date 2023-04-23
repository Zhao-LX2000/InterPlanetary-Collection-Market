package geth

import (
	"github.com/IPAM/sol"
)

var CollectionTokenService *sol.CollectionTokenService

var OriginService *sol.OriginService

func Init() {

	OriginService = sol.NewOriginService()

	CollectionTokenService = OriginService.NewCollectionTokenService()
}

//func Init() {
//	var err error
//
//	Client, err = ethclient.Dial(consts.GethURL)
//	if err != nil {
//		log.Fatal(err)
//		panic(err)
//	}
//
//	initCollectionToken()
//}
//
//func initCollectionToken() {
//	var err error
//
//	TokenClient, err = sol.NewCollectionToken(common.HexToAddress(consts.TokenSOlAddress), Client)
//	if err != nil {
//		log.Fatal(err)
//		panic(err)
//	}
//}
//
//func CloseClient(client *ethclient.Client) {
//	client.Close()
//}
