package database

import (
	"fmt"
	"net/url"
	"time"

	"github.com/asciiflix/server/config"
	"github.com/asciiflix/server/model"
	"github.com/gofrs/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var global_db *gorm.DB

func ConnectToDatabase() {

	dsn := url.URL{
		User:     url.UserPassword(config.Database.User, config.Database.Password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%s", config.Database.Host, config.Database.Port),
		Path:     config.Database.Db,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{
		//Disabling any log output from gorm
		//Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic("Can't connect to database!")
	}
	config.Log.Info("Connected to DB")
	global_db = db

	db.AutoMigrate(&model.User{}, &model.Video{}, &model.Comment{}, &model.Like{})
	uuid, _ := uuid.NewV4()
	db.Create(&model.User{
		Model:  gorm.Model{ID: 1},
		Name:   "Bob",
		Videos: nil,
	})
	db.Create(&model.Video{
		Model:          gorm.Model{ID: 1},
		UUID:           uuid,
		VideoContentID: "afds",
		Title:          "Title",
		Description:    "Desc",
		UploadDate:     time.Now(),
		Views:          10,
		UserID:         1,
		Comments: []model.Comment{{
			UserID:  1,
			VideoID: 1,
			Content: "Hello",
		}},
		Likes: nil,
	})

}
