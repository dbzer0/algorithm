package dbllist

type Item struct {
	// значение, хранимое списком
	Value interface{}

	// указатели на следующий и предыдущие значения
	next, prev *Item
}

// возвращает следующий item или nil
func (i *Item) Next() *Item {
	return i.next
}

// возвращает предыдущий item или nil
func (i *Item) Prev() *Item {
	return i.prev
}

