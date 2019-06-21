package Db

import (
	"database/sql"
	"fmt"
	"regexp"
	"ruby_article/log"
	"ruby_article/module"
	"strconv"
	"strings"
)

type Db struct{
}

var (
	dbw *sql.DB
	err error
	a = make([]module.Articles,0)
	test_domain = "jfinfo-test.oss-cn-beijing.aliyuncs.com"
	pro_domain = "jfinfo.oss-cn-beijing.aliyuncs.com"
	default_img = ``
	//default_img = `<img src="https://jfinfo.oss-cn-beijing.aliyuncs.com/uploads/default/WX20190618-150448.png" />`
)

func init() {
	//dbw, err = sql.Open("mysql", "jfinfo_qa_ruby:ffc+ZUvzh35p@tcp(rm-2ze4ic7pv96x86o526o.mysql.rds.aliyuncs.com:3306)/jfinfo_qa?charset=utf8")
	dbw, err = sql.Open("mysql", "jfinfo_ruby:yH40l=Lydtet@tcp(rm-2zehm9x9dn6dbwbrm763.mysql.rds.aliyuncs.com:3306)/jfinfo?charset=utf8mb4")
	if err != nil {
		log.FileLog(err.Error())
		panic(err)
	}
}

func UpdateContentById(id int, content string) bool {
	content = strings.Replace(content, "'", `"`, -1)
	sql := "UPDATE articles SET content='" + content + "' WHERE id=" + strconv.Itoa(id)
	fmt.Printf("%s\n", sql)
	_, err := dbw.Exec(sql)
	if err != nil {
		log.FileLog(err.Error())
		return false
	}
	return true
}

func Batch(size int64, lastId int) {
	a, last := GetArticlesById(size, lastId)
	if len(a) <= 0 {
		return
	}
	for _, v := range a  {
		if v.Content.Valid {
			log.FileLog(strconv.Itoa(v.Id))
			b := []byte(v.Content.String)
			matching := `<img.*?src=\"(http|https)?(://)?.*?(?:>|\/>)`
			reg := regexp.MustCompile(matching)
			str := reg.FindAll(b,-1)
			if str != nil {
				jsq := false
				//替换新的SRC内容
				for _, img := range str {
					if strings.Index(string(img), test_domain) > 0 || strings.Index(string(img), pro_domain) > 0 {
						continue
					}
					new_content := ""
					//处理原有img保证多次执行程序无异常
					imgs := strings.Replace(string(img), "<img", "<ximg", 1)
					new_content += "<!--" + imgs + "-->" + default_img
					v.Content.String = strings.ReplaceAll(v.Content.String, string(img), new_content)
					jsq = true
				}
				if jsq {
					//更新操作
					if UpdateContentById(v.Id, v.Content.String) {
						id := strconv.Itoa(v.Id)
						log.FileLog("-------------------" + id + " OLD--------------------")
						log.FileLog(string(b))
						log.FileLog("-------------------" + id + " NEW--------------------")
						log.FileLog(v.Content.String)
					}
				}
			}
		}
	}
	Batch(size, last)
}

func GetArticlesById(size int64, lastId int) ([]module.Articles, int){
	var articles = make([]module.Articles, 0)
	last := 0
	sql := "select id,title,content from articles where published_at < '2018-7-31 23:59:59' AND id < " + strconv.Itoa(lastId) + " ORDER BY id desc LIMIT " + strconv.FormatInt(size, 10)
	rows, err := dbw.Query(sql)
	if err != nil {
		log.FileLog(err.Error())
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		ar := new(module.Articles)
		err = rows.Scan(&ar.Id, &ar.Title, &ar.Content)
		if err != nil {
			log.FileLog(err.Error())
			panic(err)
		}
		last = ar.Id
		articles = append(articles, *ar)
	}
	return articles, last
}