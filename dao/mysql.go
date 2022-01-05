package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"regexp"
	"strings"
	"time"
	"updateLnineDownloadMorepic/consts"
	"updateLnineDownloadMorepic/models"
)

var DB *gorm.DB

func InitMysqlConn(dns string) (err error) {

	DB, err = gorm.Open("mysql", dns)
	if err != nil {
		fmt.Println("数据库链接出问题", err)
	}

	err = DB.DB().Ping()

	DB.DB().SetMaxIdleConns(64)

	DB.DB().SetMaxOpenConns(64)

	DB.DB().SetConnMaxLifetime(600 * time.Second)

	return
}

//电脑软件模块
//http://www.lnine.net/xiazai/16030.html
func QueryComputerSoftwareID() (urls []string) {
	var bzSeoEcmsDownload []*models.BzSeoEcmsDownload
	DB.Table("Bz_Seo_ecms_download").Where("id>?", 0).Scan(&bzSeoEcmsDownload)
	for i, _ := range bzSeoEcmsDownload {
		urls = append(urls, fmt.Sprintf(consts.ComputerUrlModel, bzSeoEcmsDownload[i].ID))
	}
	return urls
}

//分批获得电脑的数据
func QueryComputerMessage(id int64) (bzSeoEcmsDownload []*models.BzSeoEcmsDownload) {

	DB.Table("Bz_Seo_ecms_download").Where("id >= ? and id <= ?", id, id+300).Scan(&bzSeoEcmsDownload)

	return
}

//获得电脑软件的最大id
func QueryComputerMaxID() (id int64) {
	var bzSeoEcmsDownload []*models.BzSeoEcmsDownload
	DB.Table("Bz_Seo_ecms_download").Where("id>?", 0).Last(&bzSeoEcmsDownload)
	if len(bzSeoEcmsDownload) > 0 {
		return bzSeoEcmsDownload[0].ID
	}
	return 0
}

//更新电脑软件的softsay
func UpdateComputerSoftsay(newSoftsay string, id int64) {
	DB.Table("Bz_Seo_ecms_download").Where("id=?", id).Update("softsay", newSoftsay)
}

//安卓软件模块
// http://www.lnine.net/az/2.html
func QueryAzSoftwareID(id int) (urls []string) {
	var bzSeoEcmsAzData1 []*models.BzSeoEcmsAzData1
	DB.Table("Bz_Seo_ecms_az_data_1").Where("id<=? and id>= ?", id+1000, id).Scan(&bzSeoEcmsAzData1)
	for i, _ := range bzSeoEcmsAzData1 {
		urls = append(urls, fmt.Sprintf("http://www.lnine.net/az/%d.html", bzSeoEcmsAzData1[i].ID))
	}

	return
}

//分批获得安卓的数据
func QueryAZMessage(id int64) (bzSeoEcmsAzData1 []*models.BzSeoEcmsAzData1) {

	DB.Table("Bz_Seo_ecms_az_data_1").Where("id >= ? and id <= ?", id, id+300).Scan(&bzSeoEcmsAzData1)

	return
}

//获得安卓最大的id
func QueryAZMaxID() (id int64) {
	var bzSeoEcmsAzData1 []*models.BzSeoEcmsAzData1
	DB.Table("Bz_Seo_ecms_az_data_1").Where("id>?", 0).Last(&bzSeoEcmsAzData1)
	if len(bzSeoEcmsAzData1) > 0 {
		return bzSeoEcmsAzData1[0].ID
	}
	return 0
}

func QueryAzTitleByData1ID(id int64) (title string) {
	var bzSeoEcmsAz []*models.BzSeoEcmsAz
	DB.Table("Bz_Seo_ecms_az").Where("id= ?", id).Scan(&bzSeoEcmsAz)
	return bzSeoEcmsAz[0].Title
}

func UpdateAzNewsText(newstext string, id int64) {
	DB.Table("Bz_Seo_ecms_az_data_1").Where("id=?", id).Update("newstext", newstext)
}

