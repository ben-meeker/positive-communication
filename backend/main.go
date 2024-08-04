package main

import "github.com/gin-gonic/gin"

func main() {
	// a gin router to handle requests
	router := gin.Default()
	// insert request handlers
	router.POST("/analyzeMessage", cors, checkAuth, analyzeMessage)
	// listen with cert for comdotcomdotcomdotcomdotcomdotcom.com
	router.RunTLS(":443", "fullchain.pem", "privkey.pem")

	// client := initiateChatGPTConnection()

	// prompt, err := getUserPrompt()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// response, err := sendJSON(client, prompt)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// log.Println(response.Choices[0].Message.Content)
}
