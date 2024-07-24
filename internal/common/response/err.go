package response

import (
	"net/http"
)

type ResponseErr struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

var (
	ErrParsingJSON = ResponseErr{
		Message: "failed to parse json",
		Code:    http.StatusBadRequest,
	}

	ErrParsingQueryString = ResponseErr{
		Message: "failed to parse query string",
		Code:    http.StatusBadRequest,
	}

	ErrParsingURLParam = ResponseErr{
		Message: "failed to parse URL param",
		Code:    http.StatusBadRequest,
	}

	ErrGetSessionID = ResponseErr{
		Message: "failed to get session id from token",
		Code:    http.StatusBadRequest,
	}

	ErrInternal = ResponseErr{
		Message: "internal server error",
		Code:    http.StatusInternalServerError,
	}

	ErrTokenNotExist = ResponseErr{
		Message: "error auth token not exist",
		Code:    http.StatusBadRequest,
	}

	ErrInvalidToken = ResponseErr{
		Message: "error invalid token",
		Code:    http.StatusBadRequest,
	}

	ErrInvalidSession = ResponseErr{
		Message: "error invalid session",
		Code:    http.StatusBadRequest,
	}
)
