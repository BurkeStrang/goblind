package linkedlist

import (
	"reflect"
	"testing"
)


func TestAddTwoNumbers_Example1(t *testing.T) {
	l1 := BuildList([]int{3, 5, 3})
	l2 := BuildList([]int{5, 6, 4})
	want := []int{8, 1, 8}
	got := AddTwoNumbers(l1, l2)
	if !reflect.DeepEqual(listToSlice(got), want) {
		t.Errorf("AddTwoNumbers(%v, %v) = %v, want %v", []int{2, 4, 3}, []int{5, 6, 4}, listToSlice(got), want)
	}
}

func TestAddTwoNumbers_Example2(t *testing.T) {
	l1 := BuildList([]int{0})
	l2 := BuildList([]int{0})
	want := []int{0}
	got := AddTwoNumbers(l1, l2)
	if !reflect.DeepEqual(listToSlice(got), want) {
		t.Errorf("AddTwoNumbers(%v, %v) = %v, want %v", []int{0}, []int{0}, listToSlice(got), want)
	}
}

func TestAddTwoNumbers_Example3(t *testing.T) {
	l1 := BuildList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := BuildList([]int{9, 9, 9, 9})
	want := []int{8, 9, 9, 9, 0, 0, 0, 1}
	got := AddTwoNumbers(l1, l2)
	if !reflect.DeepEqual(listToSlice(got), want) {
		t.Errorf("AddTwoNumbers(%v, %v) = %v, want %v", []int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, listToSlice(got), want)
	}
}
