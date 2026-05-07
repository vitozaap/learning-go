package slices

import (
	"errors"
	"fmt"
)

func Slices() error {
	test := []string{"t", "e", "s", "t"}

	names := []string{"Jonh", "Sophia", "Victor"}
	//Define slice com a função make, que recebe o tipo, length e se precisar, capacity
	var testCopy []string = make([]string, cap(test))

	if len(test) <= 0 {
		return errors.New("The slice's length is less or equal to 0")
	}

	//Copia test para newSlice
	copy(testCopy, test)

	names = append(names, "Phelipyz")
	fmt.Println(names)
	test = append(names[:2], "Joao")

	fmt.Println(test)
	//Mostra a length do slice com a função len()
	fmt.Printf("test slice is of length: %d \n", len(test))
	//Mostra a capacity do slice com a função cap()
	fmt.Printf("test slice is of capacity: %d \n", cap(test))
	return nil
}
