package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
)

func DoAnalysis(path string) (string, error) {
	/*
		Open and read source files:Generate fileSet and fileStruct
	*/
	//Read file
	file, err := os.Open(path)
	if err != nil {
		typ, event = "IOError", "Unable to open file"
		log.WithFields(log.Fields{
			"type":   typ,
			"arg":    path,
			"errMsg": err,
		}).Fatal(event)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			typ, event = "IOError", "Unable to close file"
			log.WithFields(log.Fields{
				"type":   typ,
				"arg":    path,
				"errMsg": err,
			}).Fatal(event)
		}
	}(file)

	content, err := ioutil.ReadAll(file)
	if err != nil {
		typ, event = "IOError", "Unable to read file"
		log.WithFields(log.Fields{
			"type":   typ,
			"arg":    path,
			"errMsg": err,
		}).Fatal(event)
	}

	//Generating AST tree
	fileSet := token.NewFileSet() // positions are relative to fileSet
	fileStruct, err := parser.ParseFile(fileSet, path, content, 0)
	//fileStruct是一个ast.File类型结构体:
	/*
		type File struct {
		        Doc        *CommentGroup   // associated documentation; or nil
		        Package    token.Pos       // position of "package" keyword
		        Name       *Ident          // package name
		        Decls      []Decl          // top-level declarations; or nil
		        Scope      *Scope          // package scope (this file only)
		        Imports    []*ImportSpec   // imports in this file
		        Unresolved []*Ident        // unresolved identifiers in this file
		        Comments   []*CommentGroup // list of all comments in the source file
		}
	*/
	if err != nil {
		typ, event = "ASTError", "Unable to parse file"
		log.WithFields(log.Fields{
			"type":   typ,
			"arg":    path,
			"errMsg": err,
		}).Fatal(event)
	}

	ast.Inspect(fileStruct, func(node ast.Node) bool {
		var statementType string
		var s string
		switch x := node.(type) {
		case *ast.BasicLit:
			statementType = "BasicLit" //
			s = x.Value
		case *ast.Ident:
			statementType = "Ident" //关键字
			s = x.Name
		case *ast.FuncDecl:
			statementType = "FuncDecl"

		}
		if s != "" {
			typ = "ASTScan"
			log.WithFields(log.Fields{
				"type":     typ,
				"position": fileSet.Position(node.Pos()),
				"nodeType": statementType,
				"arg":      path,
			}).Info(event)
			fmt.Printf("%s:\t%s\n", fileSet.Position(node.Pos()), s)
		}
		return true
	})

	return "", nil

}

//func inspectNodes()
