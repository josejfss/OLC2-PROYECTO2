package igualdad

import (
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
)

type OperacionIgualdad struct {
	OpIzquierda     inter.Expresion
	OperadorDerecha inter.Expresion
}

func Noperacionigualdad(oi inter.Expresion, od inter.Expresion) OperacionIgualdad {
	igual := OperacionIgualdad{OpIzquierda: oi, OperadorDerecha: od}
	return igual
}

func (igual OperacionIgualdad) Opti_Expresion(block *bloque.BloquesOpt) bloque.ObjetoBloque {
	operadorizq := igual.OpIzquierda.Opti_Expresion(block)
	operadorder := igual.OperadorDerecha.Opti_Expresion(block)

	return bloque.ObjetoBloque{Tipo: bloque.OPERACION,
		Asignacion: "",
		OpeIz:      operadorizq.Valor,
		Ope:        "==",
		OpeDe:      operadorder.Valor,
		Valor:      operadorizq.Valor + "==" + operadorder.Valor}
}
