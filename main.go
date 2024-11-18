package main

import "github.com/gin-gonic/gin"

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func servidor1(c *gin.Context) {
	c.JSON(200, gin.H{
		"siape": "123456",
		"nome":  "Gasparzinho",
	})
}
func servidor2(c *gin.Context) {
	c.JSON(200, gin.H{
		"siape": "234567",
		"nome":  "Assombroso",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/servidor/1", servidor1)
	r.GET("/servidor/2", servidor2)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
