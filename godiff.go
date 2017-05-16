package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

// Read file and return it as arrays of strings
func ReadFile(fname string) []string {
	inArr := []string{}
	if file, err := os.Open(fname); err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			// trim leading & trailing spaces
			inArr = append(inArr, strings.TrimSpace(scanner.Text()))
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
	// Sort arrays
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { defer wg.Done(); sort.Strings(arr1) }()
	go func() { defer wg.Done(); sort.Strings(arr2) }()
	wg.Wait()
	// Find differences in sorted arrays
	arr := []string{} // not in first
	rra := []string{} // not in second

	// 1. len(arr1) == 0 | nothing in first file
	// 2. len(arr2) == 0 | nothing in second file
	// 3. len(arr1) & len(arr2) != 0

	if len(arr1) == 0 { // 1
		rra = arr2
	} else if len(arr2) == 0 { // 2
		arr = arr1
	} else { // 3
		i := 0
		j := 0
		// main loop, go through sorted arrays
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
		// collect tails
		if i < len(arr1) {
			rra = append(rra, arr1[i:]...)
		}
		if j < len(arr2) {
			arr = append(arr, arr2[j:]...)
		}
	}

	return arr, rra
}

// Write differences into files
func WriteFile(arr []string, fname string) {
	if fout, err := os.Create(fname); err == nil {
		defer fout.Close()
		for _, str := range arr {
			fout.WriteString(str + "\n")
		}
	} else {
		fmt.Println(err)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	if len(os.Args) != 3 {
		fmt.Println("Usage: godiff <first file> <second file>")
		os.Exit(0)
	}

	start := time.Now()
	first := []string{}
	second := []string{}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { defer wg.Done(); first = ReadFile(os.Args[1]) }()
	go func() { defer wg.Done(); second = ReadFile(os.Args[2]) }()
	wg.Wait()

	res, ser := NotInSecondWithSort(first, second)

	wg.Add(2)
	go func() {
		defer wg.Done()
		WriteFile(res, "diff_f_s.txt")
		fmt.Printf("\nFound %v strings from second file that is not in first (saved in diff_f_s.txt)\n", len(res))
	}()
	go func() {
		defer wg.Done()
		WriteFile(ser, "diff_s_f.txt")
		fmt.Printf("Found %v strings from first file that is not in second (saved in diff_s_f.txt)\n", len(ser))
	}()
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("\nDone in %v\n", elapsed)
}
