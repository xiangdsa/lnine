package dao

import (
	"fmt"
	"testing"
)

//https://www.daque.cn/softdown/43468.html

func TestGetHtmlCode(t *testing.T) {
	//GetByProxy("https://www.daque.cn/softdown/43468.html")
	index,_:=GetHTMLCode("https://www.daque.cn/")
	fmt.Println(index)
}

