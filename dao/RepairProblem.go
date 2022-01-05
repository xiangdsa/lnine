package dao

import "updateLnineDownloadMorepic/utils"

//修复一些404链接和外链 全部替换成 http://www.lnine.net/zhuanti.html

const DefaultUrl = "http://www.lnine.net/zhuanti.html"

//修复 电脑软件softsay字段中的外链或者错误链接
func RepairComputer404AndForeginUrl(foreginFilepath ,url404path string) () {
	ForeginUrls := utils.ReaderTxt(foreginFilepath, "\n")

	for i,_:=range ForeginUrls{
		Repair404OrForeginByUrl(ForeginUrls[i])
	}

	url404 := utils.ReaderTxt(url404path, "\n")

	for i,_:=range ForeginUrls{
		Repair404OrForeginByUrl(url404[i])
	}

}
