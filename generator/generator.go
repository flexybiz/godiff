package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"strconv"
	"time"
)

func writeFiles(toDel []int, arr []string) {
	f1, err := os.Create("../example/first.txt")
	defer f1.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	f2, err := os.Create("../example/second.txt")
	defer f2.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i, el := range arr {
		_, err := f1.WriteString(el + "\n")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if excludeString(i, toDel) {
			continue
		}
		_, err = f2.WriteString(el + "\n")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func genRandomNumbers(val int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := []int{}
	// exclude random number of random numbers
	el := rand.Intn(20)
	for i := 0; i < el; i++ {
		arr = append(arr, rand.Intn(val))
	}
	return arr
}

func excludeString(el int, arr []int) bool { // arr must be sorted!
	for i := 0; i < len(arr); i++ {
		if el < arr[i] {
			break
		}
		if el == arr[i] {
			return true
		}
	}
	return false
}

func main() {
	runtime.GOMAXPROCS(4)
	if len(os.Args) <= 1 {
		fmt.Println("Usage: generator <num of uuids>")
		os.Exit(1)
	}
	start := time.Now()
	// Creating UUID Version 4
	val, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	toDel := genRandomNumbers(val)
	sort.Ints(toDel)
	fmt.Printf("Exclude %d elements\n", len(toDel))
	arr := []string{}
	for i := 0; i < val; i++ {
		arr = append(arr, fmt.Sprintf("%s", uuid.NewV4()))
	}
	writeFiles(toDel, arr)
	fmt.Println(time.Since(start))
}
