package analyzer_inspect

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "goprintffuncname",
	Doc:      "Checks that printf-like functions are named with `f` at the end.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	// pass.ResultOf[inspect.Analyzer] will be set if we've added inspect.Analyzer to Requires.
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{ // filter needed nodes: visit only them
		(*ast.FuncDecl)(nil),
	}

	inspector.Nodes(nodeFilter, func(node ast.Node, push bool) bool {
		callExpr := node.(*ast.FuncDecl)
		fmt.Println(callExpr.Name.Name)
		for _, body := range callExpr.Body.List {
			call2 := body.(*ast.ExprStmt).X.(*ast.CallExpr)
			fmt.Println(call2.Fun.(*ast.Ident).Name, call2.Pos())
		}
		return true
	})

	return nil, nil
}
