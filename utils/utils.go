package utils

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func CheckUrl(url []string) (url200, url404 []string) {
	for i, _ := range url {
		resp, _ := http.Get(url[i])
		if resp != nil {
			if resp.StatusCode == 200 {
				url200 = append(url200, url[i])
			} else {
				url404 = append(url404, fmt.Sprintf("链接：%s  错误码是：%d", url[i], resp.StatusCode))
			}
		} else {
			url404 = append(url404, fmt.Sprintf("链接：%s 未响应", url[i]))
		}

	}
	return
}

func WriteToTXT(urls []string, filePath string) {

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for i, _ := range urls {
		write.WriteString(urls[i] + "\n")
	}
	//Flush将缓存的文件真正写入到文件中
	err = write.Flush()
	if err != nil {
		fmt.Errorf("write flush err :%s", err)
	}
}

func RemoveDuplicates(filepath string) []string{
	urls := ReaderTxt(filepath,"\n")
	strMap := make(map[string]string, 0)
	for _, v := range urls {
		strMap[v] = v
	}
	urls=urls[0:0]
	for _,value:=range strMap {
		urls=append(urls,value)
	}

	return urls
}

func ReaderTxt(filepath string,oldFenGeFu string) (urls []string) {
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败：", err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		url, err := reader.ReadString('\n')
		urls = append(urls, strings.Replace(url,oldFenGeFu,"",-1))
		if err == io.EOF {
			break
		}
	}

	return
}

func ElementIsInSlice(element string, elements []string) (isIn bool) {
	for _, item := range elements {
		if element == item {
			isIn = true
			return
		}
	}
	return
}

//覆盖之前的文件内容  os.O_TRUNC
func ReWriter(filepath string, oldurl []string) {

	f, _ := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0600)

	defer f.Close()

	f.WriteString(strings.Join(oldurl, ""))

}
