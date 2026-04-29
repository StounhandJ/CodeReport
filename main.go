package main

import (
	"codeReport/docx"
	_interface "codeReport/interface"
	"codeReport/models"
	"codeReport/utils"
	"context"
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	defer utils.WaitForExit()
	if err := run(); err != nil {
		fmt.Println()
		fmt.Println("Error:", err)
	}
}

func run() error {
	//welcome
	utils.Welcome()
	time.Sleep(1 * time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	animationDone := make(chan struct{})
	go func() {
		defer close(animationDone)
		utils.Animation(ctx)
	}()

	//work
	pwd, files, err := utils.SearchingFiles()
	if err != nil {
		cancel()
		<-animationDone
		return err
	}
	err = generation(docx.NewSimpleDocxGeneration(pwd), files)

	//bye
	time.Sleep(3 * time.Second)
	cancel()
	<-animationDone

	return err
}

func generation(g _interface.GenerationInterface, files []models.FileInfo) error {
	table := g.CreateTable()
	table.AddRow([]string{"Номер", "Файл", "Описание", "Количество строк", "Размер (Кбайт)"})
	table.AddRow([]string{"1", "2", "3", "4", "5"})

	for n, f := range files {
		fileNumber := strconv.Itoa(n + 1)
		table.AddRow([]string{fileNumber, f.Path, "", strconv.Itoa(f.Rows), fmt.Sprint(math.Ceil(float64(f.Size) / 1000))})

		g.AddHeadingText(fmt.Sprintf("%s. %s", fileNumber, f.Path))

		if err := g.AddFileText(f.FullPath); err != nil {
			continue
		}
	}

	return g.Close()
}
