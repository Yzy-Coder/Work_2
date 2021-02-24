package main

import "fmt"

//   Receiver 函数 通过空接口和Switch来根据数据类型做出反应
func Receiver( v interface{} ) {

	switch v.(type) {	//根据接口数据类型做出不同反应

	case int:
		fmt.Println("这是一个int型数据")
	case string:
		fmt.Println("这是一个string型数据")
	case byte:
		fmt.Println("这是一个byte型数据")
	case bool:
		fmt.Println("这是一个bool型数据")
	default:
		fmt.Println("暂时无法识别该类型")
	}
}
func main() {

	var ByteTest byte

	fmt.Scanf( "%c" , &ByteTest)

	Receiver(32)	//测试Int类型

	Receiver("Yzy") //测试String类型

	Receiver( ByteTest ) //测试Byte类型

	Receiver(true)	//测试Bool类型

	Receiver( make( []int , 1 , 1)  )	//测试其他情况


}