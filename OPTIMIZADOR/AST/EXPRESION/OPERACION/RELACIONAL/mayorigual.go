package relacional

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
)

type OperacionMayorIgual struct {
	OpIzquierda     interfaz.Expresion
	OperadorDerecha interfaz.Expresion
}

func Noperacionmayorigual(oi interfaz.Expresion, od interfaz.Expresion) OperacionMayorIgual {
	mayorigual := OperacionMayorIgual{OpIzquierda: oi, OperadorDerecha: od}
	return mayorigual
}

func (mayorigual OperacionMayorIgual) Optimizar_Expresion(block *objeto.Bloque) objeto.ObjetoBloque {
	operadorizq := mayorigual.OpIzquierda.Optimizar_Expresion(block)
	operadorder := mayorigual.OperadorDerecha.Optimizar_Expresion(block)

	return objeto.ObjetoBloque{Operacion: true,
		Opiz:  operadorizq.Valor,
		Opde:  operadorder.Valor,
		Ope:   ">=",
		Valor: operadorizq.Valor + " >= " + operadorder.Valor}
}
