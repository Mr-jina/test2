package main

import (
	"fmt"
	"strconv"
)

func main(){

	//1、strconv.ParseInt(变量,进制,多少位) : string 转成 int基本数据类型
	var str string = "32"
	var i,_ = strconv.ParseInt(str,10,32)
	fmt.Println("i=",i)


	//2、strconv.ParseBool() : string 转成 bool基本数据类型
	var str2 string = "true"
	var bool,_ = strconv.ParseBool(str2)
	fmt.Printf("bool类型=%t,bool=%v",bool,bool)

}