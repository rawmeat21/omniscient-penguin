package main

import (
	"os"
	"os/exec"
	// "net"
	// "net/http"
)

func manPageGen(input string) string{
	// given a slice of bash commands as input, return a list of related man pages

	cmd:=exec.Command("man",input)
	cmd.Env = append(os.Environ(), "MANPAGER=cat", "MAN_KEEP_FORMATTING=0")
	output,err:=cmd.Output()

	if err!=nil {

	}

	return string(output)
}

