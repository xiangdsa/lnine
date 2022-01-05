package dao

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
	"updateLnineDownloadMorepic/consts"
)

func GetAllArticleUrl(url string) (articleUrls []string) {

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

		//获得所有的页数 /html/body/div[3]/div[3]/div[1]/div[4]/div/a[9]
		lastPageNumberNode := htmlquery.Find(doc, "/html/body/div[3]/div[3]/div[1]/div[4]/div/a[9]")

		reader := strings.NewReader(htmlquery.OutputHTML(lastPageNumberNode[0], true))

		homeDocs, _ := goquery.NewDocumentFromReader(reader)

		lastPageNumber, _ := homeDocs.Find("a").Attr("href")

		pages := strings.Replace(strings.Replace(lastPageNumber, "/article/index_", "", 1), ".html", "", 1)

		page, _ := strconv.Atoi(pages)

		articleUrls = GetPageArticle(page)

	})

	c.Visit(url)

	fmt.Println(len(articleUrls))
	return
}

func GetPageArticle(page int) (articleUrls []string) {

	var articlePageUrls []string

	for i := 1; i <= page; i++ {

		if i == 1 {

			articlePageUrls = append(articlePageUrls, "http://www.lnine.net/article/index.html")

		} else {

			articlePageUrls = append(articlePageUrls, fmt.Sprintf(consts.ArticlePageModel, i))

		}
	}

	c := colly.NewCollector(

		colly.AllowedDomains("www.lnine.net", "lnine.net"),

		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),

	)

	c.DetectCharset = false

	c.OnHTML("body", func(element *colly.HTMLElement) {

		// /html/body/div[3]/div[3]/div[1]/div[4]/ul/li/div/a[1]
		doc, _ := htmlquery.Parse(strings.NewReader(string(element.Response.Body)))

		nodes := htmlquery.Find(doc, "/html/body/div[3]/div[3]/div[1]/div[4]/ul/li/div/a[1]")

		for i, _ := range nodes {

			reader := strings.NewReader(htmlquery.OutputHTML(nodes[i], true))

			homeDocs, _ := goquery.NewDocumentFromReader(reader)

			articleUrl, _ := homeDocs.Find("a").Attr("href")

			articleUrls = append(articleUrls, fmt.Sprintf("http://www.lnine.net%s",articleUrl))
		}
	})

	for i, _ := range articlePageUrls {

		c.Visit(articlePageUrls[i])

	}

	return
}
