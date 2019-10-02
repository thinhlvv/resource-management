# Restful API Document for resource manager.

| Name | URL |
| -----| --- |
|[Login](#login) | `POST /user/login`|
|[Signup](#signup)|  `POST /user/signup`|
| | |
|[Create Resource](#create-resource) | `POST /resource` |
|[Get list Resources](#get-list-resources) | `GET /resource` |
|[Delete Resource](#delete-resource) | `DELETE /resource/:id` |
| | |
|[Create User](#create-user) | `POST /user` |
|[Get list User](#get-list-user) | `GET /user` |
|[Update User](#update-user) | `PUT /user/:id` | 
|[Delete User](#delete-user) | `DELETE /user/:id` |



## User 

| Role| Name  |
| 1   | User  |
| 2   | Admin |

### Login

Used to collect a Token for a registered User.

**URL** : `/user/login`

**Method** : `POST`

**Auth required** : NO

**Request**

```json
{
    "email": "[valid email address]",
    "password": "[password in plain text]"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzA1MTIxODgsImlhdCI6MTU2OTkwNzM4OCwic3ViIjoiMSIsInJvbGUiOjF9.1rs_Tq0eb3RndHRPq2hK2c_K840_aHLGHzPk9Nuq3bI"
}
```

## Error Response

**Condition** : If 'username' and 'password' combination is wrong.

**Code** : `400 BAD REQUEST`

**Content** :

```json
{
    "error": "Key: 'SignupReq.Email' Error:Field validation for 'Email' failed on the 'email' tag"
}
```




### Signup

New user registers.

**URL** : `/user/signup`

**Method** : `POST`

**Auth required** : NO

**Request**

```json
{
    "email": "[valid email address]",
    "password": "[password in plain text]",
    "role": 1
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzA1MTIxODgsImlhdCI6MTU2OTkwNzM4OCwic3ViIjoiMSIsInJvbGUiOjF9.1rs_Tq0eb3RndHRPq2hK2c_K840_aHLGHzPk9Nuq3bI"
}
```

## Error Response

**Condition** : If 'username' and 'password' combination is wrong.

**Code** : `400 BAD REQUEST`

**Content** :

```json
{
    "error": "Key: 'SignupReq.Email' Error:Field validation for 'Email' failed on the 'email' tag"
}
```






## Resource 

### Create Resource

Create new resource.

**URL** : `/resource`

**Method** : `POST`

**Auth required** : YES

**Request**

```json
{
	"name": "test"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "id": 1,
    "name": "test"
}
```

## Error Response

**Condition** : If name is duplicated

**Code** : `422 BAD REQUEST`

**Content** :

```json
{
    "error": "sql: duplicated field name"
}
```
