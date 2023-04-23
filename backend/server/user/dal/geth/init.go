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
