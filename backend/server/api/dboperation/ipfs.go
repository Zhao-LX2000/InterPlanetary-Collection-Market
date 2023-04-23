package dboperation

import (
	"context"
	"github.com/IPAM/pkg/consts"
	"gorm.io/gorm"
	"log"
)

type IPFSFile struct {
	gorm.Model
	FileName        string `json:"filename"`
	AuthorName      string `json:"authorname"`
	WorkName        string `json:"workname"`
	Owners          string `json:"owners"`
	WorkDescription string `json:"workdescription"`
	Hash            string `json:"hash"`
	Url             string `json:"url"`
	Price           int    `json:"price"`
	Type            int    `json:"type"`
}

func (u *IPFSFile) TableName() string {
	return consts.IPFSTableName
}

// MGetUsers multiple get list of user info
//func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
//	res := make([]*User, 0)
//	if len(userIDs) == 0 {
//		return res, nil
//	}
//
//	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
//		return nil, err
//	}
//	return res, nil
//}

func AddFile(ctx context.Context, files []*IPFSFile) error {
	return DB.WithContext(ctx).Create(files).Error
}

func ListFile(ctx context.Context) ([]*IPFSFile, error) {
	var vs []*IPFSFile
	err := DB.WithContext(ctx).Order("created_at desc").Limit(10).Find(&vs).Error
	log.Printf("list:%#v\n", vs)
	return vs, err
}

func QueryFile(ctx context.Context, id int) ([]*IPFSFile, error) {
	res := make([]*IPFSFile, 0)
	if err := DB.WithContext(ctx).Where("id = ?", id).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func ChangeOwner(ctx context.Context, id int, newer string) error {
	return DB.Model(&IPFSFile{}).Where("id = ?", id).UpdateColumn("owners", newer).Error
}
