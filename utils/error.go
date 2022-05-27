package utils

import "fmt"

var Reset = "\033[0m"
var Red = "\033[31m"

func Error(error string) {
	fmt.Print(Red + "Error: " + Reset + error)
}
