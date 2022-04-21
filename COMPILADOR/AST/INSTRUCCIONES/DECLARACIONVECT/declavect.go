package declaracionvect

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"strconv"
	"time"

	"github.com/colegno/arraylist"
)

type DeclaVector struct {
	Mutable       bool
	Identificador string
	Dimension     *arraylist.List
	ValExp        interfaces.Expresion
	Linea         int
	Columna       int
}

func Ndeclavector(m bool, id string, di *arraylist.List, ve interfaces.Expresion, l int, c int) DeclaVector {
	declavect := DeclaVector{Mutable: m, Identificador: id, Dimension: di, ValExp: ve, Linea: l, Columna: c}
	return declavect
}

func (dvect DeclaVector) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	resultado := dvect.ValExp.Ejecutar_Expresion(ent)
	tipvect := dvect.Dimension.GetValue(0).(interfaces.Expresion)
	tipovect := tipvect.Ejecutar_Expresion(ent)
	if resultado.Tipo == simbolos.VECTOR {
		if !ent.ExisteAcual_Variable(dvect.Identificador) {
			if !ent.ExisteAcual_ArreVect(dvect.Identificador) {
				val := resultado.Valor.(simbolos.Valor)
				obvect := val.Valor.(simbolos.Simbolo_ArreVect)
				if obvect.Lintdim.Len() != 1 {
					if dvect.Dimension.Len() == obvect.Lintdim.GetValue(0) {
						if tipovect.Tipo == obvect.TipoVect {
							ent.Guardar_ArreVect(dvect.Mutable, obvect.TipoVect, obvect.Nombre, obvect.ValorArreVect, obvect.Lintdim.Clone(), obvect.Lintcap.Clone(), dvect.Linea, dvect.Columna)
							ent.ListaTodo.Add(simbolos.SimboloTodo{Tipo: "arrevect", Nombre: dvect.Identificador})
							reportes.ReporteSimbolos(dvect.Identificador, "vector--"+reportes.ReportObteniendoSimbolos(obvect.TipoVect), ent.Nombre_Entorno, "--", strconv.Itoa(dvect.Linea), strconv.Itoa(dvect.Columna))
						} else if tipovect.Tipo == simbolos.STRUCT {
							ent.Guardar_ArreVect(dvect.Mutable, obvect.TipoVect, dvect.Identificador, obvect.ValorArreVect, obvect.Lintdim.Clone(), obvect.Lintcap.Clone(), dvect.Linea, dvect.Columna)
							ent.ListaTodo.Add(simbolos.SimboloTodo{Tipo: "arrevect", Nombre: dvect.Identificador})
							reportes.ReporteSimbolos(dvect.Identificador, "vector--"+reportes.ReportObteniendoSimbolos(tipovect.Tipo), ent.Nombre_Entorno, "--", strconv.Itoa(dvect.Linea), strconv.Itoa(dvect.Columna))
						} else {
							t := time.Now()
							reportes.Nerror("Los tipos no coinciden en los vectores.", ent.Nombre_Entorno, strconv.Itoa(dvect.Linea), strconv.Itoa(dvect.Columna), t.Format("2006-01-02 15:04:05"))
						}

					} else {
						t := time.Now()
						reportes.Nerror("Las dimensiones de los vectores no coinciden.", ent.Nombre_Entorno, strconv.Itoa(dvect.Linea), strconv.Itoa(dvect.Columna), t.Format("2006-01-02 15:04:05"))
					}
				} else {
					if dvect.Dimension.Len() == 1 {
						if tipovect.Tipo == obvect.TipoVect {
							ent.Guardar_ArreVect(dvect.Mutable, obvect.TipoVect, obvect.Nombre, obvect.ValorArreVect, obvect.Lintdim.Clone(), obvect.Lintcap.Clone(), dvect.Linea, dvect.Columna)
							ent.ListaTodo.Add(simbolos.SimboloTodo{Tipo: "arrevect", Nombre: dvect.Identificador})
							reportes.ReporteSimbolos(dvect.Identificador, "vector--"+reportes.ReportObteniendoSimbolos(obvect.TipoVect), ent.Nombre_Entorno, "--", strconv.Itoa(dvect.Linea), strconv.Itoa(dvect.Columna))
						} else if tipovect.Tipo == simbolos.STRUCT {
							ent.Guardar_ArreVect(dvect.Mutable, obvect.TipoVect, dvect.Identificador, obvect.ValorArreVect, obvect.Lintdim.Clone(), obvect.Lintcap.Clone(), dvect.Linea, dvect.Columna)
							ent.ListaTodo.Add(simbolos.SimboloTodo{Tipo: "arrevect", Nombre: dvect.Identificador})
							reportes.ReporteSimbolos(dvect.Identificador, "vector--"+reportes.ReportObteniendoSimbolos(tipovect.Tipo), ent.Nombre_Entorno, "--", strconv.Itoa(dvect.Linea), strconv.Itoa(dvect.Columna))
						} else {
							t := time.Now()
							reportes.Nerror("Los tipos no coinciden en los vectores.", ent.Nombre_Entorno, strconv.Itoa(dvect.Linea), strconv.Itoa(dvect.Columna), t.Format("2006-01-02 15:04:05"))
						}
					} else {
						t := time.Now()
						reportes.Nerror("Las dimensiones de los vectores no coinciden.", ent.Nombre_Entorno, strconv.Itoa(dvect.Linea), strconv.Itoa(dvect.Columna), t.Format("2006-01-02 15:04:05"))
					}
				}

			} else {
				t := time.Now()
				reportes.Nerror("Vectores con el mismo nombre de un arreglo o vector.", ent.Nombre_Entorno, strconv.Itoa(dvect.Linea), strconv.Itoa(dvect.Columna), t.Format("2006-01-02 15:04:05"))
			}
		} else {
			t := time.Now()
			reportes.Nerror("Variables con el mismo nombre en el entorno.", ent.Nombre_Entorno, strconv.Itoa(dvect.Linea), strconv.Itoa(dvect.Columna), t.Format("2006-01-02 15:04:05"))
		}
	} else {
		t := time.Now()
		reportes.Nerror("No se puede declarar vector.", ent.Nombre_Entorno, strconv.Itoa(dvect.Linea), strconv.Itoa(dvect.Columna), t.Format("2006-01-02 15:04:05"))
	}
	return 0
}

func (dvect DeclaVector) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {
	return 0
}
