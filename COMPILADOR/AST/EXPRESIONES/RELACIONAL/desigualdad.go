package relacional

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

type OpDesigualdad struct {
	OpIzq   interfaces.Expresion
	OpDer   interfaces.Expresion
	Linea   int
	Columna int
}

func Nopdesigualdad(oi interfaces.Expresion, od interfaces.Expresion, l int, c int) OpDesigualdad {
	mayor := OpDesigualdad{OpIzq: oi, OpDer: od, Linea: l, Columna: c}
	return mayor
}

func (desigualdad OpDesigualdad) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	//		INTEGER				FLOAT			  BOOLEAN		CHAR	   		TEXTO		&TEXTO		NULL
	relacional_domiante := [7][7]simbolos.TipoExpresion{
		//INTEGER
		{simbolos.INTEGER, simbolos.NULL, simbolos.NULL, simbolos.INTEGER, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//FLOAT
		{simbolos.NULL, simbolos.FLOAT, simbolos.NULL, simbolos.FLOAT, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//BOOLEAN
		{simbolos.NULL, simbolos.NULL, simbolos.BOOLEAN, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//CHAR
		{simbolos.INTEGER, simbolos.FLOAT, simbolos.NULL, simbolos.CHAR, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//TEXTO
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.TEXTO, simbolos.NULL, simbolos.NULL},
		//&TEXTO
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.YTEXTO, simbolos.NULL},
		//NULL
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
	}

	var operador_izq simbolos.Valor = desigualdad.OpIzq.Ejecutar_Expresion(ent)
	var operador_der simbolos.Valor = desigualdad.OpDer.Ejecutar_Expresion(ent)

	var tipo_dominante simbolos.TipoExpresion = relacional_domiante[operador_izq.Tipo][operador_der.Tipo]
	//SE VERIFICA SI EL TIPO ES ENTERO
	if tipo_dominante == simbolos.INTEGER {
		//SE HACE LAS CONVERSIONESfunc (operacion Relacional) Ejecutar(ent *interfaces.Entorno) interfaces.Simbolo
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_izq.Valor), 64)
		val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_der.Valor), 64)
		i := int(val1)
		j := int(val2)
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.BOOLEAN, Valor: i != j}
		//SE VERIFICA SI ES DECIMAL
	} else if tipo_dominante == simbolos.FLOAT {
		//SE HACE LAS CONVERSIONES
		val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_izq.Valor), 64)
		val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", operador_der.Valor), 64)
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.BOOLEAN, Valor: val1 != val2}
		//SE VERIFICA SI ES STRING
	} else if tipo_dominante == simbolos.TEXTO {
		//SE HACE LAS CONVERSIONES
		l1 := fmt.Sprintf("%v", operador_izq.Valor)
		l2 := fmt.Sprintf("%v", operador_der.Valor)
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.BOOLEAN, Valor: l1 != l2}
	} else {
		//ERROR
		t := time.Now()
		reportes.Nerror("No es posible realizar la operacion: Desigualdad", ent.Nombre_Entorno, strconv.Itoa(desigualdad.Linea), strconv.Itoa(desigualdad.Columna), t.Format("2006-01-02 15:04:05"))
	}
	//SINO SE CUMPLEN LAS CONDICIONES SE RETORNA UN SIMBOLO DE TIPO NULL CON VALOR NULL
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (desigualdad OpDesigualdad) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	//		INTEGER				FLOAT			  BOOLEAN		CHAR	   		TEXTO		&TEXTO		NULL
	relacional_domiante := [7][7]simbolos.TipoExpresion{
		//INTEGER
		{simbolos.INTEGER, simbolos.NULL, simbolos.NULL, simbolos.INTEGER, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//FLOAT
		{simbolos.NULL, simbolos.FLOAT, simbolos.NULL, simbolos.FLOAT, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//BOOLEAN
		{simbolos.NULL, simbolos.NULL, simbolos.BOOLEAN, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//CHAR
		{simbolos.INTEGER, simbolos.FLOAT, simbolos.NULL, simbolos.CHAR, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//TEXTO
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.TEXTO, simbolos.NULL, simbolos.NULL},
		//&TEXTO
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.YTEXTO, simbolos.NULL},
		//NULL
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
	}

	var operador_izq simbolos.ValoresC3D = desigualdad.OpIzq.Compilar_Expresion(ent, gen)
	var operador_der simbolos.ValoresC3D = desigualdad.OpDer.Compilar_Expresion(ent, gen)

	var tipo_dominante simbolos.TipoExpresion = relacional_domiante[operador_izq.Tipo][operador_der.Tipo]

	etiqueta_verdadera := gen.Crear_label()
	etiqueta_falsa := gen.Crear_label()

	if tipo_dominante == simbolos.INTEGER {
		gen.Agregar_Relacional(operador_izq.Valor, "!=", operador_der.Valor, etiqueta_verdadera, etiqueta_falsa)
		rdesigual := "/*--INGRESANDO OPERACION DESIGUALDAD--*/\n"
		rdesigual += "if (" + operador_izq.Valor + " " + "!=" + " " + operador_der.Valor + ") { goto " + etiqueta_verdadera + "; }\ngoto " + etiqueta_falsa + ";"
		rdesigual += "\n/*--SALIENDO OPERACION DESIGUALDAD--*/\n"
		return simbolos.ValoresC3D{Valor: rdesigual, EsTemporal: false, Tipo: simbolos.BOOLEAN, Label_verdadera: etiqueta_verdadera, Label_false: etiqueta_falsa}
	} else if tipo_dominante == simbolos.FLOAT {
		gen.Agregar_Relacional(operador_izq.Valor, "!=", operador_der.Valor, etiqueta_verdadera, etiqueta_falsa)
		rdesigual := "/*--INGRESANDO OPERACION DESIGUALDAD--*/\n"
		rdesigual += "if (" + operador_izq.Valor + " " + "!=" + " " + operador_der.Valor + ") { goto " + etiqueta_verdadera + "; }\ngoto " + etiqueta_falsa + ";"
		rdesigual += "\n/*--SALIENDO OPERACION DESIGUALDAD--*/\n"
		return simbolos.ValoresC3D{Valor: rdesigual, EsTemporal: false, Tipo: simbolos.BOOLEAN, Label_verdadera: etiqueta_verdadera, Label_false: etiqueta_falsa}
	} else if tipo_dominante == simbolos.TEXTO {
		gen.Agregar_Relacional(operador_izq.Valor, "!=", operador_der.Valor, etiqueta_verdadera, etiqueta_falsa)
		rdesigual := "/*--INGRESANDO OPERACION DESIGUALDAD--*/\n"
		rdesigual += "if (" + operador_izq.Valor + " " + "!=" + " " + operador_der.Valor + ") { goto " + etiqueta_verdadera + "; }\ngoto " + etiqueta_falsa + ";"
		rdesigual += "\n/*--SALIENDO OPERACION DESIGUALDAD--*/\n"
		return simbolos.ValoresC3D{Valor: rdesigual, EsTemporal: false, Tipo: simbolos.BOOLEAN, Label_verdadera: etiqueta_verdadera, Label_false: etiqueta_falsa}
	} else {
		fmt.Println("ERROR, tipos")
		gen.AgregarError("ERROR-TIPOS--DESIGUALDAD", strconv.Itoa(desigualdad.Linea), strconv.Itoa(desigualdad.Columna))
	}
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
