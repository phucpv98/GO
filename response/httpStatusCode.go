package response

const (
	// Success
	StatusOK = 200

	// Client error
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusMethodNotAllowed    = 405
	StatusConflict            = 409
	StatusUnprocessableEntity = 422

	// Server error
	StatusInternalServerError = 500

	// Error code
	ErrorCodeInvalidRequest = 1001
	ErrorCodeUnauthorized   = 1002
	ErrorCodeForbidden      = 1003
	ErrorCodeNotFound       = 1004
	ErrorCodeConflict       = 1005
	ErrorCodeServerError    = 1006

	ErrorCodeSuccess       = 20001 // Success
	ErrorCodeParamsInvalid = 20003 // Email is invalid
	ErrorInvalidToken      = 30001 // Invalid token

	// Register Code
	ErrorCodeUserHasExisted = 50001 // User has existed
)

var msg = map[int]string{
	ErrorCodeSuccess:       "Success",
	ErrorCodeParamsInvalid: "Email is invalid",
	ErrorInvalidToken:      "Invalid token",

	ErrorCodeUserHasExisted: "User has existed",
}
