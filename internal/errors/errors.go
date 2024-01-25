package errors

import "fmt"

const (
	KeyNotExistInMap   int = 1
	JwtTokenParseError int = 2
)

type KapibaraGeneralError struct {
	Code    int
	Message string
}

func (e *KapibaraGeneralError) Error() string {
	return fmt.Sprintf("[KapibaraGeneralError]%d: %s", e.Code, e.Message)
}
