package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const helpText = `
Usage:
%s [names...]
 
Locates the path of the provided names.
	
names: names of the commands to look for
`

func main() {
	if len(os.Args) == 1 {
		fmt.Printf(helpText, os.Args[0])
		os.Exit(1)
	}
	names := os.Args[1:]
	dirs := filepath.SplitList(os.Getenv("PATH"))
	found, not_found := which(names, dirs)
	show_locations(found)
	if len(found) > 0 {
		fmt.Println()
	}
	show_not_in_path(not_found)
}

func show_not_in_path(names []string) {
	if len(names) == 0 {
		return
	}
	fmt.Println("Not found in path commands:")
	for _, name := range names {
		fmt.Println(name)
	}
}

func show_locations(locations map[string][]string) {
	if len(locations) == 0 {
		return
	}
	fmt.Println("Commands found:")
	for name, paths := range locations {
		fmt.Println(name)
		for _, path := range paths {
			fmt.Println("\t", path)
		}
	}
}

func which(names []string, dirs []string) (map[string][]string, []string) {
	found := map[string][]string{}
	not_found := []string{}
	for _, name := range names {
		locations := search_in(name, dirs)
		if len(locations) > 0 {
			found[name] = locations
		} else {
			not_found = append(not_found, name)
		}
	}
	return found, not_found
}

func search_in(name string, dirs []string) []string {
	found_in := []string{}
	for _, dir := range dirs {
		path := filepath.Join(dir, name)
		info, err := os.Stat(path)
		if err == nil {
			mode := info.Mode()
			if mode.IsRegular() && mode.Perm()&0111 != 0 {
				found_in = append(found_in, path)
			}
		}
	}
	return found_in
}
