package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Animal interface {
	Speak() string

}

type Student struct {
	num   	int64
	name  	string
	tel 	int64
	score   int64
}

type Dog struct {

}

func (stu Student) Speak() string {
	return "I'm Student"
}

func (dog Dog) Speak() string  {
	return "I'm Dog"
}


func Adder() func(int ) int   {
	var x int
	return func(delete int) int {
		x += delete
		return x
	}
}




func index(w http.ResponseWriter, req * http.Request)  {
	req.ParseForm()
	fmt.Println("Form : " , req.Form)
	fmt.Println("Form : " , req.URL.Path)
	fmt.Println("Form : " , req.Form["a"])
	fmt.Println("Form : " , req.Form["b"])
}

func test(w http.ResponseWriter,r * http.Request)  {
	body , _ := ioutil.ReadAll(r.Body)

	body_str := string(body)

	fmt.Println(body_str)
}
func handler(w http.ResponseWriter, r * http.Request)  {
	var body json.RawMessage
	// todo

	b, _ := body.MarshalJSON()
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	w.Write(b)
}

func main()  {

	//http.HandleFunc("/",index)
	//http.HandleFunc("/test/",test)

	if err := http.ListenAndServe("0.0.0.0:8080",nil); err != nil {
		log.Fatal("listenAndService: " , err)
	}
	//err, db := controller.OpenDB()
	//if err == false {
	//	fmt.Println("open DB error")
	//}else {
	//	fmt.Println("open DB Sucess")
	//}
	//
	//controller.InsertToDB(db)/**/

}