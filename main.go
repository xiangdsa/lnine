package main

import (
	"createHtml/config"
	"createHtml/consts"
	"createHtml/dao"
	"createHtml/utils"
	"fmt"
)

type Product struct {
	Id   int64  `json:"id"` //字段一定要大写不然各种问题
	Name string `json:"name"`
}

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		fmt.Errorf("load config failed, err:", err)
	}
	//连接seo数据库
	//dns := config.GetDSN(&cfg.MysqlConfig)
	//连接test库
	dns := config.GetTestDSN(&cfg.TestMysqlConfig)
	err = dao.InitMysqlCon(dns)
	if err != nil {
		fmt.Errorf("connect downer mysql fail error :", err)
	}
	defer dao.DB.Close()

	todayTime := utils.GetTodayTime()
	tomorrowTime := utils.GetTomorrowTime()

	urls := dao.GetUrl(todayTime, tomorrowTime)

	contents := make([]string, 0, 10)
	for i, _ := range urls {
		content, _ := dao.GetHTMLCode(urls[i])
		contents = append(contents, content)
	}

	softwareID := dao.GetSoftWareID(todayTime, tomorrowTime)

	for i, _ := range softwareID {
		utils.WriteContentToHtml(fmt.Sprintf(consts.FilePath, softwareID[i]), contents[i])
	}

	index,_:=dao.GetHTMLCode("https://www.s-do.com/")
	//请求大雀
	//index,_:=dao.GetHTMLCode("https://www.daque.cn/")

	//utils.DeleteFile(" /www/wwwroot/sdo-spider/runtime/index.html")

	//
	//utils.WriteContentToHtml("/www/wwwroot/sdo-spider/runtime/index.html",index)

	//utils.WriteContentToHtml("/www/wwwroot/spider/runtime/index.html",index)
	//utils.DeleteFile("C:\\Users\\BZ\\Desktop\\大雀图片\\index.html")
	utils.WriteContentToHtml("C:\\Users\\BZ\\Desktop\\大雀图片\\index.html",index)
}
