# gofix examples

A catalog project for the analyzers shipped with `go tool fix`. Each
subdirectory under `analyzers/` is a standalone package whose source
intentionally triggers one analyzer. Files are kept in their "before" state so
the examples remain stable — use `-diff` to preview the rewrite without
mutating them.

> Toolchain: this project tracks **gotip / Go 1.27-devel**. Earlier toolchains
> are missing some analyzers below (notably `atomictypes`, `embedlit`,
> `errorsastype`, `slicesbackward`, `unsafefuncs`) and call `waitgroupgo` by
> its old name `waitgroup`. Verify with `go tool fix help`.

## Usage

```sh
go tool fix help                        # list every analyzer
go tool fix help minmax                 # details for one analyzer

go fix -diff ./...                      # preview every fix as a unified diff
go fix ./...                            # apply every fix in place
go fix -diff ./analyzers/minmax/        # preview just one analyzer's example
go fix ./analyzers/minmax/              # apply it
```

`go fix` does not accept per-analyzer flags like `-minmax` or `-minmax=false`
(those belong to the underlying `cmd/fix` tool, which the Go toolchain no
longer lets you invoke directly). To isolate a single analyzer, point `go fix`
at its package directory — each subfolder under `analyzers/` only contains
code that triggers one analyzer, so only that one can fire.

To re-run the demo after applying fixes, restore the originals from version
control (`git restore analyzers/`) or re-clone the project.

## Catalog

| Analyzer | What it rewrites |
|---|---|
| [`any`](analyzers/any) | `interface{}` → `any` |
| [`atomictypes`](analyzers/atomictypes) | `atomic.AddInt32(&x, …)` → `atomic.Int32.Add` |
| [`buildtag`](analyzers/buildtag) | malformed `//go:build` / `// +build` directives (vet-style, no fix) |
| [`embedlit`](analyzers/embedlit) | `T{U: U{x: 1}}` → `T{x: 1}` (Go 1.27) |
| [`errorsastype`](analyzers/errorsastype) | `errors.As(err, &v)` → `errors.AsType[T](err)` |
| [`forvar`](analyzers/forvar) | drop redundant `x := x` inside range loops (Go 1.22+) |
| [`hostport`](analyzers/hostport) | `fmt.Sprintf("%s:%d", host, port)` → `net.JoinHostPort` |
| [`inline`](analyzers/inline) | inline calls/refs marked with `//go:fix inline` |
| [`mapsloop`](analyzers/mapsloop) | manual map loops → `maps.Copy` / `maps.Keys` / `maps.Values` |
| [`minmax`](analyzers/minmax) | `if a > b { … } else { … }` → `min` / `max` |
| [`newexpr`](analyzers/newexpr) | `varOf(x) { return &x }` → `new(expr)` (Go 1.26) |
| [`omitzero`](analyzers/omitzero) | struct-typed `,omitempty` → `,omitzero` (Go 1.24+) |
| [`plusbuild`](analyzers/plusbuild) | drop obsolete `// +build` next to `//go:build` |
| [`rangeint`](analyzers/rangeint) | `for i := 0; i < n; i++` → `for i := range n` |
| [`reflecttypefor`](analyzers/reflecttypefor) | `reflect.TypeOf(T{})` → `reflect.TypeFor[T]()` |
| [`slicesbackward`](analyzers/slicesbackward) | reverse index loop → `slices.Backward` |
| [`slicescontains`](analyzers/slicescontains) | manual contains loop → `slices.Contains` |
| [`slicessort`](analyzers/slicessort) | `sort.Slice` → `slices.Sort` / `slices.SortFunc` |
| [`stditerators`](analyzers/stditerators) | `Len()` + `At(i)` loops → range over the type's iterator |
| [`stringsbuilder`](analyzers/stringsbuilder) | `s += …` in a loop → `strings.Builder` |
| [`stringscut`](analyzers/stringscut) | `strings.Index` + slicing → `strings.Cut` |
| [`stringscutprefix`](analyzers/stringscutprefix) | `HasPrefix` + `TrimPrefix` → `strings.CutPrefix` |
| [`stringsseq`](analyzers/stringsseq) | `range strings.Split` → `strings.SplitSeq` / `FieldsSeq` |
| [`testingcontext`](analyzers/testingcontext) | `context.WithCancel` in tests → `t.Context()` |
| [`unsafefuncs`](analyzers/unsafefuncs) | `unsafe.Pointer(uintptr(p) + uintptr(n))` → `unsafe.Add` |
| [`waitgroupgo`](analyzers/waitgroupgo) | `wg.Add(1)` + `defer wg.Done()` → `wg.Go` |
