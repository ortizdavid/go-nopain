package docgen

import (
	"bytes"
	"fmt"
	"net/http"
	"github.com/tealeg/xlsx"
)


type ExcelGenerator struct {
	File						*xlsx.File
	Sheet						*xlsx.Sheet
	HeaderColor					string
	TitleSize					int
	TitleFont					string
	CellPatterType				string
	CellFgColor					string
	CellBgColor					string
	ColMultiplier				float64
}


func NewExcelGenerator() *ExcelGenerator {
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")
	return &ExcelGenerator{
		File:        file,
		Sheet:       sheet,
		HeaderColor: "CC9F26",
		TitleSize: 		12,
		TitleFont: 		"Arial",
		CellPatterType: "solid",
		CellFgColor: 	"FF0000",
		CellBgColor: 	"FF0000",
		ColMultiplier: 1.2,// Adjust this multiplier based on your requirements
	}
}


func (eg *ExcelGenerator) AddTitle(title string) {
	row := eg.Sheet.AddRow()
	cell := row.AddCell()
	cell.Value = title
	titleStyle := xlsx.NewStyle()
	titleStyle.Font = *xlsx.NewFont(eg.TitleSize, eg.TitleFont)
	titleStyle.Font.Bold = true
	cell.SetStyle(titleStyle)
}


func (eg *ExcelGenerator) AddHeaderRow(columns ...string) {
	row := eg.Sheet.AddRow()
	for _, col := range columns {
		cell := row.AddCell()
		cell.Value = col
		cell.SetStyle(eg.getHeaderBorderStyle())
	}
	style := xlsx.NewStyle()
	fill := *xlsx.NewFill("solid", eg.HeaderColor, eg.HeaderColor) 
	style.Fill = fill
	for _, cell := range row.Cells {
		cell.SetStyle(style)
	}
}

func (eg *ExcelGenerator) AddDataRow(data ...interface{}) {
	row := eg.Sheet.AddRow()
	for idx, d := range data {
		cell := row.AddCell()
		cell.Value = fmt.Sprintf("%v", d)
		cell.SetStyle(eg.getCellBorderStyle())

		textLength := len(fmt.Sprintf("%v", d))
		eg.adjustColumnWidth(idx, textLength)
	}
}


// getHeaderBorderStyle returns a style with header cell borders.
func (eg *ExcelGenerator) getHeaderBorderStyle() *xlsx.Style {
	style := xlsx.NewStyle()
	border := xlsx.Border{
		Left:   "thin",
		Right:  "thin",
		Top:    "thin",
		Bottom: "thin",
	}
	style.Border = border
	return style
}


// getCellBorderStyle returns a style with data cell borders.
func (eg *ExcelGenerator) getCellBorderStyle() *xlsx.Style {
	style := xlsx.NewStyle()
	border := xlsx.Border{
		Left:   "thin",
		Right:  "thin",
		Top:    "thin",
		Bottom: "thin",
	}
	style.Border = border
	return style
}

// getTitleBorderStyle returns a style with title cell borders.
func (eg *ExcelGenerator) getTitleBorderStyle() *xlsx.Style {
	style := xlsx.NewStyle()
	border := xlsx.Border{
		Left:   "thin",
		Right:  "thin",
		Top:    "thin",
		Bottom: "thin",
	}
	style.Border = border
	return style
}


// adjustColumnWidth adjusts the column width based on the text length.
func (eg *ExcelGenerator) adjustColumnWidth(colIndex int, textLength int) {
	colWidth := float64(textLength) * eg.ColMultiplier 
	eg.Sheet.Col(colIndex).Width = colWidth
}


// getCellBackgroundColorStyle returns a style with a specific background color.
func (eg *ExcelGenerator) getCellBackgroundColorStyle() *xlsx.Style {
	style := xlsx.NewStyle()
	style.Fill = *xlsx.NewFill(eg.CellPatterType, eg.CellFgColor, eg.CellBgColor)
	return style
}


// SaveToFile writes the Excel file to the response
func (eg *ExcelGenerator) SaveToFile(w http.ResponseWriter, fileName string) error {
	var buf bytes.Buffer
	err := eg.File.Write(&buf)
	if err != nil {
		http.Error(w, "Error writing Excel data", http.StatusInternalServerError)
		return err
	}
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	if _, err := w.Write(buf.Bytes()); err != nil {
		http.Error(w, "Error sending Excel file", http.StatusInternalServerError)
		return err
	}
	return nil
}