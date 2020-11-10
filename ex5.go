package main

import(
        "fmt"
)

type num int
var y int
var x num

func main(){
        fmt.Println(x)
        fmt.Printf("%T\n",x)
        x=42
        fmt.Println(x)

	y = int(x)
	fmt.Println(y)
        fmt.Printf("%T\n",y)

}

