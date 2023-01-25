package colors

import "fmt"

var ColorRed = "\033[31m"

var ColorReset = "\033[0m"
var ColorGreen = "\033[32m"
var ColorYellow = "\033[33m"
var ColorBlue = "\033[34m"
var ColorPurple = "\033[35m"
var ColorCyan = "\033[36m"
var ColorWhite = "\033[37m"

func Print(colorType string, msg string) {
	fmt.Println(string(colorType) + msg)
}
