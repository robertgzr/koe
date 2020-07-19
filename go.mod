module github.com/robertgzr/koe

go 1.14

require (
	github.com/go-joe/joe v0.8.0
	github.com/peterbourgon/ff v1.7.0
	github.com/peterbourgon/ff/v3 v3.0.0
	github.com/robertgzr/joe-adapter-telegram v0.0.0-00010101000000-000000000000
	github.com/robertgzr/joe-memory-bolt v0.0.0-00010101000000-000000000000
)

replace (
	github.com/go-joe/joe => ./pkg/joe-bot
	github.com/robertgzr/joe-adapter-telegram => ./pkg/telegram-adapter
	github.com/robertgzr/joe-memory-bolt => ./pkg/bolt-memory
)
