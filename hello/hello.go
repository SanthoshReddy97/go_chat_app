package main

import (
	"fmt"

	"github.com/santhoshreddy97/go_chat_app/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
