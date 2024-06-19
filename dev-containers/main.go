package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

// todo Follow the steps below:
// 1. Ensure that the Docker service is up and running.
// 2. Open the goDevContainer directory in GoLand 2023.2 or higher.
// 3. In the Project tool window, open the '.devcontainer' folder and double-click the 'devcontainer.json' file.
// 4. Click the gutter icon and select Create Dev Container and Mount Sources.
// 5. After the development container is created, select GoLand from the drop-down list of IDE instances (for example, GoLand 2023.2 (232.8660.185) | download latest).
// 6. Click the Continue button. The IDE will be downloaded as a backend and the Client is opened. The Client is connected to the backend.

func handle(w http.ResponseWriter, r *http.Request) {
	os := runtime.GOOS
	fmt.Fprintf(w, "Hello World from [%s] container!\n", os)
}

func main() {
	http.HandleFunc("/", handle)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
