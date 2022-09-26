package duplicate

import (
	"fmt"
	"io/fs"
	"os"
)

// FindAllDuplicates find and returns file duplicates that grouped.
func FindAllDuplicates(dir string) [][]File {
	var res0 = make(map[string][]File)
	var res1 = make([][]File, 0, 10)
	var dfs = getDirFS(dir)

	if err := fs.WalkDir(dfs, ".", groupDuplicates(&res0, dir)); err != nil {
		panic(err)
	} else {
		for _, lst := range res0 {
			if len(lst) > 1 {
				res1 = append(res1, lst)
			}
		}
	}

	return res1
}

// getDirFS returns the hierarchical file system of directory.
func getDirFS(dir string) fs.FS {
	if inf, err := os.Stat(dir); err != nil {
		panic(err)
	} else if !inf.IsDir() {
		panic(fmt.Errorf("%s isn't directory", dir))
	} else {
		return os.DirFS(dir)
	}
}

// groupDuplicates returns function which groups duplicate-files.
func groupDuplicates(paths *map[string][]File, dir string) func(string, fs.DirEntry, error) error {
	return func(pth string, ent fs.DirEntry, err error) error {
		var key string

		if err != nil {
			panic(err)
		} else if ent.IsDir() {
			return nil
		} else if inf, err := ent.Info(); err != nil {
			panic(err)
		} else {
			key = fmt.Sprintf("%v::%v", inf.Name(), inf.Size())

			(*paths)[key] = append((*paths)[key], File{
				Path: fmt.Sprintf("%v/%v", dir, pth),
				Info: inf,
			})
		}

		return nil
	}
}
