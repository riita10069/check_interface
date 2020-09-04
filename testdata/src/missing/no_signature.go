package missing

type s struct{}

type SInterface interface {
	f()
}

func (this s) f() {
	return
}

type sbad struct {}

type SbadInterface interface { // want "not implemented"
	g(int64, int64)
	h()
}

func (this sbad) g() {
	return
}
