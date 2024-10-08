package main

import (
	"Telegrambot/clients/telegram"
	"flag"
	"log"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	tgClient := telegram.New(tgBotHost, mustToken())

	// fetcher =  fetcher.New(tgClient)

	// Processor =  Processor.New(tgClient)

	// consumer.Start(Fetcher,Processing)
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for tg bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specifed") //WHAT IS?
	}
	return *token

}
