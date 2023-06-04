namespace go api

struct BaseResp {
    1: i64 status_code
    2: string status_message
    3: i64 service_time
}

struct User {
    1: i64 user_id
    2: string username
    3: string avatar
    4: string keystore
}

struct Collection {
    1: string filename
    2: string workname
    3: string workdescription
    4: string url
    5: i64 price
}

struct CreateUserRequest {
    1: string username (api.form="username", api.vd="len($) > 0")
    2: string password (api.form="password", api.vd="len($) > 0")
}

struct CreateUserResponse {
    1: BaseResp base_resp
}

struct CheckUserRequest {
    1: string username (api.form="username", api.vd="len($) > 0")
    2: string password (api.form="password", api.vd="len($) > 0")
}

struct CheckUserResponse {
    1: BaseResp base_resp
}

struct TestTokenRequest {
}

struct TestTokenRespon {
    1: BaseResp base_resp
}

struct GetEthBalanceRequest {
//   1: string address
}

struct GetEthBalanceResponse {
    1: BaseResp base_resp
    2: i64 Balance
}

struct GetTokenBalanceRequest {
//   1: string address
}

struct GetTokenBalanceResponse {
    1: BaseResp base_resp
    2: i64 Balance
}

struct BuyCollectionRequest {
//   1: string address
    1: string password
    2: string to
    3: string art_hash
}

struct BuyCollectionResponse {
    1: BaseResp base_resp
    2: i64 Nonce
}

struct UploadCollectionRequest {
    1: Collection collection
}

struct UploadCollectionResponse {
    1: BaseResp base_resp
}


struct GetCollectionListRequest {

}

struct GetCollectionListResponse {
    1: list<Collection> users
    2: BaseResp base_resp
}

service ApiService {
    CreateUserResponse CreateUser(1: CreateUserRequest req) (api.post="/v1/user/register")
    CheckUserResponse CheckUser(1: CheckUserRequest req) (api.post="/v1/user/login")
    TestTokenRespon TestToken(1: TestTokenRequest req) (api.get="/v1/test/testToken")

    //获取用户拥有以太币余额（测试用4
    GetEthBalanceResponse GetEthBalance(1: GetEthBalanceRequest req) (api.get="/v1/artwork/getBalance")
    //获取用户拥有以收藏代币余额
    GetTokenBalanceResponse GetCollectionTokenBalance(1: GetTokenBalanceRequest req) (api.get="/v1/artwork/getTokenBalance")
    //用户购买收藏品（目前只是转账）
    BuyCollectionResponse BuyCollection(1: BuyCollectionRequest req) (api.post="/v1/artwork/BuyCollection")

    UploadCollectionResponse UploadCollection(1: UploadCollectionRequest req) (api.post="/v1/file/UploadCollection")

    UploadCollectionResponse AsqnUploadFile(1: UploadCollectionRequest req) (api.post="/v1/file/AsqnUploadFile")

    GetCollectionListResponse GetCollectionListCollection(1: GetCollectionListRequest req)(api.get="/v1/file/GetCollectionList")

}