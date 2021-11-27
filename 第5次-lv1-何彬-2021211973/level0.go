package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r:= gin.Default()
	r.GET("/user/:name", func(c *gin.Context) {
		name:=c.Param("name")
		c.JSON(200,name)
	})
	r.Run()
}

