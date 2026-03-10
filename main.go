package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	fmt.Println("Welcome to Omniscient Penguin!")
	fmt.Println("Type a linux command to get an explanation or type what you want")
	fmt.Println("type quit() to exit")
	fmt.Println()


	sc:=bufio.NewScanner(os.Stdin)

	for{
		fmt.Print("omnipen: ")
		sc.Scan()
		input:=sc.Text()

		if input=="quit()"{
			break
		}

		level:="h"

		for{
			fmt.Print("omnipen: How low do you want to go (h, m or l) (default= h) ? ")
			sc.Scan()
			c:=sc.Text()

			proper:=false

			switch c{
			case "h":
				level=c
				proper=true				
			case "l":
				level=c
				proper=true				
			case "m":
				level=c
				proper=true
			case "":
				proper=true
			default:
				fmt.Println("omnipen: Please enter either h, m or l")
			}

			if proper{
				break
			}
		}

		words,err:=parseText(input)
		var manpages strings.Builder

		if(err==nil){
			// input text is a linux command for sure
			fmt.Println("Its a linux command")
			for i:=0;i<len(words);i++{
				manpages.WriteString(fmt.Sprintf("Man page for %s\n\n",words[i]))
				manpages.WriteString("\n\n")
			}
		}

		output,err:=explain(input,manpages.String(),level)

		if(err!=nil) {
			fmt.Println("Sorry, there was a problem. Please try again later.")
			continue
		}

		if output=="INVALID"{
			fmt.Println("omnipen: You seem to have entered gibberish. Don't do that again.")
			continue
		}

		fmt.Println()
		fmt.Print(output)
		fmt.Println()	
	}

}