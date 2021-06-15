package database

import (
	"github.com/asciiflix/server/config"
	"github.com/asciiflix/server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Create video
func CreateVideo(video model.Video) error {
	result := global_db.Create(&video)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Get video by id
func GetVideo(videoId string) (*model.Video, error) {
	var video model.Video
	result := global_db.Where("id = ?", videoId).First(&video)
	if result.Error != nil {
		return nil, result.Error
	}
	return &video, nil
}

//Get all Videos
func GetVideos() (*[]model.Video, error) {
	var videos []model.Video
	result := global_db.Find(&videos)
	if result.Error != nil {
		return nil, result.Error
	}
	return &videos, nil
}

//Update video by id
func UpdateVideo(videoId string, video model.Video) error {
	//Get Video
	var videoToUpdate model.Video
	result := global_db.Where("id = ?", videoId).First(&videoToUpdate)
	if result.Error != nil {
		return result.Error
	}
	//Replaces non-zero fields
	result = global_db.Model(&videoToUpdate).Updates(video)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Delete video by id
func DeleteVideo(videoId string) error {
	result := global_db.Delete(&model.Video{}, videoId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

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

	//Response
	var response = map[string]interface{}{"message": "Successfully found VideoContent by ID"}
	response["content"] = videoContent
	return response
}

//Testing for UUID-GEN
/*func CreateVideo() {
	var video model.VideoStats

	video.ID, _ = uuid.NewV4()
	video.Title = "Test Title"
	video.Description = "Test Desc"

	global_db.Create(&video)
}
*/
