# Final Project Golang 2024 Spring. Go Blog with Comments and Authentication

## Author
Toleu Bahauddin 22B030598


This project is a simple blog application implemented in Go (Golang) that includes features such as CRUD operations for articles, comments, and user authentication.

## Database Structure

### Users

- `UserID` (Primary Key)
- `Username`
- `PasswordHash` (encrypted password)
- Other fields like email, role, etc.

### Posts

- `PostID` (Primary Key)
- `Title`
- `Content`
- `UserID` (Foreign Key, relation to Users)
- `CreatedAt`
- `UpdatedAt`

### Comments

- `CommentID` (Primary Key)
- `Content`
- `UserID` (Foreign Key, relation to Users)
- `PostID` (Foreign Key, relation to Posts)
- `CreatedAt`
- `UpdatedAt`

## API Structure

### Authentication and Registration

- Register a new user.
- User login with token retrieval.

### Posts

- Create a new post.
- Retrieve a list of all posts.
- Retrieve a specific post by ID.
- Update an existing post.
- Delete a post.

### Comments

- Create a new comment for a specific post.
- Retrieve all comments for a specific post.
- Update an existing comment.
- Delete a comment.

### Authorization and Access Rights

- Token validation and user authentication.
- Implementation of roles (e.g., user and administrator).
- Restrict access to certain operations based on roles.

