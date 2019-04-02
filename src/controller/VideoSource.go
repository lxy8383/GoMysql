package controller

import (
	"fmt"
)

type VideoSourceInter interface {

	//SaveVideo( name string, url string,time string,author string)

	SayHello()
}

type Video struct {
	Name	string	//名字
	Url		string	//链接
	Time	string	//时间
	Author	string	//作者
}

type Voice struct {
	Name 	string
	Url		string
	Time 	string
	Author	string
}

func (this Video) SayHello()  {

	fmt.Println("hello nihao 看电影 ")

}

func (this Video) SaveVideo( name string, url string,time string,author string)  {
	this.Name = name
	this.Url = url
	this.Time = time
	this.Author = author

	fmt.Println("video 保存的值",name,url,time,author)
}

func (this * Voice) SayHello() {

	fmt.Println("hello nihao 声音 ")

}



