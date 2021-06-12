# Documentation about the ASCIIflix Server API
Every Endpoint is well documented in this Documentation :)

## List of Endpoints
API-Endpoint   | Methods | Doc   | Description
---------------| ------- | ----- | -----------
``/status``    | GET     | -     | Get API Status 
``/register``  | POST    |[authentication](./authentication.md) | Register a User 
``/login``     | GET     |[authentication](./authentication.md) | User Login to get a JWT Token 
``/secure/my_status``  | GET    | -     | (personalized) testing Status Page (testing jwt-token)
``/secure/video/createContent``    | POST  |[videoContent](./videoContent.md) | Create/Upload Video Content, will be stored in MongoDB
``/video/getContent``       | GET   |[videoContent](./videoContent.md) | Get Video Content by VideoContentID
``/secure/video/deleteContent``    | DELETE|[videoContent](./videoContent.md) | Delete Video Content by VideoContentID

