package declaracionarre

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

type DeclaArre struct {
	Mutable       bool
	Identificador string
	Dimension     *arraylist.List
	ValExp        interfaces.Expresion
	Linea         int
	Columna       int
}

func Ndeclaarre(m bool, id string, dim *arraylist.List, va interfaces.Expresion, l int, c int) DeclaArre {
	arr := DeclaArre{Mutable: m, Identificador: id, Dimension: dim, ValExp: va, Linea: l, Columna: c}
	return arr
}

func (darre DeclaArre) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	resultado := darre.ValExp.Ejecutar_Expresion(ent)
	tipoare := darre.Dimension.GetValue(0).(interfaces.Expresion)
	tipoarre := tipoare.Ejecutar_Expresion(ent)
	darre.Dimension.RemoveAtIndex(0)
	if resultado.Tipo == simbolos.ARRAY {
		if !ent.ExisteAcual_Variable(darre.Identificador) {
			if !ent.ExisteAcual_ArreVect(darre.Identificador) {
				val := resultado.Valor.(simbolos.Valor)
				obvect := val.Valor.(simbolos.Simbolo_ArreVect)
				if obvect.Lintdim.Len() != 1 {
					if darre.Dimension.Len() == obvect.Lintdim.Len() {
						if tipoarre.Tipo == obvect.TipoVect {
							ent.Guardar_ArreVect(darre.Mutable, tipoarre.Tipo, darre.Identificador, obvect.ValorArreVect, obvect.Lintdim.Clone(), obvect.Lintcap.Clone(), darre.Linea, darre.Columna)
							ent.ListaTodo.Add(simbolos.SimboloTodo{Tipo: "arrevect", Nombre: darre.Identificador})
							reportes.ReporteSimbolos(darre.Identificador, "arreglo--"+reportes.ReportObteniendoSimbolos(obvect.TipoVect), ent.Nombre_Entorno, "--", strconv.Itoa(darre.Linea), strconv.Itoa(darre.Columna))
						} else if tipoarre.Tipo == simbolos.STRUCT {
							ent.Guardar_ArreVect(darre.Mutable, obvect.TipoVect, darre.Identificador, obvect.ValorArreVect, obvect.Lintdim.Clone(), obvect.Lintcap.Clone(), darre.Linea, darre.Columna)
							ent.ListaTodo.Add(simbolos.SimboloTodo{Tipo: "arrevect", Nombre: darre.Identificador})
							reportes.ReporteSimbolos(darre.Identificador, "arreglo--"+reportes.ReportObteniendoSimbolos(tipoarre.Tipo), ent.Nombre_Entorno, "--", strconv.Itoa(darre.Linea), strconv.Itoa(darre.Columna))
						} else {
							t := time.Now()
							reportes.Nerror("Los tipo no coinciden en los arreglos.", ent.Nombre_Entorno, strconv.Itoa(darre.Linea), strconv.Itoa(darre.Columna), t.Format("2006-01-02 15:04:05"))
						}

					} else {
						t := time.Now()
						reportes.Nerror("Las dimensiones de los arreglos no coinciden.", ent.Nombre_Entorno, strconv.Itoa(darre.Linea), strconv.Itoa(darre.Columna), t.Format("2006-01-02 15:04:05"))
					}
				} else {
					if darre.Dimension.Len() == 1 {
						if tipoarre.Tipo == obvect.TipoVect {
							ent.Guardar_ArreVect(darre.Mutable, obvect.TipoVect, darre.Identificador, obvect.ValorArreVect, obvect.Lintdim.Clone(), obvect.Lintcap.Clone(), darre.Linea, darre.Columna)
							ent.ListaTodo.Add(simbolos.SimboloTodo{Tipo: "arrevect", Nombre: darre.Identificador})
							reportes.ReporteSimbolos(darre.Identificador, "arreglo--"+reportes.ReportObteniendoSimbolos(obvect.TipoVect), ent.Nombre_Entorno, "--", strconv.Itoa(darre.Linea), strconv.Itoa(darre.Columna))
						} else if tipoarre.Tipo == simbolos.STRUCT {
							ent.Guardar_ArreVect(darre.Mutable, obvect.TipoVect, darre.Identificador, obvect.ValorArreVect, obvect.Lintdim.Clone(), obvect.Lintcap.Clone(), darre.Linea, darre.Columna)
							ent.ListaTodo.Add(simbolos.SimboloTodo{Tipo: "arrevect", Nombre: darre.Identificador})
							reportes.ReporteSimbolos(darre.Identificador, "arreglo--"+reportes.ReportObteniendoSimbolos(tipoarre.Tipo), ent.Nombre_Entorno, "--", strconv.Itoa(darre.Linea), strconv.Itoa(darre.Columna))
						} else {
							t := time.Now()
							reportes.Nerror("Los tipo no coinciden en los arreglos.", ent.Nombre_Entorno, strconv.Itoa(darre.Linea), strconv.Itoa(darre.Columna), t.Format("2006-01-02 15:04:05"))
						}
					} else {
						t := time.Now()
						reportes.Nerror("Las dimensiones de los arreglos no coinciden.", ent.Nombre_Entorno, strconv.Itoa(darre.Linea), strconv.Itoa(darre.Columna), t.Format("2006-01-02 15:04:05"))
					}
				}

			} else {
				t := time.Now()
				reportes.Nerror("Arreglos con el mismo nombre de un arreglo o vector.", ent.Nombre_Entorno, strconv.Itoa(darre.Linea), strconv.Itoa(darre.Columna), t.Format("2006-01-02 15:04:05"))
			}
		} else {
			t := time.Now()
			reportes.Nerror("Variables con el mismo nombre en el entorno.", ent.Nombre_Entorno, strconv.Itoa(darre.Linea), strconv.Itoa(darre.Columna), t.Format("2006-01-02 15:04:05"))
		}
	} else {
		t := time.Now()
		reportes.Nerror("No se puede declarar vector.", ent.Nombre_Entorno, strconv.Itoa(darre.Linea), strconv.Itoa(darre.Columna), t.Format("2006-01-02 15:04:05"))
	}
	return 0
}

func (darre DeclaArre) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {
	return 0
}
