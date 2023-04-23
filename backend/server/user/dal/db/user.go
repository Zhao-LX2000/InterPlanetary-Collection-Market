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

package db

import (
	"context"
	"github.com/IPAM/pkg/consts"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `json:"username"`
	Password  string `json:"password"`
	Keystore  string `json:"keystore"`
	PublicKey string `json:"public_key"`
}

func (u *User) TableName() string {
	return consts.UserTableName
}

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	//klog.Info("the user is ", &users)
	return DB.WithContext(ctx).Create(users).Error
	//sql := "update order_tbl set descs=? where id=?"
	//ret, err := aDB.ExecContext(ctx, sql, fmt.Sprintf("NewDescs1-%d", time.Now().UnixMilli()), 1)
	//if err != nil {
	//	fmt.Printf("update failed, err:%v\n", err)
	//	return nil
	//}
	//
	//rows, err := ret.RowsAffected()
	//if err != nil {
	//	fmt.Printf("update failed, err:%v\n", err)
	//	return nil
	//}
	//if rows == 0 {
	//	return fmt.Errorf("rows affected 0")
	//}
	//fmt.Printf("update success： %d.\n", rows)
	//return nil
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
