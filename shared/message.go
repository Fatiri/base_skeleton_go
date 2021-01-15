package shared

import (
	"encoding/json"
	"net/http"
	"strings"
)

const (
	// HTTPErrorBadRequest bad request code
	HTTPErrorBadRequest = http.StatusBadRequest
	// HTTPErrorNotFound not found code
	HTTPErrorNotFound = http.StatusNotFound
	// HTTPErrorUnauthorized unauthorized code
	HTTPErrorUnauthorized = http.StatusUnauthorized
	// HTTPErrorForbidden forbidden code
	HTTPErrorForbidden = http.StatusForbidden
	// HTTPErrorMethodNotAllowed method not allowed code
	HTTPErrorMethodNotAllowed = http.StatusMethodNotAllowed
	// HTTPErrorInternalServer internal error code
	HTTPErrorInternalServer = http.StatusInternalServerError
	// HTTPErrorUnprocessableEntity unprocessable entity code
	HTTPErrorUnprocessableEntity = http.StatusUnprocessableEntity
	// HTTPErrorTimeOut response timeout
	HTTPErrorTimeOut = http.StatusRequestTimeout // use this error if net/http request canceled/response timeout
	// HTTPErrorDataNotFound Application/data specific
	HTTPErrorDataNotFound = 100 // use this error if requested data not found

	// SQL error
	SqlErrorViolatesUniqueConstraint = "SQLSTATE 23505"

	// FIREBASE error
	FirebaseErrorEMAILEXISTS    = "EMAIL_EXISTS"
	FirebaseErrorUSERNOTFOUND   = "USER_NOT_FOUND"
	FirebaseErrorCANNOTFINDUSER = "CANNOT FIND USER"
)

var messageMap = map[int]map[string]string{
	http.StatusBadRequest: {
		"en": "Bad request",
		"id": "Permintaan buruk",
	},
	http.StatusNotFound: {
		"en": "Route not found",
		"id": "Rute tidak ditemukan",
	},
	http.StatusUnauthorized: {
		"en": "Unauthorized",
		"id": "Tidak sah",
	},
	http.StatusForbidden: {
		"en": "Forbidden",
		"id": "Tidak diizinkan",
	},
	http.StatusMethodNotAllowed: {
		"en": "Method not allowed",
		"id": "Metode tidak diperbolehkan",
	},
	http.StatusInternalServerError: {
		"en": "Internal server error",
		"id": "Terjadi kesalahan di server",
	},
	http.StatusUnprocessableEntity: {
		"en": "Invalid inputs",
		"id": "Kesalahan input",
	},
	HTTPErrorDataNotFound: {
		"en": "Data not found",
		"id": "Data tidak ditemukan",
	},
	http.StatusRequestTimeout: {
		"en": "Request time out",
		"id": "Permintaan waktu habis",
	},
}

var messageFirebaseMap = map[string]map[string]string{
	FirebaseErrorEMAILEXISTS: {
		"en": "FIREBASE : The email already registered, try to sign in",
		"id": "FIREBASE : Email sudah terdaftar, coba untuk sign in",
	},
	FirebaseErrorUSERNOTFOUND: {
		"en": "FIREBASE : User not found",
		"id": "FIREBASE : User tidak ditemukan",
	},
	FirebaseErrorCANNOTFINDUSER: {
		"en": "FIREBASE : Cannot find user",
		"id": "FIREBASE : Tidak dapat menemukan user",
	},
}

// StringFirebaseMap get message text from firebase error code and return string map contain multilang string.
// It returns empty if the code unknown.
func StringFirebaseMap(code string) map[string]string {
	return messageFirebaseMap[code]
}

// StringMap get message text from constant code and return string map contain multilang string.
// It returns empty if the code unknown.
func StringMap(code int) map[string]string {
	return messageMap[code]
}

type MultiLangError struct {
	ErrorID int               `json:"code"`
	Msg     map[string]string `json:"message"`
}

// MultiStringValidationError multi string error will contain multi lang error message.
type MultiStringValidationError MultiLangError

