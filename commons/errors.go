package commons

import "errors"

var (
	ErrorInternalServer      = errors.New("internal server error")
	ErrorProductAlreadyExist = errors.New("product already exist")
)
