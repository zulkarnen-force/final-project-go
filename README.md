# MyGram

Base URL http://localhost:8080/

## Open Endpoints

Open endpoints require no Authentication.
# RESTAPIDocs Examples

These examples were taken from projects mainly using [Django Rest
Framework](https://github.com/tomchristie/django-rest-framework) and so the
JSON responses are often similar to the way in which DRF makes responses.

Where full URLs are provided in responses they will be rendered as if service
is running on 'http://testserver/'.

## Open Endpoints

Open endpoints require no Authentication.

* Register : `POST /users/register/`
* Login  : `POST /users/login/`


## Endpoints that require Authentication

Closed endpoints require a valid Token to be included in the header of the
request. A Token can be acquired from the Login view above.
`

### Users related

Endpoints for viewing and manipulating the Accounts that the Authenticated User
has permissions to access.

* Delete User  `DELETE /users/`
* Create Account  `PUT /users/`
* Show An Account `GET /api/accounts/:pk/`
* Update An Account `PUT /api/accounts/:pk/`
* Delete An Account `DELETE /api/accounts/:pk/`


## Endpoints that require Authentication

Closed endpoints require a valid Token to be included in the header of the
request. A Token can be acquired from the Login view above.

### Current User related

Each endpoint manipulates or displays information related to the User whose
Token is provided with the request:

* Delete User  `DELETE /users/`
* Update An User `PUT /users/`
* Delete An User `DELETE /users/`

### Comments related

Endpoints for viewing and manipulating the Comments that the Authenticated User
has permissions to access.


* Create Comment `POST /comments/`
* Show Comments `PUT /comments/`
* Update a Comment `PUT /comments/:id/`
* Delete An Account `DELETE /comment/:id/`
