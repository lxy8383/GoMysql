package Reptile

import "fmt"

type Foo interface {

	fooMethod()

	twoMethod()
}

type A struct {

}

type B struct {

}

func (a  A ) fooMethod() {
	fmt.Println("foo 内容 ")
}

func (a A) twoMethod() {
	fmt.Println("Two method ")
}

func (a  A ) callBar()  {
	fmt.Println("唤醒一个bar ")
}

func (b B) fooMethod() {
	fmt.Println("B foo 的内容 ")
}
func (b B)twoMethod() {
	fmt.Println("B Two Method ")
}

func CallFoo(f Foo)  {
	f.fooMethod()
	f.twoMethod()

}






