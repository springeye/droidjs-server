package db

import (
	"github.com/google/uuid"
	_ "github.com/satori/go.uuid"
	"gorm.io/driver/sqlite" // 基于 GGO 的 Sqlite 驱动
	"log"
	"strings"
	"time"
	// "github.com/glebarez/sqlite" // 纯 Go 实现的 SQLite 驱动, 详情参考： https://github.com/glebarez/sqlite
	"gorm.io/gorm"
)

type Model struct {
	//ID        uint           `gorm:"primaryKey;autoIncrement:true" json:"id"`
	ID        string         `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (u *Model) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = strings.ReplaceAll(uuid.NewString(), "-", "")
	return
}

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}
func Setup() {
	//db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	var err error
	db, err = gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
	if err != nil {
		log.Panic("数据库链接格式参考: https://gorm.io/zh_CN/docs/connecting_to_the_database.html")
		panic("failed to connect database")
	}

	// 迁移 schema

	dst := make([]interface{}, 0)
	dst = append(dst, new(User))
	dst = append(dst, new(Token))
	dst = append(dst, new(Device))
	dst = append(dst, new(Script))
	err = db.Migrator().DropTable(dst...)
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(dst...)
	if err != nil {
		log.Panic("数据库链接格式参考: https://gorm.io/zh_CN/docs/connecting_to_the_database.html")
		panic(err)
	}
	initTestData()

}

func initTestData() {
	user := User{
		Username: "henjue",
		Password: "123456",
		Salt:     "123",
		Tokens: []Token{
			{},
		},
		Devices: []Device{
			{},
		},
	}
	if err := db.Create(&user).Error; err != nil {
		panic(err)
	}
}

// github.com/mattn/go-sqlite3
func init() {

}
