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

// OmikujiEerror it's thrown when an error happens in Omikuji, OmikujiService and OmikujiServer
type OmikujiEerror struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

// OmikujiEerror implements Error interface
func (e *OmikujiEerror) Error() string {
	return fmt.Sprintf("[OmikujiError] message: [%v], code: [%v].", e.Message, e.Code)
}

// NewOmikujiError Create new OmikujiError
func NewOmikujiError(msg, code string) error {
	return &OmikujiEerror{
		msg,
		code,
	}
}
