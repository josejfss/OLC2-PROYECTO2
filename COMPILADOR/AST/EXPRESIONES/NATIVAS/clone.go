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

type Clone struct {
	Val     interfaces.Expresion
	Linea   int
	Columna int
}

func Nclone(v interfaces.Expresion, l int, c int) Clone {
	clon := Clone{Val: v, Linea: l, Columna: c}
	return clon
}

func (clon Clone) Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor {
	resultado := clon.Val.Ejecutar_Expresion(ent)
	if resultado.Valor != nil {
		return simbolos.Valor{Tipo: resultado.Tipo, Valor: resultado.Valor}
	} else {
		t := time.Now()
		reportes.Nerror("No es posible realizar la funcion: clone()", ent.Nombre_Entorno, strconv.Itoa(clon.Linea), strconv.Itoa(clon.Columna), t.Format("2006-01-02 15:04:05"))
	}
	return simbolos.Valor{Tipo: simbolos.NULL, Valor: -1}
}

func (clon Clone) Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D {
	resultado := clon.Val.Compilar_Expresion(ent, gen)
	if resultado.Valor != "" {
		return simbolos.ValoresC3D{Valor: resultado.Valor, EsTemporal: resultado.EsTemporal, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
	} else {
		gen.AgregarError("ERROR-TIPOS--CLONE", strconv.Itoa(clon.Linea), strconv.Itoa(clon.Columna))
	}
	return simbolos.ValoresC3D{Valor: "0", EsTemporal: false, Tipo: simbolos.INTEGER, Label_verdadera: "", Label_false: ""}
}
