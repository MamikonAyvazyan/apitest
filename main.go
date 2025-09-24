package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"io"
	"net/http"

	"github.com/MamikonAyvazyan/apitest/internal/config"
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
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)

	respStruct := struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{}

	err = json.Unmarshal(respBody, &respStruct)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", respStruct)

}
