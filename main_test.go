package main

import (
	"fmt"
	"testing"
)

type snode struct{ string }

func (node snode) Data() interface{} { return node.string }

func openqExample() []Node {
	openq := new(OpenQS)
	openq.Init()
	openq.Add(snode{"foo"}, 10)
	openq.Add(snode{"five"}, 5)
	openq.Add(snode{"foo"}, 3)
	openq.Add(snode{"two"}, 2)
	res := make([]Node, 0, 3)
	for openq.Len() > 0 {
		vptr := openq.Pop()
		res = append(res, *vptr)
	}
	return res
}

func ExampleOpenqOps() {
	fmt.Print(openqExample())
	// Output: [{two} {foo} {five}]
}

func BenchmarkOpenqOps(b *testing.B) {
	for n := 0; n < b.N; n++ {
		openqExample()
	}
}

var e = Env{Ion}

func TestEnvIsComplete(t *testing.T) {
	MakeEnv()
	sector := "enioar"
	bbox := envbbox[sector]
	for x := 0; x <= bbox.w; x++ {
		for y := 0; y <= bbox.h; y++ {
			_, ok := envdata[Coord{sector, x, y}]
			if !ok {
				t.Errorf("no key for (%s,%d,%d)", sector, x, y)
			}
		}
	}
}

func ExampleNbs() {
	p := func(x, y int) {
		fmt.Println(e.Nbs(&Coord{"enioar", x, y}))
	}
	p(0, 0)
	p(0, 12)
	p(20, 0)
	p(20, 12)
	p(1, 1)
	p(17, 6)
	p(10, 12)
	p(20, 8)
	p(7, 9)
	// Output:
	// []
	// []
	// []
	// []
	// [{enioar 2 1} {enioar 2 2} {enioar 1 2} {enioar 0 2}]
	// [{enioar 16 5} {enioar 18 7} {enioar 17 7} {enioar 16 7} {enioar 16 6}]
	// [{enioar 9 11} {enioar 10 11} {enioar 11 11} {enioar 11 12}]
	// [{enioar 19 7} {enioar 20 7} {enioar 19 8}]
	// []
}

func BenchmarkNbs(b *testing.B) {
	for n := 0; n < b.N; n++ {
		e.Nbs(&Coord{"enioar", 10, 10})
	}
}

func ExampleAstar() {
	fmt.Print(Astar(e, &Coord{"enioar", 1, 1}, &Coord{"enioar", 20, 7}))
	// Output: []
}