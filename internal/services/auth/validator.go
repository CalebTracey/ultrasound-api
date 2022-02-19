package auth

import "fmt"

func Validate(username string) error {
	if username == "" {
		return fmt.Errorf("error: missing username")
	}
	return nil
}
