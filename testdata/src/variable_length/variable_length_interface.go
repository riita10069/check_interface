package variable_length

// 可変長引数（interface{}のとき）

// 正常系
type IVariableInterface interface {
	Invoke(a int, opt ...interface{}) string
}

type VariableInterface struct {
}

func (VariableInterface) Invoke(a int, opt ...interface{}) string {
	return ""
}

// 異常系
type NIVariableInterface interface { // want "not implemented"
	Invoke(a string, opt ...interface{}) string
	Invoke2(a string, opt ...interface{}) string
}

type NVariableInterface struct {
}

func (NVariableInterface) Invoke(a string, opt ...interface{}) string {
	return ""
}
