package main

import (
	"fmt"
	"goExProject/appHandle"
	"goExProject/errorHandle"
	"net/http"
)

var cha chan int

func main() {
	http.HandleFunc(appHandle.Prefix,
		errorHandle.ErrorHandler(appHandle.FileHandle))

	err := http.ListenAndServe("8888", nil)
	if err != nil {
		fmt.Println(err)
	}
}
