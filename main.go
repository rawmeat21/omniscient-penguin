package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	fmt.Println("Welcome to Omniscient Penguin!")
	fmt.Println("Type a linux command to get an explanation or type what you want")
	fmt.Println("type quit() to exit")
	fmt.Println()

	res,err:=parseText("")
	fmt.Println(res)
    fmt.Println(err)


	sc:=bufio.NewScanner(os.Stdin)

	for{
		fmt.Print("omnipen> ")
		sc.Scan()
		command:=sc.Text()

		if command=="quit()"{
			break
		}


	}

}