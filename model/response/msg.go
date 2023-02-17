package response

var MsgFlags = map[int]string {
    SUCCESS : "Successful Request",
    ERROR : "Server Fail",
    NOTFOUND : "Not found",
    FORBIDDEN: "Request Forbidden",
    BADREQUEST : "Bad request",
    ERROR_EXIST_TAG : "Tag name already existed",
    ERROR_NOT_EXIST_TAG : "Tag does not exist",
    ERROR_NOT_EXIST_ARTICLE : "Article does not exist",
    ERROR_AUTH_CHECK_TOKEN_FAIL : "Request denied (JWT Fail)",
    ERROR_AUTH_CHECK_TOKEN_NOT_FOUND : "Token not found",
    ERROR_AUTH_CHECK_TOKEN_EMPTY : "There's no Token",
    ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token time out",
    ERROR_AUTH_CHECK_TOKEN_IN_BLACK_LIST: "Token is in black list",
    ERROR_AUTH_TOKEN : "Fail to generate Token",
    ERROR_AUTH : "Token error",
}

func GetMsg(code int) string {
    msg, ok := MsgFlags[code]
    if ok {
        return msg
    }
    return MsgFlags[ERROR]
}