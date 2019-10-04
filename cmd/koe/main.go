package main

import (
	"os"
	"errors"

	"github.com/go-joe/joe"
	telegram "github.com/robertgzr/joe-adapter-telegram"
)

type Koe struct {
	*joe.Bot
}

func main() {
	b := &Koe{
		joe.New("koe",
			telegram.Adapter(os.Getenv("TELEGRAM_TOKEN")),
		),
	}

	b.Brain.RegisterHandler(b.HandleCommands)

	err := b.Run()
	if err != nil {
		b.Logger.Fatal(err.Error())
	}
}

func (b *Koe) HandleCommands(ev telegram.ReceiveCommandEvent) error {
	var err error

	switch ev.Arg0 {
	case "version":
		b.Say(ev.Channel(), version())

	default:
		err = errors.New("unknown command")
	}

	return err
}
