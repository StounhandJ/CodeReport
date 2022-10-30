package utils

import (
	"codeReport/models"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ReadDir(pwd, curDir string) []models.FileInfo {
	curFiles, err := os.ReadDir(curDir)
	files := make([]models.FileInfo, 0)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range curFiles {
		if !f.IsDir() {
			if InArray(GetFileMime(f.Name()), []string{"docx", "doc", "xlsx", "pptx", "pdf", "png", "jpeg", "gif", "mp4", "zip", "exe", "mp3"}) != -1 {
				continue
			}
			files = append(files, *models.NewFileInfo(pwd, curDir, f))
		} else {
			files = append(files, ReadDir(filepath.Join(pwd, f.Name()), filepath.Join(curDir, f.Name()))...)
		}
	}
	return files
}

func GetFileMime(name string) string {
	fileMime := strings.Split(name, ".")
	return fileMime[len(fileMime)-1]
}
