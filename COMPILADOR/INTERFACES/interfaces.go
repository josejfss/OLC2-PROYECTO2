package interfaces

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	simbolos "OLC2-PROYECTO2/COMPILADOR/SIMBOLOS"
	"reflect"
)

type Expresion interface {
	Ejecutar_Expresion(ent *entorno.Entorno) simbolos.Valor
	Compilar_Expresion(ent *entorno.Entorno, gen *generador.Generador_C3D) simbolos.ValoresC3D
}

type Instruccion interface {
	Ejecutar_Instruccion(ent *entorno.Entorno, ent2 *entorno.Entorno) interface{}
	Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{}
}

func TYPEOF(v interface{}) string {
	return reflect.TypeOf(v).String()
}
