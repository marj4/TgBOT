package main

import (
	"flag"
	"log"
)

func main() {
	mustToken()

	//tgClient= telegram.New(token)

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
	if *token == ""{
		log.Fatal("token is not specifed") //WHAT IS?
	}
	return *token

}