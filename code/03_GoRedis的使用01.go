package main

import (

	"github.com/gomodule/redigo/redis"
	"fmt"
)

func main() {
	c,err:=redis.Dial("tcp","localhost:6379")
	if err!=nil{
		panic(err)
	}
	defer c.Close()

	_,err=c.Do("set" ,"men","shalom")
	if err!=nil{
		panic(err)
	}

	r,err:=redis.String(c.Do("get" ,"men"))
	if err!=nil{
		panic(err)
	}

	fmt.Println(r)
}
