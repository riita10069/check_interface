package b

import "fmt"

type b struct {}

type bInterface interface {
	f(a int64)
}

func (this b)f(a int64) {
	fmt.Println(a)
}

