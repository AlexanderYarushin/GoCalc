package main

import (
	"GoCalc/calculator"
	"GoCalc/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		text = text[:len(text)-1]

		result, err := calculator.Calculate(text)

		if err != nil {
			utils.Error(err.Error() + "\n")
			continue
		}

		fmt.Println(result)
	}

}
