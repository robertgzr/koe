package main

import (
	"errors"
	"os"

	"github.com/go-joe/joe"
	telegram "github.com/robertgzr/joe-adapter-telegram"
	bolt "github.com/robertgzr/joe-memory-bolt"
)

type Koe struct {
	*joe.Bot
}

func main() {
	b := &Koe{
		joe.New("koe",
			telegram.Adapter(os.Getenv("TELEGRAM_TOKEN")),
			bolt.Memory(os.Getenv("DB_PATH")),
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
