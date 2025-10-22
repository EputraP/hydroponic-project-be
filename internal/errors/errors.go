package errs

import "errors"

var (
	InvalidBearerFormat = errors.New("Invalid Authorization Bearer Format")
	InvalidToken        = errors.New("Invalid Token")
	InvalidIssuer       = errors.New("Invalid Token Issuer")

	ForbiddenAccess = errors.New("user is forbidden to access this resource")

	InvalidRequestBody = errors.New("invalid request body")
)
