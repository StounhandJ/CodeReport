package docx

import (
	"bufio"
	_interface "codeReport/interface"
	"fmt"
	"github.com/Preciselyco/unioffice/document"
	"io"
	"os"
	"path/filepath"
)

type Generation struct {
	doc        *document.Document
	pwd        string
	outputPath string
}

func NewSimpleDocxGeneration(pwd string) *Generation {
	return NewDocxGeneration(pwd, fmt.Sprintf("%s.docx", filepath.Base(pwd)))
}

func NewDocxGeneration(pwd, outputPath string) *Generation {
	return &Generation{
		doc:        document.New(),
		pwd:        pwd,
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
	defer file.Close()

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
