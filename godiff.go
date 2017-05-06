package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"sync"
)

// Read file and return it as an arrays of strings
func ReadFile(fname string) []string {
	inArr := []string{}
	if file, err := os.Open(fname); err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			inArr = append(inArr, scanner.Text())
		}
		fmt.Printf("File %v has %v strings\n", fname, len(inArr))
		// check for errors
		if err = scanner.Err(); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
	return inArr
}

// Returns two slices with all strings from second(first) array
// that is not in the first(second) array
func NotInSecondWithSort(arr1 []string, arr2 []string) ([]string, []string) {
	// Sort arrays using goroutines
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		sort.Strings(arr1)
	}()
	go func() {
		defer wg.Done()
		sort.Strings(arr2)
	}()
	wg.Wait()
	// Finding differences in sorted arrays
	arr := []string{} // not in first
	rra := []string{} // not in second
	i := 0
	j := 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			rra = append(rra, arr1[i])
			i++
		} else if arr2[j] < arr1[i] {
			arr = append(arr, arr2[j])
			j++
		} else {
			i++
			j++
		}
	}
	return arr, rra
}

func main() {
	first := ReadFile("example/first.txt")
	second := ReadFile("example/second.txt")
	res, ser := NotInSecondWithSort(first, second)
	fmt.Printf("\nStrings from the second file that is not in first:\n%v\n", res)
	fmt.Printf("\nStrings from the first file that is not in second:\n%v\n", ser)
}
