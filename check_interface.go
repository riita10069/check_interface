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
	functionFilter := []ast.Node{(*ast.FuncDecl)(nil)}

	// map[シグネイチャ（メソッド名・引数・戻り値）][]struct
	signatureMap := make(map[*types.Object][]types.Type)

	// 実装してあるstructを保存する map[構造体名]実装メソッドカウント
	implements := map[types.Type]int{}

	inspect.Preorder(functionFilter, func(funcNode ast.Node) {
		switch funcNode := funcNode.(type) {
		case *ast.FuncDecl:
			signatureObj := pass.TypesInfo.ObjectOf(funcNode.Name)
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
			//var params, ret, name string
			methodList := interfaceNode.Methods.List

			//fmt.Print("implements = ", implements)

			for _, methodField := range methodList {

				signatureObj := pass.TypesInfo.ObjectOf(methodField.Names[0])

				//switch methodType := methodField.Type.(type) {
				//case *ast.FuncType:
				//if methodType.Params != nil {
				//	params = getString(methodType.Params.List)
				//}
				//if methodType.Results != nil {
				//	ret = getString(methodType.Results.List)
				//}
				//name = methodField.Names[0].Name
				//
				//signature := strings.Join([]string{name, params, ret}, "/")
				//
				//recv, ok := signatureMap[signature]
				//if !ok {
				//	// 実装されている構造体が１つもなかった場合にimplementsをnilにする
				//	continue
				//}
				//fmt.Println(methodField.Names[0].Name)
				//recvs, ok := signatureMap[signatureObj]
				//fmt.Println("recvs = ", recvs)

				for signature, _ := range signatureMap {

					if types.Identical(signatureObj.Type().(*types.Signature), (*signature).Type().(*types.Signature)) && signatureObj.Name() == (*signature).Name() {
						implements[signature]++
					}
				}
				fmt.Println(*signatureObj)
				fmt.Println(signatureMap)

				//if !ok {
				//	continue
				//}

				//for _, s := range recvs {
				//	implements[s] = implements[s] + 1
				//}
			}

			//max, _ := maxMap(implements)
			//if key != nil {
			//	pass.Reportf(interfaceNode.Pos(), "Is it %s you want to implement?", key)
			//}


			for _, impl := range implements {
				if impl < len(methodList) {
					//fmt.Println(k)
					//fmt.Println(impl)
					//fmt.Println(len(methodList))
					pass.Reportf(interfaceNode.Pos(), "not implemented")
				}
			}

			//if max < len(methodList) {
			//	fmt.Println(len(implements))
			//	fmt.Println(len(methodList))
			//	pass.Reportf(interfaceNode.Pos(), "not implemented")
			//	//_, key := maxMap(implements)
			//	//if key != nil {
			//	//	// pass.Reportf(interfaceNode.Pos(), "Is it %s you want to implement?", key)
			//	//}
			//} else {
			//	fmt.Println("OK")
			//}
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

func getString(lists []*ast.Field) string {
	str := ""
	for _, list := range lists {
		str += fmt.Sprint(list.Type)
		str += ","
	}
	return str
}
