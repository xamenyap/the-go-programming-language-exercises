package comma

import (
	"strconv"
	"testing"
)

func TestComma(t *testing.T) {
	if comma("12345") != commaUsingBuffer("12345") {
		t.FailNow()
	}

	if comma("123456") != commaUsingBuffer("123456") {
		t.FailNow()
	}

	if comma("123") != commaUsingBuffer("123") {
		t.FailNow()
	}

	if comma("1") != commaUsingBuffer("1") {
		t.FailNow()
	}
}

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = comma(strconv.Itoa(i))
	}
}

func BenchmarkCommaUsingBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = commaUsingBuffer(strconv.Itoa(i))
	}
}
