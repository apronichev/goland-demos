// `unsafefuncs`: hand-written pointer arithmetic of the form
// `unsafe.Pointer(uintptr(p) + uintptr(n))` is brittle (the uintptr can be
// observed by the GC between the cast and the next use). `unsafe.Add` (Go
// 1.17+) is the safe, intent-revealing replacement.
package unsafefuncs

import "unsafe"

func Offset(p unsafe.Pointer, n int) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + uintptr(n))
}
