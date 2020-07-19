package main

import (
	"flag"
	"log"
	"os"

	"github.com/peterbourgon/ff/v3"
	"github.com/robertgzr/koe"
)

func main() {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	var cfg koe.Config
	fs.StringVar(&cfg.Root, "root", ".", "root path for the bot, will contain memory database")
	fs.StringVar(&cfg.TelegramToken, "telegram-token", "", "api bot token")

	if err := ff.Parse(fs, os.Args[1:],
		ff.WithEnvVarPrefix("KOE"),
	); err != nil {
		log.Fatal(err)
	}

	koe.Run(cfg)
}
