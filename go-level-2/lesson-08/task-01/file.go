package main

import (
	"fmt"
	"io/fs"
	"os"
)

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

type File struct {
	Path string
	Info os.FileInfo
}

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