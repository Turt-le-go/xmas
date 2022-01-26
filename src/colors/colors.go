package colors

import (
	"time"
	"math/rand"
)

const (
	black string= "\u001b[30m"
	red			= "\u001b[31m"
	green		= "\u001b[32m"
	yellow		= "\u001b[33m"
	blue		= "\u001b[34m"
	magenta		= "\u001b[35m"
	cyan		= "\u001b[36m"
	white		= "\u001b[37m"
	reset		= "\u001b[0m"
)

func Black(s string) string {
	return black + s + reset
}

func Red(s string) string {
	return red + s + reset
}

func Green(s string) string {
	return green + s + reset
}

func Yellow(s string) string {
	return yellow + s + reset
}

func Blue(s string) string {
	return blue + s + reset
}

func Magenta(s string) string {
	return magenta + s + reset
}

func Cyan(s string) string {
	return cyan + s + reset
}

func White(s string) string {
	return white + s + reset
}

var Colors = [8]func(string)string{Black,Red,Green,Yellow,Blue,Magenta,Cyan,White}

func Random(s string) string {
	rand.Seed(time.Now().UnixMilli())
	color := Colors[rand.Intn(len(Colors))]
	return color(s)
}
