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
	"fmt"
	testuser "github.com/IPAM/kitex_gen/testUser"
	"github.com/IPAM/pkg/errno"
	"github.com/IPAM/server/user/dal/db"
)

type QueryUserService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewQueryUserService(ctx context.Context) *QueryUserService {
	return &QueryUserService{
		ctx: ctx,
	}
}

// CheckUser check user info
func (s *QueryUserService) QueryUser(req *testuser.QueryUserRequest) (*db.User, error) {
	userName := req.Username
	users, err := db.QueryUser(s.ctx, userName)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errno.AuthorizationFailedErr
	}
	u := users[0]
	fmt.Println(u)
	return u, nil
}
