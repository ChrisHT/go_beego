package models

import (
	"time"
	"go_beego/models/mongodb"
	// "golang.org/x/crypto/scrypt"
	"gopkg.in/mgo.v2"
)

type User struct {
	ID       string    `bson:"_id"      json:"_id,omitempty"`
	Name     string    `bson:"name"     json:"name,omitempty"`
	Password string    `bson:"password" json:"password,omitempty"`
	Salt     string    `bson:"salt"     json:"salt,omitempty"`
	RegDate  time.Time `bson:"reg_date" json:"reg_date,omitempty"`
}

func (u *User) FindByID(id string) (code int, err error) {
	Conn := mongo.Conn()
	defer Conn.Close()

	c := Conn.DB("").C("users")
	err = c.FindId(id).One(u)

	if err != nil {
		if err == mgo.ErrNotFound {
			code = ErrNotFound
		} else {
			code = 1
		}
	} else {
		code = 0
	}
	return
}

func (u *User) Insert() (code int, err error) {
	Conn := mongo.Conn()
	defer Conn.Close()

	c := Conn.DB("").C("users")
	err = c.Insert(u)

	if err != nil {
		if err == mgo.ErrNotFound {
			code = ErrNotFound
		} else {
			code = 1
		}
	} else {
		code = 0
	}
	return
}