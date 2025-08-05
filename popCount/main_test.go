package main

import (
	"testing"
)

func TestMypop(t *testing.T) {
	TestPop(t, mypop)
}

func BenchmarkMypop(b *testing.B) {
	BenchPop(b, mypop)
}

func TestPreload1(t *testing.T) {
	TestPop(t, preload1)
}

func BenchmarkPreload1(b *testing.B) {
	BenchPop(b, preload1)
}

func TestPreload2(t *testing.T) {
	TestPop(t, preload2)
}

func BenchmarkPreload2(b *testing.B) {
	BenchPop(b, preload2)
}

func TestClear(t *testing.T) {
	TestPop(t, clearbit)
}

func BenchmarkClear(b *testing.B) {
	BenchPop(b, clearbit)
}
