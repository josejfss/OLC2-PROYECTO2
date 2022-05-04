package division

import (
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
)

type OperacionDivision struct {
	OpIzquierda     inter.Expresion
	OperadorDerecha inter.Expresion
}

func Noperaciondivision(oi inter.Expresion, od inter.Expresion) OperacionDivision {
	divi := OperacionDivision{OpIzquierda: oi, OperadorDerecha: od}
	return divi
}

func (divi OperacionDivision) Opti_Expresion(block *bloque.BloquesOpt) bloque.ObjetoBloque {
	operadorizq := divi.OpIzquierda.Opti_Expresion(block)
	operadorder := divi.OperadorDerecha.Opti_Expresion(block)

	return bloque.ObjetoBloque{Tipo: bloque.OPERACION,
		Asignacion: "",
		OpeIz:      operadorizq.Valor,
		Ope:        "/",
		OpeDe:      operadorder.Valor,
		Valor:      operadorizq.Valor + "/" + operadorder.Valor}
}
