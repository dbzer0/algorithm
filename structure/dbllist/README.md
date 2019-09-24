# Double linked list

Содержит реализацию двусвязного списка. Дли инициализации списка существует функция `NewList()`.
Также реализованы следующие методы:

  * PushFront() - помещает элемент в начало списка
  * PopFront() - достает элемент из начала списка
  * PushBack() - помещает элемент в конец списка
  * PopBack()  - достает элемент из конца списка
  * Remove() - удаляет элемент из списка
  * Head() - возвращает элемент головы
  * Tail() - возвращает элемент хвоста
  * Len()  - возвращает длину списка (чтобы не обходить весь список и уменьшить сложность с O(n))

Пример:

	dblList := dbllist.NewList()

	dblList.PushFront(&double_linked_list.Item{Value: 0})
	dblList.PushFront(&double_linked_list.Item{Value: 1})
	dblList.PushBack(&double_linked_list.Item{Value: -1})
	
	for e := dblList.Head(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	
	i := dblList.PopBack()
	fmt.Println(i.Value)


