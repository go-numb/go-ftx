package rest

import (
	"fmt"
)

// APIError return the api error
type APIError struct {
	Status  int
	Message string
}

// Error return the error message
func (e APIError) Error() string {
	return fmt.Sprintf("APIError: status=%d, message=%s", e.Status, e.Message)
}
