package aritmetica

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"strconv"
	"time"
)

type OpNegativo struct {
	OpUnario interfaces.Expresion
	Linea    int
	Columna  int
}

func Nopnegativo(on interfaces.Expresion, l int, c int) OpNegativo {
	neg := OpNegativo{OpUnario: on, Linea: l, Columna: c}
	return neg
}

func (neg OpNegativo) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	//SE OBTIENE EL VALOR DEL OPERADOR UNARIO
	var unarias simbolos.Valor = neg.OpUnario.Ejecutar_Expresion(ent)
	//SI EL TIPO ES ENTERO
	if unarias.Tipo == simbolos.INTEGER {
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: unarias.Tipo, Valor: -1 * unarias.Valor.(int)}
		//SI EL TIPO ES DECIMAL
	} else if unarias.Tipo == simbolos.FLOAT {
		//SE RETORNA SIMBOLO
		return simbolos.Valor{Tipo: unarias.Tipo, Valor: -1 * unarias.Valor.(float64)}
	} else {
		//ERROR
		t := time.Now()
		reportes.Nerror("No es posible realizar la operacion: Negativo", ent.Nombre_Entorno, strconv.Itoa(neg.Linea), strconv.Itoa(neg.Columna), t.Format("2006-01-02 15:04:05"))
	}
	//SI NO CUMPLE SE DEVUELVE UN SIMBOLO DE TIPO NULL CON VALOR NULL
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (neg OpNegativo) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	var opneg simbolos.ValoresC3D = neg.OpUnario.Compilar_Expresion(ent, gen)
	nuevo_temporal := gen.Crear_temporal()
	if opneg.Tipo == simbolos.INTEGER {
		gen.Agregar_Expression(nuevo_temporal, opneg.Valor, "*", "-1")
		return simbolos.ValoresC3D{Valor: nuevo_temporal, EsTemporal: true, Tipo: opneg.Tipo, Label_verdadera: "", Label_false: ""}
	} else if opneg.Tipo == simbolos.FLOAT {
		gen.Agregar_Expression(nuevo_temporal, opneg.Valor, "*", "-1")
		return simbolos.ValoresC3D{Valor: nuevo_temporal, EsTemporal: true, Tipo: opneg.Tipo, Label_verdadera: "", Label_false: ""}
	} else {
		gen.AgregarError("ERROR-TIPOS--NEGATIVO", strconv.Itoa(neg.Linea), strconv.Itoa(neg.Columna))
	}
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
