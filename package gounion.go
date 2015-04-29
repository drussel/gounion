package gounion

type Item int

// Keep track of sets under union operations.
// Initially, everything is in its own set.
// not concurrency safe
type UnionFind struct {
	edges []Item
}

// add an item and return a handle
func (this *UnionFind) Add() Item {
	ret := Item(len(this.edges))
	this.edges = append(this.edges, ret)
	return ret
}

// join the sets containing A and B
func (this *UnionFind) Union(a, b Item) {
	as := this.Get(a)
	bs := this.Get(b)
	this.edges[int(as)] = bs
}

// return a represetnative for the set
func (this *UnionFind) Get(i Item) Item {
	ip := this.edges[i]
	if ip == i {
		return i
	}
	ir := this.Get(ip)
	this.edges[i] = ir
	return ir
}

func (this *UnionFind) GetAll() (ret [][]Item) {
	retm := map[Item][]Item{}
	for i := 0; i < len(this.edges); i++ {
		ii := Item(i)
		s := this.Get(ii)
		retm[s] = append(retm[s], ii)
	}
	for _, s := range retm {
		ret = append(ret, s)
	}
	return
}
