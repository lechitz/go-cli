package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Infosys Cloud CLI")
	fmt.Println("---------------------")
	fmt.Print("=>")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		userInput := scanner.Text()
		fmt.Println(userInput)
		fmt.Print("=>")

		if strings.Compare("exit", userInput) == 0 {
			os.Exit(0)
		}

		var commandString []string = strings.Split(userInput, " ")

		if strings.Compare("checkinfyconnect", commandString[0]) == 0 {
			commandOutput := execPINGCommand(commandString[1])
			fmt.Println(commandOutput)
		} else {
			fmt.Println("Command not supported by the CLI")
		}

		fmt.Print("=>")
	}
}

func execPINGCommand(cmdParam string) string {

	if c, err := exec.Command("cmd", "/c", "ping", "-n", "3", cmdParam).CombinedOutput(); err != nil {
		s := fmt.Sprintf("Error: %v", err)
		return s
	} else {
		return string(c)
	}
}
