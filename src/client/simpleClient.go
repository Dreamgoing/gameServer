package main

import (
	"encoding/binary"
	"net"
	"fmt"
	"encoding/json"
)

///@todo 考虑如何设置返回值为err
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

///一个用户向另一个用户转发的消息
type UserMsg struct {
	MsgPack InterMsg `json:"UserMsg"`
}

type InterMsg struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
	Context string `json:"context"`
}

type MatchMsg struct {
	Match `json:"Match"`
} 
type Match struct {
	Name string `json:"name"`
	Car int `json:"car"`
}

type OrderMsg struct {
	Order `json:"Order"`
}
type Order struct {
	Name string `json:"name"`
	Val int `json:"val"`
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

func sendAdmin(conn net.Conn)  {
	adm:=[]byte(`{
			"Admin":{
				"Name":"Wang"
			}
		}`)
	m:=make([]byte,2+len(adm))
	binary.BigEndian.PutUint16(m,uint16(len(adm)))
	copy(m[2:],adm)
	conn.Write(m)
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

	//fmt.Println(string(userdata))


	return sendPackage(conn,userdata)

}

func match(conn net.Conn,name string,car int) bool {
	
	userMatch:=&MatchMsg{Match{name,car}}


	data,err:=json.Marshal(userMatch)

	//fmt.Println(data)
	if err!=nil {
		panic(err)
	}
	return sendPackage(conn,data)
	
}

func order(conn net.Conn, name string, val int) bool {
	userOrder:=&OrderMsg{Order{name,val}}

	data,err:=json.Marshal(userOrder)

	//fmt.Println(data)
	if err != nil {
		panic(err)
	}
	return sendPackage(conn,data)
}

///向一个特定的用户发送,私信暂未支持加密功能
func sendMsg(conn net.Conn,src string,dst string,context string) bool {
	tmpMsg:=&UserMsg{InterMsg{src,dst,context}}

	tmpdata,err:=json.Marshal(tmpMsg)
	if err!=nil {
		panic(err)
	}
	//fmt.Println(string(tmpdata))

	return sendPackage(conn,tmpdata)
}

func signUp(conn net.Conn,name,password string)bool  {
	user:=&SignUpMsg{User{Name:name,Password:password}}

	userdata,err:=json.Marshal(user)
	if err!=nil {
		panic(err)
	}
	//fmt.Println(string(userdata))

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
			fmt.Printf("Read from Server: %v \n",string(data))
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
				fmt.Println("o is adminUser login")
				fmt.Println("q is quit command")
				fmt.Println("m is show all online user")
				fmt.Println("b is match mode")
				fmt.Println("v is order")
			}else if op == string("o") {

				signUp(conn,"wang","123")
				login(conn,"wang","123")

			}else if op==string("m"){
				sendAdmin(conn)
			}else if op==string("n"){
				fmt.Println("send msg to an existed user")
			}else if op == string("b") {

				match(conn,"wang",1)
			}else if op==string("v"){
				order(conn,"wang",12)
			}
		}
	}

}
func main() {
	simulation()
	//go simulation()

}