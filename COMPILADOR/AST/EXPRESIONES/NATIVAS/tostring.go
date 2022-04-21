package nativas

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
	reportes "OLC2-PROYECTO2/COMPILADOR/REPORTES"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"strconv"
	"time"
)

type ToString struct {
	Val     interfaces.Expresion
	Linea   int
	Columna int
}

func Ntostring(v interfaces.Expresion, l int, c int) ToString {
	tostring := ToString{Val: v, Linea: l, Columna: c}
	return tostring
}

func (tstring ToString) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	resultado := tstring.Val.Ejecutar_Expresion(ent)
	if resultado.Tipo == simbolos.YTEXTO {
		return simbolos.Valor{Tipo: simbolos.TEXTO, Valor: resultado.Valor}
	} else {
		t := time.Now()
		reportes.Nerror("No es posible realizar la funcion: to_string()", ent.Nombre_Entorno, strconv.Itoa(tstring.Linea), strconv.Itoa(tstring.Columna), t.Format("2006-01-02 15:04:05"))
	}
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (tstring ToString) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
