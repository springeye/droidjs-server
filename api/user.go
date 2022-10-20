package api

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/springeye/droidjs-server/db"
	"github.com/springeye/droidjs-server/proto"
	"gorm.io/gorm"
	"strings"
)

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func Register(c *gin.Context) {
	req := proto.RegisterRequest{}
	if err := c.Bind(&req); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"msg": err.Error()})
		return
	}
	tx := db.GetDB()
	var user db.User
	if err := tx.Where("username = ?", req.Username).Find(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
		c.AbortWithStatusJSON(400, gin.H{"msg": err.Error()})
		return
	}
	if user.Username != "" {
		c.AbortWithStatusJSON(400, gin.H{"msg": "username is already"})
		return
	}
	Salt := strings.ReplaceAll(uuid.NewString(), "-", "")
	user = db.User{
		Username: req.Username,
		Password: md5V(req.Password + Salt),
		Salt:     Salt,
	}
	if err := tx.Create(&user); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"msg": err.Error})
		return
	}
	c.JSON(200, gin.H{"code": 0})

}
