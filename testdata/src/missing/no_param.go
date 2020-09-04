package missing

type p struct{}

type PInterface interface {
	f() (int64, int64, string)
}

func (this p) f() (b int64, c int64, d string) {
	return 0, 0, ""
}

type pbad struct {}

type PbadInterface interface { // want "not implemented"
	cbad() (int64, int64, int64)
	cbad2() (int64, int64, string)
}

func (this pbad) cbad() (b, c, d int64) {
	return 0, 0, 0
}
