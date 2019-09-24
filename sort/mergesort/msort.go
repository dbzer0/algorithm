package mergesort

// less осуществляет data assertion и сравнение элементов
// между собой в зависимости от типа данных.
func less(aVal, bVal *interface{}) bool {
	switch a := (*aVal).(type) {
	case nil:
		// при nil возвращаем false?
		return false
	case int64:
		b := (*bVal).(int64)
		return a < b
	case string:
		b := (*bVal).(string)
		return a < b
	case float64:
		b := (*bVal).(float64)
		return a < b
	}

	return false
}

func MSort(data []interface{}) []interface{} {
	if len(data) <= 1 {
		return data
	}
	middle := len(data) / 2

	// Делим данные на левую и праву часть.
	// Рекурсивно разбиваем все на самые мелкие части, а затем
	// вызываем слияние (merge()) на получившиеся данные.
	return merge(MSort(data[:middle]), MSort(data[middle:]))
}

// merge - сливает данные
func merge(left, right []interface{}) []interface{} {
	result := make([]interface{}, 0)

	// пока есть элементы в левой и правой частях
	for len(left) > 0 && len(right) > 0 {
		// сравниваем их значения
		if less(&left[0], &right[0]) {
			result = append(result, left[0])
			left = left[1:] // убираем первый просмотренный элемент
		} else {
			result = append(result, right[0])
			right = right[1:] // убираем первый просмотренный элемент
		}
	}

	// добиваем остстки, если они есть, добавляем их в result
	// для левой части
	for j := 0; j < len(left); j++ {
		result = append(result, left[j])
	}
	// для правой части
	for j := 0; j < len(right); j++ {
		result = append(result, right[j])
	}

	return result
}
