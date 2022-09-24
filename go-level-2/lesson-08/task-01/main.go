package main

import (
	"fmt"
	"go-level-2/lesson-08/task-01/duplicate"
)

func main() {
	var cfg = duplicate.NewConfig()

	fmt.Println(cfg)
	fmt.Println()

	switch cfg.Action {
	case duplicate.Find:
		findDuplicates(cfg)
	case duplicate.Create:
		createDuplicates(cfg)
	case duplicate.Delete:
		deleteDuplicates(cfg)
	}
}

// findDuplicates finds and prints file duplicates.
func findDuplicates(cfg *duplicate.Config) {
	if cfg.File == "" {
		var dup = duplicate.FindAllDuplicates(cfg.Folder)
		duplicate.PrintAllDuplicates(dup)
	} else {
		var ent = duplicate.NewFile(cfg.File)
		var dup = ent.FindDuplicates(cfg.Folder)
		duplicate.PrintDuplicates(dup, "")
	}
}

// deleteDuplicates deletes duplicates and prints paths of deleted files.
func deleteDuplicates(cfg *duplicate.Config) {
	if cfg.File == "" {
		var dup = duplicate.DeleteAllDuplicates(cfg.Folder)
		duplicate.PrintAllFiles(dup)
	} else {
		var ent = duplicate.NewFile(cfg.File)
		var dup = ent.DeleteDuplicates(cfg.Folder)
		duplicate.PrintDuplicates(dup, "")
	}
}

// createDuplicates creates duplicates and prints paths of created files.
func createDuplicates(cfg *duplicate.Config) {
	if cfg.File == "" {
		fmt.Println("Please specify the original file path (flag \"-file\")")
	} else {
		fmt.Println("Functionality in development")
	}
}
