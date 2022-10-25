# MyGram

Zulkarnen
149368582101-344

Base URL HTTP http://localhost:8080/

## Open Endpoints

Open endpoints require no Authentication.

* Register : `POST /users/register/`
* Login  : `POST /users/login/`
* Documentation  : `GET /docs/index.html`


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

Endpoints for viewing and manipulating the Photo that the Authenticated User
has permissions to access.


* Create Comment `POST /photos/`
* Show Photos `GET /photos/`
* Update a Comment `PUT /photos/:id/`
* Delete An Account `DELETE /photos/:id/`



### Comments related

Endpoints for viewing and manipulating the Comments that the Authenticated User
has permissions to access.


* Create Comment `POST /comments/`
* Show Comments `GET /comments/`
* Update a Comment `PUT /comments/:id/`
* Delete An Comment `DELETE /comments/:id/`


### Social Media's related

Endpoints for viewing and manipulating the Social Media's that the Authenticated User
has permissions to access.


* Create Social Media `POST /socialmedias/`
* Show Social Medias `GET /socialmedias/`
* Update a Social Media `PUT /socialmedias/:id/`
* Delete An Social Media `DELETE /socialmedias/:id/`



