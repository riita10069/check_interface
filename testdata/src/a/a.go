package a

import "fmt"

type a struct{}

type aInterface interface {
	f(a int64) string
}

type bInterface interface {
	g(a int64)
}

type cInterface interface {
	f(a int64) string
	g(a int64)
}

type dInterface interface { // want "methods are missing in the Methods list"
}

type eInterface interface { // want "not implemented"
	f(a int64) string
	g(a int64)
	h(a int64)
}

func (this a) f(a int64) string {
	fmt.Println(a)
	return ""
}

func (this a) g(a int64) {
	fmt.Println(a)
}
