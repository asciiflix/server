package database

import (
	"fmt"
	"net/url"
	"time"

	"github.com/asciiflix/server/config"
	"github.com/asciiflix/server/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDatabase() (db *gorm.DB) {

	dsn := url.URL{
		User:     url.UserPassword(config.Database.User, config.Database.Password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%s", config.Database.Host, config.Database.Port),
		Path:     config.Database.Db,
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
