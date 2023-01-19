package main

import (
	"fmt"
	"snippets/addition"
	checktype "snippets/if_else"
)

func main(){
	fmt.Println("Hi , hello all :)")
	res := addition.Add(12,45)
	fmt.Println("result is", res)
	animal := checktype.CheckAnimal("maa")
	fmt.Println(animal)
}