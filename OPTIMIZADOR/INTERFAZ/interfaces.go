package interfaz

import (
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
	"reflect"
)

type Expresion interface {
	Optimizar_Expresion(block *objeto.Bloque) objeto.ObjetoBloque
}

type Instruccion interface {
	Optimizar_Instruccion(block *objeto.Bloque) interface{}
}

func TYPEOF(v interface{}) string {
	return reflect.TypeOf(v).String()
}
