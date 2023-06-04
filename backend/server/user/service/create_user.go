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

package service

import (
	"context"
	"crypto/md5"
	"fmt"
	testuser "github.com/IPAM/kitex_gen/testUser"
	"github.com/IPAM/pkg/errno"
	"github.com/IPAM/server/user/dal/db"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io"
	"log"
)

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *CreateUserService) CreateUser(req *testuser.CreateUserRequest) error {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	//生成keystore
	ks := keystore.NewKeyStore("./tmp/"+req.Username, keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := ks.NewAccount(req.Password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())

	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))

	//存储到数据库
	err = db.CreateUser(s.ctx, []*db.User{{
		Username:  req.Username,
		Password:  password,
		PublicKey: account.Address.Hex(),
		Keystore:  account.URL.Path,
	}})

	fmt.Println("the hex is ", account.Address.Hex())
	fmt.Println("the string is ", account.Address.String())

	//err = geth.CollectionTokenService.MineTokenForAccount(account.Address.Hex())
	//err = geth.OriginService.MineEthForAccount(account.Address.Hex())

	return err
}
