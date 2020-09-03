package b

import "fmt"

var _ BInterface = (bs)(nil)

type bs struct{} // want "bf2がたりてないよ"

type BInterface interface { // want "not implemented"
	bf(z int64) (int64, int64, string)
	bf2(z int64) (int64, int64, string)
}

func (this bs) bf(a int64) (b, c int64, d string) {
	fmt.Println(a)
	return 0, 0, ""
}
