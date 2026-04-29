package models

type FileInfo struct {
	FullPath string

	Path string

	Size int64

	Rows int
}

func NewFileInfo(path, fullPath string, size int64, rows int) *FileInfo {
	return &FileInfo{
		FullPath: fullPath,
		Path:     path,
		Size:     size,
		Rows:     rows,
	}
}
