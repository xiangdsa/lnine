package dao

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
	"net/http"
	"regexp"
	"strings"
	"updateLnineDownloadMorepic/consts"
	"updateLnineDownloadMorepic/utils"
)

//  电脑软件a标签的位置   /html/body/div[3]/div[4]/div[1]/div[2]/div[1]/p/a
//  电脑软件img标签的位置 /html/body/div[3]/div[4]/div[1]/div[2]/div[1]/p/img
// 特殊情况处理
//<a href="http://img.lnine.net/4/29/096af82b-d504-4ea4-b3e8-36cfe4935484.png" target="_blank">
//<img alt="手机万能变声器app" src="http://img.lnine.net/29/KDYwMHgp/096af82b-d504-4ea4-b3e8-36cfe4935484.png" style="height: 333px; width: 200px;"/></a>

func GetComputerUrl(url string) (computerUrls, computerUrls404, imgUrls404,imgurl []string) {

	c := colly.NewCollector(

		colly.AllowedDomains("www.lnine.net", "lnine.net"),

		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),

	)

	//c.DetectCharset = false

	c.OnHTML("body", func(element *colly.HTMLElement) {

		doc, err := htmlquery.Parse(strings.NewReader(string(element.Response.Body)))

		if err != nil {
			fmt.Println("htmlQuery.Parse error :", err)
		}
		///html/body/div[3]/div[4]/div[1]/div[2]/div[1]/p[12]
		//获得所有的超链接节点
		computerUrlNode := htmlquery.Find(doc, "/html/body/div[3]/div[4]/div[1]/div[2]/div[1]/p")

		for i, _ := range computerUrlNode {
			reader := strings.NewReader(htmlquery.OutputHTML(computerUrlNode[i], true))

			homeDocs, _ := goquery.NewDocumentFromReader(reader)

			computerUrl, _ := homeDocs.Find("a").Attr("href")

			if computerUrl != "" {
				if ok, _ := regexp.MatchString("(.*)img.lnine.net(.*)", computerUrl); ok {
					resp, _ := http.Get(computerUrl)
					if resp!=nil{
						if resp.StatusCode != 200 {
							imgUrls404 = append(imgUrls404,  fmt.Sprintf("%s 链接的 %s 错误",url,computerUrl))
						}else {
							imgurl=append(imgurl,computerUrl)
						}
					}
				} else {

					resp, _ := http.Get(computerUrl)
					if resp!=nil{
						if resp.StatusCode != 200 {
							computerUrls404 = append(computerUrls404,  fmt.Sprintf("%s 链接的 %s 错误",url,computerUrl))
						} else {
							computerUrls = append(computerUrls, computerUrl)
						}
					}else{
						computerUrls = append(computerUrls, computerUrl)
					}
				}
			}

			imgUrl, _ := homeDocs.Find("img").Attr("src")
			if imgUrl != "" {
				resp, _ := http.Get(imgUrl)
				if resp!=nil&&resp.StatusCode != 200 {
					imgUrls404 = append(imgUrls404, imgUrl)
				}else {
					imgurl=append(imgurl,imgUrl)
				}
			}

		}

	})

	resp,_:=http.Get(url)
	if resp!=nil&&resp.StatusCode==200{
		c.Visit(url)
	}else{
		utils.WriteToTXT([]string{url}, consts.Computer404Path)
	}

	if len(computerUrls) > 0 {
		utils.WriteToTXT(computerUrls, consts.ComputerPathPath)
	}
	if len(computerUrls404) > 0 {
		utils.WriteToTXT(computerUrls404, consts.Computer404Path)
	}
	if len(imgUrls404) > 0 {
		utils.WriteToTXT(imgUrls404, consts.ComputerImgPath)
	}
	utils.WriteToTXT(imgurl, consts.ComputerImg)

	return
}

