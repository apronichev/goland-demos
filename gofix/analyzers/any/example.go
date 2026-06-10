// `any`: replace `interface{}` with the predeclared `any` alias (Go 1.18+).
package any_

type Box struct {
	Value interface{}
}

func Identity(x interface{}) interface{} { return x }

func Map() map[string]interface{} { return nil }
