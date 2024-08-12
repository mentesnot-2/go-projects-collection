# Task Management API Documentation

## Overview

The Task Management API allows you to perform CRUD (Create, Read, Update, Delete) operations on tasks, as well as handle user registration, login, and authorization. The API interacts with a MongoDB database to store and retrieve task and user data, and it includes JWT-based authentication to ensure secure access.

## Table of Contents

- [MongoDB Configuration](#mongodb-configuration)
- [Authentication & Authorization](#authentication--authorization)
- [API Endpoints](#api-endpoints)
  - [User Registration](#user-registration)
  - [User Login](#user-login)
  - [Create a Task](#create-a-task)
  - [Get All Tasks](#get-all-tasks)
  - [Get a Task by ID](#get-a-task-by-id)
  - [Update a Task](#update-a-task)
  - [Delete a Task](#delete-a-task)
- [Error Handling](#error-handling)
- [Example Responses](#example-responses)

## MongoDB Configuration

### Environment Variables

To configure MongoDB, the following environment variables should be set:

- `MONGO_URI`: The MongoDB connection string URI.
- `DB_NAME`: The name of the database.
- `COLLECTION_NAME_TASKS`: The name of the tasks collection within the database.
- `COLLECTION_NAME_USERS`: The name of the users collection within the database.
- `JWT_SECRET`: The secret key used for signing JWT tokens.

### Example `.env` File

```env
MONGO_URI=mongodb+srv://username:password@cluster0.example.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0
DB_NAME=task_management_db
JWT_SECRET=your_secret_key_here
```
### Authentication & Authorization
#### JWT Authentication
The API uses JSON Web Tokens (JWT) for authentication. After logging in, the client must include the JWT in the Authorization header for all subsequent requests that require authentication.

##### Roles
User: Can view all tasks and view specific task.
Admin: Can perform all user actions, as well as manage other users.



## API Endpoints

### User Registration

#### POST /register

Registers a new user with the system.

##### Request Body
```json
{
  "username": "john_doe",
  "password": "password123"
}
```

##### Response
```json
{
    "message": "User created successfully",
    "result": {
        "id": "66b9f13a19497721d3d663d1",
        "username": "kintsugi",
        "role": "user"
    }
}
```
#### User Login

#### POST /login

Logs in an existing user and returns a JWT token and user.

##### Request Body
```json
{
  "username": "john_doe",
    "password": "password123"
}
```

##### Response
```json
{
    "User": {
        "user": {
            "id": "66b9f13a19497721d3d663d1",
            "username": "kintsugi",
            "role": "user"
        },
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciIsInVzZXJfaWQiOiI2NmI5ZjEzYTE5NDk3NzIxZDNkNjYzZDEifQ.5bFdqYp3EDN8x3f3WoGaEhDh3fP6uFEH6kOpxkqlw6Y"
    },
    "message": "Login successful"
}
```

#### Create Task

#### POST /tasks

Creates a new task. Requires authentication and must be Admin to create task.

##### Request Body
```json
{
    "title": "Complete API documentation",
    "description": "Write the detailed API documentation for the task management system.",
    "due_date": "2024-08-31T23:59:59Z",
    "status": "pending"
}

```

##### Response
```json
{
    "message": "Task created successfully",
    "result": {
        "id": "66b9f13a19497721d3d663d1",
        "title": "Complete API documentation",
        "description": "Write the detailed API documentation for the task management system.",
        "due_date": "2024-08-31T23:59:59Z",
        "status": "pending"
    }
}
```

#### Get All Tasks


#### GET /tasks

Retrieves all tasks. Requires authentication.

##### Response
```json
{
    "Number of tasks": 1,
    "tasks": [
        {
            "id": "66b9cf04c47580f08a396216",
            "title": "Complete Project Report",
            "description": "Finalize and submit the project report by end of the week asap.",
            "due_date": "2024-08-20T17:00:00Z",
            "status": "completed",
            "user_id": "66b9ce13ba45913ab6dcd1e5"
        }
    ]
}
```

#### Get a Task by ID

#### GET /tasks/:id

Retrieves a specific task by ID. Requires authentication.

##### Response
```json
{
    "Task": {
        "id": "66b9cf04c47580f08a396216",
        "title": "Complete Project Report",
        "description": "Finalize and submit the project report by end of the week asap.",
        "due_date": "2024-08-20T17:00:00Z",
        "status": "completed",
        "user_id": "66b9ce13ba45913ab6dcd1e5"
    }
}
```


#### Update a Task

#### PUT /tasks/:id

Updates an existing task by ID. Requires authentication and must be Admin to update task.

##### Request Body
```json
{
    "title": "Complete API documentation",
    "description": "Write the detailed API documentation for the task management system.",
    "due_date": "2024-08-31T23:59:59Z",
    "status": "completed"
}
```

##### Response
```json
{
    "message": "Task updated successfully",
    "result": {
        "id": "66b9f13a19497721d3d663d1",
        "title": "Complete API documentation",
        "description": "Write the detailed API documentation for the task management system.",
        "due_date": "2024-08-31T23:59:59Z",
        "status": "completed"
    }
}
```

#### Delete a Task

#### DELETE /tasks/:id

Deletes an existing task by ID. Requires authentication and must be Admin to delete task.

##### Response
```json
{
    "message": "Task deleted successfully",
    "result": {
        "id": "66b9f13a19497721d3d663d1",
        "title": "Complete API documentation",
        "description": "Write the detailed API documentation for the task management system.",
        "due_date": "2024-08-31T23:59:59Z",
        "status": "completed"
    }
}
```

### Error Handling

The API returns appropriate error messages and status codes for various scenarios, such as invalid requests, unauthorized access, and server errors.

### Example Responses

The API responses include example JSON data for each endpoint to illustrate the expected format and structure of the data returned.
