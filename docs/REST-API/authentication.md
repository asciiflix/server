# API Authentication

# Overview
- [Register](#register)
- [Login](#login)
- [Use Secure Endpoints](#use-secure-endpoints)

# Register
You can register a User at ``/register`` like this JSON Body:
````json
{
    "email": "exampleUser@example.com",
    "name": "ExampleUser",
    "password": "ExamplePassword"
}
````
## Possible Response
For a successful user registration you will get:
````JSON
{
    "message": "User successfully registered."
} 
````
If the user already exists you will get:
````JSON
{
    "message": "User already exists."
} 
````

# Login
If you want to Login to the API to access secure endpoints, you can login at ``/login`` like:
````json
{
    "email": "exampleUser@example.com",
    "password": "ExamplePassword"
}
````
## Possible Response
If the Credentials are correct, yo will get your jwt-token and a response message: <br>
To use secure Endpoints, please save your jwt-token.
````json
{
    "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjI3MzMzNDAsImlhdCI6MTYyMjcyOTc0MCwiaXNzIjoiYXB0LmFzY2lpZmxpeC50ZWNoIiwiVXNlcl9JRCI6MzQsIlVzZXJfZW1haWwiOiJKYWRhX0Jsb2NrQGFzY2lpZmxpeC50ZWNoIn0.RKIstLIF8UvlZZ6VaOA0eUVhDWu6cFfP8pcgWK06eVg",
    "message": "Successfully logged in"
}
````
When you enter wrong password you will get:
````json
{
    "message": "Wrong Password"
}
````
If you using Credentials for a User which does not exists, you will get:
````json
{
    "message": "User does not exist."
}
````

# Logout
To add a JWT-Token to our Blacklist simply call ``/secure/logout``. The used JWT token will be added to our blacklist to prevent further use.

# Use Secure Endpoints
If you want to access secure endpoints like ``/secure/my_status``, you will need a valid JWT-Token. <br>
No JWT-Token? -> How to get a JWT-Token, please read:

Now you can simply at the JWT-Token to your http-headers with the Key ``Token`` and the Value will be your JWT-Token. Here is an example:<br>
``Token:yourPrivatJWTToken``

Keep in mind do NOT share your private JWT-Token, the API can´t block an existing token, your JWT-Token is valid until the expire time is reached.

