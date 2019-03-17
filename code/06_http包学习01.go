package main

import (
	"net/http"
		"fmt"
	"io/ioutil"
)

func main() {
	resp,err:=http.Get("http://www.baidu.com")
	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()

	//data,err:=ioutil.ReadAll(resp.Body)
	//if err!=nil{
	//	panic(err)
	//}

	if resp.StatusCode!=http.StatusOK{
		fmt.Println("statu error",resp.StatusCode)
		return
	}
	data,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		panic(err)
	}

	fmt.Println(string(data))

	//data,err:=httputil.DumpResponse(resp,true)
	//if err!=nil{
	//	panic(err)
	//}
	//fmt.Println(string(data))
}
