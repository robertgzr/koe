package main

import (
	"errors"
	"os"

	"github.com/go-joe/joe"
	telegram "github.com/robertgzr/joe-adapter-telegram"
)

type Bot struct {
	*joe.Bot
}

func main() {
	b := &Bot{
		joe.New("rjoe",
			telegram.Adapter(os.Getenv("TELEGRAM_TOKEN")),
		),
	}

	b.Brain.RegisterHandler(b.HandleCommands)

	err := b.Run()
	if err != nil {
		b.Logger.Fatal(err.Error())
	}
}

func (b *Bot) HandleCommands(ev telegram.ReceiveCommandEvent) error {
	var err error

	switch ev.Arg0 {
	case "version":
		b.Say(ev.Channel(), version())

	default:
		err = errors.New("unknown command")
	}

	return err
}
