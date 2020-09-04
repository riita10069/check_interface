package missing

type r struct{}

type RInterface interface {
	f(int64, int64)
}

func (this r) f(a int64, b int64) {
	a += b
	return
}

type rbad struct {}

type RbadInterface interface { // want "not implemented"
	g(int64, int64)
	h(int64, int64)
}

func (this rbad) g(a int64, b int64) {
	a += b
	return
}
