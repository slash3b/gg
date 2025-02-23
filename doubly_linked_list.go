package gg

import (
	"iter"
)

type node[T comparable] struct {
	value T
	next  *node[T]
	prev  *node[T]
}

func newnode[T comparable](el T) *node[T] {
	return &node[T]{value: el}
}

type doublyLinkedList[T comparable] struct {
	first *node[T]
	last  *node[T]
}

func NewDoublyLinkedList[T comparable](els ...T) *doublyLinkedList[T] {
	ddl := &doublyLinkedList[T]{}

	for _, e := range els {
		ddl.Append(e)
	}

	return ddl
}

func (ddl doublyLinkedList[T]) ToSlice() []T {
	if ddl.IsEmpty() {
		return nil
	}

	res := make([]T, 0, ddl.Len())

	curr := ddl.first
	for curr != nil {
		res = append(res, curr.value)

		curr = curr.next
	}

	return res
}

func (ddl doublyLinkedList[T]) findByValue(el T) *node[T] {
	// todo: use 2 pointers to start searching from both ends
	// use special case for ddl.Len() == 1
	// a bit lame but works okay for now.
	return nil
}

// todo: add method to check that all next and previous connections are correct
func (ddl doublyLinkedList[T]) validate() bool {
	return true
}

func (ddl doublyLinkedList[T]) InsertAfter(el, newel T) bool {
	curr := ddl.first

	for curr != nil {

		if curr.value == el {

			// meaning curr is the last element
			if curr.next == nil {
				ddl.Append(newel)

				return true
			}

			nn := newnode(newel)
			cn := curr.next

			// bind current to new node
			curr.next = nn

			// bind new node with both prev and next nodes
			nn.prev = curr
			nn.next = cn

			cn.prev = nn

			return true
		}

		curr = curr.next
	}

	return false
}

// fixme: what to do if there are duplicates
func (ddl doublyLinkedList[T]) InsertBefore(el, newel T) bool {
	return true
}

// fixme: should this be empty?
// fixme: should both be equal to the same thing?
func (ddl doublyLinkedList[T]) IsEmpty() bool {
	return ddl.first == nil && ddl.last == nil
}

func (ddl doublyLinkedList[T]) Len() int {
	res := 0

	// fixme: track it internally

	curr := ddl.first
	for curr != nil {
		res++
		curr = curr.next
	}

	return res
}

func (ddl doublyLinkedList[T]) Prepend(el T) {
	return
}

func (ddl *doublyLinkedList[T]) Append(el T) {
	n := &node[T]{value: el}

	if ddl.IsEmpty() {
		ddl.first = n

		return
	}

	if ddl.last == nil {
		ddl.last = n
		ddl.first.next = ddl.last
		ddl.last.prev = ddl.first

		return
	}

	// connect last and new node
	ddl.last.next = n
	n.prev = ddl.last
	ddl.last = n
}

func (ddl doublyLinkedList[T]) Delete(el T) {
	return
}

func (ddl doublyLinkedList[T]) All() iter.Seq[T] {
	curr := ddl.first

	return func(y func(el T) bool) {
		for curr != nil {
			if !y(curr.value) {
				break
			}
			curr = curr.next
		}
	}
}

func (ddl doublyLinkedList[T]) AllReverse() iter.Seq[T] {
	curr := ddl.last
	if ddl.Len() == 1 {
		curr = ddl.first
	}

	return func(y func(el T) bool) {
		for curr != nil {
			if !y(curr.value) {
				break
			}
			curr = curr.prev
		}
	}
	return nil
}
