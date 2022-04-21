package reportes

import (
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

//MATRIZ PARA GUARDAR LOS SIMBOLOS ENCONTRADOS
var ListaSimbolos [][]string

//CONTADOR DE SIMBOLOS AGREGADOS
var contador int = 0

//METODO PARA AGREGAR SIMBOLOS A LA MATRIZ
func ReporteSimbolos(id string, tisim string, noment string, par string, li string, co string) {
	con := strconv.Itoa(contador)
	ListaSimbolos = append(ListaSimbolos, []string{con, id, tisim, noment, par, li, co})
	contador += 1
}

//FUNCION PARA DEVOLVER EL TIPO EN STRING
func ReportObteniendoSimbolos(t simbolos.TipoExpresion) string {
	switch t {
	case simbolos.INTEGER:
		{
			return "int"
		}
	case simbolos.FLOAT:
		{
			return "float"
		}
	case simbolos.BOOLEAN:
		{
			return "bool"
		}
	case simbolos.TEXTO:
		{
			return "String"
		}
	case simbolos.YTEXTO:
		{
			return "&str"
		}
	case simbolos.STRUCT:
		{
			return "struct"
		}
	}
	return ""
}

//METODO PARA CREAR EL REPORTE DE LA TABLA DE SIMBOLOS
func CreandoReporteSimbolos() {
	hdr := []string{"NO.", "ID", "T. SIMBOLO", "NOM. ENT.", "PARAMETROS", "LIN.", "COL."}
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Times", "B", 24)
	pdf.Cell(40, 10, "Reporte Tabla Simbolos")
	pdf.Ln(12)

	pdf.SetFont("Times", "B", 16)
	pdf.SetFillColor(240, 240, 240)
	for i, str := range hdr {
		if i == 0 {
			pdf.CellFormat(15, 10, str, "1", 0, "C", true, 0, "")
		} else if i == 6 {
			pdf.CellFormat(15, 10, str, "1", 0, "C", true, 0, "")
		} else if i == 5 {
			pdf.CellFormat(15, 10, str, "1", 0, "C", true, 0, "")
		} else {
			pdf.CellFormat(45, 10, str, "1", 0, "C", true, 0, "")
		}
	}
	pdf.Ln(-1)

	pdf.SetFont("Times", "", 16)
	pdf.SetFillColor(255, 255, 255)
	align := []string{"C", "C", "C", "C", "C", "C", "C", "C"}
	for _, line := range ListaSimbolos {
		for i, str := range line {
			if i == 0 {
				pdf.CellFormat(15, 10, str, "1", 0, align[i], false, 0, "")
			} else if i == 6 {
				pdf.CellFormat(15, 10, str, "1", 0, align[i], false, 0, "")
			} else if i == 5 {
				pdf.CellFormat(15, 10, str, "1", 0, align[i], false, 0, "")
			} else {
				pdf.CellFormat(45, 10, str, "1", 0, align[i], false, 0, "")
			}
		}
		pdf.Ln(-1)
	}
	pdf.OutputFileAndClose("/home/iovana/go/src/OLC2-PROYECTO2/REP/reporte_simbolos.pdf")
}
