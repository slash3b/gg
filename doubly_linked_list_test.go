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

func TestDoublyLinkedList_Delete(t *testing.T) {
	input := []int{42, 3, 0}
	t.Logf("input: %#v", input)

	ddl := gg.NewDoublyLinkedList[int](input...)

	t.Logf("delete result (3): %v", ddl.Delete(3))
	t.Logf("slice state: %#v", ddl.ToSlice())

	if !reflect.DeepEqual(ddl.ToSlice(), []int{42, 0}) {
		t.Fatalf("err: unexpected slice state: %#v", ddl.ToSlice())
	}

	res := ddl.Delete(0)
	t.Logf("delete result(0): %v", res)
	t.Logf("slice state: %#v", ddl.ToSlice())
	if !reflect.DeepEqual(ddl.ToSlice(), []int{42}) {
		t.Fatalf("err: unexpected slice state: %#v", ddl.ToSlice())
	}

	res = ddl.Delete(84)
	t.Logf("delete result(84): %v", res)
	t.Logf("slice state: %#v", ddl.ToSlice())
	if res {
		t.Fatal("err: unexpected success upon deleting foreign element")
	}

	res = ddl.Delete(42)
	t.Logf("delete result(42): %v", res)
	t.Logf("slice state: %#v", ddl.ToSlice())
	if !res {
		t.Fatalf("err: unable to delete existing element")
	}
	if !ddl.IsEmpty() {
		t.Fatal("err: queue supposed to be empty")
	}

	res = ddl.Delete(42)
	// attempt to delete already deleted
	t.Logf("delete result(42): %v", res)
	t.Logf("slice state: %#v", ddl.ToSlice())
	if res {
		t.Fatalf("err: unexpected successful deletion for not existing element")
	}
}

func TestDoublyLinkedList_Prepend(t *testing.T) {
	input := []int{42}
	t.Logf("input: %#v", input)

	ddl := gg.NewDoublyLinkedList[int](input...)

	ddl.Prepend(7)
	ddl.Prepend(-1)

	if !reflect.DeepEqual(ddl.ToSlice(), []int{-1, 7, 42}) {
		t.Fatalf("err: unexpected slice state: %#v", ddl.ToSlice())
	}
}

func TestDoublyLinkedList_PrependToEmpty(t *testing.T) {
	ddl := gg.NewDoublyLinkedList[int]()

	ddl.Prepend(1)
	ddl.Prepend(2)
	ddl.Prepend(3)

	if !reflect.DeepEqual(ddl.ToSlice(), []int{3, 2, 1}) {
		t.Fatalf("err: unexpected slice state: %#v", ddl.ToSlice())
	}
}
