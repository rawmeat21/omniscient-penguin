package main

import (
	// "io"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func manPageGen(input string) string{
	// given a bash command as input, return the man page

	// if man page is not found, return an empty string (=> use the LLM to generate some random bullshit)

	// lets try to find the man page on the system
	cmd:=exec.Command("man",input)
	cmd.Env = append(os.Environ(), "MANPAGER=cat", "MAN_KEEP_FORMATTING=0")
	output,err:=cmd.Output()

	if err!=nil{
		// no man found, look for man pages on man7.org
		
		for i:=1;i<=8;i++{
			url:=fmt.Sprintf("https://man7.org/linux/man-pages/man%d/%s.%d.html", i,input,i)
			res,err:=http.Get(url)

			if err==nil && res.StatusCode==200 {
				doc,err:=goquery.NewDocumentFromReader(res.Body)
				if err==nil{
					var content strings.Builder
					doc.Find("h2,pre").Each(func(i int,s *goquery.Selection){
						content.WriteString(s.Text())
						content.WriteString("\n")
					})

					return content.String()
				} 
			}
		}
	}

	return string(output)
}

