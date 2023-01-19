package main

import (
	"fmt"
	"snippets/addition"
)

func main(){
	fmt.Println("Hi , hello all :)")
	res := addition.Add(12,45)
	fmt.Println("result is", res)
}