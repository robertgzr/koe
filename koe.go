package koe

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/go-joe/joe"
	telegram "github.com/robertgzr/joe-adapter-telegram"
	bolt "github.com/robertgzr/joe-memory-bolt"
	"github.com/robertgzr/koe/version"
)

type Koe struct {
	*joe.Bot
}

type Config struct {
	Root          string
	TelegramToken string
}

func Run(cfg Config) {
	b := &Koe{
		joe.New("koe",
			telegram.Adapter(cfg.TelegramToken),
			bolt.Memory(filepath.Join(cfg.Root, "memory.db")),
		),
	}

	b.Brain.RegisterHandler(b.HandleCommands)
	b.Brain.RegisterHandler(b.HandleCallbackQueries)
	b.Respond("ama", b.HowAreYou)
	b.Respond("help", b.Usage)

	if err := b.Run(); err != nil {
		b.Logger.Fatal(err.Error())
	}
}

func (b *Koe) Usage(m joe.Message) error {
	b.Say(m.Channel, `
commands:
/version

triggers:
ama     asks how you are feeling
help    show this help
`)
	return nil
}

func (b *Koe) HandleCommands(ev telegram.ReceiveCommandEvent) error {
	var err error

	switch ev.Arg0 {
	case "version":
		b.Say(ev.Channel(), fmt.Sprintf("%v version %v (%v)",
			version.Package, version.Version, version.Revision))

	default:
		err = errors.New("unknown command")
	}

	return err
}

func (b *Koe) HandleCallbackQueries(ev telegram.ReceiveCallbackQeryEvent) error {
	var err error

	switch ev.Data.Data {
	case "amazing":
		b.Say(ev.Channel(), "> amazing\n:)")
	case "ok":
		b.Say(ev.Channel(), "> ok\nit's getting better...")
	case "meh":
		b.Say(ev.Channel(), "> meh\nneed to talk?")
	default:
		return nil
	}
	return err
}

func (b *Koe) HowAreYou(msg joe.Message) error {
	adp2, ok := b.Adapter.(joe.Adapter2)
	if !ok {
		return fmt.Errorf("Adapter does not implement Self()")
	}
	tg, ok := adp2.Self().(*telegram.TelegramAdapter)
	if !ok {
		return fmt.Errorf("Adapter not Telegram")
	}
	return tg.SendButtons(msg.Channel, "How are you?",
		tg.NewButton("amazing"),
		tg.NewButton("ok"),
		tg.NewButton("meh"))
}
