package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	router.Run("0.0.0.0:8080")
}
