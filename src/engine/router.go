package engine

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "./public")

	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/admin/menu")
	})

	r.GET("/admin", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/admin/menu")
	})

	return r
}
