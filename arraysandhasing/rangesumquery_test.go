package arraysandhasing

import (
	"testing"
)

func TestRangeSumQuery_Example1(t *testing.T) {
	numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})
	if got := numArray.SumRange(0, 2); got != 1 {
		t.Errorf("SumRange(0,2) = %d; want 1", got)
	}
	if got := numArray.SumRange(2, 5); got != -1 {
		t.Errorf("SumRange(2,5) = %d; want -1", got)
	}
	if got := numArray.SumRange(0, 5); got != -3 {
		t.Errorf("SumRange(0,5) = %d; want -3", got)
	}
	if got := numArray.SumRange(0, 0); got != -2 {
		t.Errorf("SumRange(0,0) = %d; want -2", got)
	}
}

func TestRangeSumQuery_Example2(t *testing.T) {
	numArray := Constructor([]int{1, 2, 3, 4, 5, 6})
	if got := numArray.SumRange(0, 2); got != 6 {
		t.Errorf("SumRange(0,2) = %d; want 6", got)
	}
	if got := numArray.SumRange(2, 5); got != 18 {
		t.Errorf("SumRange(2,5) = %d; want 18", got)
	}
	if got := numArray.SumRange(0, 5); got != 21 {
		t.Errorf("SumRange(0,5) = %d; want 21", got)
	}
	if got := numArray.SumRange(0, 0); got != 1 {
		t.Errorf("SumRange(0,0) = %d; want 1", got)
	}
}
