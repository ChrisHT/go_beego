package redis

import (
    "fmt"
    "github.com/astaxie/goredis"
    "github.com/astaxie/beego"
)

var client goredis.Client

func SetStr(key string, value string) {
	client.Set(key, []byte(value))
}

func GetStr(key string) string {
	val, _ := client.Get(key)
	return string(val)
}

func Del(key string) {
	client.Del(key)
}

func init() {
	url := beego.AppConfig.String("redis::url")
	client.Addr = url
}