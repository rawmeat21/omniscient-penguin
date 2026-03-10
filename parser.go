package main

import (
	"strings"
	"mvdan.cc/sh/v3/syntax"
)

func parseText(input string) ([]string,error){
	// take a text as input, parse the text through sh package, if it succeeds return a slice of strings, else return error
	
	// the text may be bad, in which case return error
	reader:=strings.NewReader(input)
	parser:=syntax.NewParser()

	file,err:=parser.Parse(reader,"")

	if err!=nil {
		return nil,err
	}

	commands:=[]string{}



	syntax.Walk(file,func (node syntax.Node) bool {
		
		switch n:= node.(type){
		case *syntax.CallExpr:
			if len(n.Args)>0 {
				cmd:=n.Args[0].Lit()
				commands=append(commands, cmd)
			}
		}

		return true
	})

	return commands,nil
}



