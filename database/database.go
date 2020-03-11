package database

import (
	"fmt"
	"github.com/OhMinsSup/story-server/database/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pborman/uuid"
	"os"
	"reflect"
	"strings"
)

// Initialize 데이터베이스 초기화
func Initialize() (*gorm.DB, error) {
	dbConfig := os.Getenv("DB_CONFIG")
	db, err := gorm.Open("postgres", dbConfig)

	// Logs SQL
	db.LogMode(true)
	db.Set("gorm:table_options", "charset=utf8")
	// created uuid
	db.Callback().Create().Before("gorm:create").Register("my_plugin:before_create", BeforeCreateUUID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	Migrate(db)

	return db, err
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.EmailAuth{})
	//db.Model(&AuthToken{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	fmt.Println("Auto Migration has beed processed")
}

func Inject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func BeforeCreateUUID(scope *gorm.Scope) {
	reflectValue := reflect.Indirect(reflect.ValueOf(scope.Value))
	if strings.Contains(string(reflectValue.Type().Field(0).Tag), "uuid") {
		uuid.SetClockSequence(-1)
		scope.SetColumn("id", uuid.NewUUID().String())
	}
}
