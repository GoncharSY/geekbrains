package duplicate

func ExampleFindAllDuplicates() {
	/*
		For example, we have tree like this.
		tmp
			|-file1
			|-file2
			|-dir1
				|-file1
				|-file2
			|-dir2
				|-file2
				|-file3
	*/

	var dir = "./tmp"
	var dup = FindAllDuplicates(dir)

	/*
		...will be returned:
		[
			["./tmp/file1", "./tmp/dir1/file1"]
			["./tmp/file2", "./tmp/dir1/file2", "./tmp/dir2/file2"]
			["./tmp/file3", "./tmp/dir2/file3", "./tmp/dir3/file3"]
			["./tmp/file4", "./tmp/dir3/file4"]
		]
	*/

	PrintAllDuplicates(dup)

	/*
		...will be printed:
		1:
		- 1.1: ./tmp/file1
		- 1.2: ./tmp/dir1/file1
		2:
		- 2.1: ./tmp/file2
		- 2.2: ./tmp/dir1/file2
		- 2.3: ./tmp/dir2/file2
		3:
		- 3.1: ./tmp/file3
		- 3.2: ./tmp/dir2/file3
		- 3.3: ./tmp/dir3/file3
		4:
		- 4.1: ./tmp/file4
		- 4.2: ./tmp/dir3/file4
	*/
}

func ExamplePrintDuplicates() {
	/*
		For example, we have tree like this.
		tmp
			|-file1
			|-file2
			|-dir1
				|-file1
				|-file2
			|-dir2
				|-file2
				|-file3
	*/

	var file = NewFile("./tmp/file2")
	var dir = "./tmp"
	var dup = file.FindDuplicates(dir)

	/*
		...will be returned:
		["./tmp/file2", "./tmp/dir1/file2", "./tmp/dir2/file2"]
	*/

	PrintDuplicates(dup, "")

	/*
		1: ./tmp/file2
		2: ./tmp/dir1/file2
		3: ./tmp/dir2/file2
	*/
}
