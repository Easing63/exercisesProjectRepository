package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//用来存放分割后的字符串
var m  = make(map[int]string)
var s []string

func main() {
	reader := bufio.NewReader(os.Stdin)
	line,_,err := reader.ReadLine()
	if err != nil {
		fmt.Println(err.Error())
	}

	//line是一个byte数组

	/*for _,str := range line {
		fmt.Print(str)
	}*/
	addStringToCollection(string(line),",")

	//去查询出缓存中的所有字符串
	selectAll(m)

	//模糊查询
	findStrFromMap("go")
}

//将输入的字符串缓存在集合中
func addStringToCollection(str string,sep string) {
	//通过分隔符获取字符串
	strSlice := strings.Split(str,sep)
	fmt.Println(strSlice)
	//将字符串切片插入map中
	for i,item := range strSlice {
		//map[i] = item
		m[i] = item
	}
	fmt.Println("添加了元素到集合中后的结果")
	fmt.Println(m)
}

//可以根据字符串模糊匹配出缓存的字符串
func findStrFromMap(findStr string) []string{
	//判断是否找到
	var flag bool = false
	//循环map
	for _,item := range m {
		if strings.Contains(findStr,item){
			flag = true
			s = append(s,item)
			fmt.Print("满足模糊查询额是：" + item)
		}
	}
	if flag == false {
		fmt.Println("没有查询到")
	}
	return s
}


//直接查询出所有缓存的字符串
func selectAll(m map[int]string) {
	fmt.Println("查询出map集合中的所有字符串：")
	for _,item := range m{
		fmt.Println(item)
	}
}