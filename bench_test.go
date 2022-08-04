package uint32sets

import (
	"math/rand"
	"testing"
)

func BenchmarkAll(b *testing.B) {
	tests := map[string]func() set{
		"mapSet":          newMapSet,
		"slotSet":         newSlotSet,
		"slotPreallocSet": newSlotPreallocSet,
		"slotBitmapSet":   newSlotBitmapSet,
		"listSet":         newListSet,
	}

	numberSet := []uint32{}
	for i := 0; i < 20_000; i++ {
		numberSet = append(numberSet, uint32(rand.Int()))
	}

	containsCheckNumber := numberSet[len(numberSet)/2]

	allNumbers := []uint32{}
	for i := 0; i < 500_000; i++ {
		allNumbers = append(allNumbers, numberSet[rand.Int()%len(numberSet)])
	}

	for testName, cons := range tests {
		b.Run(testName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s := cons()
				for _, num := range allNumbers {
					s.add(num)
				}
				if !s.contains(containsCheckNumber) {
					b.Fatalf("number %d was added but could not be found", containsCheckNumber)
				}
				if l := s.length(); l != 20_000 {
					b.Fatalf("expected length of 20000, got %d", l)
				}
			}
		})
	}
}
