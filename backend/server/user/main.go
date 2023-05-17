package main

import (
	"github.com/IPAM/kitex_gen/testUser/userservice"
	"github.com/IPAM/pkg/consts"
	"github.com/IPAM/pkg/mw"
	"github.com/IPAM/server/user/dal"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func Init() {
	dal.Init()
	// klog init
	klog.SetLevel(klog.LevelInfo)
}

func main() {
	//svr := testuser.NewServer(new(UserServiceImpl))

	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	ip, err := consts.GetOutBoundIP()
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr(consts.TCP, ip+consts.UserServiceAddr)
	if err != nil {
		panic(err)
	}

	Init()

	svr := userservice.NewServer(new(UserServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.UserServiceName}),
	)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
