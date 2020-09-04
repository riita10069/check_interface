package short_hand

// 戻り値省略系

// 正常系
type IShortResult interface {
	Invoke(string) (a, b, c int, d, e string)
}

type ShortResult struct {
}

func (sp ShortResult) Invoke(string) (a, b, c int, d, e string) {
	return 0, 0, 0, "", ""
}

// 異常系
type NIShortResult interface { // want "not implemented"
	Invoke(int) (a, b, c int, d, e string)
	Invoke2(string) (a, b, c int, d, e string)
}

type NShortResult struct {
}

func (sp NShortResult) Invoke(int) (a, b, c int, d, e string) {
	return 0, 0, 0, "", ""
}
