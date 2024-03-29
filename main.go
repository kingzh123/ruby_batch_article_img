package main

import (
	"github.com/gohouse/converter"
	"ruby_article/log"
	"ruby_article/module"
	"ruby_article2/db"
	"sync"
)

var (
	a = make([]module.Articles,0)
	test_domain = "jfinfo-test.oss-cn-beijing.aliyuncs.com"
	pro_domain = "jfinfo.oss-cn-beijing.aliyuncs.com"
	default_img = ``
	wg sync.WaitGroup
	size = 1000
)

func mBatch()  {
	log.FileLog("run star task")
	Db.Batch(int64(size), 2547105)
	log.FileLog("success updated database!")
	wg.Done()
}

func main() {
	wg.Add(1)
	go mBatch()
	wg.Wait()
}

func runModule(table string)  {
	cvt := converter.NewTable2Struct()
	cvt.Config(&converter.T2tConfig{
		RmTagIfUcFirsted: false,
		TagToLower: false,
		UcFirstOnly: false,
	})
	err := cvt.Table("articles").
		EnableJsonTag(true).
		PackageName("module").
		TagKey("orm").
		RealNameMethod("articles").
		SavePath("./module/articles.go").
		Dsn("").
		// 执行
		Run()
	if err != nil {
		panic(err)
	}
}