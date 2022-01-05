package consts

import (
	"fmt"
	"regexp"
	"testing"
)

func TestZhengZe(t *testing.T) {
	r:=regexp.MustCompile("<img [\u4e00-\u9fa5 =\"\\sA-Za-z0-9]{0,40}src=\"http://img.lnine.net/[0-9a-zA-Z-_/]+.[a-zA_Z]{3,5}\"/>")
	//r:=regexp.MustCompile(AZImgTagModel2)

	matchs:=r.FindAllStringSubmatch("<img src=\"http://img.lnine.net/5/16/108e87ca-a407-4d24-a864-85d755074501.jpg\"/>",-1)
	fmt.Println(matchs)

}

