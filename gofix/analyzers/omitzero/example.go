// `omitzero`: on struct-typed fields, `,omitempty` has no effect — a struct is
// never the JSON zero value `""`/`0`/`nil`. `,omitzero` (Go 1.24+) actually
// omits the field when the struct value is zero.
//
// Note: behavior change. Read the analyzer help before applying blindly.
package omitzero

import "time"

type Event struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Source    Source    `json:"source,omitempty"`
}

type Source struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
