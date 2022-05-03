package reportes

import (
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

//MATRIZ PARA GUARDAR LOS SIMBOLOS ENCONTRADOS
var ListaOptimizacion [][]string

//CONTADOR DE SIMBOLOS AGREGADOS
var contador_opt int = 1

//METODO PARA AGREGAR SIMBOLOS A LA MATRIZ
func ReporteOpti(tip string, regla string, expor string, exprop string, li string) {
	con := strconv.Itoa(contador_opt)
	ListaOptimizacion = append(ListaOptimizacion, []string{con, tip, regla, expor, exprop, li})
	contador_opt += 1
}

//METODO PARA CREAR EL REPORTE DE LA TABLA DE SIMBOLOS
func CreandoReporteOptimizacion() {
	hdr := []string{"NO.", "TIPO OPTIMIZACION", "REGLA", "EXPR. ORIGINAL.", "EXPR. OPTIMIZADA", "LIN."}
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Times", "B", 24)
	pdf.Cell(40, 10, "Reporte Tabla Optimizacion")
	pdf.Ln(12)

	pdf.SetFont("Times", "B", 16)
	pdf.SetFillColor(240, 240, 240)
	for i, str := range hdr {
		if i == 0 || i == 5 {
			pdf.CellFormat(20, 10, str, "1", 0, "C", true, 0, "")
		} else if i == 1 {
			pdf.CellFormat(65, 10, str, "1", 0, "C", true, 0, "")
		} else if i == 3 || i == 4 {
			pdf.CellFormat(65, 10, str, "1", 0, "C", true, 0, "")
		} else {
			pdf.CellFormat(45, 10, str, "1", 0, "C", true, 0, "")
		}
	}
	pdf.Ln(-1)

	pdf.SetFont("Times", "", 16)
	pdf.SetFillColor(255, 255, 255)
	align := []string{"C", "C", "C", "C", "C", "C", "C", "C"}
	for _, line := range ListaOptimizacion {
		for i, str := range line {
			if i == 0 || i == 5 {
				pdf.CellFormat(20, 10, str, "1", 0, align[i], false, 0, "")
			} else if i == 1 {
				pdf.CellFormat(65, 10, str, "1", 0, align[i], false, 0, "")
			} else if i == 3 || i == 4 {
				pdf.CellFormat(65, 10, str, "1", 0, align[i], false, 0, "")
			} else {
				pdf.CellFormat(45, 10, str, "1", 0, align[i], false, 0, "")
			}
		}
		pdf.Ln(-1)
	}
	pdf.OutputFileAndClose("/home/iovana/go/src/OLC2-PROYECTO2/REP/reporte_optimizacion.pdf")
}
