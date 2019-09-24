package binary_search

func Find(values []int64, value int64) *int64 {
	l := len(values) / 2
	pivot := values[l]

	// нашли значение!
	if pivot == value {
		return &pivot
	}

	// эта проверка должна быть после проверки найденного значения
	if l == 0 {
		// не найдено
		return nil
	}

	switch pivot < value {
	case false:
		return Find(values[:l], value)
	case true:
		return Find(values[l:], value)
	}

	return nil
}
