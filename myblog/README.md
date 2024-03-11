# Final Project Golang 2024 Spring: Go Blog with Comments and Authentication

This project is a simple blog application implemented in Go (Golang) that includes features such as CRUD operations for articles, comments, and user authentication.

## Author

- Toleu Bahauddin 22B030598

## Database Structure

### Users

- `UserID` (Primary Key)
- `Created_At` (TIMESTAMP WITH TIME ZONE NOT NULL)
- `Update_At` (TIMESTAMP WITH TIME ZONE NOT NULL)
- `Delete_At` (TIMESTAMP WITH TIME ZONE)
- `Username`
- `PasswordHash` (encrypted password)
- `Followers` (INT NOT NULL)

### Posts

- `PostID` (Primary Key)
- `Created_At` (TIMESTAMP WITH TIME ZONE NOT NULL)
- `Update_At` (TIMESTAMP WITH TIME ZONE NOT NULL)
- `Delete_At` (TIMESTAMP WITH TIME ZONE)
- `Content` (TEXT NOT NULL)
- `Likes` (INT NOT NULL)
- `UserID` (Foreign Key, relation to Users)

### Comments

- `CommentID` (Primary Key)
- `Created_At` (TIMESTAMP WITH TIME ZONE NOT NULL)
- `Update_At` (TIMESTAMP WITH TIME ZONE NOT NULL)
- `Delete_At` (TIMESTAMP WITH TIME ZONE)
- `Content` (TEXT NOT NULL)
- `Likes` (INT NOT NULL)
- `UserID` (Foreign Key, relation to Users)
- `PostID` (Foreign Key, relation to Posts)

## API Structure

### Authentication

#### Signup

- **Method**: POST
- **Endpoint**: `/signup`
- **Body**:

  ```json
  {
    "username": "your_username",
    "password": "your_password"
  }
  ```

#### Signin

- **Method**: POST
- **Endpoint**: `/signin`
- **Body**:

```json
{
  "username": "your_username",
  "password": "your_password"
}
```

#### Users

- **Method**: GET
- **Endpoint**: `/users`

#### Posts

Create post(protected)

- **Method**: POST
- **Endpoint**: `/post`
- **Body**:

```json
{
  "content": "your_content"
}
```

List posts

- **Method**: GET
- **Endpoint**: `/post`

GET post by ID

- **Method**: GET
- **Endpoint**: `/post/:id`

Update post(protected)

- **Method**: PUT
- **Endpoint**: `/post/:id`

```json
{
  "content": "your_updated_content"
}
```

Delete post(protected)

- **Method**: DELETE
- **Endpoint**: `/post/:id`

#### Comments

Create comment(protected)

- **Method**: POST
- **Endpoint**: `/comment/:id`
- **Body**:

```json
{
  "content": "your_content"
}
```

List comments

- **Method**: GET
- **Endpoint**: `/comment`

Comment comment by ID

- **Method**: GET
- **Endpoint**: `/comment/:id`

Update comment(protected)

- **Method**: PUT
- **Endpoint**: `/comment/:id`

```json
{
  "content": "your_updated_content"
}
```

Delete comment(protected)

- **Method**: DELETE
- **Endpoint**: `/comment/:id`
