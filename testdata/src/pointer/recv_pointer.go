package pointer

// レシーバにポインタを使用しているパターン

// 正常系
type IRecvPinter interface {
	Invoke(int64) string
}

type RecvPointer struct {
}

func (rp *RecvPointer) Invoke(int64) string {
	return ""
}

// 異常系
type NIRecvPinter interface { // want "not implemented"
	Invoke(int64) int32
	Invoke2(int64) string
}

type NRecvPointer struct {
}

func (rp *NRecvPointer) Invoke(int64) int32 {
	return 0
}
