package main

import "github.com/gin-gonic/gin"

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func servidor(c *gin.Context) {
	c.JSON(200, gin.H{
		"siape": "123456",
		"nome":  "gasparzinho",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/servidor", servidor)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
