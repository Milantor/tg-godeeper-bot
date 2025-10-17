package ai

import "fmt"

func Init(api_key string) {
	if api_key == "" {
		fmt.Println("Error: API key is empty.")
		return
	}
	fmt.Println("AI chat init succesfully!")
}
