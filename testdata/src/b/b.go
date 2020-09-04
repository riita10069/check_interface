package b

import "fmt"

type bs struct{}

type BInterface interface {
	bf(z int64) (int64, int64, string)
	bf2(z int64) (string, string, string)
}

func (this bs) bf(a int64) (b, c int64, d string) {
	fmt.Println(a)
	return 0, 0, ""
}

func (this bs) bf2(a int64) (b, c , d string) {
	return "0", "0", ""
}
