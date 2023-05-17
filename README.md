# InterPlanetary-Collection-Market

## Introduction

这是一个数字版权平台项目，简称IPAM，与市面上NFT平台类似，IPAM可以让用户上传自己的作品（包括图片、视频、音乐等）转换成数字资产，基于以太坊实现交易，使用IPFS存储用户的作品数据。

项目后端采用微服务架构，基于Kitex RPC框架、Hertz 高性能Web框架开发，采用ETCD作为服务注册与发现中心。拆分出路由、用户、文件、藏品四个微服务:

| Service    | Usage                      | Framework   | protocol   | Path                   |
| ---------- | -------------------------- | ----------- | ---------- | ---------------------- |
| api        | http interface             | kitex/hertz | http       | backend/server/api     |
| user       | user data management       | kitex/gorm  | thrift rpc | backend/server/user    |
| file       | file upload and management | kitex/gorm  | thrift rpc | backend/server/file    |
| collection | geth client                | kitex       | thrift rpc | backend/server/artwork |

项目前端采用Vue3开发，组件库使用element plus和naive UI。
![前端](https://github.com/Zhao-LX2000/InterPlanetary-Collection-Market/blob/master/frontend/src/assets/2.png)

项目仍处于开发阶段，不定期更新

### Feature

* 平台使用实现ECR20标准的智能合约发行IPC代币
* 新注册的用户，矿主会向其转一定量的比特币和IPC代币，通过上传作品基于用户代币奖励
* 网页上展示的缩略图采用根据作者提供的作品描述AI生成，以FTP服务形式存储在服务器上，通过Nginx访问
* 用户真实的作品数据采用IPFS点对点大型网络存储，更具安全性
* 所有类型的数据都可以转换成NFT（目前借助于IPFS的hash，后期考虑采用ECR721实现）
* 支持容器化运行，内置每个服务的Dockerfile，后可通过YAML文件在Kubernetes集群上运行

## api

### User Service

##### /v1/user/register

用户注册

##### /v1/user/login

用户登录

### File Service

##### /v1/file/GetCollectionList

获取所有上传的作品信息

##### /v1/file/UploadCollection

上传新作品

### Collection Service

##### /v1/artwork/BuyCollection

购买作品

##### /v1/artwork/getBalance

获取比特币余额

##### /v1/artwork/getTokenBalance

获取代币余额

## deployment

### back-end

modify your server IP and account path

    backend/pkg/consts/consts.go

### front-end

download denpencies：

```
npm i
```

modify your server IP

```
frontend/src/axios/index.ts
```

run project:

```
npm run dev
```

