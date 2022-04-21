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

type ToOwned struct {
	Val     interfaces.Expresion
	Linea   int
	Columna int
}

func Ntoowned(v interfaces.Expresion, l int, c int) ToOwned {
	towned := ToOwned{Val: v, Linea: l, Columna: c}
	return towned
}

func (towned ToOwned) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	resultado := towned.Val.Ejecutar_Expresion(ent)
	if resultado.Tipo == simbolos.YTEXTO {
		return simbolos.Valor{Tipo: simbolos.TEXTO, Valor: resultado.Valor}
	} else {
		t := time.Now()
		reportes.Nerror("No es posible realizar la funcion: to_string()", ent.Nombre_Entorno, strconv.Itoa(towned.Linea), strconv.Itoa(towned.Columna), t.Format("2006-01-02 15:04:05"))
	}
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (towned ToOwned) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
