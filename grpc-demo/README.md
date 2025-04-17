# Running the gRPC Demo

1. Run the server in `grpc-demo/main.go` using the `main` function.  
   The server listens for requests on two ports:
    - `50051` for gRPC
    - `8080` for HTTP and serves the OpenAPI spec `grpc.yaml`

2. Go to **View | Tool Windows | Endpoints**.
3. In the **Endpoints** tool window, click **Add OpenAPI Specifications**.
4. In the **Remote Specifications** table, click **Add**, then paste the following URL:  
   `http://localhost:8080/grpc.yaml`
5. Click **OK** or **Apply Changes**.  
   The endpoints will appear automatically in the **Endpoints** tool window.

---

## Running HTTP Methods from the Endpoints Tool Window

1. Go to **View | Tool Windows | Endpoints**.
2. Click the `POST` method. In the **HTTP Client** tab, enter values for the `name` and `email` fields.
3. Click **Submit Request**.

**Alternatively:**

1. Double-click any method in the **Endpoints** tool window to open the OpenAPI specification.
2. In the visual editor (WYSIWYG), you can run requests using example data.

---

## Running gRPC Requests

1. Open `user.proto` in `grpc-demo/proto/`.
2. In the `UserService` definition, click the gutter icon next to a method  
   (e.g., `rpc CreateUser (CreateUserRequest) returns (CreateUserResponse);`).
3. Set the port to `50051`.
4. Paste the request body. For example:

    ```http
    GRPC http://localhost:50051 user.UserService/CreateUser
    Content-Type: application/grpc

    {
      "name": "Alice",
      "email": "alice@example.com"
    }
    ```

---

## Examples of gRPC Request Bodies

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