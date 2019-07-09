package common

var (
	SuccessResult = &Result{
		Success: true,
		Error:   "",
		Code:    0,
	}
)

type ErrorCode interface {
	Code() int32
}

func CreateErrorResult(err interface{}) *Result {
	result := &Result{
		Success: false,
		Error:   "",
		Code:    1,
	}

	if t, ok := err.(error); ok {
		result.Error = t.Error()
	}

	if t, ok := err.(string); ok {
		result.Error = t
	}

	if t, ok := err.(ErrorCode); ok {
		result.Code = t.Code()
	}

	return result
}

var EmptyMessages = make([]*Message, 0, 0)