//IOS软件模块
// http://www.lnine.net/az/2.html
func QueryIOSSoftwareID() (urls []string) {
	var bzSeoEcmsIos []*models.BzSeoEcmsIos
	DB.Table("Bz_Seo_ecms_ios").Where("id>0").Scan(&bzSeoEcmsIos)
	for i, _ := range bzSeoEcmsIos {
		urls = append(urls, fmt.Sprintf(consts.IOSUrlModel, bzSeoEcmsIos[i].ID))
	}
	return
}
func QueryIOSMaxID() (id int64) {
	var bzSeoEcmsAzData1 []*models.BzSeoEcmsIos
	DB.Table("Bz_Seo_ecms_az_data_1").Where("id>?", 0).Last(&bzSeoEcmsAzData1)
	if len(bzSeoEcmsAzData1) > 0 {
		return bzSeoEcmsAzData1[0].ID
	}
	return 0
}
//资讯  文章软件模块
// http://www.lnine.net/article/1229-4145.html    http://www.lnine.net/zixun/1261.html
func QueryNewsSoftwareID() (urls []string) {
	var bzSeoEcmsNews []*models.BzSeoEcmsNews
	DB.Table("Bz_Seo_ecms_news").Where("id>0").Scan(&bzSeoEcmsNews)
	for i, _ := range bzSeoEcmsNews {
		urls = append(urls, fmt.Sprintf(consts.NewsZiXunModel, bzSeoEcmsNews[i].Titleurl))
	}
	return
}

//视频模块
// http://www.lnine.net/shipin/2.html
func QueryVideoSoftwareID() (urls []string) {
	var bzSeoEcmsVideo []*models.BzSeoEcmsVideo
	DB.Table("Bz_Seo_ecms_video").Where("id>0").Scan(&bzSeoEcmsVideo)
	for i, _ := range bzSeoEcmsVideo {
		urls = append(urls, fmt.Sprintf(consts.VideoUrlModel, bzSeoEcmsVideo[i].Titleurl))
	}
	return
}

//软件分类模块
//http://www.lnine.net/%s
func QueryFenleiClasspath() (urls []string) {
	var bzSeoEnewsclass []*models.BzSeoEnewsclass
	DB.Table("Bz_Seo_enewsclass").Where("classid>0").Scan(&bzSeoEnewsclass)
	for i, _ := range bzSeoEnewsclass {
		urls = append(urls, fmt.Sprintf(consts.FenleiSoftwareModel, bzSeoEnewsclass[i].Classpath))
	}
	return
}

//软件专题模块
//http://www.lnine.net/%s
func QueryZhuanTiZtpath() (urls []string) {
	var bzSeoEnewszt []*models.BzSeoEnewszt
	DB.Table("Bz_Seo_enewszt ").Where("ztid>0").Scan(&bzSeoEnewszt)
	for i, _ := range bzSeoEnewszt {
		urls = append(urls, fmt.Sprintf(consts.ZhuanTiModel, bzSeoEnewszt[i].Ztpath))
	}
	return
}

//首页最下面的一些链接
//../../xiazai/category.html
func QueryFirstPageUrls() (urls []string) {
	var bzSeoEnewspage []*models.BzSeoEnewspage
	DB.Table("Bz_Seo_enewspage ").Where("id>0").Scan(&bzSeoEnewspage)
	for i, _ := range bzSeoEnewspage {
		url := strings.Replace(bzSeoEnewspage[i].Path, "../../", "", 1)
		urls = append(urls, fmt.Sprintf(consts.FirstPageModel, url))
	}
	return
}

func UpdateMorepic(newMorepic string, id int64) {
	DB.Table("Bz_Seo_ecms_download").Where("id=?", id).Update("morepic", newMorepic)
}

func Repair404OrForeginByUrl(url string) {
	var bzSeoEcmsDownload []*models.BzSeoEcmsDownload
	DB.Table("Bz_Seo_ecms_download").Where("softsay like %?% ", url).Scan(&bzSeoEcmsDownload)
	if len(bzSeoEcmsDownload) > 0 {
		newSoftsay := strings.Replace(bzSeoEcmsDownload[0].Softsay, url, DefaultUrl, -1)
		DB.Table("Bz_Seo_ecms_download").Where("id=?", bzSeoEcmsDownload[0].ID).Update("softsay", newSoftsay)
	}
}

func GetLnineImgUrl(content string) (imgUrls []string) {

	r := regexp.MustCompile(consts.LnineImgUrl1)
	matchs := r.FindAllStringSubmatch(content, -1)
	for i, _ := range matchs {
		if len(matchs[i]) > 1 {
			imgUrls = append(imgUrls, matchs[i][1])
		}
	}

	r2 := regexp.MustCompile(consts.LnineImgUrl2)
	matchs2 := r2.FindAllStringSubmatch(content, -1)
	for i, _ := range matchs2 {
		if len(matchs2[i]) > 1 {
			imgUrls = append(imgUrls, matchs2[i][1])
		}
	}

	return imgUrls
}
