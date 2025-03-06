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
	len   int
}

func NewDoublyLinkedList[T comparable](els ...T) *doublyLinkedList[T] {
	ddl := &doublyLinkedList[T]{}

	for _, e := range els {
		ddl.Append(e)
	}

	ddl.len = len(els)

	return ddl
}

func (ddl *doublyLinkedList[T]) ToSlice() []T {
	if ddl.Isempty() {
		return nil
	}

	res := make([]T, 0, ddl.len)

	curr := ddl.first
	for curr != nil {
		res = append(res, curr.value)

		curr = curr.next
	}

	return res
}

func (ddl *doublyLinkedList[T]) findByValue(el T) (*node[T], bool) {
	if ddl.Isempty() {
		return nil, false
	}

	if ddl.len == 1 {
		if ddl.first.value == el {
			return ddl.first, true
		}

		return nil, false
	}

	f, l := ddl.first, ddl.last
	for {
		// check if value is present

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
	// should I validate?
	return true
}

func (ddl *doublyLinkedList[T]) IsPresent(el T) bool {
	_, ok := ddl.findByValue(el)

	return ok
}

func (ddl *doublyLinkedList[T]) InsertAfter(el, newel T) bool {
	node, ok := ddl.findByValue(el)
	if !ok {
		return false
	}

	nn := newnode(newel)
	ddl.len++

	// meaning curr is the last element
	if node.next == nil {
		ddl.Append(newel)

		return true
	}

	p := node.next

	// bind current to new node
	node.next = nn

	// bind new node with both prev and next nodes
	nn.prev = node
	nn.next = p

	// bind prev node back to new node
	p.prev = nn

	return true
}

func (ddl *doublyLinkedList[T]) InsertBefore(el, newel T) bool {
	// check if empty
	node, ok := ddl.findByValue(el)
	if !ok {
		return false
	}

	nn := newnode(newel)
	ddl.len++

	// if first
	if ddl.first == node {
		f := ddl.first
		f.prev = nn

		ddl.first = nn
		nn.next = f

		return true
	}

	// if last
	if ddl.last == node {
		l := ddl.last

		// tie second to last node
		l.prev.next = nn
		nn.prev = l.prev

		nn.next = l
		l.prev = nn

		return true
	}

	node.prev.next = nn
	nn.prev = node.prev

	node.prev = nn
	nn.next = node

	return true
}

func (ddl *doublyLinkedList[T]) Isempty() bool {
	return ddl.first == nil && ddl.last == nil
}

func (ddl *doublyLinkedList[T]) Len() int {
	return ddl.len
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
	n := newnode(el)

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

func (ddl *doublyLinkedList[T]) Delete(el T) bool {
	node, ok := ddl.findByValue(el)
	if !ok {
		return ok
	}

	ddl.len--

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
