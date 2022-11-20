package search

import "testing"

func TestBinaryInt(t *testing.T) {
	arrOne := []int{-9, 0, 1, 5, 10, 12}
	arrtwo := []int{0, 2, 4, 9, 23, 70}

	indexArrOne := BinaryInt(arrOne, 5)
	indexArrTwo := BinaryInt(arrtwo, 23)

	if indexArrOne != 3 {
		t.Errorf("Wanted 3 got %d", indexArrOne)
	}

	if indexArrTwo != 4 {
		t.Errorf("Wanted 4 got %d", indexArrTwo)
	}
}
