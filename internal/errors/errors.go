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

	ErrorOnCreatingNewAsset = errors.New("error on creating new asset")
	InvalidPlantId          = errors.New("invalid plant id")
	InvalidUoMId            = errors.New("invalid uom id")
	InvalidAssetTypeId      = errors.New("invalid asset type id")

	ErrorOnCreatingNewPlantGrowthRecord = errors.New("error on creating new plant growth record")
	InvalidAssetId                      = errors.New("invalid asset id")
	InvalidProcessId                    = errors.New("invalid process id")

	InvalidTowerId = errors.New("invalid tower id")
)
