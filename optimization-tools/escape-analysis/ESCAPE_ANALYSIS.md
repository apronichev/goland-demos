# Escape Analysis in Go: why and how

## What it is

**Escape analysis** is a compiler pass in Go that decides, for every variable,
whether it can live on the **stack** (cheap: freed automatically when the
function returns, no GC involvement) or must *escape* to the **heap**.

A variable escapes when the compiler cannot prove it stops being needed by the
time the function returns — for example, a pointer to it outlives the call.

## Why it matters

Every heap allocation is work at runtime **and** future work for the garbage
collector. Reducing escapes gives you:

- **fewer allocations** → less GC pressure, shorter and rarer GC pauses;
- **lower latency and steadier p99** in services under load;
- **higher throughput** in hot loops;
- **better data locality** (values sit together instead of scattered on the heap).

This is not a cosmetic micro-optimization: on hot paths the difference is often
**multiplicative** (see the example below — 1000 allocations vs 0).

## How to inspect it

```bash
# basic escape-analysis decisions
go build -gcflags='-m' ./...

# more detail: reasons and flow chains, plus inlining decisions
go build -gcflags='-m=2' ./...
```

How to read the output:

- `&T{...} escapes to heap` / `moved to heap: x` — the variable went to the heap;
- `... does not escape` — it stayed on the stack (good);
- `can inline f` / `inlining call to f` — the function was inlined; inlining lets
  escape analysis "see through" the call and often keep arguments on the stack.

In GoLand the same data is available via **Go Performance Optimization →
Escape Analysis** (editor highlighting + results tree), without parsing
`-gcflags` text by hand.

Measure the effect with a benchmark and `-benchmem`:

```bash
go test -bench=. -benchmem ./...
# look at the allocs/op column
```

## Why GoLand analyzes it better

Raw `go build -gcflags='-m'` works, but it is a flat wall of text you have to
grep and cross-reference by hand. GoLand runs the same compiler diagnostics
(`-m=2` plus the compiler's structured JSON output) and turns them into
something you can actually act on:

- **Mapped onto the code, not the terminal.** Each decision is shown as a
  gutter icon and inline highlight at the exact line/column, so you see *which*
  expression escapes without matching `file:line:col` strings yourself.
- **Categorized and aggregated.** Results are grouped by kind — *escapes to
  heap*, *moved to heap*, *leaking param*, *can inline*, *inlining call* — and
  by function/file, with counts, instead of one undifferentiated stream.
- **Flow chains you can follow.** For an escape, the compiler's *why* (the
  `escflow` chain: `~r0 ← &x`, `from &x (address-of)`, `from return &x`) is
  rendered as a navigable explanation, not buried lines you must reassemble.
- **Inlining shown together with escapes.** Because inlining is what lets the
  analysis "see through" calls, GoLand surfaces `can inline` / `inlining call`
  alongside escape decisions — so you understand *why* something stayed on or
  left the stack.
- **One-click navigation.** Jump from a finding straight to the source; filter,
  group, and search the results; rerun after an edit — no rebuild-and-re-grep
  loop.
- **Same source of truth.** It is the compiler's own data, so what you see
  matches `-gcflags` exactly; GoLand only makes it readable and navigable.

In short: the raw flag tells you *what* the compiler decided; GoLand shows
*where* in your code and *why*, and lets you act on it without leaving the editor.

## Common reasons a variable escapes

| Reason | Example |
|---|---|
| Returning a pointer to a local | `func f() *T { x := T{}; return &x }` |
| Boxing a value into an interface | `fmt.Println(x)`, `var any interface{} = x` |
| Capturing a variable in a closure that outlives the function | `return func() int { return x }` |
| Storing a pointer in a slice / map / channel | `s = append(s, &x)`, `m[k] = &x`, `ch <- &x` |
| Starting a goroutine that captures the variable | `go func() { use(x) }()` |
| Object too large for the stack | a big `[N]T` array |
| A call that cannot be inlined | large function → the compiler stays conservative |

## Patterns to remove escapes

- **Slice of values instead of slice of pointers**: `[]T` over `[]*T` when
  elements need not be shared by pointer (see the `allocreduction` package).
- **Fill a caller-provided buffer/struct** instead of returning a fresh pointer
  (`func fill(dst *T)` over `func make() *T`).
- **Concrete types instead of interfaces** on the hot path (avoid boxing).
- **Reuse buffers** (`sync.Pool`, pre-sized slices with `cap`).
- **Keep functions small enough to inline.**

## Real-world use cases

- **Hot loops and batch processing** — parsing lines/logs, serialization,
  numeric simulations: eliminate per-element allocations.
- **High-load services** — a request handler shouldn't litter the heap on every
  request; fewer GC pauses → steadier latency.
- **Libraries** (codecs, marshalers, HTTP routers) — escape analysis helps keep
  a "zero-allocation" path that is then locked in with a benchmark.
- **Profiling/optimization** — after `pprof` flags allocations, escape analysis
  explains *why* an object is on the heap and points to the fix.

## Worked example in this project

The [`allocreduction`](allocreduction) package is a minimal but realistic
change (`[]*Particle` → `[]Particle`) that removes almost all allocations:

```bash
go build -gcflags='-m=2' ./allocreduction 
go test -bench=. -benchmem ./allocreduction
```

Measured result (N = 1000 elements):

| Version | ns/op | B/op | allocs/op |
|---|---|---|---|
| `BuildSlow` (`[]*Particle`) | ~11960 | 24000 | **1000** |
| `BuildFast` (`[]Particle`)  | ~1380  | 0     | **0**    |

One mechanical change — an order-of-magnitude drop in allocations and ~8× in time.
