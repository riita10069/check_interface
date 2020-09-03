package b

import "fmt"

type b struct{}

type BInterface interface {
	bf(a int64) (int64, string, string)
}

func (this b) bf(a int64) (b, c int64, d string) {
	fmt.Println(a)
	return 0, 0, ""
}
