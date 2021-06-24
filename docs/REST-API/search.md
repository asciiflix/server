# Search
Endpoint to Search for Users and Videos.

# Overview

## Usage
To perform a search just call the endpoint ``/search?query=SEARCH`` with a get request. You can enter any String into the http-parameter ``query``. <br>
The backend searches for users with these properties: name and description, for videos it will search for: title, description and the uuid.
## Response
As a response you will get a json object with a user array and a video array. If the api doesn't find anything for users or videos it will return a null array.
### User Response
````json
{
    "Users": [
        {
            "UserID": 106,
            "Name": "newUserNameforDemetris_Mann63",
            "Description": "Profile of Demetris_Mann63",
            "Picture_ID": "",
            "Videos": [
                {
                    "UUID": "a2a7b5a3-b640-40d4-8d88-ff8bc6fa9346",
                    "Title": "New title",
                    "Description": "New description",
                    "UploadDate": "2021-06-23T19:45:13.902284Z",
                    "Views": 0,
                    "Likes": 0,
                    "UserID": 106
                }
            ]
        }
    ],
    "Videos": null
}
````
### Video Response
````json
{
    "Users": null,
    "Videos": [
        {
            "UUID": "4264ba08-df6e-4f1a-b9ef-b0abb0c69e8d",
            "Title": "Test",
            "Description": "Desc",
            "UploadDate": "2021-06-23T18:58:46.540328Z",
            "Views": 0,
            "Likes": 0,
            "UserID": 105
        },
        {
            "UUID": "8ae0397b-de3d-4b31-9798-41ea2b4bdadd",
            "Title": "Test",
            "Description": "Desc",
            "UploadDate": "2021-06-23T19:34:02.735069Z",
            "Views": 4,
            "Likes": 0,
            "UserID": 105
        },
        {
            "UUID": "f03cad39-def3-44d1-9b54-3baa99aad297",
            "Title": "Test",
            "Description": "Desc",
            "UploadDate": "2021-06-23T19:39:30.307268Z",
            "Views": 0,
            "Likes": 0,
            "UserID": 105
        }
    ]
}
````
