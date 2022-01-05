package utils

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func GetTodayTime() string {
	time.Now().Add(-time.Minute * 10)

	//golang的time包里面有个AddDate方法
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, 0)
	toDay := yesTime.Format("2006-1-2")
	return toDay
}

func GetTomorrowTime() string {
	time.Now().Add(-time.Minute * 10)

	//golang的time包里面有个AddDate方法
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, 1)
	tomorrowDay := yesTime.Format("2006-1-2")
	return tomorrowDay
}

func WriteContentToHtml(filepath, content string) {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file error=%v\n", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	writer.WriteString(content)

	writer.Flush()
}

func DeleteFile(file string) {

	err := os.Remove(file) //删除文件
	if err != nil {
		//如果删除失败则输出 file remove Error!
		fmt.Println("file remove Error!")
		//输出错误详细信息
		//fmt.Printf("%s", err)
	}
}


func Input() string{
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("域名示例：https://www.daque.cn/softdown")
	fmt.Println("请输入域名:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return "输入错误 请重新输入"
	}

	fmt.Printf("开始存储域名为 %s", input,"的html文件")
	return input
}