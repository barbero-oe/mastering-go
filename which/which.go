package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a single command to search")
		os.Exit(1)
	}
	command := os.Args[1]
	path, err := which(command)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(path)
}

func which(name string) (string, error) {
	dirs := filepath.SplitList(os.Getenv("PATH"))
	for _, dir := range dirs {
		commandPath := filepath.Join(dir, name)
		info, err := os.Stat(commandPath)
		if err == nil {
			mode := info.Mode()
			if mode.IsRegular() && mode.Perm()&0111 != 0 {
				return commandPath, nil
			}
		}
	}
	return "", errors.New("command not found")
}
