package test

import "testing"
import "gopkg.in/mgo.v2"


func TestConnDB(t *testing.T) {

	session,err:=mgo.Dial("localhost")
	if err != nil {

		///出现了错误
		return
	}
	anotherSession:=session.Copy()
	defer anotherSession.Close()
	session.DB("game").C("col")

}
