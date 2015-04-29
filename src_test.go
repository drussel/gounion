package gounion

import "testing"

func TestOneSet(t *testing.T) {
	uf := UnionFind{}
	for i := 0; i < 10; i++ {
		uf.Add()
	}
	for i := 9; i > 0; i-- {
		uf.Union(Item(i-1), Item(i))
	}
	if len(uf.GetAll()) != 1 {
		t.Error("Bad sets", uf.GetAll())
	}
}

func TestTwoSet(t *testing.T) {
	uf := UnionFind{}
	for i := 0; i < 10; i++ {
		uf.Add()
	}
	for i := 1; i < 5; i++ {
		uf.Union(Item(i-1), Item(i))
	}
	for i := 6; i < 10; i++ {
		uf.Union(Item(i-1), Item(i))
	}
	if len(uf.GetAll()) != 2 {
		t.Error("Bad sets", uf.GetAll())
	}
}
