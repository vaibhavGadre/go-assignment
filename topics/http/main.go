package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func main()  {
	res, err := http.Get("https://www.thunderclient.com/welcome")

	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode)
	defer res.Body.Close()
	bytes, _ :=ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))
}