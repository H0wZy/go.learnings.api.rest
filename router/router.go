package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()
	initializeRoutes(router)
	_ = router.Run(":8080")
}
