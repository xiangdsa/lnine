package dao

import (
	"fmt"
	"regexp"
	"strings"
	"updateLnineDownloadMorepic/consts"
	"updateLnineDownloadMorepic/models"
)

//安卓替换需要两张表
func AZImgTagNumber(bzSeoEcmsAzData1 models.BzSeoEcmsAzData1) (newContent string) {

	title:=QueryAzTitleByData1ID(bzSeoEcmsAzData1.ID)

	var oldImgTags, newImgTags []string

	azModels := []string{consts.AZImgTagModel1,consts.AZImgTagModel2,consts.AZImgTagModel3}

	oldImgTags = GetOldImgTag(azModels, bzSeoEcmsAzData1.Newstext)

	var resultCode int
	var altName string
	for i, _ := range oldImgTags {
		resultCode = CheckAlt(oldImgTags[i])
		if resultCode == 2 {
			r := regexp.MustCompile("alt=\"([\u4e00-\u9fa5 A-Za-z0-9]+)\"")
			matchs := r.FindAllStringSubmatch(oldImgTags[i], 1)
			if len(matchs) > 0 {
				altName = matchs[0][1]
			}
			break
		}
	}

	//resultCode 0：无alt 1：alt=“”  2：alt=“此处是标题”
	newImgTags = ReplaceContent(resultCode, oldImgTags, title, altName)

	newContent = ReturnNewContent(oldImgTags, newImgTags, bzSeoEcmsAzData1.Newstext)

	fmt.Println(newContent)

	return newContent
}

//直接把标签换成新的标签样式
func ComputerImgTagNumber(bzSeoEcmsDownload models.BzSeoEcmsDownload) (newContent string) {
	var oldImgTags, newImgTags []string

	computerModels := []string{consts.ComputerImgTagModel1, consts.ComputerImgTagModel2}

	oldImgTags = GetOldImgTag(computerModels, bzSeoEcmsDownload.Softsay)

	var resultCode int
	var altName string
	for i, _ := range oldImgTags {
		resultCode = CheckAlt(oldImgTags[i])
		if resultCode == 2 {
			r := regexp.MustCompile("alt=\"([\u4e00-\u9fa5 A-Za-z0-9]+)\"")
			matchs := r.FindAllStringSubmatch(oldImgTags[i], 1)
			if len(matchs) > 0 {
				altName = matchs[0][1]
			}
			break
		}
	}

	newImgTags = ReplaceContent(resultCode, oldImgTags, bzSeoEcmsDownload.Title, altName)

	newContent = ReturnNewContent(oldImgTags, newImgTags, bzSeoEcmsDownload.Softsay)

	fmt.Println("加上序号后的content：")
	fmt.Println(newContent)
	return newContent
}

//传入alt的状态码 老的tag字符串切片 标题 标签名称（可能为空）  返回新的标签字符串切片
func ReplaceContent(resultCode int, oldImgTags []string, title string, altName string) (newImgTags []string) {
	for i, _ := range oldImgTags {
		if resultCode == 2 {
			//此处的正则是因为每个oldImgTag中的altname可能都不一样所以每次都需要单独使用正则得到altname去替换
			altName1 := fmt.Sprintf("%s(图%d)", title, i+1)
			r := regexp.MustCompile("alt=\"([\u4e00-\u9fa5 A-Za-z0-9]+)\"")
			matchs := r.FindAllStringSubmatch(oldImgTags[i], 1)
			var altName2 string
			if len(matchs) > 0 {
				altName2 = matchs[0][1]
			}
			newImgTags = append(newImgTags, strings.Replace(oldImgTags[i], altName2, altName1, 1))
		} else if resultCode == 0 {
			altName = fmt.Sprintf("alt=\"%s(图%d)\"", title, i+1)
			newImgTags = append(newImgTags, strings.Replace(oldImgTags[i], "/>", " "+altName, 1))
		} else {
			altName = fmt.Sprintf("alt=\"%s(图%d)\"", title, i+1)
			altName1 := fmt.Sprintf("%s(图%d)", title, i+1)
			r := regexp.MustCompile("alt=\"([\u4e00-\u9fa5 A-Za-z0-9\\(\\)]+)\"")
			matchs := r.FindAllStringSubmatch(oldImgTags[i], 1)
			var altName2 string
			if len(matchs) > 0 {
				altName2 = matchs[0][1]
			}
			if altName2 != "" {
				newImgTags = append(newImgTags, strings.Replace(oldImgTags[i], altName2, altName1, 1))
			} else {
				newImgTags = append(newImgTags, strings.Replace(oldImgTags[i], "alt=\"\"", altName, 1))
			}
		}
	}
	return
}

//返回0 没有alt 返回1 有Alt但是alt为空 返回2 有Alt且alt不为空
func CheckAlt(oldImgTag string) (haveAlt int) {
	if ok, _ := regexp.MatchString("(.*)alt=\"\"(.*)", oldImgTag); ok {
		return 1
	}

	if ok, _ := regexp.MatchString("(.*)alt(.*)", oldImgTag); ok {
		return 2
	}

	return 0
}

func GetOldImgTag(imgTagModel []string, content string) (oldImgTags []string) {
	for i, _ := range imgTagModel {
		r := regexp.MustCompile(imgTagModel[i])
		matchs := r.FindAllStringSubmatch(content, -1)
		for i1, _ := range matchs {
			if len(matchs[i1]) > 0 {
				oldImgTags = append(oldImgTags, matchs[i1][0])
			}
		}
	}

	return
}

func ReturnNewContent(oldImgTags, newImgTags []string, content string) string {
	var remainContent string
	remainContent = content
	for i, _ := range oldImgTags {

		content = strings.Replace(content, oldImgTags[i], newImgTags[i], 1)
		if content == "" {
			content = remainContent
		}

	}

	return content
}

//隐患 在全部加完标签后 检查是否存在两个标签的情况  例如：alt=“” alt=“此处是标题”
