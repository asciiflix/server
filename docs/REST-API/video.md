# Video

Keep in mind the create/update/delete video Endpoints are behind `/secure`, so you will need a valid JWT-Token.

# Overview

| API-Endpoint                | Methods | Doc                          | Description             |
| --------------------------- | ------- | ---------------------------- | ----------------------- |
| `/video/getVideo`           | GET     | [getVideo](#get-video)       | Get video by UUID       |
| `/video/getVideos`          | GET     | [getVideos](#get-videos)     | Get all videos          |
| `/secure/video/createVideo` | POST    | [createVideo](#create-video) | Create video by UUID    |
| `/secure/video/deleteVideo` | DELETE  | [deleteVideo](#update-video) | Delete video by UUID    |
| `/secure/video/updateVideo` | PUT     | [updateVideo](#delete-video) | Edit video data by UUID |

# Get video

To get data of a video use the Endpoint `/video/getVideo` with a GET request.
Here is an Example:

## Usage

Just call `/video/getVideo?id={{uuid}}`, obviously you have to enter the correct UUID in the HTTP Parameters.

## Response

```json
{
  "ID": 65,
  "CreatedAt": "2021-06-20T12:42:57.628378Z",
  "UpdatedAt": "2021-06-20T12:42:57.628378Z",
  "DeletedAt": null,
  "UUID": "bb2d47c5-b1b0-453e-ac98-b270bc5e9ee0",
  "VideoContentID": "60cf37d1bc25ac1f9b51c30d",
  "Title": "Title",
  "Description": "Desc",
  "UploadDate": "0001-01-01T00:00:00Z",
  "Views": 0,
  "UserID": 65,
  "Comments": [],
  "Likes": null
}
```

# Get videos

To get data of a video use the Endpoint `/video/getVideos` with a GET request.
Here is an Example:

## Usage

Just call `/video/getVideo`.

# Create video

You can upload/create a Video at the the Endpoint `/secure/video/createVideo` with a POST request. Use a json Object with the key `VideoStats`
for video stats and the key `VideoContent` for video content. <br>
Here is an Example:

## Usage

```json
{
  "VideoStats": {
    "Title": "Video Title",
    "Description": "This is a video description"
  },
  "VideoContent": {
    "Video": {
      "Rows": ["______", "|----|", "|----|", "______"],
      "height": 4,
      "width": 6
    }
  }
}
```

## Response

You will get a UUID which is really important. Please save the `videoID` to properly connect the other endpoints later.

```json
{
  "videoID": "bb2d47c5-b1b0-453e-ac98-b270bc5e9ee0"
}
```

# Delete video

To delete video use the Endpoint `/video/deleteVideo` with a DELETE request.
Here is an Example:

## Usage

Just call `/video/deleteVideo?id={{uuid}}`, obviously you have to enter the correct UUID in the HTTP Parameters.

# Update video

To delete video use the Endpoint `/video/updateVideo` with a PUT request.
Here is an Example:

## Usage

Just call `/video/updateVideo?id={{uuid}}`, obviously you have to enter the correct UUID in the HTTP Parameters.
Send the updated values as json.

```json
{
  "Title": "Updated title",
  "Description": "Updated description"
}
```

# Get Recomendations

To get Reomendations (optionally for a user), returning a list of videos

## Usage

Simply call ``/video/getRecomendations?limit={{limit}}`` or ``/video/getUserRecomendations?limit={{limit}}`` with a given JWT token as a header.

The response will then be an array (the size of the limit, or less) for Video objects. 
If no videos exist, the response will be null, however that really shouldn't ever be the case.