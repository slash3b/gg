package gg

import (
	"iter"
)

// set is a set implementation in golang
// set is unexported specifically to hide internal
// implementation and data structure — map.
type set[v comparable] struct {
	s map[v]struct{}
}

func NewSet[v comparable]() set[v] {
	return set[v]{
		s: make(map[v]struct{}),
	}
}

func (s set[v]) Add(e v) bool {
	if _, ok := s.s[e]; ok {
		return false
	}

	s.s[e] = struct{}{}

	return true
}

func (s set[v]) Len() int {
	return len(s.s)
}

func (s set[v]) IsEmpty() bool {
	return len(s.s) == 0
}

func (s set[v]) Lookup(el v) bool {
	_, ok := s.s[el]

	return ok
}

func (s set[v]) All() iter.Seq[v] {
	return func(y func(el v) bool) {
		for k, _ := range s.s {
			if !y(k) {
				break
			}
		}
	}
}

func (s set[v]) Union(other set[v]) set[v] {
	news := NewSet[v]()

	for el := range s.All() {
		news.Add(el)
	}

	for el := range other.All() {
		news.Add(el)
	}

	return news
}

func (s set[v]) Intersection(other set[v]) set[v] {
	news := NewSet[v]()

    for v := range other.All() {
        if _, ok := s.s[v]; ok {
            news.Add(v)
        }
    }

    return news
}

func (s set[v]) Difference(other set[v]) set[v] {
	news := NewSet[v]()

	tmp := make(map[v]int)

	for v := range s.s {
		tmp[v]++
	}

	for v := range other.All() {
		tmp[v]++
	}

	for k, cnt := range tmp {
		if cnt == 1 {
			news.Add(k)
		}
	}

	return news
}

func (s set[v]) Subset(other set[v]) bool {
	for k, _ := range other.All() {
		if _, ok := s.s[k]; !ok {
			return false
		}
	}

	return true
}

// fixme: marshall and unmarshall to implement
