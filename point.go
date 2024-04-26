package main

import "fmt"

func main() {
	mystring := "hi"
	//取指针
	mypointer := &mystring
	//取值
	mystring2 := *mypointer

	fmt.Println(mystring, mypointer, mystring2)

}
