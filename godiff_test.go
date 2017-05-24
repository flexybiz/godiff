package main

import "testing"

func BenchmarkNotInSecondWithSort(b *testing.B) {
	arr1 := []string{"list 1"}
	arr2 := []string{"list 3", "list 2", "list 1"}
	for i := 0; i < b.N; i++ {
		NotInSecondWithSort(arr1, arr2)
	}
}
