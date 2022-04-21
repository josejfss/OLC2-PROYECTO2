package reportes

import (
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

type Error struct {
	Contador    string
	Descripcion string
	Ambito      string
	Linea       string
	Columna     string
	Fecha       string
}

//¡ERROR! Argumentos no validos en la impresion"
var Data [][]string
var Contando int = 0

func Nerror(des string, am string, li string, col string, f string) {
	con := strconv.Itoa(Contando + 1)
	Data = append(Data, []string{con, des, am, li, col, f})
	Contando += 1
}

//METODO PARA CREAR REPORTE ERRORES
func Reporte_Errores() {
	hdr := []string{"NO.", "DESCRIPCION", "NOM. ENT", "LIN.", "COL.", "FECHA Y HORA"}
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Times", "B", 24)
	pdf.Cell(40, 10, "Reporte Errores")
	pdf.Ln(12)

	pdf.SetFont("Times", "B", 16)
	pdf.SetFillColor(240, 240, 240)
	for i, str := range hdr {
		if i == 0 {
			pdf.CellFormat(15, 10, str, "1", 0, "C", true, 0, "")
		} else if i == 1 {
			pdf.CellFormat(150, 10, str, "1", 0, "C", true, 0, "")
		} else if i == 5 {
			pdf.CellFormat(50, 10, str, "1", 0, "C", true, 0, "")
		} else if i == 3 {
			pdf.CellFormat(15, 10, str, "1", 0, "C", true, 0, "")
		} else if i == 4 {
			pdf.CellFormat(15, 10, str, "1", 0, "C", true, 0, "")
		} else {
			pdf.CellFormat(35, 10, str, "1", 0, "C", true, 0, "")
		}

	}
	pdf.Ln(-1)

	pdf.SetFont("Times", "", 16)
	pdf.SetFillColor(255, 255, 255)
	align := []string{"L", "L", "L", "C", "C", "C"}
	for _, line := range Data {
		for i, str := range line {
			if i == 0 {
				pdf.CellFormat(15, 10, str, "1", 0, align[i], false, 0, "")
			} else if i == 1 {
				pdf.CellFormat(150, 10, str, "1", 0, align[i], false, 0, "")
			} else if i == 5 {
				pdf.CellFormat(50, 10, str, "1", 0, align[i], false, 0, "")
			} else if i == 4 {
				pdf.CellFormat(15, 10, str, "1", 0, "C", true, 0, "")
			} else if i == 3 {
				pdf.CellFormat(15, 10, str, "1", 0, align[i], false, 0, "")
			} else {
				pdf.CellFormat(35, 10, str, "1", 0, align[i], false, 0, "")
			}

		}
		pdf.Ln(-1)
	}
	pdf.OutputFileAndClose("/home/iovana/go/src/OLC2-PROYECTO2/REP/reporte_errores.pdf")
}
