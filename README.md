# MyGram

Base URL HTTP http://localhost:8080/

## Open Endpoints

Open endpoints require no Authentication.

* Register : `POST /users/register/`
* Login  : `POST /users/login/`


## Endpoints that require Authentication

Closed endpoints require a valid Token to be included in the header of the
request. A Token can be acquired from the Login view above.


## Endpoints that require Authentication

Closed endpoints require a valid Token to be included in the header of the
request. A Token can be acquired from the Login view above.

### Current User related

Each endpoint manipulates or displays information related to the User whose
Token is provided with the request:

* Delete User  `DELETE /users/`
* Update An User `PUT /users/`
* Delete An User `DELETE /users/`


### Photos related

Endpoints for viewing and manipulating the Photos that the Authenticated User
has permissions to access.


* Create Photo `POST /photos/`
* Show Photos `PUT /photos/`
* Update a Photo `PUT /photos/:id/`
* Delete An Photo `DELETE /photos/:id/`


### Comments related

Endpoints for viewing and manipulating the Comments that the Authenticated User
has permissions to access.


* Create Comment `POST /comments/`
* Show Comments `PUT /comments/`
* Update a Comment `PUT /comments/:id/`
* Delete An Account `DELETE /comment/:id/`


