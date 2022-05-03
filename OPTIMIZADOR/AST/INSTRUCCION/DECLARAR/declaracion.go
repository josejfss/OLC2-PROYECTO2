package declarar

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
)

type Declaracion struct {
	Temporal string
	Valor    interfaz.Expresion
}

func Ndeclaracion(tem string, val interfaz.Expresion) Declaracion {
	decla := Declaracion{Temporal: tem, Valor: val}
	return decla
}

func (decla Declaracion) Optimizar_Instruccion(block *objeto.Bloque) interface{} {
	valopt := decla.Valor.Optimizar_Expresion(block)
	simbdecla := objeto.ObjetoBloque{
		Operacion: valopt.Operacion,
		Temporal:  decla.Temporal,
		Opiz:      valopt.Opiz,
		Opde:      valopt.Opde,
		Valor:     valopt.Valor,
	}
	block.Guardar_Declaracion(decla.Temporal, simbdecla)
	return 0
}
