package aritmetica

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
)

type OperacionDivision struct {
	OpIzquierda     interfaz.Expresion
	OperadorDerecha interfaz.Expresion
}

func Noperaciondivision(oi interfaz.Expresion, od interfaz.Expresion) OperacionDivision {
	divi := OperacionDivision{OpIzquierda: oi, OperadorDerecha: od}
	return divi
}

func (division OperacionDivision) Optimizar_Expresion(block *objeto.Bloque) objeto.ObjetoBloque {
	operadorizq := division.OpIzquierda.Optimizar_Expresion(block)
	operadorder := division.OperadorDerecha.Optimizar_Expresion(block)

	return objeto.ObjetoBloque{Operacion: true, Opiz: operadorizq.Valor, Opde: operadorder.Valor, Valor: operadorizq.Valor + " / " + operadorder.Valor}
}
