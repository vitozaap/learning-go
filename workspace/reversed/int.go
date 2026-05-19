package reversed

import (
	"errors"
	"strconv"
	"golang.org/x/example/hello/reverse"
)

func ReverseInt(i int) (string, error) {
	if i < 9 {
		return "", errors.New("Number should have at least two numerals.")
	}
	
	var reversedInt string = reverse.String(strconv.Itoa(i));
	return reversedInt, nil
}