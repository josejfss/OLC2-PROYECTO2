package menor

import (
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
)

type OperacionMenor struct {
	OpIzquierda     inter.Expresion
	OperadorDerecha inter.Expresion
}

func Noperacionmenor(oi inter.Expresion, od inter.Expresion) OperacionMenor {
	menor := OperacionMenor{OpIzquierda: oi, OperadorDerecha: od}
	return menor
}

func (menor OperacionMenor) Opti_Expresion(block *bloque.BloquesOpt) bloque.ObjetoBloque {
	operadorizq := menor.OpIzquierda.Opti_Expresion(block)
	operadorder := menor.OperadorDerecha.Opti_Expresion(block)

	return bloque.ObjetoBloque{Tipo: bloque.OPERACION,
		Asignacion: "",
		OpeIz:      operadorizq.Valor,
		Ope:        "<",
		OpeDe:      operadorder.Valor,
		Valor:      operadorizq.Valor + "<" + operadorder.Valor}
}
