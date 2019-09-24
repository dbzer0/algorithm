package btree

import "testing"

func TestNewBTree(t *testing.T) {
	bt := NewBTree()
	bt.Insert(5, "котик")
	bt.Insert(2, "ыыпка")
	bt.Insert(7, "кооошка")
	bt.Insert(6, "хомячооок")
	bt.Insert(4, "воробей")

	if bt.Root.Key != 5 {
		t.Errorf("Root.Key (%d) must be = 5", bt.Root.Key)
	}
	if bt.Root.Value != "котик" {
		t.Errorf("Root.Value (%d) must be = котик", bt.Root.Value)
	}

	if bt.Root.Right.Key != 7 {
		t.Errorf("Root.Right.Key (%d) must be = 7", bt.Root.Right.Key)
	}

	if bt.Root.Left.Key != 2 {
		t.Errorf("Root.Left.Key (%d) must be = 2", bt.Root.Left.Key)
	}

	bt.Remove(-1)
	bt.Remove(5)
	bt.Remove(5)
	bt.Remove(5)

	if bt.Root.Key != 6 {
		t.Errorf("Root.Key (%d) must be = 6", bt.Root.Key)
	}
	if bt.Root.Value != "хомячооок" {
		t.Errorf("Root.Value (%d) must be = хомячооок", bt.Root.Value)
	}

	if bt.Root.Right.Key != 7 {
		t.Errorf("Root.Right.Key (%d) must be = 7", bt.Root.Right.Key)
	}

	if bt.Root.Left.Key != 2 {
		t.Errorf("Root.Left.Key (%d) must be = 2", bt.Root.Left.Key)
	}

	if bt.Has(1) == true {
		t.Error("Has(1) must be false")
	}
	if bt.Has(2) == false {
		t.Error("Has(2) must be true")
	}

	if bt.Search(9) != nil {
		t.Error("Search(9) must be nil")
	}
}
