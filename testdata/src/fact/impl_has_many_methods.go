package fact

// 実装側のメソッドが多い時

// 正常系
type IImplHas interface {
	Invoke(int) string
	Invoke2(int) string
}

type ImplHas struct {
}

func (ImplHas) Invoke(int) string {
	return ""
}

func (ImplHas) Invoke2(int) string {
	return ""
}

func (ImplHas) Invoke3(int) string {
	return ""
}

type NIImplHas interface { // want "not implemented"
	Invoke(int32) string
	Invoke2(int32) string
	Invoke3()
}

type NImplHas struct {
}

func (NImplHas) Invoke(int32) string {
	return ""
}

func (NImplHas) Invoke2(int32) string {
	return ""
}

func (NImplHas) Invoke3(int32) string {
	return ""
}
