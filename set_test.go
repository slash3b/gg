package gg_test

import (
	"fmt"
	"gg"
	"reflect"
	"slices"
	"testing"
)

func Exampleset_Add() {
	s := gg.NewSet[int]()

	s.Add(1)
	s.Add(2)

	for v := range s.All() {
		fmt.Printf("%#v\n", v)
	}

	// Output:
	// 1
	// 2
}

func TestUnion(t *testing.T) {
	a := gg.NewSet[int]()
	b := gg.NewSet[int]()

	a.Add(1)
	a.Add(2)
	a.Add(3)
	a.Add(4)

	b.Add(1)
	b.Add(2)
	b.Add(3)
	b.Add(5)

	u := a.Union(b)

	res := make([]int, 0)

	for v := range u.All() {
		res = append(res, v)
	}

	expected := []int{1, 2, 3, 4, 5}

	slices.Sort(res)
	slices.Sort(expected)

	if u.Len() != len(expected) {
		t.Fatalf("expected %d number of elements, got %d", len(expected), len(res))
	}

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected: %#v, got %#v", expected, res)
	}
}

func TestIntersection(t *testing.T) {
	a := gg.NewSet[int]()
	b := gg.NewSet[int]()

	a.Add(1)
	a.Add(2)
	a.Add(3)
	a.Add(5)

	b.Add(1)
	b.Add(7)
	b.Add(8)
	b.Add(5)

	u := a.Intersection(b)

	res := make([]int, 0)

	for v := range u.All() {
		res = append(res, v)
	}

	expected := []int{1, 5}

	slices.Sort(res)
	slices.Sort(expected)

	if u.Len() != len(expected) {
		t.Fatalf("expected %d number of elements, got %d", len(expected), len(res))
	}

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected: %#v, got %#v", expected, res)
	}
}
