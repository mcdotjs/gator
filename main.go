package main

import (
	"database/sql"
	"fmt"
	"github.com/mcdotjs/blog_aggregator/internal/config"
	"github.com/mcdotjs/blog_aggregator/internal/database"
	"log"
	"os"
)
import _ "github.com/lib/pq"

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	fileContent, err := config.Read()
	if err != nil {
		log.Fatalln("Problem with reading file")
	}

	db, err := sql.Open("postgres", fileContent.DbURL)
	dbQueries := database.New(db)
	globalState := &state{
		cfg: &fileContent,
		db:  dbQueries,
	}

	fmt.Println("first read", (*globalState).cfg.CurrentUserName)
	avialableCommands := commands{
		value: make(map[string]func(*state, command) error),
	}

	avialableCommands.register("login", handlerLogin)
	avialableCommands.register("register", handlerRegister)
	avialableCommands.register("reset", handlerReset)
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
	fmt.Println("second read", (*globalState).cfg.CurrentUserName)

}
