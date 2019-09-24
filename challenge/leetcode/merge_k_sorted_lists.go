/*
 * Merge k sorted linked lists and return it as one sorted list.
 * Analyze and describe its complexity.
 */
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	var prev, head *ListNode

	for {
		nextExists := false
		minNode := &ListNode{Val: MaxInt, Next: nil}
		minIndex := 0

		// цикл для поиска минимума
		for i, list := range lists {
			if list != nil {
				nextExists = true
			} else {
				lists[i] = nil
				continue
			}

			if minNode.Val > list.Val {
				minNode = list
				minIndex = i
			}
		}

		// проверим, не остались ли мы на условности
		if minNode.Val == MaxInt {
			if prev != nil {
				prev.Next = nil
			}
			break
		}
		fmt.Println("минимальная нода в проходе:", minNode, "в списке номер", minIndex)

		if prev == nil {
			prev = &ListNode{Val: minNode.Val, Next: prev}
		} else {
			prev.Next = &ListNode{Val: minNode.Val, Next: prev}
			prev = prev.Next
		}

		if head == nil {
			// голова результирующей ноды
			head = prev
		}

		// смещаемся на указатель минимального элемента в найденом списке
		if minNode.Next != nil {
			lists[minIndex] = minNode.Next
		} else {
			lists[minIndex] = nil
		}

		// если все просмотренные списки == nil, выходим
		if !nextExists {
			prev.Next = nil
			break
		}
	}

	return head
}

func main() {
	var lists []*ListNode

	//c := initList(1, nil)
	//lists = append(lists, c)
	//result := mergeKLists(lists)

	//lists = append(lists, nil)
	//result := mergeKLists(lists)

	c := initList(3, &[]int{8, 9, 9})
	lists = append(lists, c)
	o1 := initList(2, &[]int{2, 3})
	lists = append(lists, o1)
	o2 := initList(2, &[]int{9, 1})
	lists = append(lists, o2)
	o3 := initList(0, nil)
	lists = append(lists, o3)

	result := mergeKLists(lists)

	//l1 := initList(1, &[]int{0})
	//lists = append(lists, l1)
	//l2 := initList(1, &[]int{1})
	//lists = append(lists, l2)
	//result := mergeKLists(lists)

	for l := result; l != nil; l = l.Next {
		fmt.Printf("%v → ", l)
	}

}
