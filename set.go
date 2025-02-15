package gg

import (
	"iter"
)

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
