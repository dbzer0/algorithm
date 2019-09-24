package queue

import "testing"

func TestNewLIFO(t *testing.T) {
	lifo := NewLIFO()
	lifo.Push(1)
	lifo.Push(2)
	lifo.Push(3)

	i := 4
	for lifo.Len() != 0 {
		i--
		if lifo.Pop() != i {
			t.Errorf("invalid element must be %d", i)
		}
	}
}
