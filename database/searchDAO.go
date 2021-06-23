package database

import (
	"errors"

	"github.com/asciiflix/server/model"
)

func GetSearchResult(query string) (*model.SearchResult, error) {
	if query == "" {
		return nil, errors.New("empty query")
	}

	//Search for Users (name, description)
	var tmpUsers []model.User
	db_result := global_db.Preload("Videos").Where("name LIKE ?", "%"+query+"%").Or("description LIKE ?", "%"+query+"%").Find(&tmpUsers)
	if db_result.Error != nil {
		return nil, db_result.Error
	}

	//Search for Videos (title, uuid, description)
	var tmpVideos []model.Video
	db_result = global_db.Where("title LIKE ?", "%"+query+"%").Or("uuid LIKE ?", "%"+query+"%").Or("description LIKE ?", "%"+query+"%").Find(&tmpVideos)
	if db_result.Error != nil {
		return nil, db_result.Error
	}

	//Build Results
	var result model.SearchResult
	var resultUsers []model.UserDetailsPublic
	var resultVideos []model.VideoPublic

	for _, users := range tmpUsers {
		resultUsers = append(resultUsers, users.GetPublicUser())
	}

	for _, videos := range tmpVideos {
		resultVideos = append(resultVideos, model.GetPublicVideo(videos))
	}

	result.Users = resultUsers
	result.Videos = resultVideos

	return &result, nil
}
