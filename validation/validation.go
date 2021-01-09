// Currently the token array would continue to grow for as long as the server
// was on. Purging them every once in a while would probably be needed or
// a different method for validating the token.

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

// validLoginCredentials confirms the login credentials are valid
func validLoginCredentials(hash string) bool {
	_, valid := credentialsAndLogins[hash]
	return valid
}

// validToken confirms the token is the correct format
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

// tokenExists checks if the token was already used for the login hash
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

// contains checks if the slice contains the given element
func contains(tokenSlice []string, token string) bool {
	for _, val := range tokenSlice {
		if val == token {
			return true
		}
	}
	return false
}
