package koe

import (
	"fmt"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-joe/joe"
	bolt "github.com/robertgzr/joe-bolt-memory"
	telegram "github.com/robertgzr/joe-telegram-adapter"
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
	b.Respond("help", b.Usage)
	b.Respond("ama", b.HowAreYou)

	if err := b.Run(); err != nil {
		b.Logger.Fatal(err.Error())
	}
}

func (b *Koe) Usage(m joe.Message) error {
	b.Say(m.Channel, `
commands:
/version
/dump
/fin <desc> <category> <amount>

triggers:
ama     asks how you are feeling
help    show this help
`)
	return nil
}

func (b *Koe) HandleCommands(ev telegram.ReceiveCommandEvent) error {
	switch ev.Arg0 {
	case "version":
		b.Say(ev.Channel(), fmt.Sprintf("%v version %v (%v)",
			version.Package, version.Version, version.Revision))

	case "dump":
		b.Say(ev.Channel(), "dumping next message")
		if err := b.Store.Set("dump_messages", true); err != nil {
			return err
		}
		b.Brain.RegisterHandler(func(ev joe.ReceiveMessageEvent) error {
			var dump bool
			ok, err := b.Store.Get("dump_messages", &dump)
			if err != nil {
				return err
			}
			if !ok || !dump {
				return nil
			}
			b.Say(ev.Channel, spew.Sdump(ev.Data))
			if _, err := b.Store.Delete("dump_messages"); err != nil {
				return err
			}
			return nil
		})

	case "fin":
		return b.FinHandler(ev)

	}
	return nil
}

func (b *Koe) HowAreYou(msg joe.Message) error {
	tg, ok := b.Adapter.(*telegram.TelegramAdapter)
	if !ok {
		return fmt.Errorf("Adapter not Telegram")
	}
	return tg.SendButtons(msg.Channel, "How are you?",
		tg.NewButton("amazing", func(channel string) error {
			b.Say(channel, "you said: amazing")
			return nil
		}),
		tg.NewButton("ok", func(channel string) error {
			b.Say(channel, "you said: ok")
			return nil
		}),
		tg.NewButton("meh", func(channel string) error {
			b.Say(channel, "you said: meh")
			return nil
		}))
}
