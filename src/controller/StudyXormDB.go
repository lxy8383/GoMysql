package controller

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type User struct {

	UserId		int64
	Name		string
	Balance		float64
	Time 		int64
	CreateTime 	int64
}

var x *xorm.Engine

func Init()  {

	var err error
	//"root:12345678@tcp(127.0.0.1:3306)/helloOne?charset=utf8"
	x , err = xorm.NewEngine("mysql","root:12345678@tcp(127.0.0.1:3306)/helloOne?charset=utf8")

	if err != nil {
		fmt.Println("数据库连接失败")
	} else {
		fmt.Println("数据库连接成功")
	}


	//建表
	err =  x.Sync2(new(User))
	if err != nil {
		fmt.Println("建表失败")
	}
}

// 插入
func Insert(name string, balance float64) (int64, bool) {

	var user User
	user.Name = name
	user.Balance = balance

	affected , err := x.Insert(user)

	if err != nil {
		return affected, false
	}

	return affected, true
}

//删除
func Del(id int64)  {
	var user User
	x.ID(id).Delete(user)
}

// 改
func Update(id int64 , user User ) bool {
	affected , err := x.ID(id).Update(user)
	if err != nil {
		fmt.Println("错误", err)
	}
	if affected == 0 {
		return false
	}

	return true
}


//查
func  Getinfo(id int64) User  {

	user := User{UserId:id}

	is , _ := x.Get(user)

	if !is {
		fmt.Println("搜索结果不存在")
	}
	return user


}


