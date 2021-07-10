package database

import (
	"errors"

	"github.com/asciiflix/server/config"
	"github.com/asciiflix/server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func getVideoCollection() *mongo.Collection {
	return global_mongo_client.Database("asciiflix").Collection("videos")
}

func CreateVideoContent(video *model.VideoContent) map[string]interface{} {
	//MongoDB Insert
	result, err := getVideoCollection().InsertOne(global_mongo_context, video)

	//Error Handling
	if err != nil {
		config.Log.Error(err)
		return map[string]interface{}{"message": "Error in MongoDB"}
	}

	//Response
	var response = map[string]interface{}{"message": "Successfully created VideoContent"}
	response["_id"] = result.InsertedID
	return response
}

func DeleteVideoContent(contentID primitive.ObjectID) map[string]interface{} {
	//Try to Delete VideoContent by ID
	result, err := getVideoCollection().DeleteOne(global_mongo_context, bson.M{"_id": contentID})

	//Error Handling
	if err != nil {
		config.Log.Error(err)
		return map[string]interface{}{"message": "Error in MongoDB"}
	}
	if result.DeletedCount == 0 {
		return map[string]interface{}{"message": "ID does not exist."}
	}

	//Response
	var response = map[string]interface{}{"message": "Successfully deleted VideoContent by ID"}
	response["result"] = result
	return response
}

func GetVideoContent(contentID primitive.ObjectID) map[string]interface{} {
	var videoContent model.VideoContent

	//Search by ContentID for VideoContent Entry
	err := getVideoCollection().FindOne(global_mongo_context, bson.M{"_id": contentID}).Decode(&videoContent)

	//Error Handlong
	if err != nil {
		config.Log.Error(err)
		return map[string]interface{}{"message": "ID does not exist."}
	}

	//Increment views
	result := global_db.Table("videos").Where("video_content_id = ?", contentID.Hex()).UpdateColumn("views", gorm.Expr("views + ?", 1))
	if result.Error != nil {
		config.Log.Error(result.Error)
	}
	//Response
	var response = map[string]interface{}{"message": "Successfully found VideoContent by ID"}
	response["content"] = videoContent
	return response
}

func GetVideoThumbnail(contentID primitive.ObjectID) (interface{}, error) {
	var videoContent model.VideoContent

	//Search by ContentID for VideoContent Entry
	err := getVideoCollection().FindOne(global_mongo_context, bson.M{"_id": contentID}).Decode(&videoContent)

	//Error Handling
	if err != nil {
		config.Log.Error(err)
		return nil, errors.New("not found")
	}

	var thumbnail interface{}
	if frame, ok := videoContent.Video["Frames"].(primitive.A); ok {
		thumbnail = frame[0]
	}

	return thumbnail, nil
}
