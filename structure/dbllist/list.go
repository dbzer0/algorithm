package dbllist

type List struct {
	head, tail *Item

	// будем считать длину, чтобы не считать ее со сложностью O(n)
	len        int32
}

func NewList() *List {
	return &List{head: nil, tail: nil, len: 0}
}

func (l *List) Head() *Item {
	return l.head
}

func (l *List) Tail() *Item {
	return l.tail
}

func (l *List) Len() int32 {
	return l.len
}

func (l *List) PushFront(i *Item) {
	e := Item{Value: i.Value}

	// новая голова указывает на старую
	e.next = l.head

	// меняем указатель старой головы на текущий элемент
	if l.head != nil {
		l.head.prev = &e
	}

	// если хвоста не было, то текущий элемент становится хвостом
	if l.tail == nil {
		l.tail = &e
	}

	l.head = &e
	l.len++
}

func (l *List) PushBack(i *Item) {
	e := Item{Value: i.Value}

	// новый хвост указывает на старый
	e.prev = l.tail

	// меняем указатель старого хвоста на текущий элемент
	if l.tail != nil {
		l.tail.next = &e
	}

	// если головы не было, то текущий элемент становится головой
	if l.head == nil {
		l.head = &e
	}

	l.tail = &e
	l.len++
}

func (l *List) Remove(i *Item) {
	// если головы нет
	if i.prev == nil {
		i.next.prev = nil
		l.head = i.next
	} else {
		i.prev.next = i.next
	}

	// если хвоста нет
	if i.next == nil {
		i.prev.next = nil
		l.tail = i.prev
	} else {
		i.next.prev = i.prev
	}

	// чтобы исключить утечку памяти, уберем указатели
	// у удаляемого элемента
	i.next = nil
	i.prev = nil
	i.Value = nil

	l.len--
}

func (l *List) Has(i *Item) bool {
	for e := l.head; e != l.tail; e = e.next {
		if i.Value == e.Value {
			return true
		}
	}

	return false
}

func (l *List) PopFront() *Item {
	e := l.head

	if l.head != nil && l.head.next != nil {
		l.head.next.prev = nil
		l.head = l.head.next

	}

	l.len--

	return e
}

func (l *List) PopBack() *Item {
	e := l.tail

	if l.tail != nil && l.tail.prev != nil {
		l.tail.prev.next = nil
		l.tail = l.tail.prev
	}

	l.len--

	return e
}
