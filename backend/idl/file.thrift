namespace go file

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

struct Collection {
    1: i64 id
    2: string filename
    3: string authorname
    4: string owners
    5: string workname
    6: string workdescription
    7: string hash
    8: string url
    9: i64 price
    10:string createat
}

struct UploadCollectionRequest {
    1: string filename
    2: string authorname
    3: string owners
    4: string workname
    5: string workdescription
    6: string hash
    7: string url
    8: i64 price
}

struct UploadCollectionResponse {
    1: BaseResp base_resp
}


struct GetCollectionListRequest {

}

struct GetCollectionListResponse {
    1: list<Collection> collections
    2: BaseResp base_resp
}

service FileService {
    UploadCollectionResponse UploadCollection(1: UploadCollectionRequest req)
    GetCollectionListResponse GetCollectionListCollection(1: GetCollectionListRequest req)
}