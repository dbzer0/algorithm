/*
	Даны два односвязных списка представленные как два не отрицательных целых.
	Числа хранятся в обратном порядке и каждая из нод содержит одно число.
	Сложите два значения каждой ноды и верните результирующий список.

	You may assume the two numbers do not contain any leading zero, except the number 0 itself.

	Вход:  (2 -> 4 -> 3) + (5 -> 6 -> 4)
	Выход: 7 -> 0 -> 8
*/
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	sum := l1.Val + l2.Val

	// делаем новую результирующую ноду списка с результатом суммирования
	result := &ListNode{Val: sum % 10, Next: nil}

	// создаем рекурсией следующий элемент
	result.Next = addTwoNumbers(l1.Next, l2.Next)

	// если сумма была больше 10 то прибавляем разряд к следующему элементу
	if sum >= 10 {
		result.Next = addTwoNumbers(result.Next, &ListNode{1, nil})
	}

	return result
}

// Инициализатор списка, вспомогательная функция для тестов.
// Параметры:
//   size - размер односвязного списка
//   vals - указатель на массив целых чисел
// Если vals равен nil, то все элементы инициализируются с нулевым
// значением.
func initList(size int, vals *[]int) *ListNode {
	if size < 0 {
		return nil
	}

	var prev *ListNode

	if vals != nil {
		prev = &ListNode{Val: (*vals)[0], Next: nil}
	} else {
		prev = &ListNode{Val: 0, Next: nil}
	}

	head := &prev // запоминаем указатель на первый элемент

	for i := 0; i < size-1; i++ {
		if vals != nil {
			prev = &ListNode{Val: (*vals)[i+1], Next: prev}
		} else {
			prev = &ListNode{Val: 0, Next: prev}
		}
	}

	return *head
}

func main() {
	c1 := initList(3, &[]int{8, 9, 9})
	c2 := initList(1, &[]int{2})
	for item := addTwoNumbers(c1, c2); item != nil; item = item.Next {
		fmt.Printf("%v → ", item.Val)
	}
	fmt.Println()
	fmt.Println("==================================")

	o1 := initList(1, &[]int{1})
	o2 := initList(2, &[]int{9, 9})
	for item := addTwoNumbers(o1, o2); item != nil; item = item.Next {
		fmt.Printf("%v → ", item.Val)
	}
	fmt.Println()
	fmt.Println("==================================")

	m1 := initList(3, &[]int{2, 4, 3})
	m2 := initList(3, &[]int{5, 6, 4})
	for item := addTwoNumbers(m1, m2); item != nil; item = item.Next {
		fmt.Printf("%v → ", item.Val)
	}
	fmt.Println()
	fmt.Println("==================================")

	l1 := initList(2, &[]int{8, 9})
	l2 := initList(1, &[]int{1})
	for item := addTwoNumbers(l1, l2); item != nil; item = item.Next {
		fmt.Printf("%v → ", item.Val)
	}
	fmt.Println()
	fmt.Println("==================================")

	j1 := initList(9, &[]int{5, 3, 8, 7, 6, 8, 3, 1, 1})
	j2 := initList(10, &[]int{4, 7, 8, 7, 5, 6, 4, 6, 1, 1})
	for item := addTwoNumbers(j1, j2); item != nil; item = item.Next {
		fmt.Printf("%v → ", item.Val)
	}
	fmt.Println()
	fmt.Println("==================================")

	k1 := initList(3, &[]int{2, 4, 3})
	k2 := initList(3, &[]int{5, 6, 4})
	for item := addTwoNumbers(k1, k2); item != nil; item = item.Next {
		fmt.Printf("%v → ", item.Val)
	}
	fmt.Println()
}
