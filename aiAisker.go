package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("write a question!!")
		return
	}
	args := os.Args[1:]

	var q string
	for _, i := range args {
		q += i
	}

}

func reqToAI(q string) (string, error) {
	client := http.DefaultClient
	mes := "{\"messages\":[{\"role\": \"user\", \"content\":\"" + q + "\"}]}"

	body := bytes.NewReader([]byte(mes))
	req, err := http.NewRequest("POST", "https://ai.hackclub.com/chat/completions", body)
	req.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}

	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("just sent request:")
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("couldnt send")
		return "", err
	}

	type Response struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	var fin Response
	err = json.Unmarshal(responseBody, &fin)

	if err != nil {
		fmt.Println("trying to unmarshal:")
		fmt.Println(err)
		fmt.Println(string(responseBody))
		fmt.Println("above")
		return "", err
	}
	fmt.Println(string(fin.Choices[0].Message.Content))
	return string(fin.Choices[0].Message.Content), nil
}
