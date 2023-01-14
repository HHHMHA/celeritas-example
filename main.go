package main

import (
	"fmt"

	"github.com/HHHMHA/celeritas"
)

func main() {
	result := celeritas.TestFunc(1, 1)
	fmt.Println(result)

	result = celeritas.TestFunc2(2, 1)
	fmt.Println(result)
}
