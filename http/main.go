package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// sb := string(body)
	// log.Println(sb)

	var result map[string]any
	json.Unmarshal(body, &result)

	address := result["address"].(map[string]any)
	fmt.Println(address["city"])

}
