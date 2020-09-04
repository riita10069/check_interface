package short_hand

// シグネイチャ省略系

// 正常系
type IShortSignature interface {
	Invoke(f, g string, h, i int) (a, b, c int, d, e string)
}

type ShortSignature struct {
}

func (sp ShortSignature) Invoke(f, g string, h, i int) (a, b, c int, d, e string) {
	return 0, 0, 0, "", ""
}

// 異常系
type NIShortSignature interface { // want "not implemented"
	Invoke(f, g string, h, i int32) (a, b, c int, d, e string)
	Invoke2(f, g string, h, i int32) (a, b, c int, d, e string)
}

type NShortSignature struct {
}

func (sp NShortSignature) Invoke(f, g string, h, i int32) (a, b, c int, d, e string) {
	return 0, 0, 0, "", ""
}
