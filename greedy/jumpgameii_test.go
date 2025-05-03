package greedy

import "testing"

func TestJump_SimpleCase(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	expected := 2
	result := Jump(nums)
	if result != expected {
		t.Errorf("For nums %v, expected %d but got %d", nums, expected, result)
	}
}

func TestJump_WithZeroInMiddle(t *testing.T) {
	nums := []int{2, 3, 0, 1, 4}
	expected := 2
	result := Jump(nums)
	if result != expected {
		t.Errorf("For nums %v, expected %d but got %d", nums, expected, result)
	}
}

func TestJump_IncreasingSequence(t *testing.T) {
	nums := []int{1, 2, 3}
	expected := 2
	result := Jump(nums)
	if result != expected {
		t.Errorf("For nums %v, expected %d but got %d", nums, expected, result)
	}
}

func TestJump_AllOnes(t *testing.T) {
	nums := []int{1, 1, 1, 1}
	expected := 3
	result := Jump(nums)
	if result != expected {
		t.Errorf("For nums %v, expected %d but got %d", nums, expected, result)
	}
}
