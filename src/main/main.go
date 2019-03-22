package main

import (
	"fmt"
	"io"
	"net/http"
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

func Hello( x int, y int )  {
	sum := x + y

}


func Adder() func(int ) int   {
	var x int
	return func(delete int) int {
		x += delete
		return x
	}



}


func helloHandle(w http.ResponseWriter, req * http.Request)  {
	io.WriteString(w , "hello , World ! \n ")
}


func main()  {
	//http.HandleFunc("/",helloHandle)
	//log.Println("启动了")
	//err := http.ListenAndServe(":9000",nil)
	//if err != nil {
	//	log.Fatal("list 9000")
	//}

	var stu1 Student
	stu1.name = "1111"
	stu1.num = 123

	var stu2 Student
	stu2.name = "1111"
	stu2.num = 123

	arr := []Student{stu1,stu2}

	fmt.Println(arr)


}