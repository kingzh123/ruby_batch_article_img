package log

import (
"bytes"
"fmt"
"log"
"os"
)

//log Print
func byteLog() {
	var byte bytes.Buffer
	var logger = log.New(&byte, "prefix:", log.Lshortfile)
	logger.Print("logxxxxxxxx")
	logger.Print("logxxxxxxxx2")
	fmt.Printf("%s\n", &byte)
}

//log Output
func outputLog() {
	var (
		byte   bytes.Buffer
		logger = log.New(&byte, "INFO:", log.Lshortfile)
		infof  = func(info string) {
			logger.Output(2, info)
		}
	)
	infof("hello world")
	infof("hello world2")
	fmt.Print(&byte)
}

func FileLog(msg string) {
	//打开文件
	file, _ := os.OpenFile("./log/info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	defer file.Close()
	logger := log.New(file, "INFO:", log.Lshortfile)
	//设置日志的flag log.ldate:在左侧显示完整时间 Y-m-d H:i:s  ltime: 显示 H:i:s Lshortfile:只显示文件名称及日志
	logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	infof := func(info string) {
		logger.Output(2, info)
	}
	infof(msg)
}