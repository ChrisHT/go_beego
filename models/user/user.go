package userModel

import (
	"go_beego/models/mongodb"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"io"
	"golang.org/x/crypto/scrypt"
)

type User struct {
	//ID       string    `bson:"_id"      json:"_id,omitempty"`
	Name     string    `bson:"name"     json:"name,omitempty"`
	Password string    `bson:"password" json:"password,omitempty"`
	Salt     string    `bson:"salt"     json:"salt,omitempty"`
	Date     time.Time `bson:"date"     json:"date,omitempty"`
}

const pwHashBytes = 64

func generateSalt() (salt string, err error) {
	buf := make([]byte, pwHashBytes)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", buf), nil
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

func NewUser(u *User, t time.Time) (u *User, err error) {
	salt, err := generateSalt()
	if err != nil {
		return nil, err
	}
	hash, err := generatePassHash(u.Password, salt)
	if err != nil {
		return nil, err
	}

	user := User{
		Name:     u.Name,
		Password: hash,
		Salt:     salt,
		Date:  t}

	return &user, nil
}

func (u *User) InSert() (code int, err error) {
	mConn := mymongo.Conn()
	defer mConn.Close()

	c := mConn.DB("test").C("user")
	err = c.Insert(u)

	if err != nil {
		code = 1
	} else {
		code = 0
	}
	return
}

func generatePassHash(password string, salt string) (hash string, err error) {
	h, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, pwHashBytes)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h), nil
}

func (u *User) CheckPass(password string) (ok bool, err error) {
	hash, err := generatePassHash(password, u.Salt)
	if err != nil {
		return false, err
	}
	return u.Password == hash, nil
}