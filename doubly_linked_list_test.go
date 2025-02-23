package gg_test

import (
	"reflect"
	"slices"
	"testing"

	"gg"
)

func TestDoublyLinkedList_All(t *testing.T) {
	input := []int{42, 24, 0, 2}

	ddl := gg.NewDoublyLinkedList[int](input...)

	if ddl.Len() != len(input) {
		t.Fatalf("expected %d, got %d elements in dlinked list", input, ddl.Len())
	}

	if !reflect.DeepEqual(input, ddl.ToSlice()) {
		t.Fatalf("expected %v, but got %v", input, ddl.ToSlice())
	}
}

func TestDoublyLinkedList_Backward(t *testing.T) {
	input := []int{42, 24, 0, 2}

	ddl := gg.NewDoublyLinkedList[int](input...)

	if ddl.Len() != len(input) {
		t.Fatalf("expected %d, got %d elements in dlinked list", input, ddl.Len())
	}

	res := []int{}
	for v := range ddl.AllReverse() {
		res = append(res, v)
	}

	slices.Reverse(input)

	if !reflect.DeepEqual(input, res) {
		t.Fatalf("expected %v, but got %v", input, ddl.ToSlice())
	}

	t.Logf("reversed linked list %#v", res)
}

func TestDoublyLinkedList_Backward_SingleElement(t *testing.T) {
	input := []int{42}

	ddl := gg.NewDoublyLinkedList[int](input...)

	if ddl.Len() != len(input) {
		t.Fatalf("expected %d, got %d elements in dlinked list", input, ddl.Len())
	}

	res := []int{}
	for v := range ddl.AllReverse() {
		res = append(res, v)
	}

	if !reflect.DeepEqual(input, res) {
		t.Fatalf("expected %v, but got %v", input, ddl.ToSlice())
	}

	t.Logf("reversed linked list %#v", res)
}

func TestDoublyLinkedList_InsertAfter(t *testing.T) {
	input := []int{42, 3, 0}

	ddl := gg.NewDoublyLinkedList[int](input...)

	if ddl.Len() != len(input) {
		t.Fatalf("expected %d, got %d elements in dlinked list", input, ddl.Len())
	}

	ddl.InsertAfter(3, 7)

	if !reflect.DeepEqual(ddl.ToSlice(), []int{42, 3, 7, 0}) {
		t.Fatalf("value 7 be after 3, got: %#v", ddl.ToSlice())
	}

	t.Logf("slice from linked list: %#v\n", ddl.ToSlice())
}
