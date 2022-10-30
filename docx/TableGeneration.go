package docx

import (
	"github.com/Preciselyco/unioffice/color"
	"github.com/Preciselyco/unioffice/document"
	"github.com/Preciselyco/unioffice/measurement"
	"github.com/Preciselyco/unioffice/schema/soo/wml"
)

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
		row.AddCell().AddParagraph().AddRun().AddText(cell)
	}
}
