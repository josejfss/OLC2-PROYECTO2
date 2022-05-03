package aritmetica

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
)

type OperacionResta struct {
	OpIzquierda     interfaz.Expresion
	OperadorDerecha interfaz.Expresion
}

func Noperacionresta(oi interfaz.Expresion, od interfaz.Expresion) OperacionResta {
	resta := OperacionResta{OpIzquierda: oi, OperadorDerecha: od}
	return resta
}

func (resta OperacionResta) Optimizar_Expresion(block *objeto.Bloque) objeto.ObjetoBloque {
	operadorizq := resta.OpIzquierda.Optimizar_Expresion(block)
	operadorder := resta.OperadorDerecha.Optimizar_Expresion(block)

	return objeto.ObjetoBloque{Operacion: true, Opiz: operadorizq.Valor, Opde: operadorder.Valor, Valor: operadorizq.Valor + " - " + operadorder.Valor}
}
