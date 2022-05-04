package multiplicacion

import (
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
)

type OperacionMultiplicacion struct {
	OpIzquierda     inter.Expresion
	OperadorDerecha inter.Expresion
}

func Noperacionmultiplicacion(oi inter.Expresion, od inter.Expresion) OperacionMultiplicacion {
	multi := OperacionMultiplicacion{OpIzquierda: oi, OperadorDerecha: od}
	return multi
}

func (multi OperacionMultiplicacion) Opti_Expresion(block *bloque.BloquesOpt) bloque.ObjetoBloque {
	operadorizq := multi.OpIzquierda.Opti_Expresion(block)
	operadorder := multi.OperadorDerecha.Opti_Expresion(block)

	return bloque.ObjetoBloque{Tipo: bloque.OPERACION,
		Asignacion: "",
		OpeIz:      operadorizq.Valor,
		Ope:        "*",
		OpeDe:      operadorder.Valor,
		Valor:      operadorizq.Valor + "*" + operadorder.Valor}
}
