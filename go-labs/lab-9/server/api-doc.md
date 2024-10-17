# Users
### Api to manage users on server <br><br>
### Endpoints
- [/users](#user-content-get-users)
- [/users/{id}](#user-content-get-user-by-user-id) <br>

### Methods 
- GET  get users [/users/](#user-content-get-users)
- GET  get user by id [/users/{id}](#user-content-get-user-by-user-id)
- POST create user [/users/](#user-content-create-user)
- PUT  update user [/users/{id}](#user-content-update-user-by-id)
- DELETE delete user by id [/users/{id}](#user-content-delete-user-by-user-id)

<br>
<hr>
<br>

## Get Users
> GET **/users** <br>
## Request
### QUERY PARAMS
| Param | Required | Data type | Description |
| ---   | ---     | ---       | ---         |
| **page** | optional | int | Count of displayed pages |
| **limit** | optional | int | Limit if displayed objects on page|
| **name** | optional | string | User name filter |
| **age** | optional | int | User age filter |

### Request example
>  /users?page=2&limit=5&name=John&age=30

## Response 
> 500 server error <br />
> 400 inccorrect input <br />
> 200 users found

### Response example
#### json

```json
[
    {
        "id": "605c72efb3e3c1b8b56c1a23",
        "name": "John Doe",
        "age": 30
    },
    {
        "id": "605c72efb3e3c1b8b56c1a24",
        "name": "Jane Smith",
        "age": 25
    }
]
```

### Response schema
| Object | Description | required |
| --- | :--- | --- |
| id | **string** <br> The id of the user | req |
| name | **string** <br> The name of the user | req |
| age | **integer** <br> The age of the user | req |

<br>
<hr>
<br>

## Create User
> POST **/users/** <br>
## Request
### BODY PARAMS
| Param | Required | Data type | Description |
| ---   | ---     | ---       | ---         |
| **Name** | Yes | string | User name |
| **Age** | Yes | int | User age |

### Request example
> curl -X POST http://localhost:8000/users \
-H "Content-Type: application/json" \
-d '{"username":"usr","email":"usr@test.com","passwd":"secpassword"}'

## Response 
> 500 server error <br />
> 400 inccorrect input <br />
> 201 user created

### Response example
#### json

```json
"id": "605c72efb3e3c1b8b56c1a23",
"name": "John Doe",
"age": 30
```

### Response schema
| Object | Description | required |
| --- | :--- | --- |
| id | **string** <br> The id of the user | req |
| name | **string** <br> The name of the user | req |
| age | **integer** <br> The age of the user | req |

<br>
<hr>
<br>

## Get User by User ID
> GET **/users/{id}** <br>
## Request
### PATH PARAMS
| Param | Required | Data type | Description |
| ---   | ---     | ---       | ---         |
| **id** | Yes | string | User id |

### Request example
> /users/60b6c8f1f1e2b1c3d4e5f6a7

## Response 
> 500 server error <br />
> 400 inccorrect input <br />
> 200 user found

### Response example
#### json

```json
"id":"60b6c8f1f1e2b1c3d4e5f6a7",
"name":"SUPER-TEST",
"age":55
```

### Response schema
| Object | Description | required |
| --- | :--- | --- |
| id | **string** <br> The id of the user which was found by id in request | req |
| name | **string** <br> The name of the user | req |
| age | **integer** <br> The age of the user | req |

<br>
<hr>
<br>

## Update User by ID
> PUT **/users/{id}** <br>
## Request
### PATH PARAMS
| Param | Required | Data type | Description |
| ---   | ---     | ---       | ---         |
| **id** | Yes | string | User id |<br>

### BODY PARAMS
| Param | Required | Data type | Description |
| ---   | ---     | ---       | ---         |
| **name** | Optional | string | New name of the User |
| **age** | Optional | int | New age of the User |

### Request example
> curl -X POST http://localhost:8000/users \
     -H "Content-Type: application/json" \
     -d '{
           "name": "John Doe",
           "age": 30
         }'

## Response 
> 500 server error <br />
> 400 inccorrect input <br />
> 200 user updated

### Response example
#### json

```json
"id": "60b6c8f1f1e2b1c3d4e5f6a7",
"name": "John Doe",
"age": 30
```

### Response schema
| Object | Description | required |
| --- | :--- | --- |
| id | **string** <br> The id of the user | req |
| name | **string** <br> The name of the user | req |
| age | **int** <br> The age of the user | req |

<br>
<hr>
<br>

## Delete User by User ID
> DELETE **/users/{id}** <br>
## Request
### PATH PARAMS
| Param | Required | Data type | Description |
| ---   | ---     | ---       | ---         |
| **id** | Yes | string | User id |

### Request example
> /users/60b6c8f1f1e2b1c3d4e5f6a7

## Response 
> 500 server error <br />
> 400 inccorrect input <br />
> 204 user deleted
