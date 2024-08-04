package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"theboyscommittee.com/positivecommunication/analyzed_response"
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

	var analyzedResponse analyzed_response.AnalyzedResponse

	content := strings.ReplaceAll(response.Choices[0].Message.Content, "\n", "")

	if err := json.Unmarshal([]byte(content), &analyzedResponse); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Error parsing AI response")
		return
	}

	c.JSON(http.StatusOK, analyzedResponse)
}
