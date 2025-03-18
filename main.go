package main

import (
	"fmt"
	"github.com/mcdotjs/blog_aggregator/internal/config"
	"log"
	"os"
)
import _ "github.com/lib/pq"

type state struct {
	value *config.Config
}

func main() {
	fileContent, err := config.Read()
	if err != nil {
		log.Fatalln("Problem with reading file")
	}

	globalState := &state{
		value: &fileContent,
	}

	fmt.Println("first read", (*globalState).value.CurrentUserName)
	avialableCommands := commands{
		value: make(map[string]func(*state, command) error),
	}

	avialableCommands.register("login", handlerLogin)
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	words := os.Args

	args := []string{}
	if len(words) > 1 {
		args = words[1:]
	}
	currentCommand := command{
		Name: args[0],
		Args: args[1:],
	}
	avialableCommands.run(globalState, currentCommand)
	fileContent, err = config.Read()
	if err != nil {
		fmt.Println("Problem with reading file")
	}
	fmt.Println("second read", (*globalState).value.CurrentUserName)

}