// 安卓A标签与img标签的位置 /html/body/div[3]/div[2]/div[1]/div[2]/div[2]/p
//这个特殊情况需要处理
//<a href="http://img.lnine.net/12/9/dcf7b351-5c9f-4c05-9179-86054747e745.jpg" target="_blank" rel="nofollow">
//<img alt="PK XD国际服" src="http://img.lnine.net/9/KDYwMHgp/dcf7b351-5c9f-4c05-9179-86054747e745.jpg" style="width: 420px; height: 236px;"></a>
func GetAZUrl(url string) (azUrls, azUrls404, imgUrls404,imgurls []string) {
	c := colly.NewCollector(

		colly.AllowedDomains("www.lnine.net", "lnine.net"),

		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),

	)

	c.DetectCharset = false

	c.OnHTML("body", func(element *colly.HTMLElement) {

		doc, err := htmlquery.Parse(strings.NewReader(string(element.Response.Body)))

		if err != nil {
			fmt.Println("htmlQuery.Parse error :", err)
		}

		computerUrlNode := htmlquery.Find(doc, "/html/body/div[3]/div[2]/div[1]/div[2]/div[2]/p")

		for i, _ := range computerUrlNode {
			reader := strings.NewReader(htmlquery.OutputHTML(computerUrlNode[i], true))

			homeDocs, _ := goquery.NewDocumentFromReader(reader)

			azUrl, _ := homeDocs.Find("a").Attr("href")

			if azUrl != "" {

				if ok, _ := regexp.MatchString("(.*)img.lnine.net(.*)", azUrl); ok {
					resp, _ := http.Get(azUrl)
					if resp!=nil&&resp.StatusCode != 200 {
						imgUrls404 = append(imgUrls404,  fmt.Sprintf("%s 链接的 %s 错误",url,azUrl))
					}else{
						imgurls=append(imgurls,azUrl)
					}
				} else {

					resp, _ := http.Get(azUrl)
					if resp.StatusCode != 200 {
						azUrls404 = append(azUrls404, fmt.Sprintf("%s 链接的 %s 错误",url,azUrl))
					} else {
						azUrls = append(azUrls, azUrl)
					}
				}

			}

			imgUrl, _ := homeDocs.Find("img").Attr("src")
			if imgUrl != "" {
				resp, _ := http.Get(imgUrl)
				if resp!=nil&&resp.StatusCode != 200 {
					imgUrls404 = append(imgUrls404,  fmt.Sprintf("%s 链接的 %s 错误",url,imgUrl))
				}else{
					imgurls=append(imgurls,imgUrl)
				}
			}

		}

	})

	resp,_:=http.Get(url)
	if resp!=nil&&resp.StatusCode==200{
		c.Visit(url)
	}else{
		utils.WriteToTXT([]string{url}, consts.Az404Path)
	}

	utils.WriteToTXT(azUrls, consts.AzPathPath)
	utils.WriteToTXT(azUrls404, consts.Az404Path)
	utils.WriteToTXT(imgUrls404, consts.AzImgPath)
	utils.WriteToTXT(imgurls, consts.AzImg)

	return
}

// ios 无A标签         img标签的位置 /html/body/div[3]/div[2]/div[1]/div[3]/div[2]/div[1]/div[1]/ul/li
func GetIosUrl(url string) (computerUrls, imgUrls []string) {
	c := colly.NewCollector(

		colly.AllowedDomains("www.lnine.net", "lnine.net"),

		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),

	)

	c.DetectCharset = false

	c.OnHTML("body", func(element *colly.HTMLElement) {

		doc, err := htmlquery.Parse(strings.NewReader(string(element.Response.Body)))

		if err != nil {
			fmt.Println("htmlQuery.Parse error :", err)
		}

		computerUrlNode := htmlquery.Find(doc, "/html/body/div[3]/div[2]/div[1]/div[3]/div[2]/div[1]/div[1]/ul/li")

		for i, _ := range computerUrlNode {
			reader := strings.NewReader(htmlquery.OutputHTML(computerUrlNode[i], true))

			homeDocs, _ := goquery.NewDocumentFromReader(reader)

			imgUrl, _ := homeDocs.Find("img").Attr("src")
			if imgUrl != "" {
				imgUrls = append(imgUrls, imgUrl)
			}

		}

	})

	c.Visit(url)

	return
}
