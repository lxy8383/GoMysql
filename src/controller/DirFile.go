package controller

import (
	"fmt"
	"os"
)

// 创建文件

func WriteNewFile()  {
	fileName := "/Users/liuxy/Documents/StudyGo/test_write.text"

	file , err := os.OpenFile(fileName, os.O_CREATE | os.O_WRONLY, 0755)

	if err != nil {
		fmt.Println("error ", err)

		os.Exit(1)
	}

	defer file.Close()

	fileString := "Today very happy "

	file.WriteString(fileString)

}