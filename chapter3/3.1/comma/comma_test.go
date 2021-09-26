package comma

import (
	"strconv"
	"testing"
)

func TestComma(t *testing.T) {
	if comma("12345") != anotherComma("12345") {
		t.FailNow()
	}

	if comma("123456") != anotherComma("123456") {
		t.FailNow()
	}

	if comma("123") != anotherComma("123") {
		t.FailNow()
	}

	if comma("1") != anotherComma("1") {
		t.FailNow()
	}
}

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = comma(strconv.Itoa(i))
	}
}

func BenchmarkAnotherComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = anotherComma(strconv.Itoa(i))
	}
}
