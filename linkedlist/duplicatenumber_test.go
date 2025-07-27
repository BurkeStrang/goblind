package linkedlist

import (
	"testing"
)

func TestFindDuplicate_Case1(t *testing.T) {
	nums := []int{1, 3, 4, 2, 2}
	expected := 2
	result := FindDuplicate(nums)
	if result != expected {
		t.Errorf("FindDuplicate(%v) = %d; want %d", nums, result, expected)
	}
}

func TestFindDuplicate_Case2(t *testing.T) {
	nums := []int{3, 1, 3, 4, 2}
	expected := 3
	result := FindDuplicate(nums)
	if result != expected {
		t.Errorf("FindDuplicate(%v) = %d; want %d", nums, result, expected)
	}
}

func TestFindDuplicate_Case3(t *testing.T) {
	nums := []int{3, 3, 3, 3, 3}
	expected := 3
	result := FindDuplicate(nums)
	if result != expected {
		t.Errorf("FindDuplicate(%v) = %d; want %d", nums, result, expected)
	}
}

func TestFindDuplicate_Case4(t *testing.T) {
	nums := []int{1, 1, 1}
	expected := 1
	result := FindDuplicate(nums)
	if result != expected {
		t.Errorf("FindDuplicate(%v) = %d; want %d", nums, result, expected)
	}
}

func TestFindDuplicate_Case5(t *testing.T) {
	nums := []int{2, 2, 2, 2, 2}
	expected := 2
	result := FindDuplicate(nums)
	if result != expected {
		t.Errorf("FindDuplicate(%v) = %d; want %d", nums, result, expected)
	}
}
