package response

import "net/http"

const (
    SUCCESS = http.StatusOK
    ERROR = http.StatusInternalServerError
    BADREQUEST = http.StatusBadRequest
    NOTFOUND = http.StatusNotFound
    FORBIDDEN = http.StatusForbidden

    ERROR_EXIST_TAG = 501
    ERROR_NOT_EXIST_TAG = 502
    ERROR_NOT_EXIST_ARTICLE = 503
    ERROR_AUTH_CHECK_TOKEN_EMPTY = 511
    ERROR_AUTH_CHECK_TOKEN_FAIL = 512
    ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 513
    ERROR_AUTH_CHECK_TOKEN_NOT_FOUND = 514
    ERROR_AUTH_CHECK_TOKEN_IN_BLACK_LIST = 515
    ERROR_AUTH_TOKEN = 516
    ERROR_AUTH = 517
)