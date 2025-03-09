package config

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("config")
}

func Read() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	fmt.Print(home)

	return Config{}, nil
}
