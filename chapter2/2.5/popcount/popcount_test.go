package popcount

import (
	"testing"
)

func TestAlternatePopCount(t *testing.T) {
	var i uint64 = 0
	for ; i < 1000; i++ {
		val1 := PopCount(i)
		val2 := AlternatePopCount(i)

		if val1 != val2 {
			t.Fatal("expect ", val1, ", got ", val2)
		}
	}
}

func BenchmarkAlternatePopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = AlternatePopCount(uint64(i))
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PopCount(uint64(i))
	}
}
