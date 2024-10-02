# Blogging API

# Requirements
* Design and implement a RESTful API using any Programmatic language
    - Completed

* Include endpoint(s) for retrieving all posts, retrieving a single post by ID, creating a new post, updating an existing post, and deleting a post.
    - Completed

* Implement filtering for posts based on author and creation date.
    - Completed

* Ensure that the API follows RESTful principles and returns appropriate HTTP status codes and error messages.
    - Completed

* Implement authentication and authorization(Read/Write/Administer)  mechanisms.
    - Completed

* Implement pagination for retrieving a large number of posts.
    - Completed but doesn't fully work, needs testing

* Design/Define database architecture (SQL or NoSQL) to persist the data.
    - Not completed

* Write unit tests to validate the sanity of the application
    - Completed but doesn't work as expected

* Provide API documentation for the same.
    - Done

## Authentication
JWT-based authentication has been used where the user can
1. Register themselves
2. Generate token based on credentials
3. Use the generated token for REST services

### POST /register
- **Request**: 
```json
{
  "username": "user1",
  "password": "password1",
  "role": "write"
}
```

### POST /login
- **Request**: { "username": "user1", "password": "password1" }
- **Response**: { "token": "<JWT_TOKEN>" }

## Posts Endpoints

### GET /posts
- Retrieves all posts
- Query parameters used: 
  - page (for pagination, default 1)
  - limit (number of posts per page, default 10)

### GET /posts/:id
- **Response** : Retrieves a post by ID

### POST /posts
- **Request**: 
```json
{
  "title": "Post Title",
  "content": "Post content",
  "author": "Author"
}
```

### 3. **Steps to Run and Test**:

1. **Run the application**:
   ```bash
   go run main.go
