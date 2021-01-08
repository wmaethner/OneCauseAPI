package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GinLoginHandler sets up the login handler using the gin library
func GinLoginHandler(c *gin.Context) {
	resp := ValidateLoginPost(c.Request.Body)

	c.JSON(http.StatusOK, resp)
}
