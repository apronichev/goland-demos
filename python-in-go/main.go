// main.go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func waitForServer(maxAttempts int) error {
	client := http.Client{
		Timeout: 1 * time.Second,
	}

	for i := 0; i < maxAttempts; i++ {
		_, err := client.Get("http://localhost:8000/api/hello/")
		if err == nil {
			return nil
		}
		fmt.Printf("Waiting for server to start (attempt %d/%d)...\n", i+1, maxAttempts)
		time.Sleep(1 * time.Second)
	}
	return fmt.Errorf("server failed to start after %d attempts", maxAttempts)
}

func startDjangoServer() (*exec.Cmd, error) {
	cmd := exec.Command("python3", "manage.py", "runserver")
	cmd.Dir = "pyapi"

	// Set up output to both capture it and show it in real-time
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Copy the current environment and set up Python paths
	env := os.Environ()
	env = append(env, "PYTHONPATH=/Users/jetbrains/myProjects/demos/.venv/lib/python3.13/site-packages:"+os.Getenv("PYTHONPATH"))
	cmd.Env = env

	err := cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("failed to start Django server: %v", err)
	}

	// Wait for the server to become available
	if err := waitForServer(10); err != nil {
		cmd.Process.Kill()
		return nil, err
	}

	fmt.Println("\nDjango server is running!")
	fmt.Println("Click the Stop button to stop the server and exit the program.\n")

	return cmd, nil
}

func callDjangoAPI() (string, error) {
	resp, err := http.Get("http://localhost:8000/api/hello/")
	if err != nil {
		return "", fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	return string(body), nil
}

func main() {
	// Start Django server
	cmd, err := startDjangoServer()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Set up signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Wait for interrupt signal
	<-sigChan

	fmt.Println("\nStopping Django server...")
	if err := cmd.Process.Kill(); err != nil {
		fmt.Printf("Error stopping server: %v\n", err)
	}
	cmd.Wait()
	fmt.Println("Server stopped")
}
