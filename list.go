package uint32sets

import "sort"

type listSet struct {
	list []uint32
}

func newListSet() set {
	return &listSet{
		list: []uint32{},
	}
}

func (s *listSet) add(i uint32) {
	s.list = append(s.list, i)
}

func (s *listSet) contains(i uint32) bool {
	sort.Sort(s)
	idx := sort.Search(len(s.list), func(idx int) bool {
		return s.list[idx] >= i
	})
	return idx < len(s.list) && s.list[idx] == i
}

func (s listSet) length() int {
	if len(s.list) == 0 {
		return 0
	}

	sort.Sort(s)

	unique := 1
	v := s.list[0]

	for _, i := range s.list {
		if v != i {
			unique++
			v = i
		}
	}

	return unique
}

func (s listSet) Len() int {
	return len(s.list)
}

func (s listSet) Less(i, j int) bool {
	return s.list[i] < s.list[j]
}

func (s listSet) Swap(i, j int) {
	s.list[i], s.list[j] = s.list[j], s.list[i]
}
