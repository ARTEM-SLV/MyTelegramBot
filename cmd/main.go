package main

import (
	"flag"
	"fmt"
	"log"
)

const tgBotHost = "api.telegram.org"

func main() {
	t := mustToken()
	tgClient := telegram.New(tgBotHost, t)

	fmt.Println(t)

}

func mustToken() string {
	token := flag.String(
		"t",
		"",
		"token for my test bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is required")
	}

	return *token
}
