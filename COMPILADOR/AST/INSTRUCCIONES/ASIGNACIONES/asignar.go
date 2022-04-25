package asignaciones

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"fmt"
	"strconv"
	"time"
)

type Asignacion struct {
	Identificador string
	Valor         interfaces.Expresion
	Linea         int
	Columna       int
}

func Nasignacion(i string, v interfaces.Expresion, l int, c int) Asignacion {
	asig := Asignacion{Identificador: i, Valor: v, Linea: l, Columna: c}
	return asig
}

func (as Asignacion) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {

	if ent.Existe_Variable(as.Identificador) {
		simbolo := ent.Obtener_Variable(as.Identificador)
		exp := as.Valor.Ejecutar_Expresion(ent)
		if simbolo.Mutable {
			if simbolo.TipoVariable == exp.Tipo {
				simbolo.ValorVariable = exp.Valor
				ent.Modificar_Variable(as.Identificador, simbolo)
			}
		} else {
			t := time.Now()
			reportes.Nerror("No se puede asignar valor a una variable no mutable", ent.Nombre_Entorno, strconv.Itoa(as.Linea), strconv.Itoa(as.Columna), t.Format("2006-01-02 15:04:05"))
		}

	} else {
		t := time.Now()
		reportes.Nerror("No se puede asignar valor a una variable no declarada", ent.Nombre_Entorno, strconv.Itoa(as.Linea), strconv.Itoa(as.Columna), t.Format("2006-01-02 15:04:05"))
	}
	return 0
}

func (as Asignacion) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {
	gen.Agregar_Comentario("COMENZANDO ASIGNACION")
	if ent.Existe_Variable(as.Identificador) {
		simbolovar := ent.Obtener_Variable(as.Identificador)
		val := as.Valor.Compilar_Expresion(ent, gen)
		if simbolovar.TipoVariable == val.Tipo {
			if val.Tipo == simbolos.YTEXTO || val.Tipo == simbolos.TEXTO {
				posi := simbolovar.PosicionTabla
				temp1 := gen.Crear_temporal()
				temp2 := gen.Crear_temporal()
				impr := temp1 + " = HP;\n"
				for _, txt := range val.Valor {
					f := int(txt)
					impr += "HEAP[int(HP)] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt) + "\n"
					impr += "HP = HP + 1;\n"
				}
				impr += "HEAP[int(HP)] = -1;\nHP = HP + 1;\n"
				gen.Agregar_Logica(impr)
				gen.Agregar_Logica(temp2 + " = SP + " + strconv.Itoa(posi) + ";\t\t// POSICION: " + as.Identificador)
				gen.Agregar_Stack(temp2, temp1)
			} else {
				posi := simbolovar.PosicionTabla
				temp := gen.Crear_temporal()
				gen.Agregar_Logica(temp + "= SP + " + strconv.Itoa(posi) + ";\t\t// POSICION: " + as.Identificador)
				gen.Agregar_Stack(temp, val.Valor)
			}
		}
	}
	gen.Agregar_Comentario("FINALIZANDO ASIGNACION")
	gen.LiberarTodosTemporales()
	return 0
}
