package main

import (
	"MyTelegramBot/pkg/clients/telgram"
	"flag"
	"log"
)

const tgBotHost = "api.telegram.org"

func main() {
	t := mustToken()
	tgClient := telgram.New(tgBotHost, t)

	// fetcher = fetcher.New(tgClient)

	// processor = processor.New(tgClient)

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
