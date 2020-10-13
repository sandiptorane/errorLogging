package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	file,err := os.OpenFile("abc.text",os.O_CREATE,0666)
	if err!=nil{
		log.Fatal(err)
		return
	}
	fmt.Println(file.Name(),"opened successfully")

}
