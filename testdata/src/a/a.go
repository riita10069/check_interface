package a

import (
	"fmt"
)

type a struct {}

type AInterface interface { // want "not implemented"
	f(a int64)
	g(a int64)
}

func (this a)f(a int64, b string) (int64, error) {
	fmt.Println(a)
	return a, nil
}


