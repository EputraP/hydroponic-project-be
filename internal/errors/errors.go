package errs

import "errors"

var (
	InvalidBearerFormat    = errors.New("Invalid Authorization Bearer Format")
	InvalidToken           = errors.New("Invalid Token")
	InvalidIssuer          = errors.New("Invalid Token Issuer")
	InvalidParam           = errors.New("Invalid Parameter")
	InvalidUUIDParamFormat = errors.New("Invalid UUID Parameter Format")

	ForbiddenAccess = errors.New("user is forbidden to access this resource")

	InvalidRequestBody = errors.New("invalid request body")

	CantFindAnySubModules = errors.New("can't find any sub modules")
)
