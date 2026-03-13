package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()
	fmt.Println(ColourPrompt+"Welcome to Omniscient Penguin!"+ColourReset)
	fmt.Println("Type a linux command to get an explanation or type what you want")
	fmt.Println("Type quit() to exit")
	fmt.Println()


	sc:=bufio.NewScanner(os.Stdin)

	promptText:=ColourPrompt+"omnipen: "+ColourReset
	for{
		fmt.Print(promptText)
		sc.Scan()
		input:=sc.Text()

		if input=="quit()"{
			break
		}

		if input=="clear()"{
			cmd:=exec.Command("clear")
			cmd.Stdout=os.Stdout	
			cmd.Run()
			continue
		}

		level:="h"

		for{
			fmt.Print(promptText)
			fmt.Print("How low do you want to go (h, m or l) (default= h) ? ")
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
				fmt.Print(promptText)
				fmt.Println(ColourWarning+"Please enter either h, m or l"+ColourReset)
			}

			if proper{
				break
			}
		}

		words,err:=parseText(input)
		var manpages strings.Builder

		if(err==nil){
			// input text is a linux command for sure
			for i:=0;i<len(words);i++{
				manpages.WriteString(fmt.Sprintf("Man page for %s\n\n",words[i]))
				manpages.WriteString("\n\n")
			}
		}

		output,err:=explain(input,manpages.String(),level)

		if(err!=nil) {
			fmt.Print(promptText)
			fmt.Println(ColourError+"Sorry, there was a problem. Please try again later."+ColourReset)
			continue
		}

		if output=="INVALID"{
			fmt.Print(promptText)
			fmt.Println(ColourWarning+"You seem to have entered gibberish. Don't do that again."+ColourReset)
			continue
		}

		fmt.Println()
		fmt.Println()
		fmt.Print("\033[93m"+output+ColourReset)
		fmt.Println()
		fmt.Println()	
	}

}