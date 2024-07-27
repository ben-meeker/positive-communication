package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// a gin router to handle requests
	router := gin.Default()
	// insert request handlers
	// Listen at http://localhost:8080

	router.POST("/analyzeMessage", checkAuth, analyzeMessage)

	router.Run(":8080")

	//client := initiateChatGPTConnection()

	// prompt, err := getUserPrompt()
	// if err != nil {
	// 	log.Println(err)
	// }

	// response, err := sendJSON(client, prompt)
	// if err != nil {
	// 	log.Println(err)
	// }

	// log.Println(response.Choices[0].Message.Content)

	//log.Println(response)
}
