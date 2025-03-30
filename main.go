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

	avialableCommands := commands{
		value: make(map[string]func(*state, command) error),
	}

	avialableCommands.register("login", handlerLogin)
	avialableCommands.register("register", handlerRegister)
	avialableCommands.register("reset", handlerReset)
	avialableCommands.register("users", handlerGetUsers)
	avialableCommands.register("agg", handlerFeeds)
	avialableCommands.register("feeds", getFeeds)
	avialableCommands.register("addfeed", middlewareLoggedIn(addFeed))
	avialableCommands.register("follow", middlewareLoggedIn(follow))
	avialableCommands.register("unfollow", middlewareLoggedIn(unfollow))
	avialableCommands.register("following", middlewareLoggedIn(following))

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

}
