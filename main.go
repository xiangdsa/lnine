package main

import (
	"fmt"
	"net/http"
	"time"
	"updateLnineDownloadMorepic/config"
	"updateLnineDownloadMorepic/consts"
	"updateLnineDownloadMorepic/controller"
	"updateLnineDownloadMorepic/dao"
	"updateLnineDownloadMorepic/utils"
)

//http://www.lnine.net/az/1132.html
//小九软件站 清洗注意事项 数据存放在不同的表中 电脑表 安卓表 每个表都是从不同的地方采集的
func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		fmt.Println("初始化配置信息失败")
	}
	//链接正式服
	dns := config.GetDSN(&cfg.MysqlConfig)

	//链接测试服
	//dns := config.GetTestDSN(&cfg.TestMysqlConfig)

	err = dao.InitMysqlConn(dns)
	if err != nil {
		fmt.Errorf("connect downer mysql fail error :%s", err)
	}

	defer dao.DB.Close()

	fmt.Println("开始：", time.Now().Format("2006-01-02 15:04:05"))

	//扫描
	//SaoMiao()

	//图片添加alt
	//2022-01-05 出现新的情况时只需要添加对应的正则表达式即可添加标签完成
	//controller.AddComputerImgTagNumber()
	controller.AddAzImgTagNumber()

	//修改电脑软件
	//dao.RepairComputer404AndForeginUrl(consts.ComputerPathPath, consts.Computer404Path)
}

func SaoMiao() {
	//扫描电脑
	//controller.CheckComputerUrl()
	//computerUrls:=utils.ReaderTxt(consts.ComputerPath,"\n")
	//for i,_:=range computerUrls{
	//	fmt.Println(computerUrls[i])
	//	dao.GetComputerUrl(computerUrls[i])
	//}
	//fmt.Println("电脑软件扫描完成", time.Now().Format("2006-01-02 15:04:05"))

	//扫描安卓
	for i := 0; i <= 67593; i = i + 1000 {
		urls := dao.QueryAzSoftwareID(i)
		SaoMiaoAZ(urls)
	}

	//controller.CheckIosUrl()
	//fmt.Println("ios软件扫描完成", time.Now().Format("2006-01-02 15:04:05"))

	//controller.CheckNewsUrl()
	//fmt.Println("资讯链接扫描完成", time.Now().Format("2006-01-02 15:04:05"))

	//controller.CheckVideoUrl()
	//fmt.Println("视频链接扫描完成", time.Now().Format("2006-01-02 15:04:05"))

	//controller.CheckFenleiUrl()
	//fmt.Println("分类链接扫描完成", time.Now().Format("2006-01-02 15:04:05"))

	//controller.CheckZhuantiUrl()
	//fmt.Println("专题链接扫描完成", time.Now().Format("2006-01-02 15:04:05"))

	//controller.CheckFirstPageUrl()
	//fmt.Println("首页链接扫描完成", time.Now().Format("2006-01-02 15:04:05"))
}

func SaoMiaoAZ(urls []string) {

	var zhengqueUrls []string
	var cuowuUrls []string
	for i, _ := range urls {
		resp, _ := http.Get(urls[i])
		if resp.StatusCode == 200 {
			zhengqueUrls = append(zhengqueUrls, urls[i])
			dao.GetAZUrl(urls[i])
		} else {
			cuowuUrls = append(cuowuUrls, urls[i])
		}
		fmt.Println(urls[i])
	}
	utils.WriteToTXT(zhengqueUrls, consts.AzPath)
	utils.WriteToTXT(cuowuUrls, consts.Az404Path)
}
