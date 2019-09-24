package hash_table

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("element not found")

type Element struct {
	Key   interface{}
	Value interface{}
	next  *Element
	prev  *Element
}

// соотношение хэша с указателем на первый элемент
type Bucket struct {
	Element *Element
}

type HashTable struct {
	lock     sync.RWMutex
	hashFunc HashFunc
	buckets  []Bucket
}

func NewHashTable(f HashFunc) *HashTable {
	buckets := make([]Bucket, 256)
	return &HashTable{hashFunc: f, buckets: buckets}
}

func (h *HashTable) getElement(key interface{}) *Element {
	h.lock.Lock()
	defer h.lock.Unlock()

	var e *Element
	hash := h.hashFunc(key)

	// запрошенный элемент не найден если элемент = nil
	// либо совпал хэш, но ключ различается
	if e = h.buckets[hash].Element; e == nil {
		return nil
	}

	// если следующего элемента нет, то возвращаем значение сразу
	// это условие не критично, но оно отработает быстрее, чем
	// рекурсивный вызов findExists()
	if e.next == nil {
		return e
	}

	return h.findExists(e, key)
}

func (h *HashTable) Get(key interface{}) (interface{}) {
	var e *Element
	if e = h.getElement(key); e == nil {
		return nil
	}
	return e.Value
}

func (h *HashTable) Add(key, value interface{}) {
	h.lock.Lock()
	defer h.lock.Unlock()

	var cur *Element
	hash := h.hashFunc(key)

	// новый элемент
	e := Element{
		Key:   key,
		Value: value,
		next:  nil,
		prev:  nil,
	}

	// если такого элемента нет в бакете, то сразу вствляем его
	if cur = h.buckets[hash].Element; cur == nil {
		// новый бакет
		h.buckets[hash] = Bucket{Element: &e}
		return
	}

	// ищем существующий, либо последний ключ
	cur = h.findExists(cur, key)
	if cur.Key == key {
		// если элемент не последний, но нашелся такой же по
		// ключу, то переносим указатель на следующий элемент
		e.next = cur.next
		*cur = e

		return
	}

	// элемент оказался последним, создаем указатель на новый
	cur.next = &e
}

func (h *HashTable) findExists(cur *Element, key interface{}) *Element {
	if cur.Key == key {
		return cur
	}
	if cur.next == nil {
		return cur
	}
	return h.findExists(cur.next, key)
}

func (h *HashTable) Remove(key interface{}) error {
	h.lock.Lock()
	defer h.lock.Unlock()

	var e *Element
	hash := h.hashFunc(key)

	e = h.buckets[hash].Element
	if e == nil {
		// запрошенный элемент не найден если элемент = nil
		return ErrNotFound
	}

	if e.Key == key && e.next == nil {
		e = nil                       // чтобы GC его забрал
		h.buckets[hash].Element = nil // это единственный элемент, затираем
	}

	return nil
}
