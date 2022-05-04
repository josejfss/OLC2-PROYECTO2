package mayorigual

import (
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
)

type OperacionMayorIgual struct {
	OpIzquierda     inter.Expresion
	OperadorDerecha inter.Expresion
}

func Noperacionmayorigual(oi inter.Expresion, od inter.Expresion) OperacionMayorIgual {
	mayor := OperacionMayorIgual{OpIzquierda: oi, OperadorDerecha: od}
	return mayor
}

func (mayor OperacionMayorIgual) Opti_Expresion(block *bloque.BloquesOpt) bloque.ObjetoBloque {
	operadorizq := mayor.OpIzquierda.Opti_Expresion(block)
	operadorder := mayor.OperadorDerecha.Opti_Expresion(block)

	return bloque.ObjetoBloque{Tipo: bloque.OPERACION,
		Asignacion: "",
		OpeIz:      operadorizq.Valor,
		Ope:        ">=",
		OpeDe:      operadorder.Valor,
		Valor:      operadorizq.Valor + ">=" + operadorder.Valor}
}
