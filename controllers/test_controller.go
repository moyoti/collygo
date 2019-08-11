package controllers

import (
	"encoding/csv"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gocolly/colly"
	"os"
	"strconv"
	"strings"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) URLMapping() {
	c.Mapping("Test", c.Test)
}

// @router /tt [get]
func (c *TestController) Test() {
	info := c.Input().Get("testInfo")
	c.Data["json"] = info
	c.ServeJSON()
}

// @router /download [get]
func (c *TestController) DownloadFile() {

	c.Ctx.Output.Download("./files/data.csv")
}

// @router /col [get]
func (c *TestController) CollectData() {
	url := c.Input().Get("url")
	if checkFileIsExist("./files/data.csv") {
		err := os.Remove("./files/data.csv")
		if err != nil {
			c.Data["json"] = "原文件未删除成功或不存在"
			c.ServeJSON()
		}
	}
	success := FindData(url)
	if success {
		c.Data["json"] = "success"
		c.ServeJSON()
	} else {
		c.Data["json"] = "failed"
		c.ServeJSON()
	}
}
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func FindData(url string) bool {
	c := colly.NewCollector()
	colly.MaxDepth(0)
	c.AllowedDomains = []string{"www.taptap.com"}
	c.OnHTML("li[class|='taptap-review-item collapse in']", func(e *colly.HTMLElement) {
		//cmtiList:=list.New()
		username := e.Attr("data-user")
		comment := e.DOM.Find("div[class|='item-text-body']")
		score := e.DOM.Find("div[class|='item-text-score']").Find("i[class=colored]")
		gameLength := e.DOM.Find("div[class|='item-text-score']").Find("span[class=text-score-time]")
		//link := e.Attr("href")
		sr, exists := score.Attr("style")
		var rescore float64
		if exists && strings.Contains(sr, "width") {
			rslice := strings.Split(sr, "width:")
			//pxpos:=strings.Index(strings.TrimSpace(rslice[1]),"px")
			//fmt.Println(pxpos)
			re, err_float := strconv.ParseFloat(strings.Split(strings.TrimSpace(rslice[1]), "px")[0], 10)
			if err_float != nil {
				fmt.Println(err_float)
			} else {
				fmt.Print("score:")
				rescore = re / 14
				fmt.Println(rescore)
			}
		}
		fmt.Print("username:")
		fmt.Println(username)
		fmt.Print("gameLength:")
		regame := strings.Replace(gameLength.Text(), "游戏时长 ", "", 1)
		fmt.Println(regame)
		sc := comment.Eq(0).Text()
		fmt.Print("comment:")
		recomment := strings.TrimSpace(sc)
		fmt.Println(recomment)
		fmt.Println("-----------------------------------------------------------------------")
		var data [][]string
		//nowtime:=time.Now()
		fileName := "./files/data.csv" /* + string(nowtime.Year()) + "-" + string(nowtime.Month()) + "-" + string(nowtime.Day())+".csv"*/
		var fp *(os.File)
		var err error
		if checkFileIsExist(fileName) {
			fp, err = os.OpenFile(fileName, os.O_APPEND, 0666)
			data = [][]string{{username, strconv.FormatFloat(rescore, 'f', -1, 64), regame, recomment}}
		} else {
			fp, err = os.Create(fileName)
			data = [][]string{{"用户名", "评分", "游戏时长", "评论"}, {username, strconv.FormatFloat(rescore, 'f', -1, 64), regame, recomment}}
		}
		if err != nil {
			fmt.Println("创建或读取文件失败")
		}
		defer fp.Close()
		fp.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
		w := csv.NewWriter(fp)         //创建一个新的写入文件流
		w.WriteAll(data)
		w.Flush()
	})
	c.OnHTML("div[class|='main-body-footer']", func(e *colly.HTMLElement) {
		pageNev := e.DOM.Find("a[rel|='nofollow']")
		//fmt.Println(pageNev.Text())
		nextPage := pageNev.Eq(pageNev.Size() - 1)
		if nextPage.Text() == ">" {
			nextURL, exists := nextPage.Attr("href")
			if exists {
				fmt.Println(nextURL)
				c.Visit(nextURL)
			} else {
				fmt.Println("下一页评论获取错误")
			}
		}
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit(url)
	return true
}
