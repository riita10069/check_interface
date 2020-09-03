package check_interface

import (
	"fmt"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"strings"
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
	functionFilter := []ast.Node{ (*ast.FuncDecl)(nil), }
	signatureMap := make(map[*types.Signature][]types.Type)
	var signatureObj *types.Signature

	inspect.Preorder(functionFilter, func(funcNode ast.Node) {
		switch funcNode := funcNode.(type) {
		case *ast.FuncDecl:
			signatureObj = pass.TypesInfo.ObjectOf(funcNode.Name).Type().(*types.Signature)
			recv := signatureObj.Recv().Type()
			if v, ok := signatureMap[signatureObj]; ok {
				signatureMap[signatureObj] = append(v, recv)
			}
		}
	})

	interfaceFilter := []ast.Node{ (*ast.InterfaceType)(nil), }
	inspect.Preorder(interfaceFilter, func(interfaceNode ast.Node) {
		switch interfaceNode := interfaceNode.(type) {
		case *ast.InterfaceType:
			var params, ret, name string
			methodList := interfaceNode.Methods.List
			// 実装してあるstructを保存する map[構造体名]実装しているか
			implements := map[string]int{}
			fmt.Print("implements = ", implements)
			for _, methodField := range methodList {
				switch methodType := methodField.Type.(type) {
				case *ast.FuncType:
					if methodType.Params != nil {
						params = getString(methodType.Params.List)
					}
					if methodType.Results != nil {
						ret = getString(methodType.Results.List)
					}
					name = methodField.Names[0].Name

					signature := strings.Join([]string{name, params, ret}, "/")

					recv, ok := signatureMap[signature]
					if !ok {
						// 実装されている構造体が１つもなかった場合にimplementsをnilにする
						continue
					}

					for _, s := range recv {
						_, notInit := implements[s]
						if notInit {
							implements[s] = implements[s] + 1
						} else {
							implements[s] = 1
						}
					}
				}
			}

			if len(implements) < len(methodList) {
				pass.Reportf(interfaceNode.Pos(), "not implemented")
				_, key := maxMap(implements)
				pass.Reportf(interfaceNode.Pos(),"Is it %s you want to implement?", key)
			} else {
				fmt.Println("OK")
			}
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
