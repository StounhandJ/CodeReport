package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"codeReport/models"
)

const MaxScannedFileSize int64 = 20 * 1024 * 1024

func SearchingFiles() (string, []models.FileInfo, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", nil, err
	}

	files, err := readDir("", pwd, NewIgnoreMatcher(pwd))
	if err != nil {
		return "", nil, err
	}

	return pwd, files, nil
}

func readDir(pwd, curDir string, ignore *IgnoreMatcher) ([]models.FileInfo, error) {
	curFiles, err := os.ReadDir(curDir)
	files := make([]models.FileInfo, 0)

	if err != nil {
		return nil, err
	}

	for _, f := range curFiles {
		relPath := filepath.Join(pwd, f.Name())
		if ignore.ShouldIgnore(relPath, f.IsDir()) {
			continue
		}

		if !f.IsDir() {
			fileInfo, err := f.Info()
			if err != nil || fileInfo.Size() > MaxScannedFileSize {
				continue
			}

			fullPath := filepath.Join(curDir, f.Name())
			if ok, err := IsTextFile(fullPath); err != nil || !ok {
				continue
			}

			rows, err := CountRows(fullPath)
			if err != nil {
				continue
			}

			files = append(files, *models.NewFileInfo(relPath, fullPath, fileInfo.Size(), rows))
		} else {
			nestedFiles, err := readDir(relPath, filepath.Join(curDir, f.Name()), ignore)
			if err != nil {
				return nil, err
			}

			files = append(files, nestedFiles...)
		}
	}

	return files, nil
}

func IsTextFile(path string) (bool, error) {
	file, err := os.Open(path)

	if err != nil {
		return false, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}()

	buffer := make([]byte, 8192)

	n, err := file.Read(buffer)
	if err != nil && !errors.Is(err, io.EOF) {
		return false, err
	}

	for _, b := range buffer[:n] {
		if b == 0 {
			return false, nil
		}
	}

	return true, nil
}

func CountRows(path string) (int, error) {
	file, err := os.Open(path)

	if err != nil {
		return 0, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}()

	reader := bufio.NewReader(file)

	rows := 0
	for {
		line, err := reader.ReadString('\n')
		if line != "" {
			rows++
		}

		if err == nil {
			continue
		}

		if err == io.EOF {
			break
		}

		return 0, err
	}

	return rows, nil
}
