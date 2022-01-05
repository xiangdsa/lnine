package dao

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetLnineImgUrl(t *testing.T) {
	//GetLnineImgUrl("<img src=\"http://img.lnine.net/up/2006/2020612141412875970.jpg\"<img border=\"0\" src=\"http://img.lnine.net/uploadfiles/2017-07-11/20170711_092721_300.jpg\"/>")
	//GetAllArticleUrl("http://www.lnine.net/article/index.html")

	//resp, _ := http.Get("http://www.lnine.net/az/38121.html")
	//fmt.Println(resp.StatusCode)
	//
	//_, _, _, imgurl := GetComputerUrl("http://www.lnine.net/xiazai/1.html")
	//fmt.Println(imgurl)
	//GetAZUrl("http://www.lnine.net/az/11.html")
	//utils.RemoveDuplicates("C:\\Users\\BZ\\Desktop\\seo游当网\\1111.txt")
	//strMap :=make(map[string]string)
	//strSlice := []string {"slice","int","string","int","boolean","string"}
	//for _,v:= range strSlice{
	//	strMap[v] = v
	//}
	//fmt.Println(strMap)
	//resp, _ := http.Get("https://chrome.google.com/webstore/detail/neatdownloadmanager-exten/cpcifbdmkopohnnofedkjghjiclmhdah")
	//fmt.Println(resp.StatusCode)
	//sre:=[]rune("<img 你好alt=\"\" src=\"http://img.lnine.net/1/8/2f6bbc4f-ab76-455d-81fb-82020eba39a8.jpg\" style=\"width: 500px; height: 314px;\"/>")
	//fmt.Println(len(sre))
	//s:="<img 你好alt=\"\" src=\"http://img.lnine.net/1/8/2f6bbc4f-ab76-455d-81fb-82020eba39a8.jpg\" style=\"width: 500px; height: 314px;\"/>"
	//fmt.Println(len(s))

	resp,_:=http.Get("http://www.dyjqd.com/api-baizhu.php?flag=2&id=18073")
	body,_:=ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
