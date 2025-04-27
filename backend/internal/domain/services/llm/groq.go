package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GroqClient struct {
	apiKey string
}

func NewGroqClient(apiKey string) *GroqClient {
	return &GroqClient{apiKey: apiKey}
}

type GroqRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

type GroqResponse struct {
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (client *GroqClient) Chat(model string, prompt string) (string, error) {
	reqBody := GroqRequest{
		Model: model,
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{"user", prompt},
		},
	}

	url := "https://api.groq.com/openai/v1/chat/completions"
	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+client.apiKey)

	clientHTTP := &http.Client{}
	resp, err := clientHTTP.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request failed with status: %s", resp.Status)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var groqResponse GroqResponse
	if err := json.Unmarshal(respBody, &groqResponse); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if len(groqResponse.Choices) > 0 {
		return groqResponse.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response choices found")
}