package main

import (
	"flag"
	"fmt"
	//"os"
)

var version = "v0.0.0"

var (
	isHelp    bool   //helper info
	isVersion bool   //version(seems useless)
	filename  string //entry path
)

func init() {
	//Bind vars and args
	flag.BoolVar(&isHelp, "h", false, "Show help info")
	flag.BoolVar(&isVersion, "v", false, "Show Version")
	flag.StringVar(&filename, "f", "./main.go", "Entry path")
}
func showVersion() {
	//Print version info
	fmt.Println(version)
}
func main() {
	//Main entry
	fmt.Println("Test text")
	flag.Parse()
	if isHelp {
		flag.Usage()
	}
	if isVersion {
		showVersion()
	}
	fmt.Println(filename)
	//TODO:Log module
}
