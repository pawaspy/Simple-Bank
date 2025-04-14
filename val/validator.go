package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidateUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidateFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateString(value string, min, max int) error {
	if len(value) < min && len(value) > max {
		return fmt.Errorf("the characters should be between %d-%d", min, max)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidateUsername(value) {
		return fmt.Errorf("characters should be lower case, number or underscore")
	}
	return nil
}

func ValidateFullName(value string) error {
	if err := ValidateString(value, 10, 100); err != nil {
		return err
	}
	if !isValidateFullName(value) {
		return fmt.Errorf("characters should be lower or upper case and spaces")
	}
	return nil
}

func ValidatePassword(value string) error {
	if err := ValidateString(value, 6, 100); err != nil {
		return err
	}
	return nil
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return err
	}
	return nil
}

func ValidateEmailID(value int64) error {
	if value <= 0 {
		return fmt.Errorf("must be positive integer")
	}

	return nil
}

func ValidateSecretCode(value string) error {
	return ValidateString(value, 32, 128)
}