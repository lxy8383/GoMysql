package main

import (
	"controller"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
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
func sayHelloName(w http.ResponseWriter, r * http.Request)  {

	p := Product {}
	p.Name = "lxy"
	p.ProductID = 1111
	p.Number = 222
	p.Pirce = 1234

	data ,_ := json.Marshal(p)

	fmt.Println(string(data))


}




type Teacher struct {
	Name 	string
	TelNum 	int
}

func judge( i interface{} )  {

	switch i.(type) {

	case int:

		fmt.Println("整形")

	case string:

		fmt.Println("字符串")

	}
}

func sayBoy(s string)  {
	for i:= 0 ; i < 5 ; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func add( a int , b int )  {

	c := a + b

	fmt.Printf("a = %d  , b = %d , c = %d \n " ,a , b , c )

}

func addFunc( a []int , b int  )  {
	defer func() {
		err := recover()
		if err != nil{
			fmt.Println("失败")
		}
	}()
	a[b] = b
	fmt.Println(b)
}

var wg sync.WaitGroup

func addMethod(x int )  {

	fmt.Println(" i = " ,x)
	wg.Done()
}

func deleteWait()  {
	defer wg.Done()
	fmt.Println("Done ")

}

type Product struct {

	Name  string
	
	ProductID int64

	Number int

	Pirce float64


}


func main()  {

	//http.HandleFunc("/",index)
	//http.HandleFunc("/test/",test)

	//if err := http.ListenAndServe("0.0.0.0:8080",nil); err != nil {
	//	log.Fatal("listenAndService: " , err)
	//}
	//err, db := controller.OpenDB()
	//if err == false {
	//	fmt.Println("open DB error")
	//}else {
	//	fmt.Println("open DB Sucess")
	//}
	//
	//controller.InsertToDB(db)/**/

	//names := [] string {"lilei","wangfei","honghai","4","5","6",}
	//
	//for i ,v := range names {
	//	fmt.Println(i , v )
	//}

	// *********调试接口
	//var inter  controller.VideoSourceInter
	//var vid = controller.Video{"liuxy", "www,baidu.com", "1992", "milin"}
	//inter = vid
	//inter.SayHello()
	//
	//var voice = controller.Voice{"小鸡快跑","wwww.baidu.com","1922","wanglei"}
	//inter = &voice
	//inter.SayHello()

	//for i := 0 ;i < 10 ; i++ {
	//	add(10,20)
	//}
	//
	//fmt.Printf("zouni ")

	//array := make([]int , 4)
	//
	//for x := 0 ; x < 10 ; x ++ {
	//
	//	go addFunc(array, x )
	//
	//}
	//time.Sleep(time.Second * 2)

	// ********goroutin 与add 判断同步的作用
	//for i := 0 ; i < 10 ; i ++ {
	//	wg.Add(1)
	//	addMethod(i)
	//}
	//
	//wg.Wait()

	//上传图片
	//http.HandleFunc("/",controller.ReceiveFile)  // 设置访问的路由
	//
	//err := http.ListenAndServe(":9090",nil)
	//
	//if err != nil {
	//	log.Fatal("ListenAndServe " , err)
	//}

	///Users/liuxy/Documents

	// 读写文件
	//controller.WriteNewFile()

	//// 反射
	//var x float64 = 3.4
	//fmt.Println("反射类型:", reflect.ValueOf(x))

	//t := Product{"liuxy",1111,1111,12}
	//
	////s := reflect.ValueOf(&t).Elem()
	//
	//s := reflect.ValueOf(t)
	//
	//typeOfT := s.Type()
	//
	//for i := 0 ; i < s.NumField(); i++ {
	//
	//	f := s.Field(i)
	//
	//	fmt.Println(f)
	//
	//	fmt.Println("typeOfT=",typeOfT.Field(i).Name)
	//
	//	//fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	//}

	// 联系interface
	//var a Reptile.A
	//Reptile.CallFoo(a)
	//
	//var b Reptile.B
	//Reptile.CallFoo(b)



	// xorm 数据库
	// 初始化数据库
	controller.Init()

	//controller.Insert("王龙1",32)
	//
	controller.Del(1)

	//controller.Insert("王龙2",32)



}
