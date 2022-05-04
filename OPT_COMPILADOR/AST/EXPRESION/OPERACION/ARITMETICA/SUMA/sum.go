package suma

import (
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
)

type OperacionSuma struct {
	OpIzquierda     inter.Expresion
	OperadorDerecha inter.Expresion
}

func Noperacionsuma(oi inter.Expresion, od inter.Expresion) OperacionSuma {
	suma := OperacionSuma{OpIzquierda: oi, OperadorDerecha: od}
	return suma
}

func (suma OperacionSuma) Opti_Expresion(block *bloque.BloquesOpt) bloque.ObjetoBloque {
	operadorizq := suma.OpIzquierda.Opti_Expresion(block)
	operadorder := suma.OperadorDerecha.Opti_Expresion(block)

	return bloque.ObjetoBloque{Tipo: bloque.OPERACION,
		Asignacion: "",
		OpeIz:      operadorizq.Valor,
		Ope:        "+",
		OpeDe:      operadorder.Valor,
		Valor:      operadorizq.Valor + "+" + operadorder.Valor}
}
