package btree

import (
	"reflect"
	"sync"
)

// может быть любым типом
type Item int

type Node struct {
	Value interface{}
	Key   Item
	Left  *Node
	Right *Node
}

type BTree struct {
	Root *Node
	mu   sync.RWMutex
}

func NewBTree() *BTree {
	return &BTree{Root: nil}
}

func (t *BTree) Insert(key Item, value interface{}) {
	node := &Node{value, key, nil, nil}

	t.mu.Lock()
	defer t.mu.Unlock()

	if t.Root == nil {
		t.Root = node
	} else {
		t.InsertNode(t.Root, node)
	}
}

func (t *BTree) InsertNode(parent, child *Node) {
	if t.gt(parent.Key, child.Key) {
		if parent.Left == nil {
			parent.Left = child
		} else {
			t.InsertNode(parent.Left, child)
		}
	} else {
		if parent.Right == nil {
			parent.Right = child
		} else {
			t.InsertNode(parent.Right, child)
		}
	}
}

func (t *BTree) Has(key Item) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.has(t.Root, key)
}

func (t *BTree) has(node *Node, key Item) bool {
	if node == nil {
		return false
	}

	if t.gt(node.Key, key) {
		return t.has(node.Left, key)
	}
	if t.gt(key, node.Key) {
		return t.has(node.Right, key)
	}

	return true
}

func (t *BTree) Search(key Item) interface{} {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.search(t.Root, key)
}

func (t *BTree) search(node *Node, key Item) interface{} {
	if node == nil {
		return nil
	}

	if t.gt(node.Key, key) {
		return t.search(node.Left, key)
	}
	if t.gt(key, node.Key) {
		return t.search(node.Right, key)
	}

	return node.Value
}

func (t *BTree) Remove(key Item) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	t.remove(t.Root, key)
}

func (t *BTree) remove(node *Node, key Item) *Node {
	if node == nil {
		return nil
	}

	if t.gt(node.Key, key) {
		// ищем ключ в левой ветке
		node.Left = t.remove(node.Left, key)
		return node
	}
	if t.gt(key, node.Key) {
		// ищем ключ в правой ветке
		node.Right = t.remove(node.Right, key)
		return node
	}

	if node.Left == nil {
		return node.Right
	} else if node.Right == nil {
		return node.Left
	}

	// ищем наименьшее значение в древе
	tmpNode := t.minValueNode(node)

	node.Key, node.Value = tmpNode.Key, tmpNode.Value
	node.Right = t.remove(node.Right, tmpNode.Key)

	return node
}

// minValueNode ищет наименьшее значение в древе
func (t *BTree) minValueNode(node *Node) *Node {
	tmpNode := node.Right
	for {
		if tmpNode != nil && tmpNode.Left != nil {
			tmpNode = tmpNode.Left
		} else {
			break
		}
	}

	return tmpNode
}

func (t *BTree) removeWorked(node *Node, key Item) *Node {
	if node == nil {
		return nil
	}

	// ищем ключ в правой ветке
	if t.gt(key, node.Key) {
		node.Right = t.remove(node.Right, key)
		return node
	}

	// ищем ключ в левой ветке
	if t.gt(node.Key, key) {
		node.Left = t.remove(node.Left, key)
		return node
	}

	// если обе ветки кончились - то считаем, что
	// не нашли результат
	if node.Left == nil && node.Right == nil {
		node = nil
		return nil
	}

	// выбираем одну из оставшихся нод
	if node.Left == nil {
		node = node.Right
		return node
	}
	if node.Right == nil {
		node = node.Left
		return node
	}

	rNode := node.Right
	for {
		// ищем наименьшее значение в древе
		if rNode != nil && rNode.Left != nil {
			rNode = rNode.Left
		} else {
			break
		}
	}

	// rNode гарантировано содержит минимальное значение
	// заменяем ноду с минимальным значением на удаленную
	node.Key, node.Value = rNode.Key, rNode.Value
	node.Right = t.remove(node.Right, node.Key)

	return node
}

func (t *BTree) gt(lv, rv interface{}) bool {
	// кастуем
	l := reflect.ValueOf(lv)
	r := reflect.ValueOf(rv)

	switch l.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return l.Int() > r.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return l.Uint() > r.Uint()
	case reflect.Float32, reflect.Float64:
		return l.Float() > r.Float()
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return false
	}

	return false
}
