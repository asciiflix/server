# User
Keep in mind the getPrivate/update/delete user Endpoints are behind `/secure`, so you will need a valid JWT-Token.

# Get User
To get Public Information about an user, you will just need the UserID. Call the Endpoint ``/user/getUser`` with a GET request, and the UserID as an parameter.
## Usage
Just call ``/user/getUser?id={{userID}}`` with a GET Request.
## Response
````json
{
    "UserID": 62,
    "Name": "Rachelle4",
    "Description": "",
    "Picture_ID": "",
    "Videos": null
}
````

# Get Private User
To get more Information about an user, you can use the Endpont ``/secure/user/getUser`` with a GET request and Informations like email etc. Only the user itself can access his private Information. The JWT Token has to be linked to the requested user.
## Usage
Just call ``/secure/user/getUser?id={{userID}}`` with a GET Request and the JWT-Token in the Header ``Token``.
## Response
````json
{
    "UserID": 62,
    "Name": "Rachelle4",
    "Email": "Rachelle4@asciiflix.tech",
    "Description": "",
    "Picture_ID": "",
    "Videos": null,
    "Comments": null,
    "Likes": null
}
````

# Get All Users
To get all existing Users, you can simply call ``/user/getUsers`` with a GET request. 
## Usage
Call ``/user/getUsers``
## Response
````json
[
    {
        "UserID": 1,
        "Name": "Laury_Becker",
        "Description": "",
        "Picture_ID": "",
        "Videos": null
    },
    {
        "UserID": 2,
        "Name": "Garnett.Baumbach",
        "Description": "",
        "Picture_ID": "",
        "Videos": null
    },
    {
        "UserID": 3,
        "Name": "Etha_Bartoletti89",
        "Description": "",
        "Picture_ID": "",
        "Videos": null
    },
    {
        "UserID": 4,
        "Name": "Shanna44",
        "Description": "",
        "Picture_ID": "",
        "Videos": null
    },
    {
        "UserID": 5,
        "Name": "Bob",
        "Description": "",
        "Picture_ID": "",
        "Videos": null
    },
    {
        "UserID": 38,
        "Name": "Jammie.Torp",
        "Description": "",
        "Picture_ID": "",
        "Videos": null
    },
    {
        "UserID": 42,
        "Name": "sadasdas",
        "Description": "",
        "Picture_ID": "",
        "Videos": null
    },
    {
        "UserID": 61,
        "Name": "aspdjasdjalsdjasjd",
        "Description": "",
        "Picture_ID": "",
        "Videos": null
    },
    {
        "UserID": 63,
        "Name": "newUserNameforAlvena.Schaefer",
        "Description": "Profile of Alvena.Schaefer",
        "Picture_ID": "",
        "Videos": null
    }
]
````

# Update User
To update an user, you call ``/secure/user/updateUser`` with a PUT request.
## Usage
Call ``/secure/user/updateUser?id={{userID}}`` with the users linked JWT-Token and the correct UserID.
The Body can have almost every option to update:
````json
{
    "Name": "newName",
    "Description": "newDescription",
    "EMail": "",
    "Password": "",
    "Picture_ID": ""
}
````
## Response
The Reponse Body is empty, you will get a ``202`` Code

# Delete User
To delete a user, call the endpoint ``/secure/user/deleteUser`` with a DELETE request.
## Usage
Call ``/secure/user/deleteUser?id={{userID}}`` with the users linked JWT-Token and the correct UserID, to delete the user.
## Response
The Reponse Body is empty, you will get a ``204`` Code