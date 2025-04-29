package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func runPythonScript() (string, error) {
	cmd := exec.Command("python3", "script.py")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func main() {
	// Run the Python script
	output, err := runPythonScript()
	if err != nil {
		fmt.Printf("Error running Python script: %v\n", err)
		return
	}
	fmt.Printf("Output from Python: %s\n", output)
}
