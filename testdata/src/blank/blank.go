package blank

type g struct{}

type GInterface interface {
	gf(a int64) (int64, int64, string)
}

func (this g) gf(_ int64) (a int64, c int64, d string) {
	return 0, 0, ""
}

type gbad struct {}

type GbadInterface interface { // want "not implemented"
	gbad(z int64) (int64, int64, int64)
	gbad2(z int64) (int64, int64, string)
}

func (this g) gbad(_ int64) (b, c, d int64) {
	return 0, 0, 0
}

//  blank variable in returns
type g2 struct{}

type G2Interface interface {
	g2f(a int64) (int64, int64, string)
}

func (this g2) g2f(b int64) (_ int64, c int64, d string) {
	return 0, 0, ""
}

type g2bad struct {}

type G2badInterface interface { // want "not implemented"
	cbad(z int64) (int64, int64, int64)
	cbad2(z int64) (int64, int64, string)
}

func (this g2) g2bad(_ int64) (_, c, d int64) {
	return 0, 0, 0
}
