package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Book struct {
	Title string
	Detail string
}
//type Person struct {
//	id int `json:"id"`
//	name string `json:"name"`
//	age uint `json:"age"`
//}

const URL = "localhost"
func main() {
	session, err := mgo.Dial(URL) //连接服务器
	if err != nil {
		panic(err)
	}

	c := session.DB("mydb").C("account") //选择ChatRoom库的account表

	c.Insert(map[string]interface{}{"id": 7, "name": "tongjh", "age": 25}) //增

	//objid := bson.ObjectIdHex("55b97a2e16bc6197ad9cad59")

	//c.RemoveId(objid) //删除

	//c.UpdateId(objid, map[string]interface{}{"id": 8, "name": "aaaaa", "age": 30}) //改

	//var one map[string]interface{}
	//c.FindId(objid).One(&one) //查询符合条件的一行数据
	//fmt.Println(one)

	var result []map[string]interface{}
	c.Find(nil).All(&result) //查询全部
	var man Person
	c.Find(bson.M{"age":25}).One(&man)
	fmt.Println(result)
	fmt.Println(man)
}