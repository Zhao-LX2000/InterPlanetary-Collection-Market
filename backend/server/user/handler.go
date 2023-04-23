package main

import (
	"context"
	"github.com/IPAM/kitex_gen/testUser"
	"github.com/IPAM/pkg/errno"
	"github.com/IPAM/server/user/pack"
	"github.com/IPAM/server/user/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *testuser.CreateUserRequest) (resp *testuser.CreateUserResponse, err error) {
	resp = new(testuser.CreateUserResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *testuser.MGetUserRequest) (resp *testuser.MGetUserResponse, err error) {
	resp = new(testuser.MGetUserResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	users, err := service.NewMGetUserService(ctx).MGetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Users = users
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *testuser.CheckUserRequest) (resp *testuser.CheckUserResponse, err error) {
	resp = new(testuser.CheckUserResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	user, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.PublicKey = user.PublicKey
	resp.AccountName = user.Username
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// QueryUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUser(ctx context.Context, req *testuser.QueryUserRequest) (resp *testuser.QueryUserResponse, err error) {
	resp = new(testuser.QueryUserResponse)

	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	user, err := service.NewQueryUserService(ctx).QueryUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	respu := new(testuser.User)
	respu.UserId = int64(user.ID)
	respu.Username = user.Username
	respu.PublicKey = user.PublicKey
	respu.Keystore = user.Keystore
	resp.User = respu
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
