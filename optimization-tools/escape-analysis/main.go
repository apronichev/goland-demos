package main

import (
	"escape-analysis/subpkg"
	"fmt"
)

// Escape analysis demo

// variable does NOT escape (stays on stack)
func stackOnly() int {
	x := 42
	return x
}

// returning pointer causes variable to ESCAPE to heap
func escapeViaReturn() int {
	return 42

}

// storing pointer in slice causes ESCAPE
func escapeViaSlice() []*int {
	var result []*int
	for i := 0; i < 3; i++ {
		result = append(result, new(i))
	}
	return result
}

// main.go:24:30: .autotmp_3 escapes to heap

// closure captures variable, causes ESCAPE
func escapeViaClosure() func() int {
	x := 100
	return func() int {
		return x
	}
}

// value ESCAPES due to boxing
func escapeViaInterface() interface{} {
	x := 999
	return x
}

// arguments ESCAPE (interface{})
func escapeViaPrintln() {
	x := "hello"
	fmt.Println(x)
}

// large object may escape due to size
func bigArrayEscapes() *[10000]int {
	var arr [10000]int
	return &arr
}

// small object does NOT escape (inline optimization)
func smallStructNoEscape() int {
	type point struct{ x, y int }
	p := point{1, 2}
	return p.x + p.y
}

// Container holds a pointer to data
type Container struct {
	data *int
	a    byte
	abc  int
}

// pointer stored in struct causes ESCAPE
func escapeViaStruct() Container {
	return Container{data: new(42)}
}

// used in goroutine causes ESCAPE
func escapeViaGoroutine() {
	//x := 123
	go func() {
		fmt.Println("42")
	}()
}

// sent to channel causes ESCAPE
func escapeViaChannel(ch chan *int) {
	ch <- new(55)
}

// pointer parameter does NOT escape (only read)
func noEscapeParam(p *int) int {
	return *p + 1
}

// keys and values stored in map ESCAPE
func escapeViaMap() map[string]*int {
	m := make(map[string]*int)
	m["answer"] = new(42)
	return m
}

// newVsLocal - new() vs local variable escape behavior
func newVsLocal() (*int, int) {
	a := new(int)
	b := 10
	*a = 20
	return a, b
}

// recursiveEscape - recursion with pointers
func recursiveEscape(n int) *int {
	if n <= 0 {
		return new(0)
	}
	return recursiveEscape(n - 1)
}

// MyProcessor implements Processor interface from helpers.go
type MyProcessor struct {
	value int
}

func (p MyProcessor) Process() int {
	return p.value * 2
}

func main() {
	_ = stackOnly()
	_ = escapeViaReturn()
	y := escapeViaSlice()
	_ = escapeViaClosure()
	_ = escapeViaInterface()
	escapeViaPrintln()
	_ = bigArrayEscapes()
	_ = smallStructNoEscape()
	_ = escapeViaStruct()
	escapeViaGoroutine()

	print(y)

	ch := make(chan *int, 1)
	escapeViaChannel(ch)

	_ = noEscapeParam(new(5))

	_ = escapeViaMap()
	_, _ = newVsLocal()
	_ = recursiveEscape(3)

	// === Cross-file escape analysis examples ===

	// SmallFunc is inlined - compiler sees through the call
	// This allows better escape analysis at call site
	x := 10
	_ = SmallFunc(x) // x does NOT escape (SmallFunc is inlined)

	// BigFunc is NOT inlined - compiler can't see inside
	// This is more conservative
	//y := 20
	//_ = BigFunc(y) // y does NOT escape (passed by value)

	// LeakingPointer stores in global - causes escape at call site
	LeaksPointer(new(30)) // z ESCAPES because LeakingPointer stores it

	// ContainsPointer only reads - no escape
	_ = ContainsPointer(new(40)) // w does NOT escape (only read)

	// Interface causes boxing - value escapes
	proc := MyProcessor{value: 50}
	_ = ProcessAny(proc) // proc ESCAPES (boxed to interface)

	// RegisterProcessor stores in slice - can't devirtualize
	proc2 := MyProcessor{value: 55}
	RegisterProcessor(proc2) // proc2 ESCAPES (stored in interface slice)
	_ = RunAllProcessors()

	// Slice sink causes escape
	var slice []*int
	SliceSink(&slice, new(60)) // val1 ESCAPES (stored in slice)

	// Map sink causes escape
	m := make(map[string]*int)
	MapSink(m, "key", new(70)) // val2 ESCAPES (stored in map)

	// Return pointer causes escape
	_ = ReturnPointer(new(80)) // val3 ESCAPES (returned)

	// NoEscapeWrapper doesn't leak
	val4 := 90
	_ = NoEscapeWrapper(val4) // val4 does NOT escape

	_ = subpkg.SpillInt()
	a := subpkg.NoEscape(1)
	fmt.Println("%d", a)
}

type I interface {
	foobar()
}

type T struct{}

func (T) foo() {}

func bar(i I) {
	i.foobar()
}
