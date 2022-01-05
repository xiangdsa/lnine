package utils

import (
	"createHtml/dao"
	"fmt"
	"testing"
)

func TestWriteContentToHtml(t *testing.T) {
	content,_:=dao.GetHTMLCode("https://www.daque.cn/")
	//content,_:=dao.GetHTMLCode("https://www.daque.cn/softdown/46.html")
	fmt.Println(content)
	//WriteContentToHtml("./itemcode.csv",content)
	WriteContentToHtml("C:\\Users\\BZ\\Desktop\\大雀图片\\12.html",content)
}


func TestDeleteFile(t *testing.T) {
	//for i:=43884;i<=44102;i++{
	//	file:=fmt.Sprintf("F:\\GitProject\\createHtml\\%d.html",i)
	//	DeleteFile(file)
	//}

	DeleteFile("C:\\Users\\BZ\\Documents\\WeChat Files\\wxid_jdcks9589nwl22\\FileStorage\\File\\2021-11\\ff0ba502b3814fade2123c99e72c34cc_t.gif11")
}

func TestWriteContentToHtml2(t *testing.T) {

}