package b

import "fmt"

type b struct {}


type BInterface interface {
	f(a int64)
}

func (this b)f(a int64) {
	fmt.Println(a)
}
