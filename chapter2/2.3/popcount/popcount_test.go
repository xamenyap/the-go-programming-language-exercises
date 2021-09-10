package popcount

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PopCount(uint64(i))
	}
}

func BenchmarkLoopPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = LoopPopCount(uint64(i))
	}
}
