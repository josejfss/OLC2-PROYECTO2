package resta

import (
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
)

type OperacionResta struct {
	OpIzquierda     inter.Expresion
	OperadorDerecha inter.Expresion
}

func Noperacionresta(oi inter.Expresion, od inter.Expresion) OperacionResta {
	resta := OperacionResta{OpIzquierda: oi, OperadorDerecha: od}
	return resta
}

func (resta OperacionResta) Opti_Expresion(block *bloque.BloquesOpt) bloque.ObjetoBloque {
	operadorizq := resta.OpIzquierda.Opti_Expresion(block)
	operadorder := resta.OperadorDerecha.Opti_Expresion(block)

	return bloque.ObjetoBloque{Tipo: bloque.OPERACION,
		Asignacion: "",
		OpeIz:      operadorizq.Valor,
		Ope:        "-",
		OpeDe:      operadorder.Valor,
		Valor:      operadorizq.Valor + "-" + operadorder.Valor}
}
