package inter

import (
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	generando "OLC2-PROYECTO2/OPT_COMPILADOR/GENERANDO"
	"reflect"
)

type Expresion interface {
	Opti_Expresion(block *bloque.BloquesOpt) bloque.ObjetoBloque
}

type Instruccion interface {
	Opti_Instruccion(block *bloque.BloquesOpt, genopt *generando.GenerandoOpti) interface{}
}

func TYPEOF(v interface{}) string {
	return reflect.TypeOf(v).String()
}