// NewMultiStringValidationError create new multi string error struct
func NewMultiStringValidationError(code int, msg map[string]string) *MultiStringValidationError {
	return &MultiStringValidationError{
		ErrorID: code,
		Msg:     msg,
	}
}

func (c *MultiStringValidationError) Code() int {
	return c.ErrorID
}

func (c *MultiStringValidationError) Message() map[string]string {
	return c.Msg
}

// Error comply with error interface
func (c *MultiStringValidationError) Error() string {
	// nolint
	b, _ := json.Marshal(c)

	return string(b)
}

// IsMultiStringValidationError check whether error is MultiStringValidationError pointer
func IsMultiStringValidationError(err error) bool {
	switch err.(type) {
	case *MultiStringValidationError:
		return true
	default:
		return false
	}
}

// MultiStringBadRequestError multi string bad request error will contain multi lang error message.
type MultiStringBadRequestError MultiLangError

// NewMultiStringBadRequestError create new multi string bad request error
func NewMultiStringBadRequestError(code int, msg map[string]string) *MultiStringBadRequestError {
	return &MultiStringBadRequestError{
		ErrorID: code,
		Msg:     msg,
	}
}

// DefaultMultiStringBadRequestError ...
func DefaultMultiStringBadRequestError() *MultiStringBadRequestError {
	return &MultiStringBadRequestError{
		ErrorID: http.StatusBadRequest,
		Msg:     StringMap(http.StatusBadRequest),
	}
}

func (c *MultiStringBadRequestError) Code() int {
	return c.ErrorID
}

func (c *MultiStringBadRequestError) Message() map[string]string {
	return c.Msg
}

// Error comply with error interface
func (c *MultiStringBadRequestError) Error() string {
	// nolint
	b, _ := json.Marshal(c)

	return string(b)
}

// IsMultiStringBadRequestError check if it was bad request error
func IsMultiStringBadRequestError(err error) bool {
	switch err.(type) {
	case *MultiStringBadRequestError:
		return true
	default:
		return false
	}
}

// MultiStringUnauthorizedError multi string unauthorized request error will contain multi lang error message.
type MultiStringUnauthorizedError MultiLangError

// NewMultiStringUnauthorizedError create new multi string bad request error
func NewMultiStringUnauthorizedError(code int, msg map[string]string) *MultiStringUnauthorizedError {
	return &MultiStringUnauthorizedError{
		ErrorID: code,
		Msg:     msg,
	}
}

func DefaultMultiStringUnauthorizedError() *MultiStringUnauthorizedError {
	return &MultiStringUnauthorizedError{
		ErrorID: http.StatusUnauthorized,
		Msg:     StringMap(HTTPErrorUnauthorized),
	}
}

func (c *MultiStringUnauthorizedError) Code() int {
	return c.ErrorID
}

func (c *MultiStringUnauthorizedError) Message() map[string]string {
	return c.Msg
}

// Error comply with error interface
func (c *MultiStringUnauthorizedError) Error() string {
	// nolint
	b, _ := json.Marshal(c)

	return string(b)
}

// IsMultiStringUnauthorizedError check if it was bad request error
func IsMultiStringUnauthorizedError(err error) bool {
	switch err.(type) {
	case *MultiStringUnauthorizedError:
		return true
	default:
		return false
	}
}

// MultiStringForbiddenError multi string forbidden request error will contain multi lang error message.
type MultiStringForbiddenError MultiLangError

// NewMultiStringForbiddenError create new multi string forbidden request error
func NewMultiStringForbiddenError(code int, msg map[string]string) *MultiStringForbiddenError {
	return &MultiStringForbiddenError{
		ErrorID: code,
		Msg:     msg,
	}
}

func DefaultMultiStringForbiddenError() *MultiStringForbiddenError {
	return &MultiStringForbiddenError{
		ErrorID: http.StatusForbidden,
		Msg:     StringMap(http.StatusForbidden),
	}
}

func (c *MultiStringForbiddenError) Code() int {
	return c.ErrorID
}

func (c *MultiStringForbiddenError) Message() map[string]string {
	return c.Msg
}

