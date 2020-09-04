package pointer

// 戻り値にポインターを含む

// 正常系
type IResultPointer interface {
	Invoke(int64) *string
}

type ResultPointer struct {
}

func (rp ResultPointer) Invoke(int64) *string {
	str := ""
	return &str
}

// 異常系
type NIResultPointer interface { // want "not implemented"
	Invoke(int64) *int32
	Invoke2(int64) *int32
}

type NResultPointer struct {
}

func (rp NResultPointer) Invoke(int64) *int32 {
	str := int32(0)
	return &str
}
