package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	// new a default router
	router := gin.Default()

	// use http bind router function
	router.GET("/test", func(c *gin.Context) {
		// return a json
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello! My name is GoHello!",
		})
	})

	return router
}
