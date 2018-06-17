package main

import "fmt"

func main() {

	// Working with slices
	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)

	fmt.Println(mySlice[2:4])          // slicing a slice
	fmt.Println(mySlice[1])            // index access
	fmt.Println("myString"[2])         // 83
	fmt.Println(string("myString"[2])) // S

	// Part 2
	fmt.Printf("\nPart2\n")

	// Different ways to create slices
	// make([]T, length, capacty)
	_ = make([]int, 50, 100)
	_ = new([100]int)[0:50] // "slicing"
	_ = make([]int, 50)     // makes both the length and capacity 50

	//continued()
	appending()
}

func continued() {
	mySlice := make([]int, 0, 5) //
	fmt.Println("-----------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice)) // print length
	fmt.Println(cap(mySlice)) // print capacity of slice
	fmt.Println("----------")

	// appending a slice
	// capacity doubles when it reaches the max
	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("Len: %v \t Capacity: %v \t Value: %v\n", len(mySlice), cap(mySlice), mySlice[i])
	}
}

func appending() {
	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...) // append a whole slice to another whole slice

	fmt.Println(mySlice) // [Monday Tuesday Wednesday Thursday Friday]

	mySlice = append(mySlice[:2], mySlice[3:]...) // remove an item from a slice
	fmt.Println(mySlice)                          // [Monday Tuesday Thursday Friday]
}
