package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))

	client := initiateChatGPTConnection()

	prompt, err := getUserPrompt()
	if err != nil {
		log.Println(err)
	}

	response, err := sendJSON(client, prompt)
	if err != nil {
		log.Println(err)
	}

	log.Println(response.Choices[0].Message.Content)
	//log.Println(response)
}
