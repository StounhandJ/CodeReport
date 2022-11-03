package main

import (
	"codeReport/docx"
	_interface "codeReport/interface"
	"codeReport/models"
	"codeReport/utils"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {

	utils.Animation()

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Current directory: ", pwd)

	files := utils.ReadDir("", pwd)

	generation(docx.NewSimpleDocxGeneration(pwd), files)

	fmt.Println("-----------------")
	fmt.Println("Everything is done!!!!!!!")
	fmt.Println("bye <3")
	time.Sleep(2 * time.Second)
}

func generation(g _interface.GenerationInterface, files []models.FileInfo) {
	table := g.CreateTable()
	table.AddRow([]string{"Модули", "Описание", "Количество строк кода", "Размер (в Кбайтах)"})
	table.AddRow([]string{"1", "2", "3", "4"})

	for n, f := range files {
		table.AddRow([]string{f.Path, "", strconv.Itoa(f.Rows), fmt.Sprint(math.Ceil(float64(f.Size) / 1000))})

		g.AddHeadingText(fmt.Sprintf("%d. %s", n+1, f.Path))

		content, err := os.ReadFile(f.FullPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		g.AddText(string(content))
	}

	g.Close()
}
