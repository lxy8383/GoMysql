package controller

import (
	"encoding/json"
	"net/http"
)



type Result struct {
	Code int64
	Success bool

	DataArr []string   //返回数组
}



func Login(w http.ResponseWriter, r * http.Request)  {


	profile := Result{1,true,[]string{"liuxy","mikeyu"}}

	js , err := json.Marshal(profile)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.Write(js)


	////判断传递方式
	//
	//if r.Method == "GET" {
	//	//加载界面模板
	//	t , _ := template.ParseFiles("./view/login.ctpl")
	//
	//	// 将解析好的模板应用到data上
	//	t.Execute(w,nil)
	//
	//}else if r.Method == "POST" {
	//
	//	fmt.Println("userName" , r.Form["username"])
	//	fmt.Println("userName" , r.Form["password"])
	//
	//}
}