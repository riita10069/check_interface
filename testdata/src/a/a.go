package a

import "fmt"

type a struct {}

type aInterface interface {
	f(a int64)
	g(a int64) // want "NG"
}

func (this a)f(a int64) {
	fmt.Println(a)
}


