package short_hand

// 引数省略系

// 正常系
type IShortParam interface {
	Invoke(a, b, c int, d, e string) string
}

type ShortParam struct {
}

func (sp ShortParam) Invoke(a, b, c int, d, e string) string {
	return ""
}

// 異常系
type NIShortParam interface { // want "not implemented"
	Invoke(a, b, c int, d, e string) int
	Invoke2(a, b, c int, d, e string) int
}

type NShortParam struct {
}

func (sp NShortParam) Invoke(a, b, c int, d, e string) int {
	return 0
}
