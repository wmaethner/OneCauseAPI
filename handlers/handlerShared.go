package handlers

import (
	"encoding/json"
	"io"

	"github.com/wmaethner/OneCause/API/internal"
	"github.com/wmaethner/OneCause/API/validation"
)

// ValidateLoginPost shared code to validate the login attempt
func ValidateLoginPost(body io.ReadCloser) models.Response {
	var creds models.Credentials
	_ = json.NewDecoder(body).Decode(&creds)

	result, err := validation.ValidateLogin(creds.Hash, creds.Token)
	resp := models.Response{
		Result:       result,
		ErrorMessage: err,
		RedirectURL:  redirectURL(result),
	}

	return resp
}

func redirectURL(loginResult bool) string {
	if loginResult {
		return "http://onecause.com"
	}
	return ""
}
