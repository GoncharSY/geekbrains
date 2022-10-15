package main

// В этом файле хранится код, который мне просто стало жалко удалять.
// В процессе вникания в тему, я забыл о чем была задача и написал код,
// который делает не совсем то, что нужно по ее условиям. Когда перечитал
// условия задачи, плюную, выругался про себя и все переписал.

// func analyzeGoCalling(file, function string) (int, error) {
// 	var fileSet = token.NewFileSet()
// 	var fileAst *ast.File
// 	var err error
// 	var res int

// 	if fileAst, err = parser.ParseFile(fileSet, fileName, nil, 0); err != nil {
// 		return 0, fmt.Errorf("parse error: %w", err)
// 	}

// 	for _, dec := range fileAst.Decls {
// 		if fdec, ok := dec.(*ast.FuncDecl); !ok {
// 			continue
// 		} else if cnt, err := analyzeFunc(fdec, function); err != nil {
// 			return 0, fmt.Errorf("func analyzing: %w", err)
// 		} else {
// 			res += cnt
// 		}
// 	}

// 	return res, nil
// }

// func analyzeFunc(dec *ast.FuncDecl, fnc string) (int, error) {
// 	var res int = 0

// 	for _, stt := range dec.Body.List {
// 		if gst, ok := stt.(*ast.GoStmt); !ok {
// 			continue
// 		} else if idt, ok := gst.Call.Fun.(*ast.Ident); !ok {
// 			continue
// 		} else if idt.Name == fnc {
// 			res++
// 		}
// 	}

// 	return res, nil
// }
