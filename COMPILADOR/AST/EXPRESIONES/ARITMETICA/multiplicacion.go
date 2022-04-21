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

type OpMultiplicacion struct {
	OpIzquierda interfaces.Expresion
	OpDerecha   interfaces.Expresion
	Linea       int
	Columna     int
}

func Nopmultiplicacion(oi interfaces.Expresion, od interfaces.Expresion, l int, c int) OpMultiplicacion {
	multi := OpMultiplicacion{OpIzquierda: oi, OpDerecha: od, Linea: l, Columna: c}
	return multi
}

func (multi OpMultiplicacion) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	//INTERGER			FLOAT			BOOLEAN			CHAR		TEXTO			YTEXTO			NULL
	multiplicacion_dominante := [7][7]simbolos.TipoExpresion{
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

	var operador_izquierda simbolos.Valor = multi.OpIzquierda.Ejecutar_Expresion(ent)
	var operador_derecha simbolos.Valor = multi.OpDerecha.Ejecutar_Expresion(ent)

	var tipo_dominante simbolos.TipoExpresion = multiplicacion_dominante[operador_izquierda.Tipo][operador_derecha.Tipo]
	//SE VERIFICA QUE SEA ENTERO
	if tipo_dominante == simbolos.INTEGER {
		//SE HACE CONVERSIONES
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_izquierda.Valor), 64)
		val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_derecha.Valor), 64)
		i := int(val1)
		j := int(val2)
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: tipo_dominante, Valor: i * j}
		//SE VERIFICA QUE SEA DECIMAL
	} else if tipo_dominante == simbolos.FLOAT {
		//SE HACE CONVERSIONES
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_izquierda.Valor), 64)
		val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_derecha.Valor), 64)
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: tipo_dominante, Valor: val1 * val2}
	} else {
		//ERROR
		t := time.Now()
		reportes.Nerror("No es posible realizar la opreacion: Multiplicacion", ent.Nombre_Entorno, strconv.Itoa(multi.Linea), strconv.Itoa(multi.Columna), t.Format("2006-01-02 15:04:05"))
	}
	//SI NO CUMPLE SE DEVUELVE UN SIMBOLO DE TIPO NULL CON VALOR NULL
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (multi OpMultiplicacion) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	//INTERGER			FLOAT			BOOLEAN			CHAR		TEXTO			YTEXTO			NULL
	multiplicacion_dominante := [7][7]simbolos.TipoExpresion{
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

	var operador_izquierda simbolos.ValoresC3D = multi.OpIzquierda.Compilar_Expresion(ent, gen)
	var operador_derecha simbolos.ValoresC3D = multi.OpDerecha.Compilar_Expresion(ent, gen)

	var tipo_dominante simbolos.TipoExpresion = multiplicacion_dominante[operador_izquierda.Tipo][operador_derecha.Tipo]

	nuevo_temporal := gen.Crear_temporal()

	if tipo_dominante == simbolos.INTEGER {
		gen.Agregar_Expression(nuevo_temporal, operador_izquierda.Valor, "*", operador_derecha.Valor)
		return simbolos.ValoresC3D{Valor: nuevo_temporal, EsTemporal: true, Tipo: tipo_dominante, Label_verdadera: "", Label_false: ""}
	} else if tipo_dominante == simbolos.FLOAT {
		gen.Agregar_Expression(nuevo_temporal, operador_izquierda.Valor, "*", operador_derecha.Valor)
		return simbolos.ValoresC3D{Valor: nuevo_temporal, EsTemporal: true, Tipo: tipo_dominante, Label_verdadera: "", Label_false: ""}
	} else {
		fmt.Println("ERROR, tipos")
	}
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
