package gorthack_test

import (
	"testing"
	"unsafe"

	"github.com/zhouzuoji/gorthack"
)

type unsafePtr = unsafe.Pointer

func TestMain(m *testing.M) {
	m.Run()
}

type IAdder interface {
	Add(a, b int) int
}

type Calculator struct{}

func (self *Calculator) Add(a, b int) int {
	return a + b
}

type Adder func(a, b int) int
type PureAdder func(self unsafePtr, a, b int) int

func Benchmark_Add_Direct(b *testing.B) {
	c := Calculator{}
	for n := 0; n < b.N; n++ {
		c.Add(n, n)
	}
}

func Benchmark_Add_DirectFunc(b *testing.B) {
	c := Calculator{}
	f := c.Add
	for n := 0; n < b.N; n++ {
		f(n, n)
	}
}

func Benchmark_Add_DirectFuncCopy(b *testing.B) {
	c := Calculator{}
	var f Adder
	gorthack.CastFunc(c.Add, &f)
	for n := 0; n < b.N; n++ {
		f(n, n)
	}
}

func Benchmark_Add_ByInterface(b *testing.B) {
	c := Calculator{}
	var f IAdder = &c
	for n := 0; n < b.N; n++ {
		f.Add(n, n)
	}
}

func Benchmark_Add_RefelctFunc(b *testing.B) {
	c := Calculator{}
	var f Adder
	gorthack.CastMethod(&c, &f, "Add")
	for n := 0; n < b.N; n++ {
		f(n, n)
	}
}

func Benchmark_Add_RefelctPureFunc(b *testing.B) {
	c := Calculator{}
	var f PureAdder
	gorthack.MethodToPureFunc(&c, &f, "Add")
	for n := 0; n < b.N; n++ {
		f(unsafePtr(&c), n, n)
	}
}
