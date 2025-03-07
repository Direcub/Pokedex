package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		command := input[0]
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}
		strct, ok := Registry[command]
		if ok {
			err := strct.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Uknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	lowercase := strings.ToLower(text)
	splitted := strings.Fields(lowercase)
	for i := 0; i < len(splitted); i++ {
		splitted[i] = strings.TrimSpace(splitted[i])
		splitted[i] = strings.Trim(splitted[i], ",")
	}
	return splitted
}
