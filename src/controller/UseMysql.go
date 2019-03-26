package controller

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)
const shortForm = "2006-01-02 15:04:05"

//pravite
func GetMD5Hash(text string) string  {
	haser := md5.New()
	haser.Write([]byte(text))
	return hex.EncodeToString(haser.Sum(nil))
}

func GetNowtimeMD5() string  {

	t := time.Now()
	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
	return GetMD5Hash(timestamp)

}

func CheckErr(err error)  {
	if err != nil {

		fmt.Println("error ", err)
	}
}

func GetTime() string {
	t := time.Now()
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second() , t.Nanosecond(), time.Local)
	str := temp.Format(shortForm)
	fmt.Println(t)
	return str
}
// 用户名：密码@tcp(本地ip以及mysql端口号)
//const DB_Driver  =   "user:password@tcp(127.0.0.1:3306)/Database?charset=utf8"

const DB_Driver  =   "root:12345678@tcp(127.0.0.1:3306)/helloOne?charset=utf8"

func OpenDB() (sucess bool, db *sql.DB)  {

	var isOpen bool
	db, err := sql.Open("mysql",DB_Driver)
	if err != nil {
		isOpen = false
	}else {
		isOpen = true
	}
	//CheckErr(err)
	return isOpen , db

}

// 插入
func InsertToDB(db *sql.DB)  {


	stmt , err := db.Prepare("insert Student set id = ? ,name = ? , number = ? , sex =? ")

	res, err := stmt.Exec("3","王龙","123456","男")

	id , err := res.LastInsertId()

	if err != nil{
		fmt.Println("插入数据失败")
	}else {
		fmt.Println("插入成功",id)
	}

}


// 查询
func QueryFromDB(db *sql.DB)  {
	rows, err := db.Query("select * from testDatabase")

	if err != nil {
		fmt.Println("query Error" , err)
	}

	//for rows.Next() {
	//	var uid string
	//	var username string
	//	var  departmentname  string
	//	var created string
	//	var password string
	//	var autid string
	//	err = rows.Scan(&uid, &username, &departmentname, &created, &password, &autid)
	//	fmt.Println(autid)
	//	fmt.Println(username)
	//	fmt.Println(departmentname)
	//	fmt.Println(created)
	//	fmt.Println(password)
	//	fmt.Println(uid)
	//
	//}
	for rows.Next() {
		var id string
		var name string

		fmt.Println(id)
		fmt.Println(name)
	}
}

// 更新
func UpdateDB(db *sql.DB, uid string) {
	stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	CheckErr(err)
	res, err := stmt.Exec("zhangqi", uid)
	affect, err := res.RowsAffected()
	fmt.Println("更新数据：", affect)
	CheckErr(err)
}
// 删除
func DeleteFromDB(db *sql.DB, autid int) {
	stmt, err := db.Prepare("delete from userinfo where autid=?")
	CheckErr(err)
	res, err := stmt.Exec(autid)
	affect, err := res.RowsAffected()
	fmt.Println("删除数据：", affect)
}