// Error comply with error interface
func (c *MultiStringForbiddenError) Error() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// IsMultiStringForbiddenError check if it was forbidden request error
func IsMultiStringForbiddenError(err error) bool {
	switch err.(type) {
	case *MultiStringForbiddenError:
		return true
	default:
		return false
	}
}

type MultiStringInternalServerError MultiLangError

func DefaultMultiStringInternalServerError() *MultiStringInternalServerError {
	return &MultiStringInternalServerError{
		ErrorID: http.StatusInternalServerError,
		Msg:     StringMap(http.StatusInternalServerError),
	}
}

func (c *MultiStringInternalServerError) Code() int {
	return c.ErrorID
}

func (c *MultiStringInternalServerError) Message() map[string]string {
	return c.Msg
}

// Error comply with error interface
func (c *MultiStringInternalServerError) Error() string {
	b, _ := json.Marshal(c)
	return string(b)
}

type MultiStringRouteNotFoundError MultiLangError

func DefaultMultiStringRouteNotFoundError() *MultiStringRouteNotFoundError {
	return &MultiStringRouteNotFoundError{
		ErrorID: http.StatusNotFound,
		Msg:     StringMap(http.StatusNotFound),
	}
}

func (c *MultiStringRouteNotFoundError) Code() int {
	return c.ErrorID
}

func (c *MultiStringRouteNotFoundError) Message() map[string]string {
	return c.Msg
}

type MultiStringHTTPErrorDataNotFound MultiLangError

func DefaultMultiStringHTTPErrorDataNotFound() *MultiStringHTTPErrorDataNotFound {
	return &MultiStringHTTPErrorDataNotFound{
		ErrorID: HTTPErrorDataNotFound,
		Msg:     StringMap(HTTPErrorDataNotFound),
	}
}

// IsMultiStringHTTPErrorDataNotFound check if it was forbidden request error
func IsMultiStringHTTPErrorDataNotFound(err error) bool {
	switch err.(type) {
	case *MultiStringHTTPErrorDataNotFound:
		return true
	default:
		return false
	}
}

func (c *MultiStringHTTPErrorDataNotFound) Code() int {
	return c.ErrorID
}

func (c *MultiStringHTTPErrorDataNotFound) Message() map[string]string {
	return c.Msg
}

// Error comply with error interface
func (c *MultiStringHTTPErrorDataNotFound) Error() string {
	b, _ := json.Marshal(c)
	return string(b)
}

type MultiStringMethodNotAllowedError MultiLangError

func DefaultMultiStringMethodNotAllowedError() *MultiStringMethodNotAllowedError {
	return &MultiStringMethodNotAllowedError{
		ErrorID: http.StatusMethodNotAllowed,
		Msg:     StringMap(http.StatusMethodNotAllowed),
	}
}

func (c *MultiStringMethodNotAllowedError) Code() int {
	return c.ErrorID
}

func (c *MultiStringMethodNotAllowedError) Message() map[string]string {
	return c.Msg
}

type MultiStringHTTPErrorTimeOut MultiLangError

func DefaultMultiStringHTTPErrorTimeOut() *MultiStringHTTPErrorTimeOut {
	return &MultiStringHTTPErrorTimeOut{
		ErrorID: http.StatusRequestTimeout,
		Msg:     StringMap(http.StatusRequestTimeout),
	}
}

// IsMultiStringHTTPErrorTimeOut check if it was timeout request error
func IsMultiStringHTTPErrorTimeOut(err error) bool {
	if strings.ContainsAny(err.Error(), "timeout") {
		return true
	}
	return false
}

func (c *MultiStringHTTPErrorTimeOut) Code() int {
	return c.ErrorID
}

func (c *MultiStringHTTPErrorTimeOut) Message() map[string]string {
	return c.Msg
}

// Error comply with error interface
func (c *MultiStringHTTPErrorTimeOut) Error() string {
	b, _ := json.Marshal(c)
	return string(b)
}

// NewMultiStringFirebaseRequestError create new firebase multi string bad request error
func NewMultiStringFirebaseRequestError(code int, msg string) *MultiStringBadRequestError {
	return &MultiStringBadRequestError{
		ErrorID: code,
		Msg:     StringFirebaseMap(msg),
	}
}
