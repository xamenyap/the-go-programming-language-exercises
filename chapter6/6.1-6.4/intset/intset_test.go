package intset

import (
	"fmt"
	"testing"
)

func TestIntSet_Len(t *testing.T) {
	s := IntSet{}
	s.Add(1)
	s.Add(2)
	s.Add(10)
	s.Add(50)
	s.Add(123)

	if s.Len() != 5 {
		t.FailNow()
	}

	s.Add(9)
	s.Add(11)
	if s.Len() != 7 {
		t.FailNow()
	}

	s.Remove(2)
	if s.Len() != 6 {
		t.FailNow()
	}
}

func TestIntSet_Remove(t *testing.T) {
	s := IntSet{}
	s.Add(1)
	s.Add(2)
	s.Add(10)
	s.Add(50)
	s.Add(123)

	s.Remove(50)
	if s.String() != "{1 2 10 123}" {
		fmt.Println(s.String())
		t.FailNow()
	}

	s.Remove(99)
	if s.String() != "{1 2 10 123}" {
		t.FailNow()
	}

	s.Remove(1)
	if s.String() != "{2 10 123}" {
		t.FailNow()
	}
}

func TestIntSet_Clear(t *testing.T) {
	s := IntSet{}
	s.Add(1)
	s.Add(2)
	s.Add(10)
	s.Add(50)
	s.Add(123)

	s.Clear()
	if s.String() != "{}" {
		t.FailNow()
	}
}

func TestIntSet_Copy(t *testing.T) {
	s := IntSet{}
	s.Add(1)
	s.Add(2)
	s.Add(10)
	s.Add(50)
	s.Add(123)

	cp := s.Copy()
	if cp.String() != "{1 2 10 50 123}" {
		t.FailNow()
	}
}

func TestIntSet_Elems(t *testing.T) {
	s := IntSet{}
	s.Add(1)
	s.Add(2)
	s.Add(10)
	s.Add(50)
	s.Add(123)

	elems := s.Elems()
	elemsString := fmt.Sprintf("%+v", elems)
	if elemsString != "[1 2 10 50 123]" {
		t.FailNow()
	}
}

func TestIntSet_IntersectWith(t *testing.T) {
	// same number of elements
	s1 := IntSet{}
	s1.Add(9000)
	s1.Add(800)
	s1.Add(70)
	s1.Add(6)

	t1 := IntSet{}
	t1.Add(9000)
	t1.Add(801)
	t1.Add(6)
	t1.Add(71)

	s1.IntersectWith(&t1)
	if s1.String() != "{6 9000}" {
		t.FailNow()
	}

	// s2 has more elements than t2
	s2 := IntSet{}
	s2.Add(9876)
	s2.Add(4321)
	s2.Add(25)
	s2.Add(99999)

	t2 := IntSet{}
	t2.Add(26)
	t2.Add(4321)

	s2.IntersectWith(&t2)
	if s2.String() != "{4321}" {
		t.FailNow()
	}

	// t3 has more elements that s3
	s3 := IntSet{}
	s3.Add(5)
	s3.Add(2121)
	s3.Add(99)

	t3 := IntSet{}
	t3.Add(98234)
	t3.Add(5)
	t3.Add(842)
	t3.Add(2121)
	t3.Add(2)

	s3.IntersectWith(&t3)
	if s3.String() != "{5 2121}" {
		t.FailNow()
	}

	// no intersect
	s4 := IntSet{}
	s4.Add(123)
	s4.Add(4444)
	s4.Add(987654321)

	t4 := IntSet{}
	t4.Add(987654320)
	t4.Add(124)

	s4.IntersectWith(&t4)
	if s4.String() != "{}" {
		t.FailNow()
	}
}

func TestIntSet_DifferenceWith(t *testing.T) {
	s1 := IntSet{}
	s1.Add(10)
	s1.Add(999)
	s1.Add(32145)
	s1.Add(622)
	s1.Add(50)

	t1 := IntSet{}
	t1.Add(32145)
	t1.Add(986341)
	t1.Add(9123)
	t1.Add(999)
	t1.Add(15)

	s1.DifferenceWith(&t1)
	if s1.String() != "{10 50 622}" {
		t.FailNow()
	}

	s2 := IntSet{}
	s2.Add(10)
	s2.Add(999)
	s2.Add(32145)
	s2.Add(622)
	s2.Add(50)

	t2 := IntSet{}
	t2.Add(99)
	t2.Add(15)
	t2.Add(32146)

	s2.DifferenceWith(&t2)
	if s2.String() != "{10 50 622 999 32145}" {
		t.FailNow()
	}

	s3 := IntSet{}
	s3.Add(10)
	s3.Add(999)

	t3 := IntSet{}
	t3.Add(999)
	t3.Add(15)
	t3.Add(32146)

	s3.DifferenceWith(&t3)
	if s3.String() != "{10}" {
		t.FailNow()
	}

	s4 := IntSet{}
	s4.Add(10)
	s4.Add(999)

	t4 := IntSet{}
	t4.Add(999)
	t4.Add(15987)
	t4.Add(10)

	s4.DifferenceWith(&t4)
	if s4.String() != "{}" {
		t.FailNow()
	}
}

func TestIntSet_SymmetricDifferenceWith(t *testing.T) {
	s1 := IntSet{}
	s1.Add(10)
	s1.Add(999)
	s1.Add(32145)

	t1 := IntSet{}
	t1.Add(999)
	t1.Add(986341)

	s1.SymmetricDifferenceWith(&t1)
	if s1.String() != "{10 32145 986341}" {
		fmt.Println(s1.String())
		t.FailNow()
	}

	s2 := IntSet{}
	s2.Add(2)
	s2.Add(32145)

	t2 := IntSet{}
	t2.Add(999)
	t2.Add(2)
	t2.Add(986341)

	s2.SymmetricDifferenceWith(&t2)
	if s2.String() != "{999 32145 986341}" {
		t.FailNow()
	}

	s3 := IntSet{}
	s3.Add(2)
	s3.Add(32145)
	s3.Add(854)

	t3 := IntSet{}
	t3.Add(32145)
	t3.Add(854)
	t3.Add(2)

	s3.SymmetricDifferenceWith(&t3)
	if s3.String() != "{}" {
		t.FailNow()
	}
}
