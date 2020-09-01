package a

import "fmt"

type a struct {}

type aInterface interface { // want "methods are missing in the Methods list"
	f(a int64)
	g(a int64)
}

func (this a)f(a int64) {
	fmt.Println(a)
}


