
# Task Management API Documentation

## Overview

The Task Management API allows you to perform CRUD (Create, Read, Update, Delete) operations on tasks. The API includes authentication and authorization features to ensure secure access to the data. Users must log in to obtain a JWT token, which is required to access protected endpoints.

## MongoDB Configuration

### Environment Variables

To configure MongoDB, the following environment variables should be set:

- `MONGO_URI`: The MongoDB connection string URI.
- `DB_NAME`: The name of the database.
- `COLLECTION_NAME`: The name of the collection within the database.
- `JWT_SECRET`: The secret key used to sign JWT tokens.

### Example `.env` File

```env
MONGO_URI=mongodb+srv://username:password@cluster0.example.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0
DB_NAME=test
COLLECTION_NAME=tasks
JWT_SECRET=your_secret_key
```


## API Endpoints
1 - Create a Task
2 - Get All Tasks
3 - Get a Task by ID
4 - Update a Task
5 - Delete a Task


### Create a Task
```http
POST /tasks
```
Creates a new task.

```json
{
    "title": "Task 1",
    "description": "This is task 1",
    "completed": false
}
```
#### Response

```json
{
    "message": "Task created successfully",
    "result" : {
        "InsertedID": "60f3b3b3b3b3b3b3b3b3b3b3"    
    }
}
```

### Get All Tasks
```http
GET /tasks
```
Retrieves all tasks.

#### Response

```json
{
    "message": "Tasks retrieved successfully",
    "result": [
        {
            "_id": "60f3b3b3b3b3b3b3b3b3b3b3",
            "title": "Task 1",
            "description": "This is task 1",
            "completed": false
        },
        {
            "_id": "60f3b3b3b3b3b3b3b3b3b3b",
            "title": "Task 2",
            "description": "This is task 2",
            "completed": true
        }
    ]
}
```

### Get a Task by ID
```http
GET /tasks/:id
```
Retrieves a task by ID.

#### Response

```json
{
    "message": "Task retrieved successfully",
    "task": {
        "_id": "60f3b3b3b3b3b3b3b3b3b3b3",
        "title": "Task 1",
        "description": "This is task 1",
        "completed": false
    }
}
```


### Update a Task
```http
PUT /tasks/:id
```
Updates a task by ID.


```json
{
    "title": "Task 1",
    "description": "This is task 1",
    "completed": true
}
```

#### Response

```json
{
    "message": "Task updated successfully",
    "result": {
        "MatchedCount": 1,
        "ModifiedCount": 1,
        "UpsertedCount": 0,
        "UpsertedID": null
    }
}
```

### Delete a Task
```http
DELETE /tasks/:id
```
Deletes a task by ID.

#### Response

```json
{
    "message": "Task deleted successfully",
    "result": {
        "DeletedCount": 1
    }
}
```

## Error Handling

The API returns appropriate error messages and status codes for various scenarios, such as invalid requests, missing data, or database errors.

### Example Error Response

```json
{
    "error": "Task not found",
}
```

### Example for Missing Data

```json
{
    "error": "title cannot be empty",
}

```

### Example for Database Error

```json
{
    "error": "Database error",
}
```

## Conclusion

The Task Management API provides a simple and efficient way to manage tasks using MongoDB as the underlying database. The API endpoints allow you to create, read, update, and delete tasks with ease. By following the API documentation, you can integrate task management functionality into your applications seamlessly.

```