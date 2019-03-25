package main

import (
	"controller"
	"fmt"

	//"fmt"
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
	err, db := controller.OpenDB()
	if err == false {
		fmt.Println("open DB error")
	}else {
		fmt.Println("open DB Sucess")
	}

	controller.InsertToDB(db)

}