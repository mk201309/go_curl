package router

import (
	"github.com/gin-gonic/gin"
)

// RouteProvider 路由提供者
func RouteProvider(r *gin.Engine) {
	// api route
	LoadRouter(r)
}
