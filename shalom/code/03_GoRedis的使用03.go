package main

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)

var pool *redis.Pool

func init() {

}

func main() {
	conn:=pool.Get()
	defer conn.Close()

	_,err:=conn.Do("Set","name","逗比")
	if err!=nil{
		panic(err)
	}

	r,err:=redis.String(conn.Do("get","name" ))
	fmt.Println(r)


}
