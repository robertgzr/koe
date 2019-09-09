module github.com/robertgzr/rjoe

go 1.12

require (
	github.com/go-joe/joe v0.8.0
	github.com/robertgzr/joe-adapter-telegram v0.0.0-00010101000000-000000000000
)

replace github.com/robertgzr/joe-adapter-telegram => ./telegram-adapter
