package phonenumber

import "strconv"

func IsValid(phoneNumber string) bool {
	// TODO - we can use regex to support +98 pattern
	if len(phoneNumber) != 11 && phoneNumber[0:2] != "09" {
		return false
	}

	if _, err := strconv.Atoi(phoneNumber[2:]); err != nil {
		return false
	}

	return true
}
