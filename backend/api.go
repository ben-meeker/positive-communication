package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"theboyscommittee.com/positivecommunication/prompt"
)

func analyzeMessage(c *gin.Context) {
	client := initiateChatGPTConnection()
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

func createSentimentAnalysisPrompt(content string) string {
	parameters := "Please analyze the follow statement, and return it in the following json format. The first field, labelled 'sentiment', that has either a 'positive', 'negative', or 'neutral' value, if it's negative or neutral, a second field labelled 'rephrased_statement', which has a rephrased version of the statement, but with positive language, and a third field labelled 'benefits', which describes the benefits of using the alternate, positive wording. Please make the rephrased statement field an empty string if the sentiment is already positive, and describe the benefits of the given statement in the beneftis section. Make sure to highlight why the rephrased statement is more beneficial than the negative or neutral statement, also be sure to preserve the original intended message. Here is the statement: "

	return parameters + content
}
