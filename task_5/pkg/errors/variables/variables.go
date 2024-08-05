package variables

type ErrorCode string

const (
	InternalError   ErrorCode = "RAT-500-001"
	BadRequestError ErrorCode = "RAT-400-001"
	NotFoundError   ErrorCode = "RAT-400-001"
	BadGatewayError ErrorCode = "RAT-500-001"
	ForbiddenError  ErrorCode = "RAT-400-001"
)
