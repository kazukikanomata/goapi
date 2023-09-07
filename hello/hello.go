package hello

import (
	"log"

	"local.package/greetings"
)

func Hello(msg string) string {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello(msg)
	if err != nil {
		log.Fatal(err)
	}
	return message
}
