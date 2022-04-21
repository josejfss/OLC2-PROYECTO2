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

type FIn struct {
	Variable      string
	Inferior      interfaces.Expresion
	Superior      interfaces.Expresion
	Instrucciones *arraylist.List
	Linea         int
	Columna       int
}

func Nfin(v string, inf interfaces.Expresion, sup interfaces.Expresion, linst *arraylist.List, li int, co int) FIn {
	forins := FIn{Variable: v, Inferior: inf, Superior: sup, Instrucciones: linst, Linea: li, Columna: co}
	return forins
}

func (fi FIn) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	inf := fi.Inferior.Ejecutar_Expresion(ent)
	sup := fi.Superior.Ejecutar_Expresion(ent)
	if inf.Tipo == simbolos.INTEGER {
		if sup.Tipo == simbolos.INTEGER {
			linf := inf.Valor.(int)
			lsup := sup.Valor.(int)
			for i := linf; i < lsup; i++ {
				entorno_funcion := entorno.Nuevo_Entorno("for_in__"+ent.Nombre_Entorno, ent)
				var decla declaracionvar.Declaracion
				decla.Mutable = false
				decla.Identificador = fi.Variable
				decla.TipoDecla = inf.Tipo
				valores := primitivo.Nuevo_Dato_Primitivo(i, simbolos.INTEGER)
				decla.Valor_exp = valores
				decla.Linea = fi.Linea
				decla.Columna = fi.Columna
				decla.Ejecutar_Instruccion(entorno_funcion, ent2)
				for j := 0; j < fi.Instrucciones.Len(); j++ {
					actual := fi.Instrucciones.GetValue(j).(interfaces.Instruccion)
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
			reportes.Nerror("El limite superior debe de ser int", ent.Nombre_Entorno, strconv.Itoa(fi.Linea), strconv.Itoa(fi.Columna), t.Format("2006-01-02 15:04:05"))
		}
	} else {
		t := time.Now()
		reportes.Nerror("El limite inferior debe de ser int", ent.Nombre_Entorno, strconv.Itoa(fi.Linea), strconv.Itoa(fi.Columna), t.Format("2006-01-02 15:04:05"))
	}
	return 0
}

func (fi FIn) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {

	return 0
}
