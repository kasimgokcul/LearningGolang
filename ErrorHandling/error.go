package errorhandling

import (
	"fmt"
	"strconv"
)

func Error() {

	a := "234"
	b := "test"

	a1, err := strconv.Atoi(a)
	if err != nil {
		fmt.Printf("Cannot convert value %s\n", a)
	} else {
		fmt.Printf("Transaction successful, new value %v\n", a1)
	}

	b1, err := strconv.Atoi(b)
	if err != nil {
		fmt.Printf("Cannot convert value %s\n", b)
	} else {
		fmt.Printf("Transaction successful, new value %v\n", b1)
	}
}
