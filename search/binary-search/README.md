# Binary search

Алгоритм двоичного поиска в заранее отсортированных данных. Для поиска элемента используется функция `Find()`.

Пример:

	list := []int64{2, 3, 5, 7, 53, 59, 61, 67, 71, 73, 89, 97}

	var value *int64
	if value = binary_search.Find(list, 5); value == nil {
		fmt.Println("Not found")
		return
	}
	fmt.Println(*value)

