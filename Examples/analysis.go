package main

import (
	log "github.com/sirupsen/logrus"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
)

func DoAnalysis(path string) (string, error) {

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
	fileSet := token.NewFileSet()                                  // positions are relative to fileSet
	fileStruct, err := parser.ParseFile(fileSet, path, content, 0) //fileStruct是一个ast.File类型结构体
	if err != nil {
		typ, event = "ASTError", "Unable to parse file"
		log.WithFields(log.Fields{
			"type":   typ,
			"arg":    path,
			"errMsg": err,
		}).Fatal(event)
	}
	/*err = ast.Print(fileSet, fileStruct)
	if err != nil {
		typ, event = "ASTError", "Unable to print AST tree"
		log.WithFields(log.Fields{
			"type":   typ,
			"arg":    path,
			"errMsg": err,
		}).Fatal(event)
	}*/

	//Jsonify fileSet and log
	jsonifyFileStruct := fileStruct

	typ, event = "ASTGeneration", "fileset struct generating"
	log.WithFields(log.Fields{
		"type":    typ,
		"arg":     path,
		"fileset": jsonifyFileStruct,
	}).Info(event)
	return "", nil
}

//func inspectNodes()
