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

	CodeSuccess            = 20001 // Success
	ErrorCodeParamsInvalid = 20003 // Email is invalid
	ErrorInvalidToken      = 30001 // Invalid token
	ErrorInvalidOTP        = 30002 // Invalid OTP
	ErrorSendEmailOTP      = 30003 // Send email OTP failed

	// Register Code
	ErrorCodeUserHasExisted = 50001 // User has existed
	ErrorCodeOtpNotExists   = 50002 // OTP exists but not registered
)

var msg = map[int]string{
	CodeSuccess:            "Success",
	ErrorCodeParamsInvalid: "Code is invalid",
	ErrorInvalidToken:      "Invalid token",
	ErrorInvalidOTP:        "OTP Error",
	ErrorSendEmailOTP:      "Send email OTP failed",

	ErrorCodeUserHasExisted: "User has existed",
	ErrorCodeOtpNotExists:   "OTP exists but not registered",
}
