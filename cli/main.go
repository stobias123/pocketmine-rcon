package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	rcon "github.com/micvbang/pocketmine-rcon"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: ./rcon address password")
		return
	}
	addr := os.Args[1]
	pass := os.Args[2]

	conn, err := rcon.NewConnection(addr, pass)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Successfully logged in at %s!\n", addr)

	prompt()
	stdin := bufio.NewReader(os.Stdin)
	input := ""
	for {
		if input, err = stdin.ReadString('\n'); err != nil {
			fmt.Println(err)
			return
		}
		input = strings.TrimSuffix(input, "\r\n")
		input = strings.Trim(input[:len(input)-1], " ")
		if input == ".exit" {
			break
		}
		if len(input) == 0 {
			prompt()
			continue
		}

		r, err := conn.SendCommand(input)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			prompt()
			continue
		}
		// this is where I gotta make it.
		fmt.Printf("Server:\n%s\n", sanitizeResponse(r))
		prompt()
	}
}

func prompt() {
	fmt.Print("Enter command:\n>")
}

// todo - maybe skip if we don't have xterm colors or are on windoze.
func sanitizeResponse(response string) string {
	mcColors := getMinecraftColors()
	for i := range mcColors {
		targetColor := getConsoleColor(mcColors[i])
		response = strings.ReplaceAll(response, mcColors[i], targetColor)
	}
	return string(response)
}
