package cforin

import (
	primitivo "OLC2-PROYECTO2/COMPILADOR/AST/EXPRESIONES/PRIMITIVO"
	declaracionvar "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/DECLARACIONVAR"
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"strconv"
	"time"

	"github.com/colegno/arraylist"
)

type FInVA struct {
	Identificador string
	Valores       interfaces.Expresion
	Instrucciones *arraylist.List
	Linea         int
	Columna       int
}

func NforinVA(id string, val interfaces.Expresion, linstr *arraylist.List, li int, col int) FInVA {
	fiva := FInVA{Identificador: id, Valores: val, Instrucciones: linstr, Linea: li, Columna: col}
	return fiva
}

func (f FInVA) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	arre := f.Valores.Ejecutar_Expresion(ent)
	if arre.Tipo == simbolos.VECTOR || arre.Tipo == simbolos.ARRAY {
		if interfaces.TYPEOF(arre.Valor) == "interfaces.SimboloArrVect" {
			arreglito := arre.Valor.(simbolos.Simbolo_ArreVect)
			for i := 0; i < len(arreglito.ValorArreVect); i++ {
				entorno_funcion := entorno.Nuevo_Entorno("for_in__"+ent.Nombre_Entorno, ent)
				var declas declaracionvar.Declaracion
				declas.Mutable = false
				declas.Identificador = f.Identificador
				vals := primitivo.Nuevo_Dato_Primitivo(arreglito.ValorArreVect[i], arreglito.TipoVect)
				declas.Valor_exp = vals
				declas.TipoDecla = arreglito.TipoVect
				declas.Linea = f.Linea
				declas.Columna = f.Columna
				declas.Ejecutar_Instruccion(entorno_funcion, ent2)
				for j := 0; j < f.Instrucciones.Len(); j++ {
					actual := f.Instrucciones.GetValue(j).(interfaces.Instruccion)
					valor_instr := actual.Ejecutar_Instruccion(entorno_funcion, ent2)

					if valor_instr != 0 {
						valores := valor_instr.(simbolos.Valor)
						if valores.Tipo == simbolos.BREAK {
							if valores.Valor == "break" {
								return 0
							} else {
								sal := valores.Valor.(simbolos.Valor)
								return sal
							}
						} else if valores.Tipo == simbolos.CONTINUE {
							break
						} else {
							valor_retorno := valor_instr.(simbolos.Valor)
							return valor_retorno
						}
					}
				}
			}
		} else if interfaces.TYPEOF(arre.Valor) == "interfaces.Valor_Expresion" {
			val := arre.Valor.(simbolos.Valor)
			obvect := val.Valor.(simbolos.Simbolo_ArreVect)
			for i := 0; i < len(obvect.ValorArreVect); i++ {
				entorno_funcion := entorno.Nuevo_Entorno("for_in__"+ent.Nombre_Entorno, ent)
				var declas declaracionvar.Declaracion
				declas.Mutable = false
				declas.Identificador = f.Identificador
				vals := primitivo.Nuevo_Dato_Primitivo(obvect.ValorArreVect[i], obvect.TipoVect)
				declas.Valor_exp = vals
				declas.TipoDecla = obvect.TipoVect
				declas.Linea = f.Linea
				declas.Columna = f.Columna
				declas.Ejecutar_Instruccion(entorno_funcion, ent2)
				for j := 0; j < f.Instrucciones.Len(); j++ {
					actual := f.Instrucciones.GetValue(j).(interfaces.Instruccion)
					valor_instr := actual.Ejecutar_Instruccion(entorno_funcion, ent2)

					if valor_instr != 0 {
						valores := valor_instr.(simbolos.Valor)
						if valores.Tipo == simbolos.BREAK {
							if valores.Valor == "break" {
								return 0
							} else {
								sal := valores.Valor.(simbolos.Valor)
								return sal
							}
						} else if valores.Tipo == simbolos.CONTINUE {
							break
						} else {
							valor_retorno := valor_instr.(simbolos.Valor)
							return valor_retorno
						}
					}
				}
			}
		} else {
			t := time.Now()
			reportes.Nerror("No se puede realizar el ciclo for", ent.Nombre_Entorno, strconv.Itoa(f.Linea), strconv.Itoa(f.Columna), t.Format("2006-01-02 15:04:05"))
		}
	} else {
		t := time.Now()
		reportes.Nerror("No se puede realizar el ciclo for", ent.Nombre_Entorno, strconv.Itoa(f.Linea), strconv.Itoa(f.Columna), t.Format("2006-01-02 15:04:05"))
	}
	return 0
}

func (w FInVA) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {

	return 0
}
