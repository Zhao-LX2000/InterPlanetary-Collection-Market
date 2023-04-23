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

package pack

import (
	"github.com/IPAM/kitex_gen/testUser"
	"github.com/IPAM/server/user/dal/db"
)

// User pack user info
func User(u *db.User) *testuser.User {
	if u == nil {
		return nil
	}

	return &testuser.User{UserId: int64(u.ID), Username: u.Username, Avatar: "test", Keystore: u.Keystore, PublicKey: u.PublicKey}
}

// Users pack list of user info
func Users(us []*db.User) []*testuser.User {
	users := make([]*testuser.User, 0)
	for _, u := range us {
		if temp := User(u); temp != nil {
			users = append(users, temp)
		}
	}
	return users
}
