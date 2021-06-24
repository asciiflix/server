# Documentation about the ASCIIflix Server API

Every Endpoint is well documented in this Documentation :)

## List of Endpoints

| API-Endpoint                  | Methods | Doc                                   | Description                                            |
| ----------------------------- | ------- | ------------------------------------- | ------------------------------------------------------ |
| `/status`                     | GET     | -                                     | Get API Status                                         |
| `/register`                   | POST    | [authentication](./authentication.md) | Register a User                                        |
| `/login`                      | GET     | [authentication](./authentication.md) | User Login to get a JWT Token                          |
| `/search`                     | GET     | [search](./search.md)                 | Search for Users and Videos                            |
| `/secure/my_status`           | GET     | -                                     | (personalized) testing Status Page (testing jwt-token) |
| `/secure/video/createContent` | POST    | [videoContent](./videoContent.md)     | Create/Upload Video Content, will be stored in MongoDB |
| `/video/getContent`           | GET     | [videoContent](./videoContent.md)     | Get Video Content by UUID                              |
| `/secure/video/deleteContent` | DELETE  | [videoContent](./videoContent.md)     | Delete Video Content by UUID                           |
| `/video/getVideo`             | GET     | [video](./video.md)                   | Get video by UUID                                      |
| `/video/getVideos`            | GET     | [video](./video.md)                   | Get all videos                                         |
| `/secure/video/createVideo`   | POST    | [video](./video.md)                   | Create video by UUID                                   |
| `/secure/video/deleteVideo`   | DELETE  | [video](./video.md)                   | Delete video by UUID                                   |
| `/secure/video/updateVideo`   | PUT     | [video](./video.md)                   | Edit video data by UUID                                |
| `/user/getUser`               | GET     | [user](./user.md)                     | Get User by UserID                                     |
| `/user/getUsers`              | GET     | [user](./user.md)                     | Get all Users                                          |
| `/secure/user/getUser`        | GET     | [user](./user.md)                     | Get Private-User Information by UserID                 |
| `/secure/user/editUser`       | PUT     | [user](./user.md)                     | Edit User by UserID                                    |
| `/secure/user/deleteUser`     | DELETE  | [user](./user.md)                     | Delete User by UserID                                  |