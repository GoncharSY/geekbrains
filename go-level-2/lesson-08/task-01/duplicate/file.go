package duplicate

import (
	"fmt"
	"io/fs"
	"os"
)

// Creates file-description object and returns pointer of to it's.
func NewFile(path string) *File {
	var file *os.File
	var err error

	if file, err = os.Open(path); err != nil {
		panic(err)
	} else {
		defer file.Close()
	}

	if stt, err := file.Stat(); err != nil {
		panic(err)
	} else {
		return &File{
			Path: path,
			Info: stt,
		}
	}
}

// The type for file description.
type File struct {
	Path string
	Info os.FileInfo
}

// FindDuplicates finds and returns files which have same name and size like the receiver.
func (file *File) FindDuplicates(dir string) []File {
	var res = make([]File, 0, 10)
	var dfs = getDirFS(dir)

	fs.WalkDir(dfs, ".", func(pth string, ent fs.DirEntry, err error) error {
		if err != nil {
			panic(err)
		} else if ent.IsDir() {
			return nil
		} else if inf, err := ent.Info(); err != nil {
			panic(err)
		} else if inf.Name() == file.Info.Name() && inf.Size() == file.Info.Size() {
			res = append(res, File{
				Path: fmt.Sprintf("%v/%v", dir, pth),
				Info: inf,
			})
		}

		return nil
	})

	return res
}

// DeleteDuplicates removes duplicates of the file and returns list of them.
func (file *File) DeleteDuplicates(dir string) []File {
	var dps = file.FindDuplicates(dir)
	var dls = make([]File, 0, len(dps))

	for i := range dps {
		if os.SameFile(file.Info, dps[i].Info) {
			continue
		} else if err := os.Remove(dps[i].Path); err != nil {
			panic(err)
		} else {
			dls = append(dls, dps[i])
		}
	}

	return dls
}
