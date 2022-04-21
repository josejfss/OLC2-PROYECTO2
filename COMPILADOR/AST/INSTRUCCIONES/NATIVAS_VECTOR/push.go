package nativasvector

import (
	entorno "OLC2-PROYECTO2/COMPILADOR/ENTORNO"
	generador "OLC2-PROYECTO2/COMPILADOR/GENERADOR"
	interfaces "OLC2-PROYECTO2/COMPILADOR/INTERFACES"
)

type NVPush struct {
	Identificador string
	Argumento     interfaces.Expresion
	Linea         int
	Columna       int
}

func Npush(id string, ar interfaces.Expresion, l int, c int) NVPush {
	push := NVPush{Identificador: id, Argumento: ar, Linea: l, Columna: c}
	return push
}

func (pus NVPush) Ejecutar_Instruccion(ent *entorno.Entorno) interface{} {
	// if ent.Existe_ArreVect(pus.Identificador) {
	// 	//nuevo := pus.Argumento.Ejecutar_Expresion(ent)

	// }
	return 0
}

func (pus NVPush) Compilar_Instruccion(ent *entorno.Entorno, gen *generador.Generador_C3D) interface{} {
	return 0
}
