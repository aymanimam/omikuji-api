package errors

import "fmt"

const (
	// OmikujiErrorCode Error code originated from Omikuji model
	OmikujiErrorCode = "100"
	// OmikujiServiceErrorCode Error code originated from OmikujiService
	OmikujiServiceErrorCode = "101"
	// OmikujiServerErrorCode Error code originated from OmikujiServer
	OmikujiServerErrorCode = "102"
	// OmikujiRecoveryErrorCode Error code originated from error handling middleware
	OmikujiRecoveryErrorCode = "103"
	// OmikujiRecoveryUnknownErrorCode Error code originated from unknown errors in the error handling middleware
	OmikujiRecoveryUnknownErrorCode = "104"
)

// OmikujiException it's thrown when an error happens in Omikuji, OmikujiService and OmikujiServer
type OmikujiException struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

// OmikujiException implements Error interface
func (e *OmikujiException) Error() string {
	return fmt.Sprintf("[OmikujiException] message: [%v], code: [%v].", e.Message, e.Code)
}

// NewOmikujiException Create new OmikujiException
func NewOmikujiException(msg, code string) *OmikujiException {
	return &OmikujiException{
		msg,
		code,
	}
}

// ThrowOmikujiException Throw an OmikujiException
func ThrowOmikujiException(msg, code string) {
	e := NewOmikujiException(msg, code)
	panic(e)
}
