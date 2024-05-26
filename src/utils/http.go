package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HandleHttpRequest[T any](req *http.Request, responseBody *T) (*T, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while sending request: %w", err)
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("error: status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading response body: %w", err)
	}

	err = json.Unmarshal(body, responseBody)
	if err != nil {
		return nil, fmt.Errorf("error while parsing response body: %w", err)
	}
	return responseBody, nil
}
