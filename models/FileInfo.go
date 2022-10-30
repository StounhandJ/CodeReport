package models

import (
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
)

type FileInfo struct {
	FullPath string

	Path string

	Size int64

	Rows int
}

func NewFileInfo(pwd, fullPath string, file fs.DirEntry) *FileInfo {
	fileInfo, _ := file.Info()
	path := filepath.Join(pwd, file.Name())

	content, _ := os.ReadFile(path)

	return &FileInfo{
		FullPath: filepath.Join(fullPath, file.Name()),
		Path:     path,
		Size:     fileInfo.Size(),
		Rows:     bytes.Count(content, []byte{'\n'}),
	}
}
