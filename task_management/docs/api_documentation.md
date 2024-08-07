# Task Management API Documentation

## Endpoints

### Get All Tasks

**GET** `/tasks`

Response:
```json
[
    {
        "id": "string",
        "title": "string",
        "description": "string",
        "due_date": "string",
        "status": "string"
    }
]
```


### Get Task by ID
GET `/tasks/:id`

Response:
```json
{
    "id": "string",
    "title": "string",
    "description": "string",
    "due_date": "string",
    "status": "string"
}
```

### Create Task
POST `/tasks`

Request:
```json
{
    "title": "string",
    "description": "string",
    "due_date": "string"
}
```

Response:
```json
{
    "id": "string",
    "title": "string",
    "description": "string",
    "due_date": "string",
    "status": "string"
}
```
### Update Task

PUT `/tasks/:id`

Request:
```json
{
    "title": "string",
    "description": "string",
    "due_date": "string",
    "status": "string"
}
```
Response:
```json
{
    "id": "string",
    "title": "string",
    "description": "string",
    "due_date": "string",
    "status": "string"
}
```

### Delete Task
DELETE `/tasks/:id`

response:
```json
{}
```


### Testing with Postman

You can test each endpoint using Postman by sending the appropriate HTTP requests with the required payloads.

### Running the Application

To run the application, use the following command:

```bash
go run main.go
```