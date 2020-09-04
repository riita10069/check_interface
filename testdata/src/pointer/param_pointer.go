package pointer

// 引数にポインターを含む

// 正常系
type IParamPointer interface {
	Invoke(*int64) string
}

type ParamPointer struct {
}

func (pp ParamPointer) Invoke(*int64) string {
	return ""
}

// 異常系
type NIParamPointer interface { // want "not implemented"
	Invoke(*int64) int32
	Invoke2(*int64) int32
}

type NParamPointer struct {
}

func (pp NParamPointer) Invoke(*int64) int32 {
	return 0
}
