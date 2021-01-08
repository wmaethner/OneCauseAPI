package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wmaethner/OneCause/API/handlers"
)

func main() {
	setupHTTPHandlerAndRun()
	//setupGinHandlerAndRun()
}

func setupHTTPHandlerAndRun() {
	http.HandleFunc("/login", handlers.HTTPLoginHandler)
	fmt.Println("Running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupGinHandlerAndRun() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.POST("/login", handlers.GinLoginHandler)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

// CORSMiddleware handles cross origin resource sharing
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
