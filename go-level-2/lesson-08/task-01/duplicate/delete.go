package duplicate

import "fmt"

// DeleteAllDuplicates deletes duplicates and returns remoed files that grouped.
func DeleteAllDuplicates(dir string) [][]File {
	var dup = FindAllDuplicates(dir)
	var dlt = make([][]File, 0, len(dup))

	for _, fls := range dup {
		if len(fls) <= 1 {
			continue
		}

		if file := selectOriginFile(fls); file != nil {
			fmt.Println("Selected:", file.Path)
			dlt = append(dlt, file.DeleteDuplicates(dir))
		} else {
			fmt.Println("Canceled")
		}

		fmt.Println()
	}

	return dlt
}

// selectOriginFile returns selected file from the `list`.
func selectOriginFile(list []File) *File {
	PrintDuplicates(list, "")

	var txt = "Select number of original file from the list (cancel: 0): "
	var idx = askForFileNumber(txt, len(list))

	if idx > 0 && int(idx) <= len(list) {
		return &list[idx-1]
	} else {
		return nil
	}
}

// askForFileNumber returns selected number of the file.
func askForFileNumber(txt string, max int) uint {
	var res uint

	for {
		fmt.Print(txt)
		if _, err := fmt.Scanln(&res); err != nil {
			fmt.Println("Incorrect input:", err)
			clearInput()
		} else if int(res) > max {
			fmt.Println("Incorrect input:", max, "is maximum")
		} else {
			break
		}
	}

	return res
}

// clearInput clears input in console.
func clearInput() {
	var num = 1
	var err error

	for !(err == nil && num == 0) {
		num, err = fmt.Scanln()
	}
}
