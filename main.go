package main

import (
	"fmt"
	"runtime"
	"strings"
)

func init() {
	var goos string = runtime.GOOS
	fmt.Printf("操作系统: %s\n", goos)
}

func main() {
	//var name string
	//var symbol string
	//var age int
	//
	//fmt.Print("请输入姓名：")
	//fmt.Scanln(&name)
	//
	//fmt.Print("请输入符号：")
	//fmt.Scanln(&symbol)
	//
	//fmt.Print("请输入年龄：")
	//fmt.Scanln(&age)
	//
	//fmt.Println("姓名："+name+symbol, "年龄："+fmt.Sprint(age))

	////生成随机数
	//fmt.Println("随机数：", rand.Intn(100))

	//输出字符串禁用转义符
	//escapedStr := "哈哈哈！\n嗨嗨嗨！"
	//escapedStr2 := `哈哈哈！\n嗨嗨嗨！`
	//fmt.Println(escapedStr)
	//fmt.Println(escapedStr2)
	//fmt.Printf("%q\n", escapedStr)
	//escapedStr3 := replaceEscapedChars(escapedStr)
	//fmt.Println(escapedStr3)

	//输出时间
	//t := time.Now()
	//fmt.Println(t.Format("02 January 2006 15:04:12"))

	//IF判断语句
	//if1 := true
	//if2 := false
	//if3 := 12
	//if4 := "hello"
	//
	//if if2 {
	//	fmt.Println("if2 is true")
	//} else if if1 {
	//	fmt.Println("if1 is true")
	//	fmt.Println("if2 is false")
	//}
	//
	//if if3 > 10 {
	//	fmt.Println("if3 > 10")
	//}
	//if len(if4) > 5 {
	//	fmt.Println("if4 > 5")
	//} else {
	//	fmt.Println("if4 < 5")
	//}
	//
	//if if4 == "hello" {
	//	fmt.Println("if4 == hello")
	//}
	//
	//if if4 != "hello" {
	//	fmt.Println("if4 != hello")
	//} else {
	//	fmt.Println("if4 == hello")
	//}
	//ifResult := "hello"
	//ifResult2 := "12"
	//ifResult3 := "14"
	//
	//if ifResult == "hello" && ifResult2 == "12" {
	//	fmt.Println("T")
	//} else {
	//	fmt.Println("F")
	//}
	//
	//if ifResult == "hello" || ifResult2 == "13" {
	//	fmt.Println("T")
	//} else {
	//	fmt.Println("F")
	//}
	//
	//if ifResult == "hello" && ifResult2 == "13" || ifResult3 == "14" {
	//	fmt.Println("T")
	//} else {
	//	fmt.Println("F")
	//}

	//switch语句
	//fmt.Println(alg(10, 0.5))
	//switch result := alg(1, 0.5); {
	//case result >= 10.5:
	//	fmt.Println("T")
	//default:
	//	fmt.Println("default")
	//}
	//mode := "T"
	//switch mode {
	//case "T":
	//	fmt.Println("T")
	//case "F":
	//	fmt.Println("F")
	//default:
	//	fmt.Println("default")
	//}

	//错误处理
	// 尝试打开一个不存在的文件
	//_, err := os.Open("non-existing.txt")
	//if err != nil {
	//	// 错误处理
	//	fmt.Println("报错！")
	//	fmt.Println(err)
	//}
	//result, err := doSomething("阿巴阿巴", 0, "F")
	//if err != nil {
	//	fmt.Println(err)
	//	fmt.Println(result)
	//} else {
	//	fmt.Println(err)
	//	fmt.Println(result)
	//}
	//fmt.Println("继续执行")

	//for循环

	//for i := 0; ; i++ {
	//	fmt.Println("Value of i is now:", i)
	//} //无限循环

	//for i := 0; i < 3; {
	//	fmt.Println("Value of i:", i)
	//} //无限循环

	//s := ""
	//for s != "aaaaa" {
	//	fmt.Println("Value of s:", s)
	//	s = s + "a"
	//} //循环5次

	//for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j, s = i+1, j+1, s+"a" {
	//	fmt.Println("Value of i, j, s:", i, j, s)
	//} //循环3次

	//for {
	//	fmt.Println("奥利给！")
	//} //无限循环

	//for i := 1; ; i++ {
	//	fmt.Println(i)
	//	if i == 5 {
	//		break //跳出循环
	//	} else if i == 4 {
	//		continue
	//	}
	//	fmt.Println(string(i) + "第二次")
	//}

	//标签
	//i := 1
	//GOTO:
	//	fmt.Println(i)
	//	i = i + 1
	//	goto GOTO
}

// 自定义错误
type MyError struct {
	Msg  string
	Code int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("问题代码: %d, 问题信息: %s", e.Code, e.Msg)
}

// 使用自定义错误
func doSomething(message string, code int, mode string) (string, error) {
	// 模拟一个错误情况
	if mode == "F" {
		return "操作失败！", &MyError{message, code}
	} else {
		// 没有错误发生，返回具体的字符串和nil
		return "操作成功！", nil
	}
}

// 禁用转义符
func replaceEscapedChars(s string) string {
	// Replace the newline escape sequence with its literal representation
	s = strings.Replace(s, "\n", "\\n", -1)
	return s
}

func alg(number int, floatNumber float64) float64 {
	result := float64(number) + floatNumber
	return result
}
