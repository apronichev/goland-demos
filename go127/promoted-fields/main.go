package main

import (
	"encoding/json"
	"fmt"
)

type AuditInfo struct {
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

type Document struct {
	AuditInfo
	Name string `json:"name"`
	Path string `json:"path"`
}

func main() {
	// Go 1.27: a key in a struct literal may be any valid field selector,
	// including fields promoted from an embedded struct. So you can set the
	// promoted fields directly in the outer literal.
	doc := Document{
		CreatedBy: "alice", // promoted from AuditInfo
		UpdatedBy: "bob",   // promoted from AuditInfo
		Name:      "report.pdf",
		Path:      "/documents/report.pdf",
	}

	// Before 1.27 you had to construct the embedded struct explicitly:
	//
	//	doc := Document{
	//		AuditInfo: AuditInfo{CreatedBy: "alice", UpdatedBy: "bob"},
	//		Name:      "report.pdf",
	//		Path:      "/documents/report.pdf",
	//	}
	//
	// Note: you may not mix a promoted field with the embedded field that
	// promotes it, and pointer-embedded paths are not supported.

	out, _ := json.MarshalIndent(doc, "", "  ")
	fmt.Println(string(out))
}
