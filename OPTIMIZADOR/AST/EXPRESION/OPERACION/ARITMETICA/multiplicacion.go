package aritmetica

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
)

type OperacionMultiplicacion struct {
	OpIzquierda     interfaz.Expresion
	OperadorDerecha interfaz.Expresion
}

func Noperacionmultiplicacion(oi interfaz.Expresion, od interfaz.Expresion) OperacionMultiplicacion {
	multi := OperacionMultiplicacion{OpIzquierda: oi, OperadorDerecha: od}
	return multi
}

func (multi OperacionMultiplicacion) Optimizar_Expresion(block *objeto.Bloque) objeto.ObjetoBloque {
	operadorizq := multi.OpIzquierda.Optimizar_Expresion(block)
	operadorder := multi.OperadorDerecha.Optimizar_Expresion(block)

	return objeto.ObjetoBloque{Operacion: true,
		EsTemporal: false,
		Constante:  false,
		Opiz:       operadorizq.Valor,
		Opde:       operadorder.Valor,
		Ope:        "*",
		Valor:      operadorizq.Valor + " * " + operadorder.Valor}
}
