package main

import (
	"reflect"
	"testing"
)

func TestEx1(t *testing.T) {
	const n = 4
	a := []int{1, 2, 4, 7}
	const k = 13

	const success = true
	sumList := []int{2, 4, 7}

	if x, y := Ex(n, a, k); x != success || !reflect.DeepEqual(y, sumList) {
		t.Errorf("Ex(%v, %v, %v) = %v, %v, want %v, %v", n, a, k, x, y, success, sumList)
	}
}

func TestEx2(t *testing.T) {
	const n = 4
	a := []int{1, 2, 4, 7}
	const k = 15

	const success = false
	sumList := []int{}

	if x, y := Ex(n, a, k); x != success || !reflect.DeepEqual(y, sumList) {
		t.Errorf("Ex(%v, %v, %v) = %v, %v, want %v, %v", n, a, k, x, y, success, sumList)
	}
}
