package main

import (
	"net/http"
		"fmt"
	"net/http/httputil"
)

//	返回手机端页面
func main() {
	request,err:=http.NewRequest(http.MethodGet,"http://imooc.com",nil)
	if err!=nil{
		panic(err)
	}

	request.Header.Add("User-Agent",
		" Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	//resp,err:=http.DefaultClient.Do(request)
	client:=http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req)
			return nil
		},
	}
	resp,err:=client.Do(request)
	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()

	data,err:=httputil.DumpResponse(resp,true)
	if err!=nil{
		panic(err)
	}
	fmt.Println(string(data))
}
