# Go 1.25: New go vet Analyzers

The go vet command includes new analyzers:

* `waitgroup`, which reports misplaced calls to `sync.WaitGroup.Add`; and
* `hostport`, which reports uses of `fmt.Sprintf("%s:%d", host, port)` to construct addresses for `net.Dial`, as these will not work with IPv6; instead it suggests using `net.JoinHostPort`.
