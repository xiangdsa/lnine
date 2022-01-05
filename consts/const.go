package consts

const (
	//<img border="0" src="http://img.lnine.net/uploadfiles/2017-07-11/20170711_092721_300.jpg"/>
	LnineImgUrl1 = "src=\"(http://img.lnine.net/[a-z]+/[0-9]{4}-[0-9]{2}-[0-9]{2}/[0-9_]+.[a-zA-Z]+)"

	//<img src="http://img.lnine.net/up/2006/2020612141412875970.jpg"
	LnineImgUrl2 = "src=\"(http://img.lnine.net/[a-z]+/[0-9]{4}/[0-9]+.[a-zA-Z]+)"

	ComputerUrlModel    = "http://www.lnine.net/xiazai/%d.html"
	AzUrlModel          = "http://www.lnine.net/az/%d.html"
	IOSUrlModel         = "http://www.lnine.net/ios/%d.html"
	NewsZiXunModel      = "http://www.lnine.net%s"
	VideoUrlModel       = "http://www.lnine.net/shipin%s"
	FenleiSoftwareModel = "http://www.lnine.net/%s"
	ZhuanTiModel        = "http://www.lnine.net/%s"
	FirstPageModel      = "http://www.lnine.net/%s"
	ArticlePageModel    = "http://www.lnine.net/article/index_%d.html"

	//SiteMapPath = "C:\\Users\\BZ\\Desktop\\seo游当网\\sitemap.txt"
	//WrongPath   = "C:\\Users\\BZ\\Desktop\\seo游当网\\错误链接.txt"
	//
	IosPath    = "C:\\Users\\BZ\\Desktop\\seo游当网\\Ios软件.txt"
	Ios404Path = "C:\\Users\\BZ\\Desktop\\seo游当网\\Ios软件404.txt"
	//
	ComputerPath     = "C:\\Users\\BZ\\Desktop\\seo游当网\\电脑软件.txt"
	ComputerPathPath = "C:\\Users\\BZ\\Desktop\\seo游当网\\电脑软件-softsay.txt"
	Computer404Path  = "C:\\Users\\BZ\\Desktop\\seo游当网\\电脑软件404.txt"
	ComputerImgPath  = "C:\\Users\\BZ\\Desktop\\seo游当网\\电脑软件图片链接.txt"
	ComputerImg      = "C:\\Users\\BZ\\Desktop\\seo游当网\\电脑图片.txt"
	//
	AzPath     = "C:\\Users\\BZ\\Desktop\\seo游当网\\安卓软件.txt"
	AzPathPath = "C:\\Users\\BZ\\Desktop\\seo游当网\\安卓软件-softsay.txt"
	Az404Path  = "C:\\Users\\BZ\\Desktop\\seo游当网\\安卓软件404.txt"
	AzImgPath  = "C:\\Users\\BZ\\Desktop\\seo游当网\\安卓软件图片链接.txt"
	AzImg      = "C:\\Users\\BZ\\Desktop\\seo游当网\\安卓图片.txt"

	FenLeiPath    = "C:\\Users\\BZ\\Desktop\\seo游当网\\分类.txt"
	FenLeiPath404 = "C:\\Users\\BZ\\Desktop\\seo游当网\\分类404.txt"

	VideoPath    = "C:\\Users\\BZ\\Desktop\\seo游当网\\视频.txt"
	VideoPath404 = "C:\\Users\\BZ\\Desktop\\seo游当网\\视频404.txt"

	ZtPath    = "C:\\Users\\BZ\\Desktop\\seo游当网\\专题.txt"
	ZtPath404 = "C:\\Users\\BZ\\Desktop\\seo游当网\\专题404.txt"

	ZiXunPath    = "C:\\Users\\BZ\\Desktop\\seo游当网\\资讯.txt"
	ZiXunPath404 = "C:\\Users\\BZ\\Desktop\\seo游当网\\资讯404.txt"

	SiteMapPath = "/www/wwwroot/clear/lnine/sitemap.txt"
	WrongPath   = "/www/wwwroot/clear/lnine/错误链接.txt"
	//ComputerPath="/www/wwwroot/clear/lnine/电脑软件.txt"
	//ComputerPathPath="/www/wwwroot/clear/lnine/电脑软件-softsay.txt"
	//Computer404Path="/www/wwwroot/clear/lnine/电脑软件404.txt"
	//ComputerImgPath="/www/wwwroot/clear/lnine/电脑软件图片链接.txt"
	//ComputerImg="/www/wwwroot/clear/lnine/电脑图片.txt"

	//AzPath="/www/wwwroot/clear/lnine/安卓软件.txt"
	//AzPathPath="/www/wwwroot/clear/lnine/安卓软件-softsay.txt"
	//Az404Path="/www/wwwroot/clear/lnine/安卓软件404.txt"
	//AzImgPath="/www/wwwroot/clear/lnine/安卓软件图片链接.txt"
	//AzImg="/www/wwwroot/clear/lnine/安卓图片.txt"

	//获得img标签
	// <img src="http://img.lnine.net/up/2010/202010131402183470.gif" alt="蒲公英MP3格式转换软件"/>
	// <img border="0" src="http://img.lnine.net/uploadfiles/2018-06-17/20180617_104359_580.jpg"/>
	ComputerImgTagModel1 = "<img src=\"http://img.lnine.net/[0-9a-zA-Z-_/]+.[a-zA_Z]{3,5}\" alt=\"([\u4e00-\u9fa5 \\sA-Za-z0-9]+)\"/>"
	ComputerImgTagModel2 = "<img border=\"0\" src=\"http://img.lnine.net/[0-9a-zA-Z-_/]+.[a-zA_Z]{3,5}\"/>"


	//<img alt="甜盐相机" src="http://img.lnine.net/2/KDYwMHgp/d0e5379b-efe3-4144-96c4-0dea63da4b2a.png" style="width: 185px; height: 400px;"/>
	//<img width="300" alt="乐呵呵app" src="http://img.lnine.net/up/1611/20161118083802421.jpg"/>
	AZImgTagModel1 = "<img [\u4e00-\u9fa5 =\"\\s A-Za-z0-9]{0,40}src=\"http://img.lnine.net/[0-9a-zA-Z-_/]+.[a-zA_Z]{3,5}\" [0-9a-zA-Z=:\";\\s]{0,60}/>"
	//<img src="http://img.lnine.net/12/9/b9ae4cf3-5c0e-4adf-91ce-fb5a43b8b91f.jpg" style="height: 271px; width: 400px;"/>
	AZImgTagModel2 = "<img [\u4e00-\u9fa5 =\"\\sA-Za-z0-9]{0,40} src=\"http://img.lnine.net/[0-9a-zA-Z-_/]+.[a-zA_Z]{3,5}\"/>"
	//	<img src="http://img.lnine.net/12/13/73a68b6d-cb9f-4453-bb2b-981d9c3d2485.jpg"/>
	AZImgTagModel3 = "<img [\u4e00-\u9fa5 =\"\\sA-Za-z0-9]{0,40}src=\"http://img.lnine.net/[0-9a-zA-Z-_/]+.[a-zA_Z]{3,5}\"/>"
)
