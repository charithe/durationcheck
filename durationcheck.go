package durationcheck

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"go/types"
	"log"
	"os"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "durationcheck",
	Doc:      "check for two durations multiplied together",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	// if the package does not import time, it can be skipped from analysis
	if !hasImport(pass.Pkg, "time") {
		return nil, nil
	}

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeTypes := []ast.Node{
		(*ast.BinaryExpr)(nil),
	}

	inspect.Preorder(nodeTypes, check(pass))
	return nil, nil
}

func hasImport(pkg *types.Package, importPath string) bool {
	for _, imp := range pkg.Imports() {
		if imp.Path() == importPath {
			return true
		}
	}

	return false
}

// check contains the logic for checking that time.Duration is used correctly in the code being analysed
func check(pass *analysis.Pass) func(ast.Node) {
	return func(node ast.Node) {
		expr := node.(*ast.BinaryExpr)
		// we are only interested in multiplication
		if expr.Op != token.MUL {
			return
		}

		// get the types of the two operands
		x, xOK := pass.TypesInfo.Types[expr.X]
		y, yOK := pass.TypesInfo.Types[expr.Y]
		if !xOK || !yOK {
			return
		}

		if isDuration(x) && isDuration(y) {
			// check that both sides are acceptable expressions
			if isUnacceptableExpr(expr.X) && isUnacceptableExpr(expr.Y) {
				pass.Reportf(expr.Pos(), "Multiplication of durations: `%s`", formatNode(expr))
			}
		}
	}
}

func isDuration(x types.TypeAndValue) bool {
	return x.Type.String() == "time.Duration"
}

// isUnacceptableExpr returns true if the argument is not an acceptable time.Duration expression
func isUnacceptableExpr(expr ast.Expr) bool {
	switch e := expr.(type) {
	case *ast.BasicLit: // constants are acceptable
		return false
	case *ast.CallExpr: // explicit casting of constants such as `time.Duration(10)` is acceptable
		return !isConstExprCastToDuration(e)
	}
	return true
}

// isConstExprCastToDuration returns true if the argument is a constant expression cast to time.Duration
func isConstExprCastToDuration(e *ast.CallExpr) bool {
	// check that there's a single argument
	if len(e.Args) != 1 {
		return false
	}

	// check that the argument is a constant expression
	if !allConstArgs(e.Args[0]) {
		return false
	}

	// check for time.Duration cast
	selector, ok := e.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	pkg, ok := selector.X.(*ast.Ident)
	if !ok {
		return false
	}

	if pkg.Name != "time" {
		return false
	}

	return selector.Sel.Name == "Duration"
}

// allConstArgs checks that the argument is a constant expression
func allConstArgs(n ast.Node) bool {
	switch e := n.(type) {
	case *ast.BasicLit:
		return true
	case *ast.BinaryExpr:
		return allConstArgs(e.X) && allConstArgs(e.Y)
	default:
		return false
	}
}

func formatNode(node ast.Node) string {
	buf := new(bytes.Buffer)
	if err := format.Node(buf, token.NewFileSet(), node); err != nil {
		log.Printf("Error formatting expression: %v", err)
		return ""
	}

	return buf.String()
}

func printAST(node ast.Node) {
	fmt.Printf(">>> %s\n", formatNode(node))
	ast.Fprint(os.Stdout, nil, node, nil)
	fmt.Println("--------------")
}
