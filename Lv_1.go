package main

import "fmt"

// 鸽子接口，具有Gugugu方法
type	Dover		interface {
	Gugugu()
}

// 复读机接口，具有Repeat方法
type	Repeater	interface {
	Repeat( Word string, Times int )
}

// 柠檬精接口,具有Envy方法
type 	Envier		interface {
	Envy()
}

// 真香怪接口,具有Zhenxiang方法
type 	Zhenxianger interface {
	Zhenxiang()
}

// 表示人的结构体
type Person struct {

	Name 			string	//姓名
	Age 			int		//年龄
	GuguguNum 		int		//运行Gugugu方法的次数
	RepeatNum 		int		//运行Repeat方法的次数
	EnvyNum			int		//运行Envy方法的次数
	ZhenxiangNum	int		//运行Zhenxiang方法的次数
}

//实现Gugugu方法
func (Person *Person) Gugugu (){

	fmt.Println(Person.Name+"鸽了一次")
	Person.GuguguNum++
}

//实现Repeat方法
func (Person *Person) Repeat ( Word string, Times int ){

	for  i := 0 ; i < Times ; i++ {
		fmt.Println( Word )
	}
	Person.RepeatNum++
}

//实现Envy方法
func (Person *Person) Envy (){

	fmt.Println(Person.Name+":啊，好酸啊")
	Person.EnvyNum++
}

//实现Zhenxiang方法
func (Person *Person) Zhenxiang (){

	fmt.Println("诶，真香嘿")
	Person.ZhenxiangNum++
}


func main() {

	Yzy := &Person{ "Yzy" , 20 ,0,0,0,0 }	//定义并初始化一个Person指针变量

	var Dover Dover = Yzy	//测试是否实现鸽子接口
	var Repeater Repeater = Yzy	///测试是否实现复读机接口
	var Envier	Envier	= Yzy	//测试是否实现柠檬精接口
	var Zhenxianger Zhenxianger = Yzy	//测试是否实现真香怪接口

	Dover.Gugugu()
	Repeater.Repeat("chy",10)
	Envier.Envy()
	Zhenxianger.Zhenxiang()

	fmt.Printf("目前Yzy的属性为%v" , *Yzy ) //打印Yzy目前的属性，判断结果
}