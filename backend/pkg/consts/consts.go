// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package consts

const (
	UserTableName       = "user"
	IPFSTableName       = "ipfs"
	SecretKey           = "secret key"
	IdentityKey         = "addr"
	AccountName         = "fileName"
	Total               = "total"
	ApiServiceName      = "demoapi"
	UserServiceName     = "testuser"
	ArtworkServiceName  = "artwork"
	FileServiceName     = "file"
	MySQLDefaultDSN     = "gorm:1119@tcp(43.136.22.7:7712)/gorm?charset=utf8&parseTime=True&loc=Local"
	IPFSMySQLDefaultDSN = "gorm:1119@tcp(43.136.22.7:7712)/gorm?charset=utf8&parseTime=True&loc=Local"
	//GethURL 填你的geth
	GethURL                   = "http://43.136.22.7:8545"
	TCP                       = "tcp"
	UserServiceAddr           = ":9000"
	ArtworkServiceAddr        = ":11000"
	FileServiceAddr           = ":12000"
	ETCDAddress               = "43.136.22.7:2379"
	CollectionTokenSOlAddress = "0xc76d37b7EfA2A1b1Dc4eC4B70EfC9bC7FF0EA2Fe"
	//MinerAddress 合约发布者的地址
	MinerAddress  = "0x16e27c766f8dE98863edc4F73F6Dcb0635ef6017"
	MinerName     = "auth"
	MinerPassword = "123456"
	ChainID       = 12345
)
