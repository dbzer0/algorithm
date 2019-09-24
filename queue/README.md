# queue

Реализация очередей FIFO и LIFO. Для инициализации списка существуют соотетствующие функции `NewFIFO()` и `NewLIFO()`.

Оба класса имеют методы:

  * Push() - поместить элемент в очередь
  * Pop() - достать элемент из очереди
  * Len() - проверить длину очереди

Пример:

	fifo := queue.NewFIFO()
	fifo.Push(1)
	fifo.Push(2)
	fifo.Push(3)

	for fifo.Len() != 0 {
		fmt.Println(fifo.Pop())
	}

