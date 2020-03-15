package main

import "fmt"

import "strconv"

func main(){

	var i string = "100"

	var b int64 

	//要求: string 转成 int基本数据类型  采用“函数”换化
	b,_ = strconv.ParseInt(i,10,64) //strconv.ParseInt(变量,进制,多少位)

	fmt.Println(b)


	//要求: string 转成 float64基本数据类型  采用“函数”换化
	var a string = "100.32"

	var c float64 

	c,_ = strconv.ParseFloat(a,64) //strconv.ParseInt(变量,多少位)

	fmt.Println(c)

	//要求: string 转成 bool基本数据类型  采用“函数”换化
	var m string = "true"

	var y bool

	y,_ = strconv.ParseBool(m) //strconv.ParseBool(变量)

	fmt.Println(y)

	//要求: string 转成 int基本数据类型  
	//采用 strconv.Atoi(h) Atoi是ParseInt(s, 10, 0)的简写。 “函数”换化
	var h string = "10000"

	var e int

	e,_ = strconv.Atoi(h) //strconv.ParseByte(变量)

	fmt.Println(e)

}