package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func Chat(model string, prompt string) string {
	payload := map[string]any{
		"model": model,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
		"stream": false,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	// Create the request
	req, err := http.NewRequest("POST", "http://localhost:11434/api/chat", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Unmarshal into a map
	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err)
	}

	return result["message"].(map[string]interface{})["content"].(string)
}
