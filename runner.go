package main

import (
	"fmt"
	"github.com/braintree/manners"
	"github.com/deevatech/runner/languages/ruby"
	. "github.com/deevatech/runner/types"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

var Languages = map[string]FuncRunCode{
	"ruby": ruby.Run,
}

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
	//defer manners.Close()

	var run RunParams

	if errParams := c.BindJSON(&run); errParams == nil {
		if result, errRun := RunCode(run); errRun == nil {
			c.JSON(http.StatusOK, result)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errRun,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errParams,
		})
	}
}

func RunCode(p RunParams) (*RunResults, error) {
	if langRunner, ok := Languages[p.Language]; ok {
		result := langRunner(p)
		return &result, nil
	} else {
		return nil, fmt.Errorf("Unknown language: %s", p.Language)
	}
}
