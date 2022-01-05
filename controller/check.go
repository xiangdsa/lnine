package controller

import (
	"updateLnineDownloadMorepic/consts"
	"updateLnineDownloadMorepic/dao"
	"updateLnineDownloadMorepic/utils"
)

//电脑软件链接
func CheckComputerUrl() {
	computerUrls := dao.QueryComputerSoftwareID()

	url200, url404 := utils.CheckUrl(computerUrls)

	utils.WriteToTXT(url200, consts.ComputerPath)

	utils.WriteToTXT(url404, consts.WrongPath)
}

//安卓软件链接
func CheckAzUrl() {
	//azUrls := dao.QueryAzSoftwareID()
	//
	//url200, url404 := utils.CheckUrl(azUrls)
	//
	//utils.WriteToTXT(url200, consts.AzPath)
	//
	//utils.WriteToTXT(url404, consts.WrongPath)

}

//ios软件链接
func CheckIosUrl() {
	iosUrls := dao.QueryIOSSoftwareID()

	url200, url404 := utils.CheckUrl(iosUrls)

	utils.WriteToTXT(url200, consts.IosPath)

	utils.WriteToTXT(url404, consts.Ios404Path)
}

//资讯链接
func CheckNewsUrl() {
	newsUrls := dao.QueryNewsSoftwareID()

	url200, url404 := utils.CheckUrl(newsUrls)

	utils.WriteToTXT(url200, consts.ZiXunPath)

	utils.WriteToTXT(url404, consts.ZiXunPath404)
}

//视频链接
func CheckVideoUrl() {
	videoUrls := dao.QueryVideoSoftwareID()

	url200, url404 := utils.CheckUrl(videoUrls)

	utils.WriteToTXT(url200, consts.VideoPath)

	utils.WriteToTXT(url404, consts.VideoPath404)
}

//分类链接
func CheckFenleiUrl() {
	fenleiUrls := dao.QueryFenleiClasspath()

	url200, url404 := utils.CheckUrl(fenleiUrls)

	utils.WriteToTXT(url200, consts.FenLeiPath)

	utils.WriteToTXT(url404, consts.FenLeiPath404)
}

//专题链接
func CheckZhuantiUrl() {
	zhuantiUrls := dao.QueryZhuanTiZtpath()

	url200, url404 := utils.CheckUrl(zhuantiUrls)

	utils.WriteToTXT(url200, consts.ZtPath)

	utils.WriteToTXT(url404, consts.ZtPath404)
}

//首页下面的一些零散链接
func CheckFirstPageUrl() {
	firstPageUrls := dao.QueryFirstPageUrls()

	url200, url404 := utils.CheckUrl(firstPageUrls)

	utils.WriteToTXT(url200, consts.ZtPath)

	utils.WriteToTXT(url404, consts.ZtPath404)
}