package quicksort

import "math/rand"

// Вспомогательная функция, осуществляющая data assertion и
// сравнение элементов между собой в зависимости от типа данных.
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

func QSort(data []interface{}) []interface{} {
	if len(data) <= 1 {
		return data
	}

	// делаем слайсы для хранения элементов меньшего по
	// значению, равному и большему
	smaller := make([]interface{}, 0, len(data))
	equal := make([]interface{}, 1, len(data))
	larger := make([]interface{}, 0, len(data))

	// нулевой элемент будет начальной точкой разделения
	pivot := data[0]

	// точка разделения всегда равна
	equal[0] = pivot

	// проходим по массиву, сравнивая и заполняя
	// соответствующие слайсы
	for i := 1; i < len(data); i++ {
		if less(&pivot, &data[i]) {
			larger = append(larger, data[i])
		} else if less(&data[i], &pivot) {
			smaller = append(smaller, data[i])
		} else {
			equal = append(equal, data[i])
		}
	}

	// возвращаем единый массив, запуская рекурсию по
	// собранным слайсам, кроме равного, так как он уже в середине
	return append(append(QSort(smaller), equal...), QSort(larger)...)
}

func Sort(l []int) []int {
	if len(l) <= 1 {
		return l
	}

	idx := rand.Intn(len(l))
	op := l[idx]

	left := make([]int, 0, len(l[:idx]))
	right := make([]int, 0, len(l[idx:])-1)
	middle := make([]int, 0, len(l)-1)

	for i, _ := range l {
		switch {
		case l[i] < op:
			left = append(left, l[i])
		case l[i] == op:
			middle = append(middle, l[i])
		case l[i] > op:
			right = append(right, l[i])
		}
	}

	left, right = Sort(left), Sort(right)

	return append(append(left, middle...), right...)
}
