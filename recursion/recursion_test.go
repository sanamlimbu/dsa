package recursion

import (
	"testing"
)

func TestSumInt(t *testing.T) {
	arrOne := []int{1, 2, 3, 4}
	arrTwo := []int{1, 1, 1, 1}
	arrThree := []int{2, 3, 1, 5}

	sumArrOne := SumInt(arrOne)
	sumArrTwo := SumInt(arrTwo)
	sumArrThree := SumInt(arrThree)

	if sumArrOne != 10 {
		t.Errorf("Wanted 10 got %d", sumArrOne)
	}

	if sumArrTwo != 4 {
		t.Errorf("Wanted 4 got %d", sumArrTwo)
	}

	if sumArrThree != 11 {
		t.Errorf("Wanted 11 got %d", sumArrThree)
	}
}

func TestCountInt(t *testing.T) {
	arrOne := []int{1, 2, 3, 4}
	arrTwo := []int{1, 2, 3}

	countOne := CountInt(arrOne)
	countTwo := CountInt(arrTwo)

	if countOne != 4 {
		t.Errorf("Wanted 4 got %d", countOne)
	}

	if countTwo != 3 {
		t.Errorf("Wanted 3 got %d", countTwo)
	}
}

func TestMaximumInt(t *testing.T) {
	arrOne := []int{0, 2, 13, 4}
	arrTwo := []int{-2, -4, -3, -5}

	maxArrOne := MaximumInt(arrOne, len(arrOne))
	maxArrTwo := MaximumInt(arrTwo, len(arrOne))

	if maxArrOne != 13 {
		t.Errorf("Wanted 13 got %d", maxArrOne)
	}

	if maxArrTwo != -2 {
		t.Errorf("Wanted -2 got %d", maxArrTwo)
	}
}

func TestBinarySearch(t *testing.T) {
	arrOne := []int{1, 2, 3, 4, 5}
	arrTwo := []int{-10, -8, -5, -2, -1}

	indexArrOne := BinarySearch(arrOne, 3, 0, len(arrOne))
	indexArrTwo := BinarySearch(arrTwo, -2, 0, len(arrTwo))

	if indexArrOne != 2 {
		t.Errorf("Wanted 2 got %d", indexArrOne)
	}

	if indexArrTwo != 3 {
		t.Errorf("Wanted 3 got %d", indexArrTwo)
	}
}
