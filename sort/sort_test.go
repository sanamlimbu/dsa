package sort

import "testing"

func TestSelectionInt(t *testing.T) {
	arrOne := []int{5, 4, 3, 2, 0, 1}
	arrTwo := []int{20, 10, 12, -5, 1}

	wantArrOne := []int{0, 1, 2, 3, 4, 5}
	wantArrTwo := []int{-5, 1, 10, 12, 20}

	gotArrOne := SelectionInt(arrOne)
	gotArrTwo := SelectionInt(arrTwo)

	if compare(gotArrOne, wantArrOne) != true {
		t.Error("Array one is not sorted")
	}

	if compare(gotArrTwo, wantArrTwo) != true {
		t.Error("Array two is not sorted")
	}
}

func compare(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
