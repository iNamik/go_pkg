package cmp_test

import (
	"github.com/iNamik/go_pkg/sort/cmp"
	"math/rand"
	"testing"
)

// TestIntEqual
func TestIntEqual(t *testing.T) {
	a := cmp.Int(1)
	b := cmp.Int(1)

	if a.Cmp(b) != cmp.EQ {
		t.Fatalf("%d <=> %d did not generate EQUAL result", a, b)
	}
}

// TestIntLessThan
func TestIntLessThan(t *testing.T) {
	a := cmp.Int(1)
	b := cmp.Int(2)

	if a.Cmp(b) != cmp.LT {
		t.Fatalf("%d <=> %d did not generate LESS_THAN result", a, b)
	}
}

// TestIntGreaterThan
func TestIntGreaterThan(t *testing.T) {
	a := cmp.Int(3)
	b := cmp.Int(2)

	if a.Cmp(b) != cmp.GT {
		t.Fatalf("%d <=> %d did not generate GREATER_THAN result", a, b)
	}
}

// TestStringEqual
func TestStringEqual(t *testing.T) {
	a := cmp.String("Apple")
	b := cmp.String("Apple")

	if a.Cmp(b) != cmp.EQ {
		t.Fatal("%s <=> %s did not generate EQUAL result", a, b)
	}
}

// TestStringLessThan
func TestStringLessThan(t *testing.T) {
	a := cmp.String("Apple")
	b := cmp.String("Bunny")

	if a.Cmp(b) != cmp.LT {
		t.Fatal("%s <=> %s did not generate LESS_THAN result", a, b)
	}
}

// TestStringGreaterThan
func TestStringGreaterThan(t *testing.T) {
	a := cmp.String("Dagger")
	b := cmp.String("Candy")

	if a.Cmp(b) != cmp.GT {
		t.Fatal("%s <=> %s did not generate GREATER_THAN result", a, b)
	}
}

func TestFunc1(t *testing.T) {
	a := cmp.Int(1)
	b := cmp.Int(2)
	f := func(c_ interface{}) int {
		c := c_.(cmp.Int)
		switch {
		case b < c:
			return cmp.LT
		case b > c:
			return cmp.GT
		default:
			return cmp.EQ
		}
	}

	var c_ cmp.T = cmp.TF(f)

	c, ok := c_.(cmp.T)
	if !ok {
		t.Fatalf("Function is not of type cmp.T")
	}

	if c.Cmp(a) != cmp.GT {
		t.Fatalf("%d <=> %d did not generate GREATER_THAN result", a, b)
	}
}

/**********************************************************************
 ** Benchmark Functions
 **********************************************************************/

const BM_SIZE = 10000 //000

// Benchmark1
func Benchmark1(b *testing.B) {
	l := rand.Perm(BM_SIZE + 1)
	var lt, eq, gt int = 0, 0, 0
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < BM_SIZE; i++ {
			switch cmp.Int(l[i]).Cmp(cmp.Int(l[i+1])) {
			case cmp.LT:
				lt++
			case cmp.EQ:
				eq++
			case cmp.GT:
				gt++
			}
		}
	}
}

// Benchmark2_PreCast
func Benchmark2_PreCast(b *testing.B) {
	l_ := rand.Perm(BM_SIZE + 1)
	l := make([]cmp.Int, BM_SIZE+1)
	for i, v := range l_ {
		l[i] = cmp.Int(v)
	}
	var lt, eq, gt int = 0, 0, 0
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < BM_SIZE; i++ {
			switch l[i].Cmp(l[i+1]) {
			case cmp.LT:
				lt++
			case cmp.EQ:
				eq++
			case cmp.GT:
				gt++
			}
		}
	}
}
