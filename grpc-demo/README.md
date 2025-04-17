# Running the gRPC demo

1. Run the server in `grpc-demo/main.go` (use the `main` function). The server listens to requests on two ports: 50051 for gRPC and 8080 for HTTP. It also serves the OpenAPI spec `grpc.yaml`. 
2. Click **View | Tool Windows | Endpoints**.
3. In the **Endpoints** window, click **Add OpenAPI specifications**.
4. In the **Remote Specifications** table, click **Add** and paste the following link: `http://localhost:8080/grpc.yaml`. 
5. Click **OK** or **Apply Changes**. Endpoints should be automatically detected and displayed in the **Endpoints** window.

## Running HTTP methods from the Endpoints tool window

1. Click **View | Tool Windows | Endpoints**.
2. Click the `POST` method. On the **HTTP Client** tab, type data into `name` and `email` fields.
3. Click **Submit Request**.

Alternatively,

1. Double-click any method in the **Endpoints** tool window. The OpenAPI specification will open.
2. In the WYSIWYG presentation of the specification, you can run requests using preconfigured example data. 


## Running gRPC requests

1. Open `user.proto` file in `grpc-demo/proto/`.
2. In `UserService` definition, click gutter icons against service methods (for example, against the line `rpc CreateUser (CreateUserRequest) returns (CreateUserResponse);`).
3. Change the port to `50051`.
4. Paste the body for the request. For the `CreateUser` method, it might be:
```
GRPC http://localhost:50051 user.UserService/CreateUser
Content-Type: application/grpc

{
  "name": "Alice",
  "email": "alice@example.com"
}
```

## Examples for body of other gRPC requests
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