package main

import (
	"flag"
	"log"
	"never_read_list/clients/telegram"
)

func main() {
	host := mustFlag("telegram-host", "Telegram API host")
	token := mustFlag("telegram-token", "Token to connect to Telegram API")
	tgClient = telegram.New(host, token)

	// fetcher = fetcher.New(tgClient)

	// processor = processor.New(tgClient)

	// consumer.Start(feetcher, processor)
}

func mustFlag(name string, hint string) string {
	val := flag.String(name, "", hint,)

	flag.Parse()

	if *val == "" {
		log.Fatal("%s was not provided", name)
	}

	return *val
}