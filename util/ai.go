package util

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type AIResponse struct {
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func OpenAISingleRound(content string) string {
	jsonRequestBody, err := json.Marshal(map[string]interface{}{
		"model": GetString("ai.model"),
		"messages": []map[string]string{
			{"role": "system", "content": GetString("ai.setting")},
			{"role": "user", "content": content},
		},
	})
	if err != nil {
		return err.Error()
	}

	req, err := http.NewRequest("POST", GetString("ai.server"), bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		return err.Error()
	}

	req.Header.Set("Authorization", "Bearer "+GetString("ai.secret"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}

	var response AIResponse
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return err.Error()
	}

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content
	} else {
		return "AI没有应答"
	}
}
