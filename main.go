package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"bufio"
	"log"

	"github.com/MamikonAyvazyan/apitest/internal/config"
	"net/http"
)

func main() {
	if err := config.Conf.Init(); err != nil {
		log.Fatal(err)
	}

	prompt := "check"

	payload := map[string]interface{}{
		"model":      config.Conf.Llmserver.Model,
		"keep_alive": 0,
		"stream":     config.Conf.Llmserver.Stream,
		"prompt":     prompt,
	}

	body, _ := json.Marshal(payload)
	resp, err := http.Post(config.Conf.Llmserver.URL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody := bufio.NewReader(resp.Body)

	respStruct := struct{
		Status int
		Message string
	}{}

	json.Unmarshal(respBody, &respStruct)

	fmt.Println(respStruct.Status)
}
