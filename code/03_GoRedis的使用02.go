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

	_,err=c.Do("hset" ,"men01","name","shalom")
	if err!=nil{
		panic(err)
	}
	_,err=c.Do("hset" ,"men01","age",19)
	if err!=nil{
		panic(err)
	}

	r,err:=redis.String(c.Do("hget" ,"men01","name"))
	if err!=nil{
		panic(err)
	}
	r01,err:=redis.Int(c.Do("hget" ,"men01","age"))
	if err!=nil{
		panic(err)
	}

	fmt.Println(r)
	fmt.Println(r01)

}
