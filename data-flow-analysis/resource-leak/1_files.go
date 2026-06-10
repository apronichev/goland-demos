package resource_leak

import (
	"math/rand"
	"os"
)

// simple case
func _(path string, flag bool) error {
	f, err := os.Open(path) // resource leak!
	if f == nil {
		return err
	}

	if flag {
		// the resource is not closed here
		return nil
	}

	// process file...

	f.Close()
	return nil
}

// helper function
func openFile() *os.File {
	f, err := os.Open("test")
	if err != nil {
		return nil
	}
	return f
}

func unknownCondition() bool {
	return rand.Int()%2 == 0
}

// reassignment
func _() {
	f := openFile() // resource leak!
	if unknownCondition() {
		// the first resource is not closed here
		f = openFile()
	}
	f.Close()
}
