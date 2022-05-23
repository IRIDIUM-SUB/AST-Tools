package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	//"path/filepath"
	//"os"
)

var version = "v0.0.0"
var typ, event, arg string
var (
	isHelp    bool   //helper info
	isVersion bool   //version(seems useless)
	filename  string //entry path
)

func init() {
	//Bind vars and args
	flag.BoolVar(&isHelp, "h", false, "Show help info")
	flag.BoolVar(&isVersion, "v", false, "Show Version")
	flag.StringVar(&filename, "f", "./sample.src", "Entry path")
}
func showVersion() {
	//Print version info
	fmt.Println(version)
}
func main() {
	//Main entry
	log.SetLevel(log.TraceLevel) // 在测试环境中设置低等级级别，记录trace及以上级别
	//log.SetLevel(log.InfoLevel)    // 在生产环境中需要考虑性能，关注关键信息，level 设置高一点
	//log.SetReportCaller(true)            // 调用者文件名与位置
	//log.SetFormatter(new(log.JSONFormatter))    // 日志格式设置成json

	flag.Parse()

	if isHelp {
		flag.Usage()
		typ, event = "argparse", "Check help"

		log.WithFields(log.Fields{
			"type": typ,
			"args": "Help",
		}).Debug(event)
		return
	}
	if isVersion {
		showVersion()
		typ, event = "argparse", "Check version"
		log.WithFields(log.Fields{
			"type": typ,
			"args": "Version",
		}).Info(event)
		return
	}
	fmt.Println(filename)
	typ, event, arg = "argparse", "Input Filename", filename
	log.WithFields(log.Fields{
		"type": typ,
		"arg":  arg,
	}).Debug(event)

	//Check if it is .go file: Not used in test mode
	/*var fileDiv = strings.Split(filename, ".")
	if fileDiv == nil || fileDiv[cap(fileDiv)-1] != "go" {
		typ, event = "Error", "Invalid suffix"
		log.WithFields(log.Fields{
			"type": typ,
			"arg":  filename,
		}).Fatal(event)
	}*/

	//Start analysis
	result, err := DoAnalysis(filename)
	if err != nil {
		typ, event, arg = "ASTError", "Unknown error", filename
		log.WithFields(log.Fields{
			"type": typ,
			"arg":  arg,
		}).Fatal(event)
	}
	fmt.Println(result)

}
