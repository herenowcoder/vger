package main

import "container/heap"

type qitem struct {
	value    Node
	priority Cost
	index    int
	t        int // 'timestamp' so that items added earlier have it lower
}
type queue []*qitem

type OpenQS struct {
	q        queue
	m        map[Node]*qitem
	tcounter int
}

// heap.Interface impl

func (q queue) Len() int { return len(q) }

func (q queue) Less(i, j int) bool {
	pri1, pri2 := q[i].priority, q[j].priority
	if pri1 == pri2 {
		return q[i].t < q[j].t
	}
	return pri1 < pri2
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *queue) Push(obj interface{}) {
	x := obj.(*qitem)
	x.index = len(*q)
	*q = append(*q, x)
}

func (q *queue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}

// external API for OpenQS

func (qs *OpenQS) Init() {
	qs.q = make(queue, 0, 10)
	qs.m = make(map[Node]*qitem)
}

func (qs *OpenQS) Add(v Node, priority Cost) {
	qs.tcounter++
	x := &qitem{v, priority, -1, qs.tcounter}
	heap.Push(&qs.q, x)
	qs.m[v] = x
}

func (qs *OpenQS) Update(v Node, newPriority Cost) {
	x, ok := qs.m[v]
	if !ok {
		panic(v)
	}
	qs.tcounter++
	x.priority = newPriority
	x.t = qs.tcounter
	heap.Fix(&qs.q, x.index)
}

func (qs *OpenQS) Pop() Node {
	v := heap.Pop(&qs.q).(*qitem).value
	delete(qs.m, v)
	return v
}

func (qs *OpenQS) Len() int { return len(qs.m) }

func (qs *OpenQS) Contains(v Node) bool {
	_, ok := qs.m[v]
	return ok
}
