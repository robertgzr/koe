package main

import "fmt"

var (
	Version   string = "0.1.0"
	BuildTime string = "undefined"
)

func version() string {
	return fmt.Sprintf("joebot %v (%s)", Version, BuildTime)
}
