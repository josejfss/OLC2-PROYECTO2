package relacional

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
)

type OperacionMenorIgual struct {
	OpIzquierda     interfaz.Expresion
	OperadorDerecha interfaz.Expresion
}

func Noperacionmenorigual(oi interfaz.Expresion, od interfaz.Expresion) OperacionMenorIgual {
	menorigual := OperacionMenorIgual{OpIzquierda: oi, OperadorDerecha: od}
	return menorigual
}

func (menorigual OperacionMenorIgual) Optimizar_Expresion(block *objeto.Bloque) objeto.ObjetoBloque {
	operadorizq := menorigual.OpIzquierda.Optimizar_Expresion(block)
	operadorder := menorigual.OperadorDerecha.Optimizar_Expresion(block)

	return objeto.ObjetoBloque{Operacion: true,
		Opiz:  operadorizq.Valor,
		Opde:  operadorder.Valor,
		Ope:   "<=",
		Valor: operadorizq.Valor + " <= " + operadorder.Valor}
}
