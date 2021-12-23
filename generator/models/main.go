package main

import (
	"bufio"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

var (
	flagSrc  string
	flagDest string
)

const prefix = `// Package for models.//
// Code generated by github.com/FlowingSPDG/resolume-go/generator/models DO NOT EDIT.
package models
`

func main() {
	flag.StringVar(&flagSrc, "src", "", "Source file")
	flag.StringVar(&flagDest, "dest", "", "Destination file")
	flag.Parse()

	// Read Source file
	src, err := os.ReadFile(flagSrc)
	if err != nil {
		panic(err)
	}

	// Remove Dest file
	if err := os.Remove(flagDest); err != nil {
		panic(err)
	}
	// Create Dest file
	dest, err := os.OpenFile(flagDest, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer dest.Close()

	// Initialize writer
	w := bufio.NewWriter(dest)
	if _, err := w.WriteString(prefix); err != nil {
		panic(err)
	}

	// Create AST
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	for _, node := range f.Decls {
		switch node.(type) {
		case *ast.GenDecl:
			genDecl := node.(*ast.GenDecl)
			for _, spec := range genDecl.Specs {
				switch spec.(type) {
				case *ast.TypeSpec:
					typeSpec := spec.(*ast.TypeSpec)
					if _, err := w.WriteString(fmt.Sprintf("type %s %s.%s\n", typeSpec.Name.Name, f.Name.Name, typeSpec.Name.Name)); err != nil {
						fmt.Println("Failed to write code :", err)
						return
					}
				}
			}
		}
	}
}
