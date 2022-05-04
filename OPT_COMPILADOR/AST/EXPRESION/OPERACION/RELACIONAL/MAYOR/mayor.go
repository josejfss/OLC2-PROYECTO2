package mayor

import (
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
)

type OperacionMayor struct {
	OpIzquierda     inter.Expresion
	OperadorDerecha inter.Expresion
}

func Noperacionmayor(oi inter.Expresion, od inter.Expresion) OperacionMayor {
	mayor := OperacionMayor{OpIzquierda: oi, OperadorDerecha: od}
	return mayor
}

func (mayor OperacionMayor) Opti_Expresion(block *bloque.BloquesOpt) bloque.ObjetoBloque {
	operadorizq := mayor.OpIzquierda.Opti_Expresion(block)
	operadorder := mayor.OperadorDerecha.Opti_Expresion(block)

	return bloque.ObjetoBloque{Tipo: bloque.OPERACION,
		Asignacion: "",
		OpeIz:      operadorizq.Valor,
		Ope:        ">",
		OpeDe:      operadorder.Valor,
		Valor:      operadorizq.Valor + ">" + operadorder.Valor}
}
