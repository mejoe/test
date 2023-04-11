package main

import "fmt"

func t20230411() {
	//println("2023-04-11")
	//createSimpleSlice()
	//modifySlice()
	sliceWithTwoLevels()
}

func sliceWithTwoLevels() {
	l := 10
	a := make([][]int, l)
	fmt.Println(a)
	for i := 0; i < l; i++ {
		a[i] = append(a[i], i)
		for x := 0; x < l; x++ {
			a[i] = append(a[i], x)
		}

		fmt.Println(a[i])
	}
	fmt.Println(a)
}

// # Today's calistentics:
// ## For each of the following data structures:
// - slice
// - map
// - struct

// ## Demonstrate at least 3 ways each of:
// - to create them
// - to add data to them
// - to delete data

// ## bonus for combining the above:
// - a slice of structures
// - a slice of maps
// - a struct with slices

func createSimpleSlice() {
	// 1. initialize with 0 values (e.g. everything is an empty string)
	s := make([]string, 5)
	fmt.Println(s)
	// add some data
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"
	s[4] = "e"

	// 2. create with var and data
	var b = []string{"sara", "jane", "Matt"}
	fmt.Println(b)

	// 3. use the "new" keyword - which allocates memory
	var c = new([10]int)[0:10]
	fmt.Println(c)
}

func modifySlice() {
	s := make([]int, 10) // init with values = 0
	fmt.Println(len(s), s)
	// update values in the slice
	for i := 0; i <= 9; i++ {
		s[i] = i
	}
	// add new values, extending the length
	s = append(s, 1, 2)
	fmt.Println(len(s), s)
	// remove a value
	s = append(s[:5], s[6:]...)
	fmt.Println(len(s), s)

	// use my new delete slice function
	newSlice := deleteSliceValue(s, 3)
	fmt.Println(len(newSlice), newSlice)
}

func deleteSliceValue(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}
