### Verify the gRPC Client Requests

Make sure your `requests.http` file is correctly set up:

#### Create requests.http

In `requests.http`, insert the following examples. You can run them only when the server is running.

```http
### Get User by ID
GRPC http://localhost:50051 user.UserService/GetUser
Content-Type: application/grpc

{
  "id": 1
}

###

### Create User
GRPC http://localhost:50051 user.UserService/CreateUser
Content-Type: application/grpc

{
  "name": "Alice",
  "email": "alice@example.com"
}

###

### Update User
GRPC http://localhost:50051 user.UserService/UpdateUser
Content-Type: application/grpc

{
  "id": 1,
  "name": "Alice Updated",
  "email": "alice.updated@example.com"
}

###

### Delete User
GRPC http://localhost:50051 user.UserService/DeleteUser
Content-Type: application/grpc

{
  "id": 1
}
```