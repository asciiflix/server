package database

import (
	"fmt"
	"net/url"

	"github.com/asciiflix/server/config"
	"github.com/asciiflix/server/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var global_db *gorm.DB

func StartDatabase() {

	dsn := url.URL{
		User:     url.UserPassword(config.Database.User, config.Database.Password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%s", config.Database.Host, config.Database.Port),
		Path:     config.Database.Db,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{
		//Disabling any log output from gorm
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic("Can't connect to database!")
	}
	config.Log.Info("DB Connected")
	global_db = db

	db.AutoMigrate(&model.User{})
}
