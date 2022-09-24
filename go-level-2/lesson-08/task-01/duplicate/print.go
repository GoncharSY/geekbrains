package duplicate

import "fmt"

// PrintDuplicates prints duplicate file paths.
func PrintDuplicates(files []File, prefix string) {
	if len(files) == 0 {
		fmt.Println("No files")
	}

	for j, file := range files {
		fmt.Printf("%v%v: %v\n", prefix, j+1, file.Path)
	}
}

// PrintAllDuplicates prints duplicate paths that are grouped.
func PrintAllDuplicates(files [][]File) {
	if len(files) == 0 {
		fmt.Println("No duplicates found")
	}

	for i, dups := range files {
		if len(dups) > 1 {
			fmt.Printf("%v:\n", i+1)
			PrintDuplicates(dups, fmt.Sprintf("- %v.", i+1))
		}
	}
}

// PrintAllFiles prints files paths that are grouped.
func PrintAllFiles(files [][]File) {
	if len(files) == 0 {
		fmt.Println("No files")
	}

	for i, dups := range files {
		fmt.Printf("%v:\n", i+1)
		PrintDuplicates(dups, fmt.Sprintf("- %v.", i+1))
	}
}
