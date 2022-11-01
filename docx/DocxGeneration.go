package docx

import (
	_interface "codeReport/interface"
	"fmt"
	"github.com/Preciselyco/unioffice/document"
	"os"
	"path/filepath"
	"strings"
)

type Generation struct {
	doc *document.Document
	pwd string
}

func NewSimpleDocxGeneration(pwd string) *Generation {
	return &Generation{
		doc: document.New(),
		pwd: pwd,
	}
}

func (g *Generation) CreateTable() _interface.TableGenerationInterface {
	return NewTableGeneration(g.doc.AddTable())
}

func (g *Generation) AddHeadingText(text string) {
	para := g.doc.AddParagraph()
	para.AddRun().AddText(text)
}

func (g *Generation) AddText(text string) {
	for _, t := range strings.Split(text, "\n") {
		para := g.doc.AddParagraph()
		run := para.AddRun()
		run.Properties().SetSize(8)
		run.AddText(t)
	}
}

func (g *Generation) Close() {
	err := g.doc.SaveToFile(fmt.Sprintf("%s.docx", filepath.Base(g.pwd)))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
