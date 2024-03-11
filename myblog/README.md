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
