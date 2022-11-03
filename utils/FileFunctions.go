package utils

import (
	"codeReport/models"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func SearchingFiles() (string, []models.FileInfo) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	files := readDir("", pwd)

	return pwd, files
}

func readDir(pwd, curDir string) []models.FileInfo {
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
			files = append(files, readDir(filepath.Join(pwd, f.Name()), filepath.Join(curDir, f.Name()))...)
		}
	}
	return files
}

func GetFileMime(name string) string {
	fileMime := strings.Split(name, ".")
	return fileMime[len(fileMime)-1]
}
