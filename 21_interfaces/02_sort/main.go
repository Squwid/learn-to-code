package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

// Sort interfaces
func main() {
	studyGroup := people{"Ben", "Zeno", "Gnome", "Jason"}
	fmt.Println(studyGroup)
	fmt.Println("Is sorted?", sort.IsSorted(studyGroup))

	// Sort the study group
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
	fmt.Println("Is sorted?", sort.IsSorted(studyGroup))

	// Sorting a string slice
	names := []string{"twitter", "zebra", "ostrich", "animal"}
	sort.Strings(names)
	// sort.StringSlice(names).Sort() this works too
	fmt.Println(names)

	// reverse
	sort.Sort(sort.Reverse(sort.StringSlice(names)))
	fmt.Println(names)
}
