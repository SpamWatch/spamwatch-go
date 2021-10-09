package spamwatch

import (
	"encoding/json"
	"fmt"
	"time"
)

func handleError(sw *SpamWatch, s int, b []byte, method string) error {
	if s != 200 && s != 201 && s != 204 {
		var e = Error{}
		err := json.Unmarshal(b, &e)
		if err != nil {
			return err
		}
		return &ErrorHandler{Err: &e, Spamwatch: sw, Method: method}
	}
	return nil
}

func (e *ErrorHandler) Error() string {
	switch e.Err.Code {
	case 401:
		return fmt.Sprintf("%s: Make sure your API Token is correct", SPAMWATCH_ERROR_PREFIX)
	case 403:
		return fmt.Sprintf("%s: Your token's permission '%s' is not high enough", SPAMWATCH_ERROR_PREFIX, e.Spamwatch.Token.Permission)
	case 429:
		return fmt.Sprintf("%s: Too Many Requests for method '%s': Try again in %d seconds", SPAMWATCH_ERROR_PREFIX, e.Method, e.Err.Until-time.Now().Unix())
	default:
		return fmt.Sprintf("%s: %s: %d-%s", SPAMWATCH_ERROR_PREFIX, e.Method, e.Err.Code, e.Err.Description)
	}
}
