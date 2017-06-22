package main

import "fmt"

var userDB map[string]string
func main()  {
	userDB=make(map[string]string)
	fmt.Println(userDB["wang"])
	if userDB["wang"] ==""{
		userDB["wang"]="123"
	}
	fmt.Println(userDB)
}
