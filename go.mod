module github.com/robertgzr/koe

go 1.14

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/go-joe/joe v0.8.0
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
	github.com/peterbourgon/ff v1.7.0
	github.com/peterbourgon/ff/v3 v3.0.0
	github.com/robertgzr/joe-bolt-memory v0.0.0-00010101000000-000000000000
	github.com/robertgzr/joe-telegram-adapter v0.0.0-00010101000000-000000000000
)

replace (
	github.com/go-joe/joe => ./pkg/joe-bot
	github.com/robertgzr/joe-bolt-memory => ./pkg/bolt-memory
	github.com/robertgzr/joe-telegram-adapter => ./pkg/telegram-adapter
)
