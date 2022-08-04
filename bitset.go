package uint32sets

type slotBitmapSet struct {
	slots [1 << 16]*bitmap
	l     int
}

type bitmap [(1 << 16) / 8]byte

func newSlotBitmapSet() set {
	return &slotBitmapSet{}
}

func (s *slotBitmapSet) add(i uint32) {
	slot := int(i >> 16)
	lowerWord := int(i & 0xFFFF)

	bm := s.slots[slot]
	if bm == nil {
		bm = new(bitmap)
		s.slots[slot] = bm
	}

	bitmapIdx := int(lowerWord / 8)
	mask := byte(1 << (lowerWord % 8))
	if (*bm)[bitmapIdx]&mask == mask {
		return
	}

	(*bm)[bitmapIdx] |= mask
	s.l++
}

func (s *slotBitmapSet) contains(i uint32) bool {
	slot := int(i >> 16)
	lowerWord := int(i & 0xFFFF)

	bm := s.slots[slot]
	if bm == nil {
		return false
	}

	bitmapIdx := int(lowerWord / 8)
	mask := byte(1 << (lowerWord % 8))
	if (*bm)[bitmapIdx]&mask == mask {
		return true
	}

	return false
}

func (s *slotBitmapSet) length() int {
	return s.l
}
