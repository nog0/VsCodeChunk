	package main

	import (
		"github.com/gin-gonic/gin"
	)

	func main() {
	r:=gin.Default()
	r.POST("/login",func(c *gin.Context){
		Username:=c.PostForm("Username")
		Password:=c.PostForm("Password")
		if Password=="123"&&Username=="321"{
			c.SetCookie("gin_Cookie","value",3600,"/","",false,true)
			c.JSON(200,gin.H{"message":"账户或密码正确"})
		}else{
			c.JSON(403,gin.H{"error":"账户或密码错误"})
		}
	})
	auth:=func(c *gin.Context){
		value,err:=c.Cookie("gin_Cookie")
		if err!=nil{
			c.JSON(403,gin.H{"message":"获取Cookie失败"})
			c.Abort()
		}else{
			c.Set("Cookie",value)
		}
	}
	r.GET("/hello",auth,func(c *gin.Context){
		cookie,_:=c.Get("Cookie")
		str:=cookie.(string)
		c.JSON(200,gin.H{"message":"hello"+" "+str})
		})
	r.Run()
	}