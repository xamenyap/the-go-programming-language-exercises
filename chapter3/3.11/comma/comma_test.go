package comma

import (
	"testing"
)

func TestComma(t *testing.T) {
	if comma("-123456789") != "-123,456,789" {
		t.FailNow()
	}

	if comma("12345678") != "12,345,678" {
		t.FailNow()
	}

	if comma("12378.9876") != "12,378.9876" {
		t.FailNow()
	}
}
