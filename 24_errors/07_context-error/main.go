package main

import (
	"fmt"
	"log"
)

// Example of using the error interface on a custrom structure

// NorgateMathError is
type NorgateMathError struct {
	lat, long string
	err       error
}

func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-1)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math redux: square root of neg number %v", f)
		return 0, &NorgateMathError{"50.22 N", "90.3 W", nme}
	}
	return 42, nil
}
