package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

var fileName = "./task-02/example.go"
var funcs = []string{"example1", "example2", "example3"}

func main() {
	for _, fnc := range funcs {
		if num, err := analyze(fileName, fnc); err != nil {
			fmt.Println("analize error:", err)
		} else {
			fmt.Printf("File name:          %v\n", fileName)
			fmt.Printf("Function name:      %v\n", fnc)
			fmt.Printf("Go-calling (times): %v\n", num)
			fmt.Println()
		}
	}
}

func analyze(file, function string) (int, error) {
	var fileSet = token.NewFileSet()
	var fileAst *ast.File
	var err error
	var res int

	if fileAst, err = parser.ParseFile(fileSet, fileName, nil, 0); err != nil {
		return 0, fmt.Errorf("parse error: %w", err)
	}

	for _, dec := range fileAst.Decls {
		if fdec, ok := dec.(*ast.FuncDecl); !ok {
			continue
		} else if fdec.Name.Name != function {
			continue
		} else if cnt, err := analyzeFunc(fdec); err != nil {
			return 0, fmt.Errorf("func analyzing: %w", err)
		} else {
			res += cnt
		}
	}

	return res, nil
}

func analyzeFunc(dec *ast.FuncDecl) (int, error) {
	var res int = 0

	for _, stt := range dec.Body.List {
		if _, ok := stt.(*ast.GoStmt); ok {
			res++
		}
	}

	return res, nil
}
