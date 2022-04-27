package declaraciones

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"strconv"
	"time"
)

type Declarar struct {
	Mutable       bool
	Identificador string
	Valor_exp     interfaces.Expresion
	Linea         int
	Columna       int
}

func Ndeclarar(mut bool, id string, vexp interfaces.Expresion, l int, c int) Declarar {
	decla := Declarar{Mutable: mut, Identificador: id, Valor_exp: vexp, Linea: l, Columna: c}
	return decla
}

func (decla Declarar) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	resultado := decla.Valor_exp.Ejecutar_Expresion(ent)

	if resultado.Tipo == simbolos.ARRAY {
		if !ent.ExisteAcual_Variable(decla.Identificador) {
			if !ent.ExisteAcual_ArreVect(decla.Identificador) {
				val := resultado.Valor.(simbolos.Valor)
				obvect := val.Valor.(simbolos.Simbolo_ArreVect)
				//ent.Guardar_ArreVect(decla.Mutable, obvect.TipoVect, obvect.Nombre, obvect.ValorArreVect, obvect.Lintdim.Clone(), obvect.Lintcap.Clone(), decla.Linea, decla.Columna)
				reportes.ReporteSimbolos(decla.Identificador, "arreglo--"+reportes.ReportObteniendoSimbolos(obvect.TipoVect), ent.Nombre_Entorno, "--", strconv.Itoa(decla.Linea), strconv.Itoa(decla.Columna))
			} else {
				t := time.Now()
				reportes.Nerror("Vectores con el mismo nombre de un arreglo o vector.", ent.Nombre_Entorno, strconv.Itoa(decla.Linea), strconv.Itoa(decla.Columna), t.Format("2006-01-02 15:04:05"))
			}
		} else {
			t := time.Now()
			reportes.Nerror("Variables con el mismo nombre que el vector.", ent.Nombre_Entorno, strconv.Itoa(decla.Linea), strconv.Itoa(decla.Columna), t.Format("2006-01-02 15:04:05"))
		}
	} else if resultado.Tipo == simbolos.VECTOR {
		if !ent.ExisteAcual_Variable(decla.Identificador) {
			if !ent.ExisteAcual_ArreVect(decla.Identificador) {
				val := resultado.Valor.(simbolos.Valor)
				obvect := val.Valor.(simbolos.Simbolo_ArreVect)
				//ent.Guardar_ArreVect(decla.Mutable, obvect.TipoVect, obvect.Nombre, obvect.ValorArreVect, obvect.Lintdim.Clone(), obvect.Lintcap.Clone(), decla.Linea, decla.Columna)
				reportes.ReporteSimbolos(decla.Identificador, "vector--"+reportes.ReportObteniendoSimbolos(obvect.TipoVect), ent.Nombre_Entorno, "--", strconv.Itoa(decla.Linea), strconv.Itoa(decla.Columna))
			} else {
				t := time.Now()
				reportes.Nerror("Vectores con el mismo nombre de un arreglo o vector.", ent.Nombre_Entorno, strconv.Itoa(decla.Linea), strconv.Itoa(decla.Columna), t.Format("2006-01-02 15:04:05"))
			}
		} else {
			t := time.Now()
			reportes.Nerror("Variables con el mismo nombre que el vector.", ent.Nombre_Entorno, strconv.Itoa(decla.Linea), strconv.Itoa(decla.Columna), t.Format("2006-01-02 15:04:05"))
		}
	} else {
		if !ent.ExisteAcual_ArreVect(decla.Identificador) {
			if !ent.ExisteAcual_Variable(decla.Identificador) {
				simbguardar := simbolos.Simbolo_Vars{Mutable: decla.Mutable, TipoVariable: resultado.Tipo, NombreVariable: decla.Identificador, ValorVariable: resultado.Valor, PosicionTabla: ent.Posicion, Linea: decla.Linea, Columna: decla.Columna}
				ent.Guardar_Variable(decla.Identificador, simbguardar)
				reportes.ReporteSimbolos(decla.Identificador, "variable--"+reportes.ReportObteniendoSimbolos(resultado.Tipo), ent.Nombre_Entorno, "--", strconv.Itoa(decla.Linea), strconv.Itoa(decla.Columna))
			} else {
				t := time.Now()
				reportes.Nerror("Variables con el mismo nombre en el entorno.", ent.Nombre_Entorno, strconv.Itoa(decla.Linea), strconv.Itoa(decla.Columna), t.Format("2006-01-02 15:04:05"))
			}
		} else {
			t := time.Now()
			reportes.Nerror("Variables con el mismo nombre de un arreglo o vector.", ent.Nombre_Entorno, strconv.Itoa(decla.Linea), strconv.Itoa(decla.Columna), t.Format("2006-01-02 15:04:05"))
		}
	}

	return 0
}

func (decla Declarar) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {

	return 0
}
