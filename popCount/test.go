
// util for testing
package main

import (
	"strconv"
	"testing"
	"runtime"
	"reflect"
)

func Funcname(f func(uint64) int) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func TestPop (t testing.TB, f func (uint64) int) {
	var cases = []struct {
		b uint64
		n int
	}{
		{ 0, 0 },
		{ 1, 1 },
		{ 3, 2 },
		{ 127, 7 },
		{ 255, 8 },
		{ 128, 1 },
	}
	for _, c := range cases {
		i := f(c.b)
		if c.n != i {
			t.Errorf("%v: %v - %v != %v", Funcname(f), c.b, c.n, i)
		}
	}
}

func bench(b *testing.B, name string, f func(*testing.B)) {
	benchmarks := []int{ 100000, 1_000_000, }
	for _, i := range benchmarks {
		b.Run(strconv.Itoa(i) + " " + name, func(b *testing.B) {
			for range i {
				f(b)
			}
		})
	}
}

func BenchPop(b *testing.B, f func(uint64) int) {
	for range b.N {
		bench(b, Funcname(f), func(b *testing.B) {
			TestPop(b, f)
		})
	}
}

