package logica

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"strconv"
	"time"
)

type OpOr struct {
	OpIzq   interfaces.Expresion
	OpDer   interfaces.Expresion
	Linea   int
	Columna int
}

func Noor(oi interfaces.Expresion, od interfaces.Expresion, l int, c int) OpOr {
	or := OpOr{OpIzq: oi, OpDer: od, Linea: l, Columna: c}
	return or
}

func (or OpOr) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	//		INTEGER				FLOAT			  BOOLEAN		CHAR	   		TEXTO		&TEXTO		NULL
	logica_domiante := [7][7]simbolos.TipoExpresion{
		//INTEGER
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//FLOAT
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//BOOLEAN
		{simbolos.NULL, simbolos.NULL, simbolos.BOOLEAN, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//CHAR
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//TEXTO
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//&TEXTO
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//NULL
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
	}

	//OPERADORES SE EJECUTAN PARA OBTENER SU VALOR
	var operador_izq simbolos.Valor = or.OpIzq.Ejecutar_Expresion(ent)
	var operador_der simbolos.Valor = or.OpDer.Ejecutar_Expresion(ent)

	//TIPO QUE TENDRA EL RESULTADO AL FINAL DE LA
	//SE OBITENE EL TIPO DE LA RELACION
	tipo_dominante := logica_domiante[operador_izq.Tipo][operador_der.Tipo]
	//SE VERIFICA SI EL TIPO ES ENTERO
	if tipo_dominante == simbolos.BOOLEAN {
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: simbolos.BOOLEAN, Valor: operador_izq.Valor.(bool) || operador_der.Valor.(bool)}
	} else {
		//ERROR
		t := time.Now()
		reportes.Nerror("No es posible realizar la operacion: Or", ent.Nombre_Entorno, strconv.Itoa(or.Linea), strconv.Itoa(or.Columna), t.Format("2006-01-02 15:04:05"))
	}

	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (or OpOr) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	//		INTEGER				FLOAT			  BOOLEAN		CHAR	   		TEXTO		&TEXTO		NULL
	logica_domiante := [7][7]simbolos.TipoExpresion{
		//INTEGER
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//FLOAT
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//BOOLEAN
		{simbolos.NULL, simbolos.NULL, simbolos.BOOLEAN, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//CHAR
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//TEXTO
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//&TEXTO
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
		//NULL
		{simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL, simbolos.NULL},
	}

	var operador_izq simbolos.ValoresC3D = or.OpIzq.Compilar_Expresion(ent, gen)
	pos := gen.Obtener_TamCodigo()
	gen.Obtener_codigo().RemoveAtIndex(pos - 1)
	var operador_der simbolos.ValoresC3D = or.OpDer.Compilar_Expresion(ent, gen)
	pos1 := gen.Obtener_TamCodigo()
	gen.Obtener_codigo().RemoveAtIndex(pos1 - 1)

	var tipo_dominante simbolos.TipoExpresion = logica_domiante[operador_izq.Tipo][operador_der.Tipo]

	//etiqueta_verdadera := gen.Crear_label()
	//etiqueta_falsa := gen.Crear_label()

	if tipo_dominante == simbolos.BOOLEAN {
		ors := "/*--INGRESANDO OPERACION OR--*/\n"
		ors += operador_izq.Valor + "\n" + operador_izq.Label_false + ":\n" + operador_der.Valor
		gen.Eliminar_label(operador_izq.Label_false)
		ors += "\n/*--SALIENDO OPERACION OR--*/\n"
		gen.Agregar_Logica(ors)
		return simbolos.ValoresC3D{Valor: ors, EsTemporal: false, Tipo: simbolos.BOOLEAN, Label_verdadera: operador_der.Label_verdadera, Label_false: operador_der.Label_false}
	} else {
		gen.AgregarError("ERROR-TIPOS--OR", strconv.Itoa(or.Linea), strconv.Itoa(or.Columna))
	}
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
