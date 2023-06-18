package Util

type CustomError struct {
	Message string `json:"error"`
}

func (e *CustomError) Error() string {
	return e.Message
}
