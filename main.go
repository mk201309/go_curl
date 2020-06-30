package main

import (
	"coreui/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	//  Enable line numbers in logging
	// log.SetFlags(log.LstdFlags | log.Lshortfile)

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[❌ Fatal❌ ]:", err)
		}
	}()

	r := gin.Default()

	// 開發時，console視窗不顯示有顏色的log
	// gin.DisableConsoleColor()

	// 載入router設定
	router.RouteProvider(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
