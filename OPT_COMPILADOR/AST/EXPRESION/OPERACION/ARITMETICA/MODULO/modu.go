package modulo

import (
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
)

type OperacionModulo struct {
	OpIzquierda     inter.Expresion
	OperadorDerecha inter.Expresion
}

func Noperacionmodulo(oi inter.Expresion, od inter.Expresion) OperacionModulo {
	modu := OperacionModulo{OpIzquierda: oi, OperadorDerecha: od}
	return modu
}

func (modu OperacionModulo) Opti_Expresion(block *bloque.BloquesOpt) bloque.ObjetoBloque {
	operadorizq := modu.OpIzquierda.Opti_Expresion(block)
	operadorder := modu.OperadorDerecha.Opti_Expresion(block)

	return bloque.ObjetoBloque{Tipo: bloque.OPERACION,
		Asignacion: "",
		OpeIz:      operadorizq.Valor,
		Ope:        "%",
		OpeDe:      operadorder.Valor,
		Valor:      operadorizq.Valor + "%" + operadorder.Valor}
}
