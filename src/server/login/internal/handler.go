package internal



import (
	"reflect"
	//"github.com/name5566/leaf/db/mongodb"

	"server/msg"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"

	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)


///@todo 登出时,修改相关的变量,在内存数据库中,可以考虑使用redis缓存


///表示当前是否使用,内存作为数据库
var useMemoryDB bool = true
///全局的UserID,表示当前在线的用户
var UserID int
///内存数据库,简单的map实现
var Userdb map[string]string


const URL  = "localhost"

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {


	if !useMemoryDB {
		handleMsg(&msg.SignUp{},handleSignUpDB)
		handleMsg(&msg.SignIn{},handleSignInDB)
	}else {
		handleMsg(&msg.SignIn{},handleSignInMem)
		handleMsg(&msg.SignUp{},handleSignUpMem)
	}


}

func handleSignInMem(args []interface{})  {
	m:=args[0].(*msg.SignIn)
	a:=args[1].(gate.Agent)

	if Userdb[m.Name] == "" {
		///不存在用户
		a.WriteMsg(&msg.State{"No exist user"})
	}else {
		if Userdb[m.Name]!=m.Password{
			///用户名密码错误
			a.WriteMsg(&msg.State{"password error"})
		}else {
			a.SetUserData(&msg.Car{CarID:UserID})
			UserID++
			a.WriteMsg(&msg.State{"sign in successfully,carID:"+string(UserID-1)})
		}
	}

}

func handleSignUpMem(args []interface{})  {


	///Comma-ok断言

	m:=args[0].(*msg.SignUp)


	a:=args[1].(gate.Agent)

	if Userdb==nil {
		Userdb=make(map[string]string)
	}
	if Userdb[m.Name]==""{
		Userdb[m.Name]=m.Password
		a.WriteMsg(&msg.State{"SignUp successfully"})
	}else {
		a.WriteMsg(&msg.State{"this user has exist in DB"})
	}
}

///@todo 考虑,具体的登录流程,


func handleSignUpDB(args []interface{}) {
	m:=args[0].(*msg.SignUp)



	///获取消息的发送者
	a:=args[1].(gate.Agent)

	log.Debug("login module: sign up%v %v",m.Name,a)

	session,err:=mgo.Dial("localhost")
	defer session.Close()
	if err!=nil{
		panic(err)
	}
	defer session.Close()

	c:=session.DB("mydb").C("try")


	//err =c.Find(bson.M{})

	///@todo 注册新用户时,可以添加加密模块和查看是否重复注册模块
	err = c.Insert(m)

	if err!=nil{
		log.Error("insert mydb fail")
		//a.WriteMsg("signUp error")
	}


	///注册成功,向客户端返回一个消息
	a.WriteMsg(&msg.SignUp{
		Name:"client",
	})


}

func handleSignInDB(args []interface{})  {
	m:=args[0].(*msg.SignIn)

	///客户端地址
	a:=args[1].(gate.Agent)

	log.Debug("login module: sign in%v %v",m.Name,a)

	session,err:=mgo.Dial(URL)
	defer session.Close()
	if err!=nil{
		panic(err)
	}
	defer session.Close()
	c:=session.DB("mydb").C("try")

	var tmpUsr map[string]interface{}


	err=c.Find(bson.M{"name":m.Name}).One(&tmpUsr)

	log.Debug("Sign in %v %v",tmpUsr["name"],tmpUsr["password"])
	log.Debug("%v %v",a.RemoteAddr().String(),a.LocalAddr().String())

	if err!=nil {
		log.Error("sign in module find fail")
		panic(err)
	}

	if tmpUsr["name"]==nil {
		///空值没找到,向客户端发送一个该用户不存在的消息
		log.Debug("sign in module user no found")

	}else{
		//log.Debug("%v %v",tmpUsr[m.Name],m.Password)
		log.Debug("%s %v",tmpUsr["password"].(string),tmpUsr["password"].(string)==m.Password)
		if tmpUsr["password"].(string)!=m.Password{
			///密码错误
			log.Debug("sign in module user name and password mismatch")
		}else {

			///向客户端发送一个成功连接的token
			///并在内存中创建一个当前用户

			if a.UserData()!=nil{
				a.WriteMsg(&msg.State{
					Name:"duplicated user",
				})
				log.Debug("sign in module duplicate signIn")
			}


			///@todo 初始化车的时候,设置相应的车的编号
			a.SetUserData(&msg.Car{CarID:UserID})
			UserID++



			//log.Debug("%v",reflect.TypeOf(a.UserData()))


			a.WriteMsg(&msg.State{
				Name:"Login Success carID:"+string(UserID-1),
			})

		}

	}

///hello
}
