package check_interface

import (
	"fmt"
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

	functionFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	interfaceFilter := []ast.Node{
		(*ast.InterfaceType)(nil),
	}

	//signatureMap := make(map[string][]string)

	var tmp *types.Signature

	inspect.Preorder(functionFilter, func(funcNode ast.Node) {
		switch funcNode := funcNode.(type) {
		case *ast.FuncDecl:

			tmp = pass.TypesInfo.ObjectOf(funcNode.Name).Type().(*types.Signature)
			fmt.Println("tmp")
			fmt.Println(tmp)
			//
			//var recv, params, ret, name string
			//if funcNode.Recv != nil {
			//	recv = fmt.Sprint(funcNode.Recv.List[0].Type)
			//}
			//if funcNode.Type.Params != nil {
			//	params = getString(funcNode.Type.Params.List)
			//}
			//if funcNode.Type.Results != nil {
			//	ret = getString(funcNode.Type.Results.List)
			//}
			//name = funcNode.Name.Name
			//
			//signature := strings.Join([]string{name, params, ret}, "/")
			//
			//if v, ok := signatureMap[signature]; ok {
			//	signatureMap[signature] = append(v, recv)
			//} else {
			//	signatureMap[signature] = []string{recv}
			//}

		}
	})

	inspect.Preorder(interfaceFilter, func(interfaceNode ast.Node) {
		switch interfaceNode := interfaceNode.(type) {
		case *ast.InterfaceType:

			//var params, ret, name string
			methodList := interfaceNode.Methods.List
			//var once sync.Once
			// 実装してあるstructを保存する map[構造体名]実装しているか
			//implements := map[string]bool{}
			for _, methodField := range methodList {

				tmp2 := pass.TypesInfo.ObjectOf(methodField.Names[0]).Type().(*types.Signature)
				fmt.Println("tmp2")
				fmt.Println(tmp2)

				fmt.Println(types.Identical(tmp, tmp2))

				//switch methodType := methodField.Type.(type) {
				//case *ast.FuncType:
				//	if methodType.Params != nil {
				//		params = getString(methodType.Params.List)
				//	}
				//	if methodType.Results != nil {
				//		ret = getString(methodType.Results.List)
				//	}
				//	name = methodField.Names[0].Name
				//
				//	signature := strings.Join([]string{name, params, ret}, "/")
				//
				//	recv, ok := signatureMap[signature]
				//	if !ok {
				//		// 実装されている構造体が１つもなかった場合にimplementsをnilにする
				//		implements = nil
				//		break
				//	}
				//
				//	// 最初のメソッドで該当する構造体をimplementsに格納
				//	once.Do(func() {
				//		for _, s := range recv {
				//			implements[s] = true
				//		}
				//	})
				//
				//	for k := range implements {
				//		for _, s := range recv {
				//			// implementとstructで同値の物がないときfalseに更新
				//			if k != s {
				//				implements[k] = false
				//			}
				//		}
				//	}
				//}
			}

			//if implements == nil {
			//	pass.Reportf(interfaceNode.Pos(), "not implemented")
			//}
			//
			//isImplement := true
			//for _, implement := range implements {
			//	isImplement = isImplement && implement
			//}
			//if !isImplement {
			//	pass.Reportf(interfaceNode.Pos(), "not implemented")
			//
			//}
		}
	})

	return nil, nil
}

func getString(lists []*ast.Field) string {
	str := ""
	for _, list := range lists {
		str += fmt.Sprint(list.Type)
		str += ","
	}
	return str
}
