package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
  "fmt"
  "os"
)

// env var helper function to default if not set
func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return defaultValue
    }
    return value
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!\n")
	})

  r.GET("/hello/", func(c *gin.Context) {
    c.String(http.StatusOK, "Hello world!\n")
  })

  r.GET("/hello/:name", func(c *gin.Context) {
    name := c.Params.ByName("name")
    c.String(http.StatusOK, "Why Hello " + name + "!\n")
  })

	return r
}

func main() {
  port := getEnv("PORT", "8080")
	r := setupRouter()

	r.Run(fmt.Sprintf(":%s", port))
}
