package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

)

type Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}

type RequestBody struct {
	Model string `json:"model"`
	MaxTokens int `json:"max_tokens"`
	Messages []Message `json:"messages"`
}

type Choice struct {
	Msg Message `json:"message"`
}

type ResponseBody struct {
	Choices []Choice `json:"choices"`
}

func explain (input string,manpages string,level string) (string,error){

	prompt:=fmt.Sprintf(`You are a Linux system admin. A user has given you the follwing input text inside the double quotes: "%s"
		Note that the text may either be a command or something that the user wants to get done on their Linux system.
		If it's neither then reply with only a 'INVALID' only. Do not say things like "This is not related to linux", just say INVALID and quit.
		If not, generate an explaination of the following level- `,input)

	switch level{
	case "h":
		prompt+=`high level
		Explain the input in a simple way, avoid complicated jargon. Be to the point, keep it as simple as possible.

		OR if the user wants to do something:

		List relevant command statemnts that get the job done (List the whole statement). Follow the above rules for explaining each command you type.
		`
	case "m":
		prompt+=`middle level (lower than high level, but higher than low level)
		Explain the input assuming user is an intermediate user. 
		Your explaination should include the following:
		
		If input text looks like a command:

		1. Explain the command by giving a small description
		2. Explain what the command does internally, but don't go too deep. Don't use overly complicated terms, rather explain complicated things
		in a simple way. AVOID USING ANALOGIES
		3. You can list files/ folders that get changed through a command, relevant terms that might be interesting to a learner

		OR if the user wants to do something:

		List relevant commands (state full commands) that get the job done and how exactly they get the job done. Follow the above rules for explaining each command you type.
		However don't describe by points, just explain it normally, like you are a sysadmin
		`
	case "l":
		prompt+=`low level
		Explain the input assuming user is an advanced user. 
		Your explaination should include the following:
		
		If input text looks like a command:

		1. Explain the command by giving a small description
		2. Explain what the command does internally at great depth. Your explaination should include changes being made at the depths of the 
		operating system, what's being modified and why. Explain each part in a sequential manner, don't skip anything
		3. If there uncommon terms in between, explain each of them at the bottom (after the explaination is complete). A brief description
		will suffice here.

		OR if the user wants to do something:

		List relevant commands (state full commands) that get the job done and how exactly they get the job done. Follow the above rules for explaining each command you type.
		However don't describe by points, just explain it normally, like you are a sysadmin
		`
	}

	prompt+=fmt.Sprintf(`Write everything in plain text (no code insertion) Here are relevant man pages that you should use to explain the input text:

	%s

	(Note that man pages may not exist, in which case you are to generate an explaination using your knowledge adhering to the rules above)
	Do not address the user EVER, explain it formally. YOUR OUTPUT MUST CONTAIN ONLY INFORMATION. 
	IT MUST NOT CONTAIN REDUNDANT NON IMPORTANT STUFF LIKE NATURE OF TEXT, WHETHER ITS A COMMAND OR PROMPT, THOSE THINGS ARE NOT IMPORTANT
	Give an example too on how to use the commands`,manpages)


	key:=os.Getenv("OMNIPEN_API_KEY")

	msgs:=[]Message{
		{Role: "system",Content: prompt},
		{Role: "user",Content: input},
	}

	body:=RequestBody{Model: "llama-3.3-70b-versatile",MaxTokens: 3000,Messages: msgs}

	jsonData,_:=json.Marshal(body)

	req,err:=http.NewRequest("POST","https://api.groq.com/openai/v1/chat/completions",bytes.NewBuffer(jsonData))
	
	if(err!=nil){
		return "",err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + key)

	
	client:= &http.Client{}

	resp,err := client.Do(req)
	
	if(err!=nil){
		return "",err
	}

	defer resp.Body.Close()

	data,err:=io.ReadAll(resp.Body)

	var result ResponseBody
	json.Unmarshal(data,&result)


	if len(result.Choices) == 0 {
		return "", fmt.Errorf("no response from API")
	}

	return result.Choices[0].Msg.Content,nil
}