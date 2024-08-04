package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// a gin router to handle requests
	router := gin.Default()
	// use CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://ben-meeker.github.io"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))
	// insert request handlers
	router.POST("/analyzeMessage", checkAuth, analyzeMessage)
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
