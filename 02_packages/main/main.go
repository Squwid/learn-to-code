package main

import (
	"fmt"

	"github.com/Squwid/learn-to-code/02_packages/sutil"
)

func main() {
	k := sutil.Reverse("this")
	fmt.Println(k)
}
