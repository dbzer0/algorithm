package queue

import (
	"sync"

	"github.com/dbzer0/algorithm/structure/dbllist"
)

type FIFO struct {
	q dbllist.List
	l sync.Mutex
}

func NewFIFO() *FIFO {
	return &FIFO{q: *dbllist.NewList()}
}

func (f *FIFO) Push(value interface{}) {
	f.l.Lock()
	f.q.PushFront(&dbllist.Item{Value: value})
	f.l.Unlock()
}

func (f *FIFO) Len() int32 {
	f.l.Lock()
	defer f.l.Unlock()
	return f.q.Len()
}

func (f *FIFO) Pop() interface{} {
	f.l.Lock()
	defer f.l.Unlock()
	return f.q.PopBack().Value
}
