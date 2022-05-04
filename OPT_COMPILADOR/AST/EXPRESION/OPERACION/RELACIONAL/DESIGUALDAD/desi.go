package desigualdad

import (
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
)

type OperacionDesigualdad struct {
	OpIzquierda     inter.Expresion
	OperadorDerecha inter.Expresion
}

func Noperaciondesigualdad(oi inter.Expresion, od inter.Expresion) OperacionDesigualdad {
	desi := OperacionDesigualdad{OpIzquierda: oi, OperadorDerecha: od}
	return desi
}

func (desi OperacionDesigualdad) Opti_Expresion(block *bloque.BloquesOpt) bloque.ObjetoBloque {
	operadorizq := desi.OpIzquierda.Opti_Expresion(block)
	operadorder := desi.OperadorDerecha.Opti_Expresion(block)

	return bloque.ObjetoBloque{Tipo: bloque.OPERACION,
		Asignacion: "",
		OpeIz:      operadorizq.Valor,
		Ope:        "!=",
		OpeDe:      operadorder.Valor,
		Valor:      operadorizq.Valor + "!=" + operadorder.Valor}
}
