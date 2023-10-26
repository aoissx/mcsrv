package config

import "github.com/fatih/color"

const (
	logo = "[mcsrv]"
)

func LogSuccess(msg string) {
	c := color.New(color.FgGreen)
	c.Println(logo, msg)
}

func LogError(msg string) {
	c := color.New(color.FgRed)
	c.Println(logo, msg)
}

func LogInfo(msg string) {
	c := color.New(color.FgBlue)
	c.Println(logo, msg)
}
