package response

var MsgFlags = map[int]string {
    SUCCESS : "ok",
    ERROR : "fail",
    NOTFOUND : "NOt found",
    INVALID_PARAMS : "error of parameters request",
    ERROR_EXIST_TAG : "tag name already existed",
    ERROR_NOT_EXIST_TAG : "tag does not exist",
    ERROR_NOT_EXIST_ARTICLE : "article does not exist",
    ERROR_AUTH_CHECK_TOKEN_FAIL : "Fail to check Token",
    ERROR_AUTH_CHECK_TOKEN_NOT_FOUND : "Token not found",
    ERROR_AUTH_CHECK_TOKEN_EMPTY : "There's no Token",
    ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token time out",
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