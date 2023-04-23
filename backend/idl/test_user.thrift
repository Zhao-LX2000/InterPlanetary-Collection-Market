namespace go testUser

enum ErrCode {
    SuccessCode                = 0
    ServiceErrCode             = 10001
    ParamErrCode               = 10002
    UserAlreadyExistErrCode    = 10003
    AuthorizationFailedErrCode = 10004
    UnaffordableErrCode        = 10005
}

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
    5: string public_key
}

struct CreateUserRequest {
    1: string username (vt.min_size = "1")
    2: string password (vt.min_size = "1")
    //
//    3: string xid
}

struct CreateUserResponse {
    1: BaseResp base_resp
}

struct MGetUserRequest {
    1: list<i64> user_ids (vt.min_size = "1")
}

struct MGetUserResponse {
    1: list<User> users
    2: BaseResp base_resp
}

struct QueryUserRequest {
    1: string username (vt.min_size = "1")
}

struct QueryUserResponse {
    1: User user
    2: BaseResp base_resp
}

struct CheckUserRequest {
    1: string username (vt.min_size = "1")
    2: string password (vt.min_size = "1")
}

struct CheckUserResponse {
    1: string public_key
    2: string account_name
    3: BaseResp base_resp
}

service UserService {
    CreateUserResponse CreateUser(1: CreateUserRequest req)
    MGetUserResponse MGetUser(1: MGetUserRequest req)
    CheckUserResponse CheckUser(1: CheckUserRequest req)
    QueryUserResponse QueryUser(1: QueryUserRequest req)
}