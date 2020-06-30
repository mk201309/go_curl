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

	r.POST("/get", service.Get)
	r.POST("/post", service.Post)
	r.POST("/put", service.Put)
	r.POST("/delete", service.Delete)

}
