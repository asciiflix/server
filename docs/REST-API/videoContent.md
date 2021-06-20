# Video Content
## The API-Endpoint has been updated! Now it´s recommended to use the Endpoints with the VideoStats. With that Update, you can no longer get/delete VideoContent with the videoContentID, just with the regular videoID like ``bb2d47c5-b1b0-453e-ac98-b270bc5e9ee0`` 
Keep in mind the createContent and deleteContent Endpoints are behind ``/secure``, so you will need a valid JWT-Token.

# Overview
- [Create VideoContent](#create-videocontent)
- [Get VideoContent](#get-videocontent)
- [Delete VideoContent](#delete-videocontent)
- [Get/Delete Non Existing Video](#getdelete-non-existing-video)
- [No HTTP Parameters](#no-http-parameters)

# Create VideoContent
You can upload/create a Video at the the Endpoint ``/secure/video/createContent`` with a POST request, you can upload any JSON Object/Array within the Key ``Video``. <br>
Here is an Example:
## Usage
````json
{
    "Video": {
        "Rows": [
        "______",
        "|----|",
        "|----|",
        "______"
        ],
        "height": 4,
        "width": 6
    }
}
````
## Response
You will get the VideoContentID which is really important and a message. Please save the ``_id``  to properly connect the VideoContent with VideoStats later.
````json
{
    "_id": "60c4a7a234cad8de5f2ea71c",
    "message": "Successfully created VideoContent"
}
````

# Get VideoContent
To get VideoContent you basically call the Endpoint ``/video/getContent`` with a GET Request
## Usage
Just call ``/video/getContent?id={{videoID}}``, obviously you have to enter the correct videoID in the HTTP Parameters.
## Response
````json
{
    "content": {
        "ID": "60c4a7a234cad8de5f2ea71c",
        "Video": {
            "Rows": [
                "______",
                "|----|",
                "|----|",
                "______"
            ],
            "height": 4,
            "width": 6
        }
    },
    "message": "Successfully found VideoContent by ID"
}
````

# Delete VideoContent
To delete VideoContent you basically call the Endpoint ``secure/video/deleteContent`` with a DELETE Request
## Usage
Just call ``/secure/video/deleteContent?id={{videoID}}``, obviously you have to enter the correct videoID in the HTTP Parameters.
## Response
````json
{
    "message": "Successfully deleted VideoContent by ID",
    "result": {
        "DeletedCount": 1
    }
}
````

# Get/Delete Non Existing Video
If you are trying to get or delete VideoContent with an Invalid ID or and ID which is already deleted, you will get a response like this:
````json
{
    "message": "ID does not exist."
}
````

# No HTTP Parameters
If you are trying to get or delete VideoContent without an ID in the HTTP Parameters, you will get:
````json
{
    "message": "No ID in parameters"
}
````