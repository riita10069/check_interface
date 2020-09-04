package slice

// 引数にスライスを含む

//正常系
type IParamSlice interface {
	Invoke([]int) string
}

type ParamSlice struct {
}

func (ParamSlice) Invoke([]int) string {
	return ""
}

//異常系
type NIParamSlice interface { // want "not implemented"
	Invoke([]int32) string
	Invoke2([]int32) string
}

type NParamSlice struct {
}

func (NParamSlice) Invoke([]int32) string {
	return ""
}
