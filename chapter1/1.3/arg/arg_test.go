package arg

import (
	"testing"
)

func generateArgs(num int) []string {
	args := make([]string, 0)
	for i := 0; i < num; i++ {
		args = append(args, "foo")
	}

	return args
}

func BenchmarkArgSimpleShort(b *testing.B) {
	args := generateArgs(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		simple(args)
	}
}

func BenchmarkArgEfficientShort(b *testing.B) {
	args := generateArgs(100)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		efficient(args)
	}
}

func BenchmarkSimpleLong(b *testing.B) {
	args := generateArgs(10000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		simple(args)
	}
}

func BenchmarkArgEfficientLong(b *testing.B) {
	args := generateArgs(10000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		efficient(args)
	}
}
