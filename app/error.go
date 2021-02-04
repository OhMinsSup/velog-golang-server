package app

import (
	"github.com/OhMinsSup/story-server/helpers"
	"net/http"
)

type ErrorStatus string

type ResponseException struct {
	Code          int          `json:"code"`
	Message       ErrorStatus  `json:"message"`
	ResultCode    int          `json:"result_code"`
	ResultMessage string       `json:"result_message"`
	Data          helpers.JSON `json:"data"`
}

const (
	AlreadyExist       = ErrorStatus("ALREADY_EXIST")
	BadRequest         = ErrorStatus("BAD_REQUEST")
	NotFound           = ErrorStatus("NOT_FOUND")
	Forbidden          = ErrorStatus("FORBIDDEN")
	InteralServer      = ErrorStatus("INTERNAL_SERVER_ERROR")
	UnAuthorized       = ErrorStatus("UNAUTHORIZED")
	NotExist           = ErrorStatus("NOT_EXIST")
	DBQueryError       = ErrorStatus("DB_QUERY_ERROR")
	TransactionsError  = ErrorStatus("TransactionsError")
	GenerateTokenError = ErrorStatus("GenerateTokenError")
)

func AlreadyExistsErrorResponse(msg string, data helpers.JSON) *ResponseException {
	exception := ResponseException{
		Code:          http.StatusOK,
		Message:       AlreadyExist,
		ResultCode:    2003,
		ResultMessage: msg,
		Data:          data,
	}
	return &exception
}

func NotExistsErrorResponse(msg string, data helpers.JSON) *ResponseException {
	exception := ResponseException{
		Code:          http.StatusOK,
		Message:       NotExist,
		ResultCode:    2002,
		ResultMessage: msg,
		Data:          data,
	}
	return &exception
}

func DBQueryErrorResponse(msg string, data helpers.JSON) *ResponseException {
	exception := ResponseException{
		Code:          http.StatusInternalServerError,
		Message:       DBQueryError,
		ResultCode:    5000,
		ResultMessage: msg,
		Data:          data,
	}
	return &exception
}

func TransactionsErrorResponse(msg string, data helpers.JSON) *ResponseException {
	exception := ResponseException{
		Code:          http.StatusInternalServerError,
		Message:       TransactionsError,
		ResultCode:    5000,
		ResultMessage: msg,
		Data:          data,
	}
	return &exception
}

func UnAuthorizedErrorResponse(msg string, data helpers.JSON) *ResponseException {
	exception := ResponseException{
		Code:          http.StatusUnauthorized,
		Message:       UnAuthorized,
		ResultCode:    -1,
		ResultMessage: msg,
		Data:          data,
	}
	return &exception
}

func BadRequestErrorResponse(msg string, data helpers.JSON) *ResponseException {
	exception := ResponseException{
		Code:          http.StatusBadRequest,
		Message:       BadRequest,
		ResultCode:    -1,
		ResultMessage: msg,
		Data:          data,
	}
	return &exception
}

func NotFoundErrorResponse(msg string, data helpers.JSON) *ResponseException {
	exception := ResponseException{
		Code:          http.StatusNotFound,
		Message:       NotFound,
		ResultCode:    -1,
		ResultMessage: msg,
	}
	return &exception
}

func ForbiddenErrorResponse(msg string, data helpers.JSON) *ResponseException {
	exception := ResponseException{
		Code:          http.StatusForbidden,
		Message:       Forbidden,
		ResultCode:    -1,
		ResultMessage: msg,
		Data:          data,
	}
	return &exception
}

func InteralServerErrorResponse(msg string, data helpers.JSON) *ResponseException {
	exception := ResponseException{
		Code:          http.StatusInternalServerError,
		Message:       InteralServer,
		ResultCode:    -1,
		ResultMessage: msg,
		Data:          data,
	}
	return &exception
}
