package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"test/model"
	"time"
	"unsafe"
)

func main() {
	/*
		整数类型
	*/
	var n1 = 100
	fmt.Printf("n1的数据类型为%T \n", n1)
	var n2 int64 = 10
	fmt.Printf("n2的数据类型为%T,n2占用的字节数是%d \n", n2, unsafe.Sizeof(n2))
	//保小不保大,使用小内存
	var age byte = 100
	fmt.Println("age的值为", age)
	// 开发中使用内存占用小的
	fmt.Println("-------------------------------------------------------------------")

	/*
		小数类型/浮点型
	*/
	var num3 float32 = -123.000901
	var num4 float64 = -123.000901
	fmt.Println("num3的值为:", num3)
	fmt.Println("num4的值为:", num4)
	var num5 = 1.1
	fmt.Printf("num5的数据类型为%T \n", num5)
	num6 := 5.1234e2
	num7 := 5.1234e2
	num8 := 5.1234e-2
	fmt.Println("num6=", num6, "num7=", num7, "num8=", num8)
	// 开发中使用float64,数值更精准
	fmt.Println("-------------------------------------------------------------------")

	/*
		字符类型
	*/
	var c1 byte = 'a'
	var c2 byte = '0'
	fmt.Println("c1=", c1, "c2=", c2)
	fmt.Printf("c1=%c c2=%c \n", c1, c2)
	var c3 int = '北'
	fmt.Printf("c3=%c c3对应的码值=%d\n", c3, c3)
	var c4 int = 22269
	fmt.Printf("c4=%c\n", c4)

	var nn1 = 10 + 'a'
	fmt.Println("nn1=", nn1)
	fmt.Println("-------------------------------------------------------------------")
	/*
		布尔类型
	*/
	var b = false
	fmt.Println("b的占用空间=", unsafe.Sizeof(b))
	fmt.Println("-------------------------------------------------------------------")
	/*
		基本数据类型转换
	*/
	var i int32 = 100
	var ni float32 = float32(i)
	fmt.Printf("i=%v ni=%v\ni的数据类型为%T ni的数据类型为%T\n", i, ni, i, ni)
	var num11 int64 = 999999
	var num12 int8 = int8(num11)
	fmt.Println("num12=", num12)
	fmt.Println("-------------------------------------------------------------------")
	// 基本类型转string类型
	var num13 int = 99
	var str string
	str = fmt.Sprintf("%d", num13)
	fmt.Printf("str type %T str=%v \n", str, str)
	var num14 int = 100
	str = strconv.FormatInt(int64(num14), 10)
	fmt.Printf("str type %T str=%v \n", str, str)
	fmt.Println("-------------------------------------------------------------------")
	//指针
	var ii int = 10
	fmt.Println("ii的地址是", &ii)
	var ptr *int = &ii
	fmt.Println("ptr=", ptr)
	fmt.Println("ptr的地址是", &ptr)
	fmt.Println("ptr指向的值是", *ptr)
	//引入本地外部包
	fmt.Println(model.HeroName)

	//接受终端输入
	var name string
	var age1 byte
	// fmt.Println("请输入姓名")
	// fmt.Scanln(&name)
	// fmt.Println("请输入年龄")
	// fmt.Scanln(&age1)
	fmt.Printf("名字是%v,年龄是%v\n", name, age1)
	//二进制
	var i2 int = 5
	fmt.Printf("%b\n", i2)

	var j int = 011
	fmt.Println("j=", j)

	var k int = 0x11
	fmt.Println("k=", k)
	//位运算
	fmt.Println(2 & 3)
	// 2 [0000 0010]
	// 3 [0000 0011]
	// 2&3 [0000 0010]
	fmt.Println(2 | 3)
	// 2 [0000 0010]
	// 3 [0000 0011]
	// 2|3 [0000 0011]
	fmt.Println(2 ^ 3)
	// 2 [0000 0010]
	// 3 [0000 0011]
	// 2^3 [0000 0001]
	fmt.Println(-2 ^ 2)
	// -2[1000 0010] [1111 1101] [1111 1110]
	// 2[0000 0010]
	// -2^2[1111 1100] [1111 1011] [1000 0100] 4
	fmt.Println(1 >> 2)
	// 1[0000 0001] [0000 0000]
	fmt.Println(1 << 2)
	// 1[0000 0001] [0000 0100]

	//循环
	str1 := "hello world!北京"
	str2 := []rune(str1)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("index=%d,val=%c\n", i, str2[i])
	}
	for index, val := range str1 {
		fmt.Printf("index=%d,val=%c\n", index, val)
	}
	//随机生成[0,100)的数
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
	//函数
	fmt.Println(cal(1, 2))
}
func cal(n1 float64, n2 float64) float64 {
	return n1 + n2
}
