package queue

import (
	"testing"
)

func TestNewFIFO(t *testing.T) {
	fifo := NewFIFO()
	fifo.Push(1)
	fifo.Push(2)
	fifo.Push(3)

	i := 0
	for fifo.Len() != 0 {
		i++
		if fifo.Pop() != i {
			t.Errorf("invalid element must be %d", i)
		}
	}
}
