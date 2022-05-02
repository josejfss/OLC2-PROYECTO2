package funciones

import (
	parametros "OLC2-PROYECTO2/COMPILADOR/AST/INSTRUCCIONES/FUNCIONES/PARAMETROS"
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	"strconv"

	"github.com/colegno/arraylist"
)

type Llamada struct {
	Nombre_Funca   string
	L_Definiciones *arraylist.List
	Linea          int
	Columna        int
}

func Nllamada(nombre string, ldef *arraylist.List, l int, col int) Llamada {
	llama := Llamada{Nombre_Funca: nombre, L_Definiciones: ldef, Linea: l, Columna: col}
	return llama
}

func (llama Llamada) Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{} {
	return 0
}

func (llama Llamada) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {
	gen.Agregar_Comentario("ENTORNO ANTICIPADO -> " + llama.Nombre_Funca)
	temppos := gen.Crear_temporal()
	gen.Agregar_Logica(temppos + " = SP + " + strconv.Itoa(ent.Tamaño) + ";")
	temp1 := gen.Crear_temporal()
	gen.Agregar_Comentario("INICIO DECLARACION DE PARAMETROS DE -> " + llama.Nombre_Funca)
	simbfunc := ent.Obtener_Funciones(llama.Nombre_Funca)
	if simbfunc.L_Para.Len() == llama.L_Definiciones.Len() {
		for i := 0; i < simbfunc.L_Para.Len(); i++ {
			actual := simbfunc.L_Para.GetValue(i).(parametros.Param)
			val := llama.L_Definiciones.GetValue(i).(interfaces.Expresion).Compilar_Expresion(ent, gen)
			gen.Agregar_Logica(temp1 + " = " + temppos + " + " + strconv.Itoa(i) + ";\t //Parametro - " + actual.Identificador)
			gen.Agregar_Logica("STACK[(int)" + temp1 + "] = " + val.Valor + ";")

		}
	}
	gen.Agregar_Comentario("FIN DECLARACION DE PARAMETROS DE -> " + llama.Nombre_Funca)
	gen.Agregar_Comentario("INICIO LLAMADA DE -> " + llama.Nombre_Funca)
	gen.Agregar_Logica("SP = SP + " + strconv.Itoa(ent.Tamaño) + ";")
	gen.Agregar_Logica(llama.Nombre_Funca + "();")
	gen.Agregar_Logica("SP = SP - " + strconv.Itoa(ent.Tamaño) + ";")
	gen.Agregar_Comentario("FIN LLAMADA DE -> " + llama.Nombre_Funca)
	return 0
}
