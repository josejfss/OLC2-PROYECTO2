package negativo

import (
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
)

type Negativo struct {
	Operador inter.Expresion
}

func Nnegativo(op inter.Expresion) Negativo {
	nega := Negativo{Operador: op}
	return nega
}

func (nega Negativo) Opti_Expresion(block *bloque.BloquesOpt) bloque.ObjetoBloque {
	neg := nega.Operador.Opti_Expresion(block)
	return bloque.ObjetoBloque{Tipo: neg.Tipo,
		Asignacion: neg.Asignacion,
		OpeIz:      "-" + neg.OpeIz,
		OpeDe:      "-" + neg.OpeDe,
		Ope:        "-",
		Valor:      "-" + neg.Valor,
		Linea:      0}
}
