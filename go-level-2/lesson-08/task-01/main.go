package main

import "fmt"

func main() {
	var cfg = NewConfig()

	fmt.Println(cfg)
	fmt.Println()

	switch cfg.Action {
	case Find:
		if cfg.File == "" {
			printAllDuplicates(FindAllDuplicates(cfg.Folder))
		} else {
			printDuplicates(NewFile(cfg.File).FindDuplicates(cfg.Folder), "")
		}
	case Create:
		fmt.Println("Functionality in development")
	case Delete:
		fmt.Println("Functionality in development")
	}
}

func printDuplicates(files []File, prefix string) {
	if len(files) == 0 {
		fmt.Println("No duplicates found")
	}

	for j, file := range files {
		fmt.Printf("%v%v: %v\n", prefix, j+1, file.Path)
	}
}

func printAllDuplicates(files [][]File) {
	if len(files) == 0 {
		fmt.Println("No duplicates found")
	}

	for i, dups := range files {
		if len(dups) > 1 {
			fmt.Printf("%v:\n", i+1)
			printDuplicates(dups, fmt.Sprintf("- %v.", i+1))
		}
	}
}
