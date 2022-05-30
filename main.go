package main
import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
	"time"
)

func main() {
	var webhook string
	var message string

	fmt.Print("Webhook: ")
	fmt.Scan(&webhook)

	fmt.Print("Message: ")
	fmt.Scan(&message)
	
	// Actual thing
	message_json := map[string]interface{}{
		"content": message,
	}
	toBytes, err := json.Marshal(message_json)
	if err != nil {
		panic(err)
	}
	for {
		site, err := http.Post(webhook, "application/json", bytes.NewBuffer(toBytes))
		if err != nil {
			panic(err)
		}
		// Doing this thing so golang doesn't cry about the site variable not being used
		var result map[string]interface{}
		json.NewDecoder(site.Body).Decode(&result)
		time.Sleep(1 * time.Millisecond)
	}
}