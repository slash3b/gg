package gg_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
	"testing"

	"gg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	t.Parallel()

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
	t.Parallel()

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

func TestDifference(t *testing.T) {
	t.Parallel()

	a := gg.NewSet[int]()
	b := gg.NewSet[int]()

	a.Add(1)
	a.Add(2)

	b.Add(2)
	b.Add(8)

	u := a.Difference(b)

	res := make([]int, 0)

	for v := range u.All() {
		res = append(res, v)
	}

	expected := []int{1, 8}

	slices.Sort(res)
	slices.Sort(expected)

	if u.Len() != len(expected) {
		t.Fatalf("expected %d number of elements, got %d", len(expected), len(res))
	}

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected: %#v, got %#v", expected, res)
	}
}

func TestSubset(t *testing.T) {
	t.Parallel()

	a := gg.NewSet[int]()
	b := gg.NewSet[int]()

	a.Add(1)
	a.Add(2)

	b.Add(2)

	if !a.Subset(b) {
		t.Fatalf("expected b to be a subset of a.  a:%#v, b:%#v", a, b)
	}
}

func TestEnumerate(t *testing.T) {
	t.Parallel()

	a := gg.NewSet[int](1, 2)

	res := a.Enumerate()
	expected := []int{1, 2}

	slices.Sort(expected)
	slices.Sort(res)

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected enumeration:%#v, got :%#v", a.Enumerate(), expected)
	}
}

func TestDelete(t *testing.T) {
	t.Parallel()

	a := gg.NewSet[int](1, 2, 3)

	a.Delete(1, 3)

	t.Log(a.Len())
	t.Log(a.Lookup(2))

	if a.Len() == 1 && a.Lookup(2) {
		return
	}

	t.Fatalf("expected set to have exactly one element, instead %#v", a)
}

func TestMarshalSetOfStructs(t *testing.T) {
	t.Parallel()

	type fu struct{ Str string }

	a := gg.NewSet[fu](fu{Str: "hello"}, fu{Str: "bye"})

	t.Logf("set of structures: \n %#v\n", a)

	b, err := json.Marshal(a)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	t.Logf("result JSON: \n %s\n", string(b))

	// why this works without sorting?
	// the next test requires sorting to work
	// since we are comparing slices — ordered collection
	assert.JSONEq(t, `[{"Str":"hello"},{"Str":"bye"}]`, string(b))
}

func TestMarshalSetOfStrings(t *testing.T) {
	t.Parallel()

	a := gg.NewSet[string]("hello", "bye")

	t.Logf("set of strings: \n %#v\n", a)

	b, err := json.Marshal(a)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	t.Logf("result JSON: \n %s\n", string(b))

	var res []string

	err = json.Unmarshal(b, &res)
	require.NoError(t, err)

	expected := []string{"hello", "bye"}

	slices.Sort(expected)
	slices.Sort(res)

	if !reflect.DeepEqual(res, expected) {
		t.Fatal("set of strings shoudl be equal")
	}
}
