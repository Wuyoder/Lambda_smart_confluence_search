package openai

import (
	"context"
	"errors"
	"fmt"
	"os"
	"smart_confluence_search/lib"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
)

type CompletionObj struct {
	Payload    string
	Query      string
	IsChatMode bool
	IsGPT4     bool
}

func (s *CompletionObj) GetOpenAIResp(ctx context.Context) ([]string, int32) {
	var result []string
	var tokenUsage int32
	var deployID string
	if s.IsGPT4 {
		deployID = os.Getenv("MODEL4_DEPLOYMENT_ID")
	} else {
		deployID = os.Getenv("MODEL3_DEPLOYMENT_ID")
	}

	if s.IsChatMode {
		fmt.Println("start to get chat completion")
		resp, err := GetChatCompletions(ctx, s.Payload, s.Query, deployID)
		lib.HandleErr(err)
		for i := range resp.Choices {
			result = append(result, *resp.Choices[i].Message.Content)
		}
		tokenUsage = *resp.Usage.TotalTokens
	} else {
		resp := GetCompletion(ctx, s.Payload, s.Query, deployID)
		for i := range resp.Choices {
			result = append(result, *resp.Choices[i].Text)
		}
		tokenUsage = *resp.Usage.TotalTokens
	}

	return result, tokenUsage
}

func GetChatCompletions(ctx context.Context, payload string, query string, deployID string) (azopenai.GetChatCompletionsResponse, error) {
	prompt := QueryWithPage(payload, query)
	resp, err := getClient().GetChatCompletions(ctx, azopenai.ChatCompletionsOptions{
		Messages: []azopenai.ChatMessage{
			{Content: &prompt, Role: to.Ptr(azopenai.ChatRoleAssistant)},
		},
		MaxTokens:   to.Ptr(int32(2048)),
		Temperature: to.Ptr(float32(0.0)),
		Deployment:  deployID,
		N:           to.Ptr(int32(1)),
	}, nil)
	lib.HandleErr(err)

	return resp, err
}

func GetCompletion(ctx context.Context, payload string, query string, deployID string) azopenai.GetCompletionsResponse {
	prompt := QueryWithPage(payload, query)
	resp, err := getClient().GetCompletions(ctx, azopenai.CompletionsOptions{
		Prompt:      []string{prompt},
		MaxTokens:   to.Ptr(int32(2048)),
		Temperature: to.Ptr(float32(0.0)),
		Deployment:  deployID,
		N:           to.Ptr(int32(1)),
	}, nil)
	lib.HandleErr(err)

	return resp
}

func QueryWithPage(payload string, query string) string {
	if payload == "" {
		return query
	}

	return query + payload
}

func getClient() *azopenai.Client {
	azureOpenAIKey := os.Getenv("AZURE_OPENAI_KEY")
	azureOpenAIEndpoint := os.Getenv("AZURE_OPENAI_ENDPOINT")

	if azureOpenAIKey == "" || azureOpenAIEndpoint == "" {
		lib.HandleErr(errors.New("environment variables missing"))
	}

	keyCredential, err := azopenai.NewKeyCredential(azureOpenAIKey)
	lib.HandleErr(err)

	client, err := azopenai.NewClientWithKeyCredential(azureOpenAIEndpoint, keyCredential, nil)
	lib.HandleErr(err)

	return client
}
