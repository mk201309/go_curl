package router

import (
	"coreui/app/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

// LoadRouter 路由控制
func LoadRouter(r *gin.Engine) {

	r.GET("/ping", func(c *gin.Context) {
		fmt.Printf("%+v\n", "12345")
		c.JSON(200, gin.H{
			"message": "pong123",
		})
	})

	api := r.Group("/api")
	{
		api.POST("/get", service.Get)
		api.POST("/post", service.Post)
		api.POST("/put", service.Put)
		api.POST("/delete", service.Delete)
	}

}
