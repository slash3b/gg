package gg

import (
	"encoding/json"
	"maps"

	"iter"
)

// set is a set implementation in golang
// set is unexported specifically to hide internal
// implementation and data structure — map.
type set[v comparable] struct {
	s map[v]struct{}
}

func NewSet[v comparable](els ...v) set[v] {
	s := set[v]{
		s: make(map[v]struct{}),
	}

	for _, val := range els {
		s.Add(val)
	}

	return s
}

func (s set[v]) Add(e v) bool {
	if _, ok := s.s[e]; ok {
		return false
	}

	s.s[e] = struct{}{}

	return true
}

func (s set[v]) Delete(els ...v) {
	for _, val := range els {
		delete(s.s, val)
	}
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
		for k := range s.s {
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
	for k := range other.All() {
		if _, ok := s.s[k]; !ok {
			return false
		}
	}

	return true
}

func (s set[v]) Enumerate() []v {
	res := make([]v, 0, s.Len())

	for k := range s.All() {
		res = append(res, k)
	}

	return res
}

func (s set[v]) MarshalJSON() ([]byte, error) {
	if s.Len() == 0 {
		return []byte{}, nil
	}

	var res []v

	for el := range maps.Keys(s.s) {
		res = append(res, el)
	}

	return json.Marshal(res)
}

func (s set[v]) UnmarshalJSON(data []byte) error {
	return nil
}

// fixme: add more methods from wiki https://en.wikipedia.org/wiki/Set_(abstract_data_type)
