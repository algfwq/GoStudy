package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func myWeb(w http.ResponseWriter, r *http.Request) {
	//交互。。。
	r.ParseForm() //它还将请求主体解析为表单，获得POST Form表单数据，必须先调用这个函数

	fmt.Println("path", r.URL.Path)

	for k, v := range r.URL.Query() {
		fmt.Println("第一个")
		fmt.Println("key:", k, ", value:", v[0])
		if k == "mode" && v[0] == "test" {
			fmt.Println("测试模式")
		}
	}

	for k, v := range r.PostForm {
		fmt.Println("第二个")
		fmt.Println("key:", k, ", value:", v[0])
	}

	//模版
	//t := template.New("index")
	//t.Parse("<div id='templateTextDiv'>Hi,{{.name}},{{.someStr}}</div>")
	t, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		fmt.Println("模版解析错误: ", err)
	}

	data := map[string]string{
		"name":    "zeta",
		"someStr": "这是一个开始",
	}

	t.Execute(w, data)

	//fmt.Fprintf(w, "这是一个开始")
}

func main() {
	http.HandleFunc("/", myWeb)
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./static"))))

	fmt.Println("服务器即将开启，访问地址 http://localhost:8081")

	//err := http.ListenAndServe(":8081", nil)
	//if err != nil {
	//	fmt.Println("服务器开启错误: ", err)
	//}
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("服务器开启错误: ", err)
	}
}
