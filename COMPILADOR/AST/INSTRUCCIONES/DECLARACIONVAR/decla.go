package declaracionvar

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

type Declaracion struct {
	Mutable       bool
	Identificador string
	TipoDecla     simbolos.TipoExpresion
	Valor_exp     interfaces.Expresion
	Linea         int
	Columna       int
}

func Ndeclaracion(m bool, id string, tip simbolos.TipoExpresion, val interfaces.Expresion, l int, col int) Declaracion {
	declas := Declaracion{Mutable: m, Identificador: id, TipoDecla: tip, Valor_exp: val, Linea: l, Columna: col}
	return declas
}

func (decla Declaracion) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	if !ent.Existe_ArreVect(decla.Identificador) {
		if !ent.Existe_Variable(decla.Identificador) {
			resultado := decla.Valor_exp.Ejecutar_Expresion(ent)
			if resultado.Tipo == decla.TipoDecla {
				if ent.Nombre_Entorno != ent2.Nombre_Entorno {
					ent2.Posicion = ent2.Posicion - 1
					simbguardar := simbolos.Simbolo_Vars{Mutable: decla.Mutable, TipoVariable: resultado.Tipo, NombreVariable: decla.Identificador, ValorVariable: resultado.Valor, PosicionTabla: ent.Posicion, Linea: decla.Linea, Columna: decla.Columna}
					ent.Guardar_Variable(decla.Identificador, simbguardar)
					ent2.Guardar_Variable(decla.Identificador, simbguardar)
					// todito := simbolos.SimboloTodo{Tipo: "declaracion", Nombre: decla.Identificador, Pos: ent2.Posicion}
					// ent2.GuardaLTodo(todito)
				} else {
					simbguardar := simbolos.Simbolo_Vars{Mutable: decla.Mutable, TipoVariable: resultado.Tipo, NombreVariable: decla.Identificador, ValorVariable: resultado.Valor, PosicionTabla: ent.Posicion, Linea: decla.Linea, Columna: decla.Columna}
					ent.Guardar_Variable(decla.Identificador, simbguardar)
					ent2.Guardar_Variable(decla.Identificador, simbguardar)
					// todito := simbolos.SimboloTodo{Tipo: "declaracion", Nombre: decla.Identificador, Pos: ent2.Posicion}
					// ent2.GuardaLTodo(todito)
				}
				reportes.ReporteSimbolos(decla.Identificador, "variable--"+reportes.ReportObteniendoSimbolos(resultado.Tipo), ent.Nombre_Entorno, "--", strconv.Itoa(decla.Linea), strconv.Itoa(decla.Columna))
			} else {
				t := time.Now()
				reportes.Nerror("No se puede declarar la variable: "+decla.Identificador+" con distinto tipo", ent.Nombre_Entorno, strconv.Itoa(decla.Linea), strconv.Itoa(decla.Columna), t.Format("2006-01-02 15:04:05"))
			}
		} else {
			t := time.Now()
			reportes.Nerror("Variables con el mismo nombre en el entorno.", ent.Nombre_Entorno, strconv.Itoa(decla.Linea), strconv.Itoa(decla.Columna), t.Format("2006-01-02 15:04:05"))
		}
	}

	return 0
}

func (decla Declaracion) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {

	if ent.Existe_Variable(decla.Identificador) {
		gen.Agregar_Comentario("COMENZANDO DECLARACION")
		simdecla := ent.Obtener_Variable(decla.Identificador)
		vals := decla.Valor_exp.Compilar_Expresion(ent, gen)
		if vals.Tipo == simbolos.YTEXTO || vals.Tipo == simbolos.TEXTO {
			posi := simdecla.PosicionTabla
			temp1 := gen.Crear_temporal()
			temp2 := gen.Crear_temporal()
			impr := temp1 + " = HP;\n"
			for _, txt := range vals.Valor {
				f := int(txt)
				impr += "HEAP[int(HP)] = " + fmt.Sprintf("%v", f) + "; //LETRA-> " + string(txt) + "\n"
				impr += "HP = HP + 1;\n"
			}
			impr += "HEAP[int(HP)] = -1;\nHP = HP + 1;\n"
			gen.Agregar_Logica(impr)
			gen.Agregar_Logica(temp2 + " = SP + " + strconv.Itoa(posi) + ";\t\t// POSICION: " + decla.Identificador)
			gen.Agregar_Stack(temp2, temp1)
		} else {
			posi := simdecla.PosicionTabla
			temp := gen.Crear_temporal()
			gen.Agregar_Logica(temp + "= SP + " + strconv.Itoa(posi) + ";\t\t// POSICION: " + decla.Identificador)
			gen.Agregar_Stack(temp, vals.Valor)
		}

		gen.Agregar_Comentario("FINALIZANDO DECLARACION")
	}
	gen.LiberarTodosTemporales()
	return 0
}
