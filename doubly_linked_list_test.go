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
