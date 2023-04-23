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
	//ArtworkTableName 还没建Artwork表
	ArtworkTableName    = "artwork"
	UserTableName       = "user"
	IPFSTableName       = "ipfs"
	SecretKey           = "secret key"
	IdentityKey         = "addr"
	AccountName         = "fileName"
	Total               = "total"
	ApiServiceName      = "demoapi"
	UserServiceName     = "testuser"
	ArtworkServiceName  = "artwork"
	MySQLDefaultDSN     = "gorm:1119@tcp(localhost:7712)/gorm?charset=utf8&parseTime=True&loc=Local"
	IPFSMySQLDefaultDSN = "gorm:1119@tcp(localhost:7712)/gorm?charset=utf8&parseTime=True&loc=Local"
	//GethURL 填你的geth
	GethURL            = "http://localhost:8545"
	TCP                = "tcp"
	UserServiceAddr    = ":9000"
	ArtworkServiceAddr = ":11000"
	NoteServiceAddr    = ":10000"
	ExportEndpoint     = ":4317"
	ETCDAddress        = "localhost:2379"
	DefaultLimit       = 10
	//CollectionTokenSOlAddress 收藏币合约的地址
	TokenSOlAddress           = "0x21dC5f46601683789d4e30c78675312d0D7c4cB8"
	CollectionTokenSOlAddress = "0x6fD056Deb2d1ED162a22C944756011EDf5cb0061"
	//MinerAddress 合约发布者的地址
	MinerAddress  = "0x721d8bf469ea9053210631eda4a6b0c4857b6df1"
	MinerName     = "auth"
	MinerPassword = "123456"
	ChainID       = 12345
)
