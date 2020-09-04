package check_interface

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "check_interface is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "check_interface",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	functionFilter := []ast.Node{(*ast.FuncDecl)(nil)}

	// map[メソッドの情報]構造体の配列(キーのメソッドを持っているもの）
	signatureMap := make(map[*types.Object][]types.Type)

	// 実装してあるstructを保存する map[構造体名]実装メソッドカウント
	implements := map[types.Type]int{}

	inspect.Preorder(functionFilter, func(funcNode ast.Node) {
		switch funcNode := funcNode.(type) {
		case *ast.FuncDecl:
			signatureObj := pass.TypesInfo.ObjectOf(funcNode.Name)
			if pass.TypesInfo.ObjectOf(funcNode.Name).Type().(*types.Signature).Recv() == nil {
				break
			}
			recv := pass.TypesInfo.ObjectOf(funcNode.Name).Type().(*types.Signature).Recv().Type()

			if v, ok := signatureMap[&signatureObj]; ok {
				signatureMap[&signatureObj] = append(v, recv)
			} else {
				signatureMap[&signatureObj] = []types.Type{recv}
			}

			// implementsをあらかじめ作っておく
			if _, ok := implements[recv]; !ok {
				implements[recv] = 0
			}
		}
	})

	interfaceFilter := []ast.Node{(*ast.InterfaceType)(nil)}
	inspect.Preorder(interfaceFilter, func(interfaceNode ast.Node) {
		switch interfaceNode := interfaceNode.(type) {
		case *ast.InterfaceType:

			// implementsの初期化
			for k, _ := range implements {
				implements[k] = 0
			}

			methodList := interfaceNode.Methods.List
			for _, methodField := range methodList {

				signatureObj := pass.TypesInfo.ObjectOf(methodField.Names[0])

				for signature, _ := range signatureMap {
					// シグネイチャと名前がsignatureMapに登録されているもののカウントを増やす
					if types.Identical(signatureObj.Type().(*types.Signature), (*signature).Type().(*types.Signature)) && signatureObj.Name() == (*signature).Name() {
						implements[(*signature).Type().(*types.Signature).Recv().Type()]++
					}
				}
			}

			max, _ := maxMap(implements)
			if max < len(methodList) {
				pass.Reportf(interfaceNode.Pos(), "not implemented")
				//if key != nil {
				//	pass.Reportf(interfaceNode.Pos(), fmt.Sprintf("Isn't this %s interface?", key.String() ))
				//}
			}
		}
	})
	return nil, nil
}

func maxMap(implements map[types.Type]int) (int, types.Type) {
	ret := 0
	var key types.Type
	for k, i := range implements {
		if i > ret {
			ret = i
			key = k
		}
	}
	return ret, key
}

