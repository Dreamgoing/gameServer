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

///@todo io交互模型设计
///@todo go routine 两个线程,一个用来输入,另一个用来读取数据

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

func simulation() {
	conn:=connect("tcp","127.0.0.1:3563")

	defer conn.Close()

	// Hello 消息（JSON 格式）
	// 对应游戏服务器 Hello 消息结构体
	login(conn,"leaf","12345")
	ret:=make([]byte,50)

	for true{
		conn.Read(ret)
		fmt.Printf("%v\n",string(ret))
		var in string
		fmt.Scanf("%s",&in)
		fmt.Println(in)
		if in==string("w"){
			sendUp(conn)
		}else if in==string("q"){

			///退出
			break
		}

	}

}

func main() {
	simulation()

	//go simulation()




}