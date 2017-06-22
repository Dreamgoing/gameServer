package main

import (
	"encoding/binary"
	"net"
	"fmt"
	"encoding/json"
)

type User struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

type SignInMsg struct {
	User `json:"SignIn"`
}

type SignUpMsg struct {
	User `json:"SignUp"`
}


/**
发送数据包以这种数据包发送
--------------
| len | data |
--------------
*/
func sendPackage(conn net.Conn, jsonData []byte) bool {

	m:=make([]byte,len(jsonData)+2)

	binary.BigEndian.PutUint16(m,uint16(len(jsonData)))

	copy(m[2:],jsonData)

	conn.Write(m)

	return true

}

func sendUp(conn net.Conn)  {
	up:=[]byte(`{
			"Up": {
				"Direction": 0
			}
		}`)
	m:=make([]byte,2+len(up))
	binary.BigEndian.PutUint16(m,uint16(len(up)))
	copy(m[2:],up)
	conn.Write(m)
}
func sendLeft(conn net.Conn)  {
	left:=[]byte(`{
			"Left":{
				"Direction": 0
			}
		}`)
	m:=make([]byte,2+len(left))
	binary.BigEndian.PutUint16(m,uint16(len(left)))
	copy(m[2:],left)
	conn.Write(m)
}

func sendRight(conn net.Conn)  {
	right:=[]byte(`{
			"Right":{
				"Direction": 0
			}
		}`)
	m:=make([]byte,2+len(right))
	binary.BigEndian.PutUint16(m,uint16(len(right)))
	copy(m[2:],right)
	conn.Write(m)

}
func connect(network string,address string) net.Conn {
	conn,err:=net.Dial(network,address)
	if err!=nil {
		panic(err)
	}
	return conn

}
func login(conn net.Conn,name, password string) bool {

	user:=&SignInMsg{User{Name:name,Password:password}}

	userdata,err:=json.Marshal(user)
	if err!=nil {
		panic(err)
	}

	fmt.Println(string(userdata))


	return sendPackage(conn,userdata)

}

func signUp(conn net.Conn,name,password string)bool  {
	user:=&SignUpMsg{User{Name:name,Password:password}}

	userdata,err:=json.Marshal(user)
	if err!=nil {
		panic(err)
	}
	fmt.Println(string(userdata))

	return sendPackage(conn,userdata)
}


func simulation() {
	conn:=connect("tcp","127.0.0.1:3389")

	defer conn.Close()

	// Hello 消息（JSON 格式）
	// 对应游戏服务器 Hello 消息结构体


	ret:=make([]byte,80)

	c1:=make(chan []byte)
	c2:=make(chan string)

	///结束之后关闭channel
	defer func() {
		close(c1)
		close(c2)
	}()
	go func() {
		for true{
			conn.Read(ret)
			c1<-ret
		}
	}()

	go func() {
		for true{
			var in string
			fmt.Scanf("%s",&in)
			c2<-in
		}
	}()
	var name string
	var password string

	fmt.Println("welcom client! press 'h' to get help")
	for true{
		select {
		case data:=<-c1:
			fmt.Printf("%v\n",string(data))
		case op:=<-c2:
			if op==string("w"){
				sendUp(conn)
			}else if op==string("a") {
				sendLeft(conn)
			}else if op==string("d") {
				sendRight(conn)
			}else if op==string("q"){
				///退出
				fmt.Println("client quit")
				return
			}else if op==string("j"){
				fmt.Println("Sign Up:")
				fmt.Print("name: ")
				fmt.Scanf("%s",&name)
				fmt.Print("password: ")
				fmt.Scanf("%s",&password)
				signUp(conn,name,password)
			}else if op==string("l"){
				fmt.Println("Sign In:")
				fmt.Print("name: ")
				fmt.Scanf("%s",&name)
				fmt.Print("password: ")
				fmt.Scanf("%s",&password)
				login(conn,name,password)
			}else if op==string("h"){
				fmt.Println("command:")
				fmt.Println("w a s d is direction command")
				fmt.Println("j is sign up command")
				fmt.Println("l is login command")
				fmt.Println("o is administer login")
				fmt.Println("q is quit command")

			}else if op == string("o") {
				signUp(conn,"wang","123")
				login(conn,"wang","123")

			}

		}
	}

}

func main() {
	simulation()
	//go simulation()

}