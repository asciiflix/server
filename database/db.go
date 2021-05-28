package database

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/asciiflix/server/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDatabase() (db *gorm.DB) {

	dbUser, dbPassword, dbName, dbHost, dbPort :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT")

	dsn := url.URL{
		User:     url.UserPassword(dbUser, dbPassword),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%s", dbHost, dbPort),
		Path:     dbName,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	time.Sleep(5 * time.Second)
	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})
	if err != nil {
		panic("Can't connect to database!")
	}
	fmt.Println("DB Connected")

	db.Debug().AutoMigrate(&model.User{})

	return
}
