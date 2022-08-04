package uint32sets

const slotSetSize = 1 << 17 // that's the sweet spot on my machine.

type slotSet struct {
	slots [slotSetSize][]uint32
	l     int
}

func newSlotSet() set {
	return &slotSet{}
}

func (s *slotSet) add(i uint32) {
	slotIdx := int(i) % len(s.slots)

	for _, v := range s.slots[slotIdx] {
		if v == i {
			return
		}
	}

	s.slots[slotIdx] = append(s.slots[slotIdx], i)
	s.l++
}

func (s *slotSet) contains(i uint32) bool {
	slotIdx := int(i) % len(s.slots)

	for _, v := range s.slots[slotIdx] {
		if v == i {
			return true
		}
	}

	return false
}

func (s *slotSet) length() int {
	return s.l
}

type slotPreallocSet struct {
	slots [slotSetSize][]uint32
	l     int
}

func newSlotPreallocSet() set {
	return &slotPreallocSet{}
}

func (s *slotPreallocSet) add(i uint32) {
	slotIdx := int(i) % len(s.slots)

	slot := s.slots[slotIdx]

	if slot == nil {
		slot = make([]uint32, 0, 8)
		s.slots[slotIdx] = slot
	} else {
		for _, v := range slot {
			if v == i {
				return
			}
		}
	}

	s.slots[slotIdx] = append(s.slots[slotIdx], i)
	s.l++
}

func (s *slotPreallocSet) contains(i uint32) bool {
	slotIdx := int(i) % len(s.slots)

	for _, v := range s.slots[slotIdx] {
		if v == i {
			return true
		}
	}

	return false
}

func (s *slotPreallocSet) length() int {
	return s.l
}
