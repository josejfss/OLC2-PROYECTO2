package aritmetica

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

type OpDivision struct {
	OpIz    interfaces.Expresion
	OpDe    interfaces.Expresion
	Linea   int
	Columna int
}

func Nopdivision(oi interfaces.Expresion, od interfaces.Expresion, l int, c int) OpDivision {
	division := OpDivision{OpIz: oi, OpDe: od, Linea: l, Columna: c}
	return division
}

func (division OpDivision) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	//INTERGER			FLOAT			BOOLEAN			CHAR		TEXTO			YTEXTO			NULL
	division_dominante := [7][7]simbolos.TipoExpresion{
		//INTEGER
		{simbolos.INTEGER, simbolos.FLOAT, simbolos.NULL, simbolos.INTEGER, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//FLOAT
		{simbolos.FLOAT, simbolos.FLOAT, simbolos.NULL, simbolos.FLOAT, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//BOOLEAN
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//CHAR
		{simbolos.INTEGER, simbolos.FLOAT, simbolos.NULL, simbolos.CHAR, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//TEXTO
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//&TEXTO
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//NULL
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
	}

	var operador_izquierda simbolos.Valor = division.OpIz.Ejecutar_Expresion(ent)
	var operador_derecha simbolos.Valor = division.OpDe.Ejecutar_Expresion(ent)

	var tipo_dominante simbolos.TipoExpresion = division_dominante[operador_izquierda.Tipo][operador_derecha.Tipo]

	//SE VERIFICA QUE SEA ENTERO
	if tipo_dominante == simbolos.INTEGER {
		//SE HACE CONVERSIONES
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_izquierda.Valor), 64)
		val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_derecha.Valor), 64)
		i := int(val1)
		j := int(val2)
		if j != 0 {
			//SE RETORNA SIMBOLO
			return simbolos.Valor{Tipo: tipo_dominante, Valor: i / j}
		} else {
			t := time.Now()
			reportes.Nerror("No es posible dividir entre 0", ent.Nombre_Entorno, strconv.Itoa(division.Linea), strconv.Itoa(division.Columna), t.Format("2006-01-02 15:04:05"))
		}
		//SE VERIFICA QUE SEA DECIMAL
	} else if tipo_dominante == simbolos.FLOAT {
		//SE HACE CONVERSIONES
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_izquierda.Valor), 64)
		val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_derecha.Valor), 64)
		if val2 != 0 {
			//SE RETORNA SIMBOLO
			return simbolos.Valor{Tipo: tipo_dominante, Valor: val1 / val2}
		} else {
			t := time.Now()
			reportes.Nerror("No es posible dividir entre 0", ent.Nombre_Entorno, strconv.Itoa(division.Linea), strconv.Itoa(division.Columna), t.Format("2006-01-02 15:04:05"))
		}
	} else {
		//ERROR
		t := time.Now()
		reportes.Nerror("No es posible realizar la operacion: Division", ent.Nombre_Entorno, strconv.Itoa(division.Linea), strconv.Itoa(division.Columna), t.Format("2006-01-02 15:04:05"))
	}

	//NO SE ACEPTO LOS TIPOS
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (division OpDivision) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	//INTERGER			FLOAT			BOOLEAN			CHAR		TEXTO			YTEXTO			NULL
	division_dominante := [7][7]simbolos.TipoExpresion{
		//INTEGER
		{simbolos.INTEGER, simbolos.FLOAT, simbolos.NULL, simbolos.INTEGER, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//FLOAT
		{simbolos.FLOAT, simbolos.FLOAT, simbolos.NULL, simbolos.FLOAT, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//BOOLEAN
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//CHAR
		{simbolos.INTEGER, simbolos.FLOAT, simbolos.NULL, simbolos.CHAR, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//TEXTO
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//&TEXTO
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//NULL
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
	}

	var operador_izquierda simbolos.ValoresC3D = division.OpIz.Compilar_Expresion(ent, gen)
	var operador_derecha simbolos.ValoresC3D = division.OpDe.Compilar_Expresion(ent, gen)

	var tipo_dominante simbolos.TipoExpresion = division_dominante[operador_izquierda.Tipo][operador_derecha.Tipo]

	nuevo_temporal := gen.Crear_temporal()
	etiqueta_verdad := gen.Crear_label()
	etiqueta_salida := gen.Crear_label()

	if tipo_dominante == simbolos.INTEGER {
		gen.Agregar_Logica("if (" + operador_derecha.Valor + " != 0) { goto " + etiqueta_verdad + ";}")
		gen.AgregarErrorMate(strconv.Itoa(division.Linea), strconv.Itoa(division.Columna))
		gen.Agregar_Logica(nuevo_temporal + " = 0;")
		gen.Agregar_Logica("goto " + etiqueta_salida + ";")
		gen.Agregar_label(etiqueta_verdad)
		gen.Eliminar_label(etiqueta_verdad)
		gen.Agregar_Expression(nuevo_temporal, operador_izquierda.Valor, "/", operador_derecha.Valor)
		gen.Agregar_label(etiqueta_salida)
		gen.Eliminar_label(etiqueta_salida)
		return simbolos.ValoresC3D{Valor: nuevo_temporal, EsTemporal: true, Tipo: tipo_dominante, Label_verdadera: "", Label_false: ""}
	} else if tipo_dominante == simbolos.FLOAT {
		gen.Agregar_Logica("if (" + operador_derecha.Valor + " != 0) { goto " + etiqueta_verdad + ";}")
		gen.AgregarErrorMate(strconv.Itoa(division.Linea), strconv.Itoa(division.Columna))
		gen.Agregar_Logica(nuevo_temporal + " = 0;")
		gen.Agregar_Logica("goto " + etiqueta_salida + ";")
		gen.Agregar_label(etiqueta_verdad)
		gen.Eliminar_label(etiqueta_verdad)
		gen.Agregar_Expression(nuevo_temporal, operador_izquierda.Valor, "/", operador_derecha.Valor)
		gen.Agregar_label(etiqueta_salida)
		gen.Eliminar_label(etiqueta_salida)
		return simbolos.ValoresC3D{Valor: nuevo_temporal, EsTemporal: true, Tipo: tipo_dominante, Label_verdadera: "", Label_false: ""}
	} else {
		fmt.Println("ERROR, tipos")
		gen.AgregarError("ERROR-TIPOS--DIVISION", strconv.Itoa(division.Linea), strconv.Itoa(division.Columna))
	}
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
