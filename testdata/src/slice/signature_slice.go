package slice

// シグネイチャにスライスを含む

//正常系
type ISignatureSlice interface {
	Invoke([]int) []string
}

type SignatureSlice struct {
}

func (SignatureSlice) Invoke([]int) []string {
	return []string{""}
}

//異常系
type NISignatureSlice interface { // want "not implemented"
	Invoke([]int32) []string
	Invoke2([]int32) []string
}

type NSignatureSlice struct {
}

func (NSignatureSlice) Invoke([]int32) []string {
	return []string{""}
}
