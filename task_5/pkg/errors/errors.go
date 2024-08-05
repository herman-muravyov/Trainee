package errors

import (
	"encoding/json"
	"fmt"
	"task/pkg/errors/variables"
)

type AppError struct {
	Err              error  `json:"-"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             string `json:"code,omitempty"`
}

func NewAppError(err error, message, developerMessage, code string) *AppError {
	return &AppError{
		Err:              err,
		Message:          message,
		DeveloperMessage: developerMessage,
		Code:             code,
	}
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Error: %v, Message: %s, DeveloperMessage: %s, Code: %s", e.Err, e.Message, e.DeveloperMessage, e.Code)
}

func (e *AppError) Unwrap() error { return e.Err }

func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return marshal
}

func ErrInvalidURL(message, developerMessage string) error {
	return NewAppError(fmt.Errorf(message), message, developerMessage, string(variables.BadRequestError))
}

func ErrConnectionFailed(message, developerMessage string) error {
	return NewAppError(fmt.Errorf(message), message, developerMessage, string(variables.BadGatewayError))
}

func ErrDownloadFailed(message, developerMessage string) error {
	return NewAppError(fmt.Errorf(message), message, developerMessage, string(variables.InternalError))
}

func ErrFileNotFound(message, developerMessage string) error {
	return NewAppError(fmt.Errorf(message), message, developerMessage, string(variables.NotFoundError))
}

func ErrUnexpectedEOF(message, developerMessage string) error {
	return NewAppError(fmt.Errorf(message), message, developerMessage, string(variables.InternalError))
}
