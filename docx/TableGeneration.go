package docx

import (
	"github.com/Preciselyco/unioffice/color"
	"github.com/Preciselyco/unioffice/document"
	"github.com/Preciselyco/unioffice/measurement"
	"github.com/Preciselyco/unioffice/schema/soo/wml"
)

const documentFontFamily = "Times New Roman"

type TableGeneration struct {
	table document.Table
}

func NewTableGeneration(table document.Table) *TableGeneration {
	table.Properties().Borders().SetAll(wml.ST_BorderSingle, color.Auto, measurement.Point)
	return &TableGeneration{
		table: table,
	}
}

func (s *TableGeneration) AddRow(cells []string) {
	row := s.table.AddRow()
	for _, cell := range cells {
		paragraph := row.AddCell().AddParagraph()
		paragraph.Properties().SetAlignment(wml.ST_JcCenter)
		run := paragraph.AddRun()
		run.Properties().SetFontFamily(documentFontFamily)
		run.AddText(cell)
	}
}
