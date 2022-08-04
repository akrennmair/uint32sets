package uint32sets

type mapSet map[uint32]struct{}

func newMapSet() set {
	return make(mapSet)
}

func (s mapSet) add(i uint32) {
	s[i] = struct{}{}
}

func (s mapSet) contains(i uint32) bool {
	_, ok := s[i]
	return ok
}

func (s mapSet) length() int {
	return len(s)
}
