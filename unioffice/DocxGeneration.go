package unioffice

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"

	_interface "codeReport/interface"

	"github.com/Preciselyco/unioffice/document"
)

type Generation struct {
	doc        *document.Document
	outputPath string
}

func NewSimpleDocxGeneration(pwd string) *Generation {
	return NewDocxGeneration(fmt.Sprintf("%s.docx", filepath.Base(pwd)))
}

func NewDocxGeneration(outputPath string) *Generation {
	return &Generation{
		doc:        document.New(),
		outputPath: outputPath,
	}
}

func (g *Generation) CreateTable() _interface.TableGenerationInterface {
	return NewTableGeneration(g.doc.AddTable())
}

func (g *Generation) AddHeadingText(text string) {
	para := g.doc.AddParagraph()
	run := para.AddRun()
	run.Properties().SetFontFamily(documentFontFamily)
	run.AddText(text)
}

func (g *Generation) AddText(text string) {
	para := g.doc.AddParagraph()
	run := para.AddRun()
	run.Properties().SetFontFamily(documentFontFamily)
	run.Properties().SetSize(8)
	run.AddText(text)
}

func (g *Generation) AddFileText(path string) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error: %s\n", err)
		}
	}()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if line != "" {
			g.AddText(line)
		}

		if err == nil {
			continue
		}

		if err == io.EOF {
			break
		}

		return err
	}

	return nil
}

func (g *Generation) Close() error {
	return g.doc.SaveToFile(g.outputPath)
}
