package internal

import "github.com/Genez-io/pulumi-genezio/provider/requests"

type ErrorCode int64

const (
	UnknownError                   ErrorCode = 0
	Unauthorized                   ErrorCode = 1
	NotFoundError                  ErrorCode = 2
	InternalServerError            ErrorCode = 3
	MethodNotAllowedError          ErrorCode = 4
	MissingRequiredParametersError ErrorCode = 5
	StatusPaymentRequired          ErrorCode = 402
	BadRequest                     ErrorCode = 6
	StatusConflict                 ErrorCode = 7
	UpdateRequired                 ErrorCode = 8
	ForbiddenSubdomain             ErrorCode = 9
)

type ErrorPayload struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

type ErrorResponse struct {
	Status requests.ResponseStatus `json:"status"`
	Error  ErrorPayload            `json:"error"`
}
