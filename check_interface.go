package check_interface

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"reflect"
	"strings"
	"sync"
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

type Signature struct {
	Name   string
	Param  []string
	Result []string
}

func (sig *Signature) Equal(signature Signature) bool {
	return reflect.DeepEqual(sig, signature)
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// map[シグネチャ]struct名
	signatureMap := map[string][]string{}

	// funcを探索してstructMap
	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:

			var signature Signature
			signature.Name = n.Name.Name
			if n.Type.Params != nil {
				for _, param := range n.Type.Params.List {
					signature.Param = append(signature.Param, fmt.Sprint(param.Type))
				}
			}
			if n.Type.Results != nil {
				for _, result := range n.Type.Results.List {
					signature.Result = append(signature.Result, fmt.Sprint(result.Type))
				}
			}
			name := fmt.Sprintf("name:%s", signature.Name)
			params := fmt.Sprintf("params:%s", strings.Join(signature.Param, ","))
			results := fmt.Sprintf("results:%s", strings.Join(signature.Result, ","))

			key := strings.Join([]string{name, params, results}, ",")

			if n.Recv != nil {
				recv := fmt.Sprint(n.Recv.List[0].Type)
				signatureMap[key] = append(signatureMap[key], recv)
			}
		}
	})

	// interfaceのみを取得
	nodeFilter = []ast.Node{
		(*ast.InterfaceType)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.InterfaceType:
			if n.Incomplete {
				pass.Reportf(n.Pos(), "methods are missing in the Methods list")
				break
			}

			var once sync.Once

			// 実装してあるstructを保存する map[構造体名]実装しているか
			implements := map[string]bool{}

			// interfaceのメソッド毎に処理
			for _, method := range n.Methods.List {
				if implements == nil {
					break
				}
				// funcTypeにアップキャスト
				switch mType := method.Type.(type) {
				case *ast.FuncType:

					var signature Signature
					signature.Name = method.Names[0].Name
					if mType.Params != nil {
						for _, param := range mType.Params.List {
							signature.Param = append(signature.Param, fmt.Sprint(param.Type))
						}
					}
					if mType.Results != nil {
						for _, result := range mType.Results.List {
							signature.Result = append(signature.Result, fmt.Sprint(result.Type))
						}
					}
					name := fmt.Sprintf("name:%s", signature.Name)
					params := fmt.Sprintf("params:%s", strings.Join(signature.Param, ","))
					results := fmt.Sprintf("results:%s", strings.Join(signature.Result, ","))

					key := strings.Join([]string{name, params, results}, ",")
					structs, ok := signatureMap[key]

					if !ok {
						// 実装されている構造体が１つもなかった場合にimplementsをnilにする
						implements = nil
						break
					}

					// 最初のメソッドで該当する構造体をimplementsに格納
					once.Do(func() {
						for _, s := range structs {
							implements[s] = true
						}
					})

					for k := range implements {
						for _, s := range structs {
							// implementとstructで同値の物がないときfalseに更新
							if k != s {
								implements[k] = false
							}
						}
					}
				}
			}

			if implements == nil {
				pass.Reportf(n.Pos(), "not implemented")
			}

			isImplement := true
			for _, implement := range implements {
				isImplement = isImplement && implement
			}
			if !isImplement {
				pass.Reportf(n.Pos(), "not implemented")
			}
		}
	})

	return nil, nil
}
