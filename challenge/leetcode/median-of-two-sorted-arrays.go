/*
	Даны два отсортированных массива nums1 и nums2 размерами m и n соответственно.
	Нужно найти медиану двух отсортированных массивов. Общая сложность выполнения должна
	быть O(log(n+m)).

	Пример 1:
		nums1 = [1, 3]
		nums2 = [2]
		Медиана 2.0

	Пример 2:
		nums1 = [1, 2]
		nums2 = [3, 4]
		Медиана (2 + 3)/2 = 2.5
*/

package main

import (
	"fmt"
)

// Считаем медиану уже у отсортированного массива.
func median(a []int) float64 {
	if l := len(a); l%2 == 0 {
		// если количество элементов четное, то считаем полусумму центральных элементов
		// с помощью побитового смещения получаем индексы центральных элементов
		return (float64(a[l>>1]) + float64(a[(l>>1)-1])) * 0.5
	} else {
		// если количество элементов нечетное, то ответом будет центральный элемент
		return float64(a[l>>1])
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	i, j, z := 0, 0, 0

	nums := make([]int, len1+len2)

	// формируем слайс
	for z < len1+len2 {
		// последний элемент не повлияет на подсчет медианы, поэтому
		// можно его не добавлять
		if i >= len1 {
			nums[z] = nums2[j]
			j++
			z++
			continue
		}
		if j >= len2 {
			nums[z] = nums1[i]
			i++
			z++
			continue
		}

		if nums1[i] < nums2[j] {
			nums[z] = nums1[i]
			i++
		} else {
			nums[z] = nums2[j]
			j++
		}
		z++
	}

	return median(nums)
}

func main() {
	r := findMedianSortedArrays([]int{1, 3}, []int{2})
	fmt.Println("result:", r)
	r = findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Println("result:", r)

	r = findMedianSortedArrays([]int{1}, []int{})
	fmt.Println("result:", r)

	r = findMedianSortedArrays([]int{}, []int{2, 3})
	fmt.Println("result:", r)
}
