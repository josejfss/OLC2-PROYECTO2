package relacional

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
)

type OperacionIgualdad struct {
	OpIzquierda     interfaz.Expresion
	OperadorDerecha interfaz.Expresion
}

func Noperacionigualdad(oi interfaz.Expresion, od interfaz.Expresion) OperacionIgualdad {
	igualdad := OperacionIgualdad{OpIzquierda: oi, OperadorDerecha: od}
	return igualdad
}

func (igualdad OperacionIgualdad) Optimizar_Expresion(block *objeto.Bloque) objeto.ObjetoBloque {
	operadorizq := igualdad.OpIzquierda.Optimizar_Expresion(block)
	operadorder := igualdad.OperadorDerecha.Optimizar_Expresion(block)

	return objeto.ObjetoBloque{Operacion: true,
		Opiz:  operadorizq.Valor,
		Opde:  operadorder.Valor,
		Ope:   "==",
		Valor: operadorizq.Valor + " == " + operadorder.Valor}
}
