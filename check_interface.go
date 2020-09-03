package check_interface

import (
	"fmt"
	"go/ast"
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

	interfaceFilter := []ast.Node {
		(*ast.InterfaceType)(nil),
	}

	hash := make(map[string]*string)

	inspect.Preorder(functionFilter, func(funcNode ast.Node) {
		switch funcNode := funcNode.(type) {
		case *ast.FuncDecl:
			var recv, params, ret string
			if funcNode.Recv != nil {
				recv = fmt.Sprint(funcNode.Recv.List[0].Type)
			}
			if funcNode.Type.Params != nil {
				params = getString(funcNode.Type.Params.List)
			}
			if funcNode.Type.Results != nil {
				ret = getString(funcNode.Type.Results.List)
			}

			signature := params + "/" + ret


			if v, ok := hash[signature]; ok {
				*hash[signature] += *v + ","
			} else {
				hash[signature] = &recv
			}

		}
	})

	inspect.Preorder(interfaceFilter, func(interfaceNode ast.Node) {
		switch interfaceNode := interfaceNode.(type) {
		case *ast.InterfaceType:
			var recv, params, ret string
			methodList := interfaceNode.Methods.List
			for _, methodField := range methodList {
				switch methodType := methodField.Type.(type) {
				case *ast.FuncType:
					if methodType.Params != nil {
						params = getString(methodType.Params.List)
					}
					if methodType.Results != nil {
						ret = getString(methodType.Results.List)
					}
					signature := params + "/" + ret

					recv = *hash[signature]
					println(recv)
				}
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

