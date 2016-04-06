package main

import (
	"fmt"
	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func init() {
	log.Println("Deeva Runner!")
}

func main() {
	router := gin.Default()
	router.POST("/run", handleRunRequest)

	port := os.Getenv("DEEVA_RUNNER_PORT")
	if len(port) == 0 {
		port = "8080"
	}

	log.Printf("Starting in %v mode on port %v\n", gin.Mode(), port)
	host := fmt.Sprintf(":%v", port)
	manners.ListenAndServe(host, router)
}

func handleRunRequest(c *gin.Context) {
	defer manners.Close()
	c.JSON(http.StatusOK, gin.H{})
}
