package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
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
			"key":  nil,
		}).Info(event)
		return
	}
	if isVersion {
		showVersion()
		typ, event = "argparse", "Check version"
		log.WithFields(log.Fields{
			"type": typ,
			"key":  nil,
		}).Info(event)
		return
	}
	fmt.Println(filename)
	typ, event, arg = "argparse", "Input Filename", filename
	log.WithFields(log.Fields{
		"type": typ,
		"arg":  arg,
	}).Trace(event)
	//TODO:Log module
}
