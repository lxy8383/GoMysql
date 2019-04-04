package controller

import (
	"Tool"
	"fmt"
	"io/ioutil"
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


// 读取文件夹以及文件夹下的所有文件 利用递归方法
func ListFile(myfolder string)  {
	files , err := ioutil.ReadDir(myfolder)

	//调用打印
	Tool.IsError(err,"读取文件")

	for _ , file := range files {
		if file.IsDir(){
			ListFile(myfolder + "/" + file.Name())
		} else {
			fmt.Println(myfolder + "/" + file.Name())
		}
	}
}






