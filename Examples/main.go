package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
	//"path/filepath"
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
	log.SetLevel(log.TraceLevel) // 在测试环境中设置低等级级别，记录trace及以上级别
	//log.SetLevel(log.InfoLevel)    // 在生产环境中需要考虑性能，关注关键信息，level 设置高一点
	//log.SetReportCaller(true)            // 调用者文件名与位置
	//log.SetFormatter(new(log.JSONFormatter))    // 日志格式设置成json

	flag.Parse()
	var typ, event, arg string
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

	//Open File
	//Check if it is .go file
	var fileDiv = strings.Split(filename, ".")
	if fileDiv == nil || fileDiv[cap(fileDiv)-1] != "go" {
		typ, event = "Error", "Invalid suffix"
		log.WithFields(log.Fields{
			"type": typ,
			"arg":  filename,
		}).Fatal(event)
	}

	filePointer, errMsg := os.Open(filename)
	if errMsg != nil {
		typ, event = "Error", "Invalid path"
		log.WithFields(log.Fields{
			"type": typ,
			"arg":  filename,
		}).Fatal(event)
	}
	defer filePointer.Close()
	contentByte, errMsg := ioutil.ReadAll(filePointer) //NOTE: 这里要在正式版本中改成按行读取节省内存
	if errMsg != nil {
		typ, event = "Error", "Unable to read file"
		log.WithFields(log.Fields{
			"type": typ,
			"arg":  filename,
		}).Fatal(event)
	}
	typ, event = "Work", "Read file"
	log.WithFields(log.Fields{
		"type": typ,
		"arg":  filename,
	}).Debug(event)
	//fmt.Println(string(content))
	var contentString = string(contentByte)

	//Start analysis
	var result = DoAnalysis(filename, contentString)
	fmt.Println(result) //TODO:完成分析后修改结果处理
}
