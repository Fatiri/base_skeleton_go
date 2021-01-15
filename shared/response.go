package shared

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	// Error struct
	Error struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	// Response struct
	Response struct {
		Success bool        `json:"success"`
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Meta    interface{} `json:"meta,omitempty"`
		Data    interface{} `json:"data,omitempty"`
		Errors  interface{} `json:"errors,omitempty"`
	}

	// Meta struct
	Meta struct {
		Page         int `json:"page,omitempty"`
		Limit        int `json:"limit,omitempty"`
		TotalRecords int `json:"totalRecords,omitempty"`
		TotalPages   int `json:"totalPages,omitempty"`
	}

	// Pagination ...
	Pagination struct {
		Data interface{}
		Meta Meta
	}
)

// JSONResponse func response
func JSONResponse(code int, message string, status bool, param interface{}) *Response {
	response := new(Response)

	switch param.(type) {
	case Meta:
		response.Meta = param
	case []*Error, Error:
		response.Errors = param
	case *Pagination:
		paginate := param.(*Pagination)
		response.Meta = paginate.Meta
		response.Data = paginate.Data
	default:
		response.Data = param
	}

	response.Success = status
	response.Code = code
	response.Message = message

	return response
}

// JSONResponse func response
func JSONMetaResponse(code int, message string, status bool, param interface{}, meta interface{}) *Response {
	response := new(Response)

	response.Data = param
	response.Meta = meta
	response.Success = status
	response.Code = code
	response.Message = message

	return response
}

// JSONSuccess ...
func JSONSuccess(message string, params interface{}) *Response {
	response := JSONResponse(http.StatusOK,
		message,
		true,
		params)
	return response
}

// JSONError ...
func JSONError(code int, message string, params ...interface{}) *Response {
	response := JSONResponse(code,
		message,
		false,
		params)
	return response
}

// ErrorBadRequest ...
func ErrorBadRequest(message string, params ...interface{}) *Response {
	response := JSONResponse(http.StatusBadRequest,
		message,
		false,
		params)
	return response
}

// ErrorDataNotFound ...
func ErrorDataNotFound(message string, params ...interface{}) *Response {
	response := JSONResponse(http.StatusNotFound,
		message,
		false,
		params)
	return response
}

func HttpError(e echo.Context, err error) error {
	if IsMultiStringValidationError(err) {
		return e.JSON(http.StatusUnprocessableEntity, err.(*MultiStringValidationError))
	}

	if IsMultiStringBadRequestError(err) {
		return e.JSON(http.StatusBadRequest, err.(*MultiStringBadRequestError))
	}

	if IsMultiStringUnauthorizedError(err) {
		return e.JSON(http.StatusUnauthorized, DefaultMultiStringUnauthorizedError())
	}

	return e.JSON(http.StatusBadRequest, DefaultMultiStringInternalServerError())
}
