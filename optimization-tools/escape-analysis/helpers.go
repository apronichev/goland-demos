package main

// helpers.go - demonstrates cross-file escape analysis effects

// SmallFunc is small enough to be inlined into callers in main.go
// When inlined, the escape analysis sees through the call
func SmallFunc(x int) int {
	return x * 2
}

// BigFunc is too big to inline (has loop), so callers can't see inside
// This affects escape decisions in callers
func wwBigFunc(x int) int {
	result := 0
	for i := 0; i < x; i++ {

	}
	return result
}

// LeaksPointer stores the pointer somewhere, causing caller's value to escape
var globalPtr *int

func LeaksPointer(p *int) {
	globalPtr = p // p escapes because it's stored in global
}

// ContainsPointer only reads the pointer, so caller's value may not escape
func ContainsPointer(p *int) int {
	return *p + 1 // p does not escape, just dereferenced
}

// Processor is an interface - passing concrete types to it causes boxing/escape
type Processor interface {
	Process() int
}

// ProcessAny takes interface - concrete values passed here will escape
func ProcessAny(p Processor) int {
	return p.Process()
}

// ProcessorRegistry stores processors - prevents devirtualization
var processorRegistry []Processor

// RegisterProcessor adds to registry - forces escape
func RegisterProcessor(p Processor) {
	//processorRegistry = append(processorRegistry, p)
}

// RunAllProcessors runs all - can't devirtualize
func RunAllProcessors() int {
	sum := 0
	for _, p := range processorRegistry {
		sum += p.Process()
	}
	return sum
}

// SliceSink appends to external slice - causes escape
func SliceSink(slice *[]*int, val *int) {
	*slice = append(*slice, val) // val escapes to heap
}

// MapSink puts into map - causes escape
func MapSink(m map[string]*int, key string, val *int) {
	m[key] = val // val escapes to heap
}

// ChannelSink sends to channel - causes escape
func ChannelSink(ch chan<- *int, val *int) {
	ch <- val // val escapes to heap
}

// ReturnPointer returns its argument - escape depends on caller
func ReturnPointer(p *int) *int {
	return p // p escapes to heap (returned)
}

// NoEscapeWrapper wraps value but doesn't leak it
func NoEscapeWrapper(x int) int {
	p := &x         // x does not escape
	return *p + 100 // just read, not stored
}
