package gg

import (
	"fmt"

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

func (ddl *doublyLinkedList[T]) ToSlice() []T {
	if ddl.Isempty() {
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

func (ddl *doublyLinkedList[T]) Debug() {
	println("debug------------------")
	if ddl.Isempty() {
		fmt.Println("ddl is empty")
		return
	}

	if ddl.first != nil {
		fmt.Println("first node", ddl.first)
	} else {
		fmt.Println("first node is empty")
	}

	if ddl.last != nil {
		fmt.Println("last node", ddl.last)
	} else {
		fmt.Println("last node is empty")
	}

	fmt.Println("all nodes:", ddl.ToSlice())
	println("------------------debug")
}

// fixme: make private again
func (ddl *doublyLinkedList[T]) findByValue(el T) (*node[T], bool) {
	if ddl.Isempty() {
		return nil, false
	}

	if ddl.Len() == 1 {
		if ddl.first.value == el {
			return ddl.first, true
		}

		return nil, false
	}

	f, l := ddl.first, ddl.last
	for {
		// check if values are present
		if f.value == el {
			return f, true
		}

		if l.value == el {
			return l, true
		}

		// check if we need to stop

		// point to each other
		if f.next == l {
			return nil, false
		}

		// point to the same element
		if f.next == l.prev {
			if f.next.value == el {
				return f.next, true
			}

			return nil, false
		}

		f = f.next
		l = l.prev
	}

	return nil, false
}

// todo: add method to check that all next and previous connections are correct
func (ddl *doublyLinkedList[T]) validate() bool {
	return true
}

func (ddl *doublyLinkedList[T]) IsPresent(el T) bool {
	return false
}

func (ddl *doublyLinkedList[T]) InsertAfter(el, newel T) bool {
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
func (ddl *doublyLinkedList[T]) InsertBefore(el, newel T) bool {
	return true
}

// fixme: should this be empty?
// fixme: should both be equal to the same thing?
func (ddl *doublyLinkedList[T]) Isempty() bool {
	return ddl.first == nil && ddl.last == nil
}

func (ddl *doublyLinkedList[T]) Len() int {
	res := 0

	// fixme: track it internally

	curr := ddl.first
	for curr != nil {
		res++
		curr = curr.next
	}

	return res
}

func (ddl *doublyLinkedList[T]) Prepend(el T) {
	nn := newnode(el)

	if ddl.first == nil {
		ddl.first = nn

		return
	}

	nn.next = ddl.first
	ddl.first.prev = nn

	ddl.first = nn
}

func (ddl *doublyLinkedList[T]) Append(el T) {
	n := &node[T]{value: el}

	if ddl.Isempty() {
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

func (ddl *doublyLinkedList[T]) isSigleNodeOnly() bool {
	return ddl.first != nil && ddl.last == nil
}

func (ddl *doublyLinkedList[T]) Delete(el T) bool {
	node, ok := ddl.findByValue(el)
	if !ok {
		return ok
	}

	p := node.prev
	n := node.next

	// matched with single element
	if p == nil && n == nil {
		ddl.first = nil

		return true
	}

	// matched with first element
	if p == nil && n != nil {
		n.prev = nil
		ddl.first = n

		return true
	}

	// matched with last element
	if p != nil && n == nil {

		// special case for 2 element queue
		if p.value == ddl.first.value {
			ddl.first.next = nil
			ddl.last = nil

			return true
		}

		p.next = nil
		ddl.last = p

		return true
	}

	p.next = n
	n.prev = p

	return true
}

func (ddl *doublyLinkedList[T]) All() iter.Seq[T] {
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

func (ddl *doublyLinkedList[T]) AllReverse() iter.Seq[T] {
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
