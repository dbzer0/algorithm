package queue

import (
	"sync"

	"github.com/dbzer0/algorithm/structure/dbllist"
)

type LIFO struct {
	q dbllist.List
	l sync.Mutex
}

func NewLIFO() *LIFO {
	return &LIFO{q: *dbllist.NewList()}
}

func (f *LIFO) Push(value interface{}) {
	f.l.Lock()
	f.q.PushBack(&dbllist.Item{Value: value})
	f.l.Unlock()
}

func (f *LIFO) Len() int32 {
	f.l.Lock()
	defer f.l.Unlock()
	return f.q.Len()
}

func (f *LIFO) Pop() interface{} {
	f.l.Lock()
	defer f.l.Unlock()
	return f.q.PopBack().Value
}
