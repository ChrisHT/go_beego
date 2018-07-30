package userModel

import (
	"go_beego/models/mongodb"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type User struct {
	//ID       string    `bson:"_id"      json:"_id,omitempty"`
	Name     string    `bson:"name"     json:"name,omitempty"`
	Password string    `bson:"password" json:"password,omitempty"`
	//Salt     string    `bson:"salt"     json:"salt,omitempty"`
	//Date     time.Time `bson:"date"     json:"date,omitempty"`
}

func (u *User) FindByName(name string) (code int, err error) {
	Conn := mongo.Conn()
	defer Conn.Close()

	c := Conn.DB("test").C("user")
	err = c.Find(bson.M{"name": name}).One(&u)
	fmt.Println(err, code)

	if err != nil {
		code = 1
	} else {
		code = 0
	}
	return
}
