package app

import (
	"github.com/OhMinsSup/story-server/helpers"
	"net/http"
)

type ErrorStatus string

type ErrorException struct {
	Code          int          `json:"code"`
	Message       ErrorStatus  `json:"message"`
	ResultCode    int          `json:"result_code"`
	ResultMessage error        `json:"result_message"`
	Data          helpers.JSON `json:"data"`
}

const (
	AlreadyExist  = ErrorStatus("ALREADY_EXIST")
	BadRequest    = ErrorStatus("BAD_REQUEST")
	NotFound      = ErrorStatus("NOT_FOUND")
	Forbidden     = ErrorStatus("FORBIDDEN")
	InteralServer = ErrorStatus("INTERNAL_SERVER_ERROR")
	UnAuthorized  = ErrorStatus("UNAUTHORIZED")
	NotExist      = ErrorStatus("NOT_EXIST")
)

func AlreadyExistsErrorResponse(msg error, data helpers.JSON) *ErrorException {
	exception := ErrorException{
		Code:          http.StatusOK,
		Message:       AlreadyExist,
		ResultCode:    2003,
		ResultMessage: msg,
		Data:          data,
	}
	return &exception
}

func NotExistsErrorResponse(msg error, data helpers.JSON) *ErrorException {
	exception := ErrorException{
		Code:          http.StatusOK,
		Message:       NotExist,
		ResultCode:    2002,
		ResultMessage: msg,
		Data:          data,
	}
	return &exception
}

func UnAuthorizedErrorResponse(msg error, data helpers.JSON) *ErrorException {
	exception := ErrorException{
		Code:          http.StatusUnauthorized,
		Message:       UnAuthorized,
		ResultCode:    -1,
		ResultMessage: msg,
		Data:          data,
	}
	return &exception
}

func BadRequestErrorResponse(msg error, data helpers.JSON) *ErrorException {
	exception := ErrorException{
		Code:          http.StatusBadRequest,
		Message:       BadRequest,
		ResultCode:    -1,
		ResultMessage: msg,
		Data:          data,
	}
	return &exception
}

func NotFoundErrorResponse(msg error, data helpers.JSON) *ErrorException {
	exception := ErrorException{
		Code:          http.StatusNotFound,
		Message:       NotFound,
		ResultCode:    -1,
		ResultMessage: msg,
	}
	return &exception
}

func ForbiddenErrorResponse(msg error, data helpers.JSON) *ErrorException {
	exception := ErrorException{
		Code:          http.StatusForbidden,
		Message:       Forbidden,
		ResultCode:    -1,
		ResultMessage: msg,
		Data:          data,
	}
	return &exception
}

func InteralServerErrorResponse(msg error, data helpers.JSON) *ErrorException {
	exception := ErrorException{
		Code:          http.StatusInternalServerError,
		Message:       InteralServer,
		ResultCode:    -1,
		ResultMessage: msg,
		Data:          data,
	}
	return &exception
}
