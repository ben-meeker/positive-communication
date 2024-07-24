package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ayush6624/go-chatgpt"
)

func getUserPrompt() (string, error) {
	var prompt string
	fmt.Print("What message would you like to analyze? ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		prompt = scanner.Text()
		fmt.Println(prompt)
		break
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	prompt = ("Please analyze the follow statement, and return it in the following json format. The first field, labelled 'sentiment', that has either a 'positive', 'negative', or 'neutral' value, if it's negative or neutral, a second field labelled 'rephrased_statement', which has a rephrased version of the statement, but with positive language, and a third field labelled 'benefits', which describes the benefits of using the alternate, positive wording. Please make the rephrased statement field an empty string if the sentiment is already positive, and describe the benefits of the given statement in the beneftis section. Here is the statement: " + prompt)

	return prompt, nil
}

func initiateChatGPTConnection() *chatgpt.Client {
	key := os.Getenv("OPENAI_KEY")

	client, err := chatgpt.NewClient(key)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func simpleSend(client *chatgpt.Client, prompt string) (*chatgpt.ChatResponse, error) {
	ctx := context.Background()

	res, err := client.SimpleSend(ctx, prompt)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Send a prompt to ChatGPT and get a JSON response.
func sendJSON(client *chatgpt.Client, prompt string) (*chatgpt.ChatResponse, error) {
	ctx := context.Background()

	message := chatgpt.ChatMessage{
		Role:    "assistant",
		Content: prompt,
	}

	var requestMessages []chatgpt.ChatMessage
	requestMessages = append(requestMessages, message)

	enableJSON := chatgpt.ResponseFormat{
		Type: "json_object",
	}

	request := chatgpt.ChatCompletionRequest{
		Model:    "gpt-3.5-turbo",
		Messages: requestMessages,
		// MaxTokens: 3000,
		Response_Format: &enableJSON,
	}

	log.Println(request)

	res, err := client.Send(ctx, &request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
