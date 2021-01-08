package validation

import (
	"strconv"
)

var credentialsAndLogins = map[string][]string{"EOCtK6aNq4iF67IjxyS3LIB3ymQb0/iP+T/ptOQaQX8=": []string{}}

// ValidateLogin validates the login hash and one time token
// 	returns true if a valid login otherwise returns false and an error message
func ValidateLogin(hash string, token string) (bool, string) {
	if !validLoginCredentials(hash) {
		return false, "Invalid username or password"
	}

	if !validToken(token) {
		return false, "Error validating login"
	}

	if tokenExists(hash, token) {
		return false, "Too many login attempts, try again later"
	}

	tokens, _ := credentialsAndLogins[hash]
	tokens = append(tokens, token)
	credentialsAndLogins[hash] = tokens

	return true, ""
}

func validLoginCredentials(hash string) bool {
	_, valid := credentialsAndLogins[hash]
	return valid
}

func validToken(token string) bool {
	// confirm token length
	if len(token) != 4 {
		return false
	}

	// confirm token is a number
	_, err := strconv.Atoi(token)
	if err != nil {
		return false
	}

	return true
}

func tokenExists(hash string, token string) bool {
	tokens, valid := credentialsAndLogins[hash]

	if !valid {
		return false
	}

	if contains(tokens, token) {
		return true
	}

	return false
}

func contains(tokenSlice []string, token string) bool {
	for _, val := range tokenSlice {
		if val == token {
			return true
		}
	}
	return false
}
