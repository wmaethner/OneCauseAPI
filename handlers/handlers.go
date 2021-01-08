package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wmaethner/OneCause/API/validation"
)

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

// HTTPLoginHandler sets up the login handler using the standard net/http library
func HTTPLoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var creds Credentials
	_ = json.NewDecoder(r.Body).Decode(&creds)

	fmt.Println(creds.Hash)
	fmt.Println(creds.Token)
	result, err := validation.ValidateLogin(creds.Hash, creds.Token)
	resp := Response{
		Result:       result,
		ErrorMessage: err,
	}

	json.NewEncoder(w).Encode(resp)
}

// GinLoginHandler sets up the login handler using the gin library
func GinLoginHandler(c *gin.Context) {
	fmt.Println("cloginhandler")
	var creds Credentials
	_ = json.NewDecoder(c.Request.Body).Decode(&creds)

	result, err := validation.ValidateLogin(creds.Hash, creds.Token)
	resp := Response{
		Result:       result,
		ErrorMessage: err,
	}

	c.JSON(http.StatusOK, resp)
}
