# Struct layout optimisation in GoLand

The order you declare struct fields in changes how much memory the struct uses.
Go keeps each field on its alignment boundary, so a small field in front of a
larger one leaves a gap that the compiler fills with padding. Put the fields in
a better order and the gaps go away, with no change to what the struct holds or
how it behaves.

GoLand finds structs that are laid out this way and reorders them for you. It's
part of the Go Performance Optimisation tool window, next to the profiler and
escape analysis.

## In short

Four beats, one thing said and one thing shown in each:

1. **The idea.** A struct's size depends on the order of its fields, because of
   alignment padding. *Show* `go run .`: `Event` is 32 bytes, `OptimizedEvent` is 24.
2. **The panel.** Analyse a scope in the Go Performance Optimisation panel and let
   GoLand reorder a struct. *Show* the tree of suboptimal structs and the
   before/after grid, then click Update.
3. **The editor.** Turn the inspection on from the panel and GoLand flags structs
   inline. *Show* the hover on `Event`, then click Update struct.
4. **The proof.** *Show* `go run .` again: `Event` is now 24.

The rest of this file is the same four beats in full.

## The same fields, two sizes

These two structs have the same four fields in a different order:

```go
type Event struct {
    OK   bool
    ID   int64
    Done bool
    Time int64
}

type OptimizedEvent struct {
    ID   int64
    Time int64
    OK   bool
    Done bool
}
```

Run the program and you can see what the order costs:

```
$ go run .
Event:          32 bytes
OptimizedEvent: 24 bytes
```

`ID` and `Time` are `int64`s, and an `int64` has to start on an 8-byte boundary.
In `Event` each one sits behind a `bool`, so the compiler inserts padding to
reach the boundary. `OptimizedEvent` puts the two `int64`s first, so there's
nothing to pad around. Same data, 8 bytes (about a quarter) smaller.

**Demo:** `go run .` → 32 vs 24

## Finding it in GoLand

You don't reorder fields by hand. GoLand finds the structs worth fixing and
reorders them. The analysis is off until you turn it on, so the place to start
is the panel.

### The Go Performance Optimisation panel

Open the panel and go to the Structs optimisation tab. Pick what to analyse (the
whole project, uncommitted files, a single file, or a custom scope) and click
Analyze.

GoLand lists the suboptimal structs it found, grouped by package and file, so you
can look through a whole project's worth in one place. Select a file or a single
finding and the right side shows the layout before and after, byte by byte: gray
squares are bytes the struct uses, red squares are padding the field order forces.
Click Update and the struct is reordered in your source, with the red gone. You
can also reorder every struct at once with Update all structs, and even on a big
project that only takes seconds.

The starting view also has an "Enable struct optimization by default" checkbox.
Turn it on and GoLand starts flagging these structs in the editor too, without
opening Settings.

**Demo:** pick scope → Analyze → select `Event` → Update → undo (Cmd+Z / git rollback) to revert

### The editor inspection

Once it's on, GoLand underlines the name of any struct that could be smaller as
you work. Hover over the underlined name:

> Struct 'Event' might be suboptimal: 8 bytes wasted (~25%)

Click Update struct in the popup and GoLand reorders the fields: the same fix as
the panel, right where you write the code.

**Demo:** tick *Enable struct optimization by default* → `Event` is now underlined → hover → Update struct → `go run .` again → 24

## Why it's opt-in

Some structs keep their field order on purpose, so GoLand never reorders on its
own. The inspection stays off until you turn it on, and the panel only changes a
struct when you click Update.
