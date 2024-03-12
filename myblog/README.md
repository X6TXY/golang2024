# Final Project Golang 2024 Spring: Go Blog with Comments and Authentication

This project is a simple blog application implemented in Go (Golang) that includes features such as CRUD operations for articles, comments, and user authentication.

## Author

<!-- - Toleu Bahauddin 22B030598 -->

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

## Database Relationships

### User - Post (One-to-Many)

A user can have many posts. This is represented by the `Posts` slice in the `User` struct, with a `foreignKey` tag pointing to `UserID` in the `Post` struct, indicating that multiple posts can belong to a single user.

### Post - Comment (One-to-Many)

A post can have many comments. This relationship is shown by the `Comments` slice in the `Post` struct, with a `foreignKey` tag pointing to `PostID` in the `Comment` struct, indicating that multiple comments can be associated with a single post.

### User - Comment (One-to-Many)

A user can have many comments. This is represented similarly to posts, where the `Comments` slice in the `User` struct has a `foreignKey` pointing to `UserID` in the `Comment` struct, indicating a user can author multiple comments.

### Post - User (Many-to-One)

Many posts belong to one user. This inverse relationship of the first point is represented by the `User` field in the `Post` struct, which points back to the owning user. The `foreignKey:UserID` indicates the association's direction.

### Comment - User (Many-to-One)

Many comments belong to one user. Similar to posts, this is the inverse relationship of the third point, where each `Comment` struct has a `User` field pointing back to the commenter.

### Post - PostLike (Many-to-Many)

Posts can have many likes from users, and users can like many posts. This is represented by the `PostLike` struct, which creates a many-to-many relationship between posts and users through a composite unique index on `UserID` and `PostID`.

### Comment - CommentLike (Many-to-Many)

Comments can have many likes, and users can like many comments, similarly managed by the `CommentLike` struct, indicating a many-to-many relationship between comments and users with a unique composite index on `UserID` and `CommentID`.

### User - User (Followers/Followings) (Many-to-Many)

This is a self-referencing many-to-many relationship where users can follow and be followed by many other users. The `Followers` and `Followings` slices in the `User` struct represent this relationship through a join table (`user_followers`). The `many2many` tag specifies the name of the join table, and `joinForeignKey`/`JoinReferences` tags specify the columns in the join table representing the following and follower users, respectively.

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

Get Users List (protected)

- **Method**: GET
- **Endpoint**: `/users`

Get User by ID (protected)

- **Method**: GET
- **Endpoint**: `/users/:id`

Delete User by ID (protected)

- **Method**: DELETE
- **Endpoint**: `/users/:id`

Follow User (protected)

- **Method**: POST
- **Endpoint**: `/users/:id/follow`

Unfollow User (protected)

- **Method**: POST
- **Endpoint**: `/users/:id/unfollow`

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

Like post(protected)

- **Method**: POST
- **Endpoint**: `/post/:id/like`

Unlike post(protected)

- **Method**: POST
- **Endpoint**: `/post/:id/unlike`

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

Like comment(protected)

- **Method**: POST
- **Endpoint**: `/comment/:id/like`

Unlike comment(protected)

- **Method**: POST
- **Endpoint**: `/comment/:id/unlike`
