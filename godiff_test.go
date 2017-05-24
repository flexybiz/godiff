package main

import "testing"

func TestDiff(t *testing.T) {
	arr1 := []string{"line 1"}
	arr2 := []string{"line 3", "line 2", "line 1"}

	res, ser := NotInSecondWithSort(arr1, arr2)
	if res[0] != "line 2" {
		t.Error(res)
	}
	if len(ser) != 0 {
		t.Error("Second array should be empty")
	}
}

func BenchmarkNotInSecondWithSort(b *testing.B) {
	arr1 := []string{"list 1"}
	arr2 := []string{"list 3", "list 2", "list 1"}
	for i := 0; i < b.N; i++ {
		NotInSecondWithSort(arr1, arr2)
	}
}
