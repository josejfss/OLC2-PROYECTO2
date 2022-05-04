package reportes

import (
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

//MATRIZ PARA GUARDAR LOS SIMBOLOS ENCONTRADOS
var ListaDB [][]string

//CONTADOR DE SIMBOLOS AGREGADOS
var contador_bas int = 1

//METODO PARA AGREGAR SIMBOLOS A LA MATRIZ
func ReporteBasesDatos(nombre string, li string) {
	con := strconv.Itoa(contador_bas)
	ListaDB = append(ListaDB, []string{con, nombre, li})
	contador_bas += 1
}

//METODO PARA CREAR EL REPORTE DE LA TABLA DE SIMBOLOS
func CreandoReporteBasesDatos() {
	hdr := []string{"NO.", "NOMBRE BASES DE DATOS", "LIN."}
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Times", "B", 24)
	pdf.Cell(40, 10, "Reporte Bases de Datos")
	pdf.Ln(12)

	pdf.SetFont("Times", "B", 16)
	pdf.SetFillColor(240, 240, 240)
	for i, str := range hdr {
		if i == 0 {
			pdf.CellFormat(20, 10, str, "1", 0, "C", true, 0, "")
		} else if i == 1 {
			pdf.CellFormat(80, 10, str, "1", 0, "C", true, 0, "")
		} else if i == 2 {
			pdf.CellFormat(20, 10, str, "1", 0, "C", true, 0, "")
		}
	}
	pdf.Ln(-1)

	pdf.SetFont("Times", "", 16)
	pdf.SetFillColor(255, 255, 255)
	align := []string{"C", "C", "C", "C", "C", "C", "C", "C"}
	for _, line := range ListaDB {
		for i, str := range line {
			if i == 0 {
				pdf.CellFormat(20, 10, str, "1", 0, align[i], false, 0, "")
			} else if i == 1 {
				pdf.CellFormat(80, 10, str, "1", 0, align[i], false, 0, "")
			} else if i == 2 {
				pdf.CellFormat(20, 10, str, "1", 0, align[i], false, 0, "")
			}
		}
		pdf.Ln(-1)
	}
	pdf.OutputFileAndClose("/home/iovana/go/src/OLC2-PROYECTO2/REP/reporte_bases_datos.pdf")
}
