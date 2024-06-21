
## Setup Instructions

1. Initialize the Go module:
    ```sh
    go mod init web-notebook
    ```

2. Install the necessary Go packages:
    ```sh
    go get modernc.org/sqlite
    ```

3. Run the Go server:
    ```sh
    go run server/main.go
    ```
    This creates an SQLite database that you can attach to the **Database** tool window.

4. Open [http://127.0.0.1:8080/](http://127.0.0.1:8080/) in a web browser to interact with the application.
