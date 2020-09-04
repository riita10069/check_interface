package slice

// 戻り値にスライスを含む

//正常系
type IResultSlice interface {
	Invoke(int) []string
}

type ResultSlice struct {
}

func (ResultSlice) Invoke(int) []string {
	return []string{""}
}

//異常系
type NIResultSlice interface { // want "not implemented"
	Invoke(int32) []string
	Invoke2(int32) []string
}

type NResultSlice struct {
}

func (NResultSlice) Invoke(int32) []string {
	return []string{""}
}
