package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	app := gin.Default()
	app.Any("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})
	app.Run(":3000")
}
