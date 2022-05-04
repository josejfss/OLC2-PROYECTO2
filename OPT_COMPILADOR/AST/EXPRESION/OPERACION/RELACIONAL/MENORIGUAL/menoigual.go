package menorigual

import (
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
)

type OperacionMenorIgual struct {
	OpIzquierda     inter.Expresion
	OperadorDerecha inter.Expresion
}

func Noperacionmenorigual(oi inter.Expresion, od inter.Expresion) OperacionMenorIgual {
	menorigual := OperacionMenorIgual{OpIzquierda: oi, OperadorDerecha: od}
	return menorigual
}

func (menorigual OperacionMenorIgual) Opti_Expresion(block *bloque.BloquesOpt) bloque.ObjetoBloque {
	operadorizq := menorigual.OpIzquierda.Opti_Expresion(block)
	operadorder := menorigual.OperadorDerecha.Opti_Expresion(block)

	return bloque.ObjetoBloque{Tipo: bloque.OPERACION,
		Asignacion: "",
		OpeIz:      operadorizq.Valor,
		Ope:        "<=",
		OpeDe:      operadorder.Valor,
		Valor:      operadorizq.Valor + "<=" + operadorder.Valor}
}
