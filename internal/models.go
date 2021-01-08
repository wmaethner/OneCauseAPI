package models

// Credentials defines the credential structure
type Credentials struct {
	Token string
	Hash  string
}

// Response defines the login attempt response
type Response struct {
	Result       bool
	ErrorMessage string
}
