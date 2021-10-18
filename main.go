package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.Int("p", 3000, "listening port")
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	app.Any("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})
	if err := app.Run(fmt.Sprintf(":%d", *port)); err != nil {
		panic(err)
	}
}
