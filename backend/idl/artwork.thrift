namespace go artwork

enum ErrCode {
    SuccessCode                = 0
    ServiceErrCode             = 10001
    ParamErrCode               = 10002
    UserAlreadyExistErrCode    = 10003
    AuthorizationFailedErrCode = 10004
}

struct BaseResp {
    1: i64 status_code
    2: string status_message
    3: i64 service_time
}

struct Artwork {
    1: i64 art_id
    2: string art_hash
    3: string art_owner_addr
}
//
//struct Payload{
//   1: string art_hash
//}

struct BuyArtworkRequest {
   1: string art_hash
   2: string address
}

struct BuyArtworkResponse {
    1: BaseResp base_resp
}

struct BuyCollectionRequest {
   1: string art_hash
   2: string from
   3: string to
   4: string password
   5: string filename
   6: i64 value
}

struct BuyCollectionResponse {
    1: BaseResp base_resp
    2: i64 nonce
}

struct GetOwnerArtworkRequest {
   1: string address
}

struct GetOwnerArtworkResponse {
    1: BaseResp base_resp
    2: list<Artwork> artwork_list
}

struct GetOwnerBalanceRequest {
    1: string address
}

struct GetOwnerBalanceResponse {
    1: BaseResp base_resp
    2: i64 Balance
}

struct GetEthBalanceRequest {
   1: string address
}

struct GetEthBalanceResponse {
    1: BaseResp base_resp
    2: i64 Balance
}

service ArtworkService {
    //购买艺术品
    BuyArtworkResponse BuyArtwork(1: BuyArtworkRequest req)
    //购买收藏品
    BuyCollectionResponse BuyCollection(1: BuyCollectionRequest req)
    //获取用户拥有艺术品
    GetOwnerArtworkResponse GetOwnerArtwork(1: GetOwnerArtworkRequest req)
    //获取用户拥有艺术币余额
    GetOwnerBalanceResponse GetOwnerBalance(1: GetOwnerBalanceRequest req)
    //获取用户拥有以太币余额（测试用4
    GetEthBalanceResponse GetEthBalance(1: GetEthBalanceRequest req)
}