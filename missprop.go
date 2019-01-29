package missprop

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "missprop",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

const Doc = "missprop is the tool to find missing props in composite literal"

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CompositeLit)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.CompositeLit:
			if len(n.Elts) == 0 {
				return
			}
			if _, ok := n.Elts[0].(*ast.KeyValueExpr); !ok {
				return
			}

			used := map[string]bool{}
			for _, elt := range n.Elts {
				used[elt.(*ast.KeyValueExpr).Key.(*ast.Ident).Name] = true
			}

			t := pass.TypesInfo.TypeOf(n).Underlying().(*types.Struct)
			for i := 0; i < t.NumFields(); i++ {
				name := t.Field(i).Name()
				if _, ok := used[name]; !ok {
					pass.Reportf(n.Pos(), "find missing props: %v", name)
				}
			}
		}
	})

	return nil, nil
}