package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const saveFileName  =  "/Users/liuxy/Documents/ImgTest/"

type Auth struct {
	Username 	string
	Pwd 		string
}

type Result struct {
	Code 	int64
	//Success bool
	Msg 	string
	//DataArr []string   //返回数组
}

func ReceiveFile(w http.ResponseWriter, r * http.Request)  {

	var result Result

	file , file_head, file_err := r.FormFile("fileName")

	if file_err != nil {
		fmt.Println("失败")
	}
	fileSave := saveFileName + file_head.Filename

	f ,f_err := os.OpenFile(fileSave,os.O_WRONLY | os.O_CREATE,0666)
	if f_err != nil {
		fmt.Fprintf(w, "file open fail:%s", f_err)
		return
	}
	//文件 copy
	_, copy_err := io.Copy(f, file)

	if copy_err != nil {
		fmt.Fprintf(w, "file copy fail:%s", copy_err)
		return
	}
	//关闭对应打开的文件
	defer f.Close()
	defer file.Close()

	result.Code = 200
	result.Msg = "上传成功"

	//fmt.Fprintf(w, fileSave)
	json.NewEncoder(w).Encode(result)
	// 之后需要将文件路劲添加到数据库中
}

func Login(w http.ResponseWriter, r * http.Request)  {

	// 获取表单
	r.ParseForm()
	username , uError := r.Form["username"]
	pwd , pError := r.Form["password"]

	fmt.Printf("%s",username)
	fmt.Printf("%s",pwd)
	var result  Result
	if !uError || !pError {
		result.Code = 401
		result.Msg = "登录失败"
	} else if username[0] == "admin" && pwd[0] == "123456" {
		result.Code = 200
		result.Msg = "登录成功"
	} else {
		result.Code = 202
		result.Msg = "账户名或密码错误"
	}
	json.NewEncoder(w).Encode(result)

}


