package main

import (
	"codeReport/docx"
	_interface "codeReport/interface"
	"codeReport/models"
	"codeReport/utils"
	"context"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	//welcome
	utils.Welcome()
	time.Sleep(1 * time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	go utils.Animation(ctx)

	//work
	pwd, files := utils.SearchingFiles()
	generation(docx.NewSimpleDocxGeneration(pwd), files)

	//bye
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}

func generation(g _interface.GenerationInterface, files []models.FileInfo) {
	table := g.CreateTable()
	table.AddRow([]string{"Модули", "Описание", "Количество строк кода", "Размер (в Кбайтах)"})
	table.AddRow([]string{"1", "2", "3", "4"})

	for n, f := range files {
		content, err := os.ReadFile(f.FullPath)
		if err != nil {
			continue
		}

		table.AddRow([]string{f.Path, "", strconv.Itoa(f.Rows), fmt.Sprint(math.Ceil(float64(f.Size) / 1000))})

		g.AddHeadingText(fmt.Sprintf("%d. %s", n+1, f.Path))

		g.AddText(string(content))
	}

	g.Close()
}
