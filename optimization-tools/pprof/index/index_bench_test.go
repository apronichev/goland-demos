package index_test

import (
	"testing"

	"github.com/goland-demos/optimization-tools/pprof/generator"
	"github.com/goland-demos/optimization-tools/pprof/index"
)

const benchUserCount = 100_000

func BenchmarkIndexing(b *testing.B) {
	users := generator.Generate(benchUserCount, 42)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		idx := index.New()
		idx.Build(users)
	}
}

func BenchmarkSearch(b *testing.B) {
	users := generator.Generate(benchUserCount, 42)
	idx := index.New()
	idx.Build(users)

	queries := []string{
		"john",
		"germany",
		"programming languages",
		"smith mail.com",
		"databases distributed",
		"maria spain",
		"interested",
		"kubernetes",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q := queries[i%len(queries)]
		_ = idx.Search(q)
	}
}
