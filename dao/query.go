package dao

import (
	"createHtml/consts"
	"createHtml/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"time"
)

var DB *gorm.DB

func InitMysqlCon(dns string) (err error) {
	DB, err = gorm.Open("mysql", dns)
	if err != nil {
		fmt.Println("链接数据库出问题:", err)
	}
	err = DB.DB().Ping()

	DB.DB().SetMaxIdleConns(64)

	DB.DB().SetMaxOpenConns(64)

	DB.DB().SetConnMaxLifetime(600 * time.Second)

	//禁用复数表名
	DB.SingularTable(true)

	return err
}

func GetHTMLCode(url string) (string, error) {
	request, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}

	resp, _ := client.Do(request)
	fmt.Println(resp)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}

func GetUrl(todayTime, tomorrowTime string) []string {
	softWares := GetSoftWareID(todayTime,tomorrowTime)
	urls := make([]string, 0)
	for i, _ := range softWares {

		url := fmt.Sprintf(consts.AddUrl,softWares[i])

		urls = append(urls, url)
	}
	fmt.Println(urls)
	return urls
}

func GetSoftWareID(todayTime, tomorrowTime string) (softWareID []int) {
	softWare := QuerySoftWare(todayTime, tomorrowTime)
	for i, _ := range softWare {
		softWareID = append(softWareID, int(softWare[i].ID))
	}
	return
}

func QuerySoftWare(todayTime, tomorrowTime string) (softWare []*models.WebsiteSoftware) {
	if todayTime == "" || tomorrowTime == "" {
		fmt.Println("today is nil")
		return nil
	}

	DB.Table("website_software").Where("update_time>=? and update_time<?", todayTime, tomorrowTime).Scan(&softWare)

	return
}

func QueryContent(todayTime, tomorrowTime string) (softWareContents []*models.WebsiteSoftwareContent) {
	if todayTime == "" || tomorrowTime == "" {
		fmt.Println("today is nil")
		return nil
	}

	softWare := make([]*models.WebsiteSoftware, 0, 10)
	DB.Table("website_software").Where("update_time>=? and update_time<?", todayTime, tomorrowTime).Scan(&softWare)

	softWareContent := make([]*models.WebsiteSoftwareContent, 0, 10)

	if softWare != nil {
		for i, _ := range softWare {
			DB.Table("website_software_content").Where("software_id=?", softWare[i].ID).Scan(&softWareContent)
			if softWareContent != nil {
				softWareContents = append(softWareContents, softWareContent[0])
			}
		}
	}

	return
}
