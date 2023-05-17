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
	"github.com/IPAM/kitex_gen/file"
	"github.com/IPAM/server/file/dal/db"
)

// User pack user info
func File(ipfs *db.IPFSFile) *file.Collection {
	if ipfs == nil {
		return nil
	}
	return &file.Collection{Id: int64(ipfs.ID),
		Filename:        ipfs.FileName,
		Authorname:      ipfs.AuthorName,
		Owners:          ipfs.Owners,
		Workname:        ipfs.WorkName,
		Workdescription: ipfs.WorkDescription,
		Hash:            ipfs.Hash,
		Url:             ipfs.Url,
		Price:           int64(ipfs.Price),
		Createat:        ipfs.CreatedAt.String(),
	}
}

// Users pack list of user info
func Files(us []*db.IPFSFile) []*file.Collection {
	files := make([]*file.Collection, 0)
	for _, u := range us {
		if temp := File(u); temp != nil {
			files = append(files, temp)
		}
	}
	return files
}
