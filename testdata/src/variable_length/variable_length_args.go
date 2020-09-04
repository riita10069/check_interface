package variable_length

// 可変長引数

// 正常系
type IVariableArgs interface {
	Invoke(a int, values ...int) string
}

type VariableArgs struct {
}

func (VariableArgs) Invoke(a int, values ...int) string {
	return ""
}

// 異常系
type NIVariableArgs interface { // want "not implemented"
	Invoke(a string, values ...int) string
	Invoke2(a string, values ...int) string
}

type NVariableArgs struct {
}

func (NVariableArgs) Invoke(a string, values ...int) string {
	return ""
}
