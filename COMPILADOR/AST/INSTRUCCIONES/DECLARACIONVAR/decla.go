package declaracionvar

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
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
	if !ent.ExisteAcual_Variable(decla.Identificador) {
		resultado := decla.Valor_exp.Ejecutar_Expresion(ent)
		if resultado.Tipo == decla.TipoDecla {
			if ent.Existe_Variable(decla.Identificador) {

				if ent.Nombre_Entorno != ent2.Nombre_Entorno {
					ent2.Posicion = ent2.Posicion - 1
					simbguardar := simbolos.Simbolo_Vars{Mutable: decla.Mutable, TipoVariable: resultado.Tipo, NombreVariable: decla.Identificador, ValorVariable: resultado.Valor, PosicionTabla: ent2.Posicion, Linea: decla.Linea, Columna: decla.Columna}
					ent.Guardar_Variable(decla.Identificador, simbguardar)
					nombres := decla.Identificador + strconv.Itoa(ent2.ContadorTodo)
					ent2.AumentarTodo()
					ent2.Guardar_Variable(nombres, simbguardar)
					todito1 := simbolos.SimboloTodo{Tipo: "declaracion", Nombre: nombres, Pos: ent2.Posicion}
					ent2.GuardaLTodo(todito1)
				} else {
					simbguardar := simbolos.Simbolo_Vars{Mutable: decla.Mutable, TipoVariable: resultado.Tipo, NombreVariable: decla.Identificador, ValorVariable: resultado.Valor, PosicionTabla: ent2.Posicion, Linea: decla.Linea, Columna: decla.Columna}
					ent.Guardar_Variable(decla.Identificador, simbguardar)
					nombres := decla.Identificador + strconv.Itoa(ent2.ContadorTodo)
					ent2.AumentarTodo()
					ent2.Guardar_Variable(nombres, simbguardar)
					todito1 := simbolos.SimboloTodo{Tipo: "declaracion", Nombre: nombres, Pos: ent2.Posicion}
					ent2.GuardaLTodo(todito1)
				}

			} else {

				if ent.Nombre_Entorno != ent2.Nombre_Entorno {
					ent2.Posicion = ent2.Posicion - 1
					simbguardar := simbolos.Simbolo_Vars{Mutable: decla.Mutable, TipoVariable: resultado.Tipo, NombreVariable: decla.Identificador, ValorVariable: resultado.Valor, PosicionTabla: ent.Posicion, Linea: decla.Linea, Columna: decla.Columna}
					ent.Guardar_Variable(decla.Identificador, simbguardar)
					ent2.Guardar_Variable(decla.Identificador, simbguardar)
					todito := simbolos.SimboloTodo{Tipo: "declaracion", Nombre: decla.Identificador, Pos: ent2.Posicion}
					ent2.GuardaLTodo(todito)
				} else {
					simbguardar := simbolos.Simbolo_Vars{Mutable: decla.Mutable, TipoVariable: resultado.Tipo, NombreVariable: decla.Identificador, ValorVariable: resultado.Valor, PosicionTabla: ent.Posicion, Linea: decla.Linea, Columna: decla.Columna}
					ent.Guardar_Variable(decla.Identificador, simbguardar)
					ent2.Guardar_Variable(decla.Identificador, simbguardar)
					todito := simbolos.SimboloTodo{Tipo: "declaracion", Nombre: decla.Identificador, Pos: ent2.Posicion}
					ent2.GuardaLTodo(todito)
				}

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
	return 0
}

func (decla Declaracion) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {

	valtodi := ent.ListaTodo.GetValue(0).(simbolos.SimboloTodo)
	if valtodi.Tipo == "declaracion" {
		if ent.Existe_Variable(valtodi.Nombre) {
			gen.Agregar_Comentario("COMENZANDO DECLARACION")
			simdecla := ent.Obtener_Variable(valtodi.Nombre)
			vals := decla.Valor_exp.Compilar_Expresion(ent, gen)
			posi := simdecla.PosicionTabla
			temp := gen.Crear_temporal()
			gen.Agregar_Logica(temp + "= SP + " + strconv.Itoa(posi) + ";\t\t// POSICION: " + decla.Identificador)
			gen.Agregar_Stack(temp, vals.Valor)
			gen.Agregar_Comentario("FINALIZANDO DECLARACION")
		}
	}
	ent.EliminarLTodo()
	gen.LiberarTodosTemporales()
	return 0
}
