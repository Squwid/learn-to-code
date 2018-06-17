package main

import (
	"fmt"
)

func main() {
	// slice of slice of string
	records := make([][]string, 0)
	// student 1
	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "74.00"
	// store the record
	records = append(records, student1)

	// student2
	student2 := make([]string, 4)
	student2[0] = "Gomez"
	student2[1] = "Lisa"
	student2[2] = "92.00"
	student2[3] = "93.00"
	//Store the record
	records = append(records, student2)

	fmt.Println(records)
}
