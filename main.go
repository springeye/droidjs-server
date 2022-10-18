package main

import (
	"github.com/gin-gonic/gin"
	"github.com/springeye/droidjs-server/api"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	public := r.Group("/api/v1/public")
	{
		ws := public.Group("/ws")
		{
			hub := api.NewHub()
			go hub.Run()
			ws.GET("", func(context *gin.Context) {

				api.ServeWs(hub, context.Writer, context.Request)

			})
		}
	}
	public.GET("/qrcode", api.Create)
	auth := api.SetupJwt(r)
	{
		device := auth.Group("/device")
		{
			device.GET("list", func(context *gin.Context) {

			})
		}
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
