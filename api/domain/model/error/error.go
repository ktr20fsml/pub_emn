package error

import "fmt"

type InvalidValueError struct {
	Field   string
	Message string
}

func (ive *InvalidValueError) Error() string {
	return fmt.Sprintf("%s is invalid. %s", ive.Field, ive.Message)
}

type DatabaseError struct {
	Method string
	Detail error
}

func (dbe *DatabaseError) Error() string {
	return fmt.Sprintf("FAILED DATABASE OPERATION: METHOD: %s, DETAIL: %s", dbe.Method, dbe.Detail)
}
