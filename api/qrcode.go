package api

import (
	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
	"net/http"
)

type QRCodeQuery struct {
	DeviceId   string `form:"device_id"`
	DeviceName string `form:"device_name"`
	Platform   string `form:"platform"`
}

func Create(c *gin.Context) {
	bind := QRCodeQuery{}
	err := c.ShouldBindQuery(&bind)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if pic, err := qrcode.Encode(bind.DeviceId, qrcode.Medium, 256); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.Data(http.StatusOK, "image/png", pic)
	}
}
