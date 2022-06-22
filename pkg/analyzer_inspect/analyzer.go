package analyzer_inspect

import (
	"fmt"
	"go/ast"
	"go/token"
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

var updateNameList = []string{"Save", "Update"}

func run(pass *analysis.Pass) (interface{}, error) {
	// pass.ResultOf[inspect.Analyzer] will be set if we've added inspect.Analyzer to Requires.
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{ // filter needed nodes: visit only them
		(*ast.FuncDecl)(nil),
	}

	inspector.Nodes(nodeFilter, func(node ast.Node, push bool) bool {
		if !push {
			return true
		}
		callExpr := node.(*ast.FuncDecl)
		hasCallWhere := false
		var WherePos token.Pos
		for _, body := range callExpr.Body.List {
			if _, ok := body.(*ast.ExprStmt); !ok {
				continue
			}
			if call2, ok := body.(*ast.ExprStmt).X.(*ast.CallExpr); ok {
				switch node := call2.Fun.(type) {
				case *ast.SelectorExpr:
					iterator(&node.X, &hasCallWhere, &WherePos, pass)
					if node.Sel != nil {
						isIdent(node.Sel, &hasCallWhere, &WherePos, pass)
					}
				case *ast.Ident:
					isIdent(node, &hasCallWhere, &WherePos, pass)
				}
			}
		}
		return true
	})

	return nil, nil
}

func iterator(expr *ast.Expr, hasCallWhere *bool, WherePos *token.Pos, pass *analysis.Pass) {
	switch temp := (*expr).(type) {
	case *ast.SelectorExpr:
		iterator(&temp.X, hasCallWhere, WherePos, pass)
		if temp.Sel != nil {
			isIdent(temp.Sel, hasCallWhere, WherePos, pass)
		}
	case *ast.CallExpr:
		iterator(&temp.Fun, hasCallWhere, WherePos, pass)
	case *ast.Ident:
		isIdent(temp, hasCallWhere, WherePos, pass)
	}
}

func in(target string, strArray []string) bool {
	for _, element := range strArray {
		if target == element {
			return true
		}
	}
	return false
}

func isIdent(ident *ast.Ident, hasCallWhere *bool, WherePos *token.Pos, pass *analysis.Pass) {
	callName := ident.Name
	if callName == "Where" {
		isTrue := true
		hasCallWhere = &isTrue
		pos := ident.Pos()
		WherePos = &pos
	}

	if in(callName, updateNameList) && (!*hasCallWhere || *WherePos > ident.Pos()) {
		fmt.Println(ident.Pos(), "use orm update must first call Where()")
		pass.Reportf(ident.Pos(), "use orm update must first call Where()")
	}
	fmt.Println(ident.Name, ident.Pos())
}
