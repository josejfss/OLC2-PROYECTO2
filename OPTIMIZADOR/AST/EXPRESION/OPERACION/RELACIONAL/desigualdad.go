package relacional

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
)

type OperacionDesigualdad struct {
	OpIzquierda     interfaz.Expresion
	OperadorDerecha interfaz.Expresion
}

func Noperaciondesigualdad(oi interfaz.Expresion, od interfaz.Expresion) OperacionDesigualdad {
	desigualdad := OperacionDesigualdad{OpIzquierda: oi, OperadorDerecha: od}
	return desigualdad
}

func (desigualdad OperacionDesigualdad) Optimizar_Expresion(block *objeto.Bloque) objeto.ObjetoBloque {
	operadorizq := desigualdad.OpIzquierda.Optimizar_Expresion(block)
	operadorder := desigualdad.OperadorDerecha.Optimizar_Expresion(block)

	return objeto.ObjetoBloque{Operacion: true,
		Opiz:  operadorizq.Valor,
		Opde:  operadorder.Valor,
		Ope:   "!=",
		Valor: operadorizq.Valor + " != " + operadorder.Valor}
}
