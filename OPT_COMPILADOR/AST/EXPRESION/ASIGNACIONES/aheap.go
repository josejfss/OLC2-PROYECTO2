package asignaciones

import (
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
)

type AsiHeap struct {
	Expresion inter.Expresion
}

func Nasiheap(exp inter.Expresion) AsiHeap {
	asig := AsiHeap{Expresion: exp}
	return asig
}

func (asig AsiHeap) Opti_Expresion(block *bloque.BloquesOpt) bloque.ObjetoBloque {
	expres := asig.Expresion.Opti_Expresion(block)

	return bloque.ObjetoBloque{
		Tipo:       bloque.TEMPORAL,
		Asignacion: "",
		OpeIz:      expres.Valor,
		OpeDe:      "",
		Ope:        "HEAP",
		Valor:      expres.Valor}
}
