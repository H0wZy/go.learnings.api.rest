package router

import "github.com/gin-gonic/gin"

func Init() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
