package database

import (
	"fmt"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/kidboy-man/simple-shop/conf"
	"github.com/kidboy-man/simple-shop/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB //database

func InitDB() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	fmt.Println(dbUri)

	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
	}

	conf.AppConfig.DbClient = conn
	conf.AppConfig.DbClient.Debug().AutoMigrate(&models.Metadata{}) //Database migration
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
