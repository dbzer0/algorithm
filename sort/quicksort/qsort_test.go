package quicksort

import (
	"testing"
)

func TestSort(t *testing.T) {
	var list []int

	if len(Sort(list)) != 0 {
		t.Error("slice must be empty")
	}

	list = append(list, 5)
	if len(Sort(list)) != 1 || Sort(list)[0] != 5 {
		t.Error("slice with 1 element must be already sorted")
	}

	list = append(list, 0, -10, 2, 16, 8)
	r := Sort(list)

	if r[0] != -10 || r[1] != 0 || r[2] != 2 || r[3] != 5 || r[4] != 8 || r[5] != 16 {
		t.Error("sorting failed")
	}
}
