package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"theboyscommittee.com/positivecommunication/prompt"
)

func analyzeMessage(c *gin.Context) {
	client, err := initiateChatGPTConnection()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, error(err))
		return
	}
	var prompt prompt.Prompt

	if err := c.BindJSON(&prompt); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Please provide a prompt value")
		return
	}

	response, err := sendJSON(client, createSentimentAnalysisPrompt(prompt.Content))
	if err != nil {
		log.Println(err)
		return
	}

	c.IndentedJSON(http.StatusOK, response.Choices[0].Message.Content)
}
