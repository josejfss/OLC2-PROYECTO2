package aritmetica

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
)

type OperacionModulo struct {
	OpIzquierda     interfaz.Expresion
	OperadorDerecha interfaz.Expresion
}

func Noperacionmodulo(oi interfaz.Expresion, od interfaz.Expresion) OperacionModulo {
	mod := OperacionModulo{OpIzquierda: oi, OperadorDerecha: od}
	return mod
}

func (modulo OperacionModulo) Optimizar_Expresion(block *objeto.Bloque) objeto.ObjetoBloque {
	operadorizq := modulo.OpIzquierda.Optimizar_Expresion(block)
	operadorder := modulo.OperadorDerecha.Optimizar_Expresion(block)

	return objeto.ObjetoBloque{Operacion: true,
		EsTemporal: false,
		Constante:  false,
		Opiz:       operadorizq.Valor,
		Opde:       operadorder.Valor,
		Ope:        "%",
		Valor:      operadorizq.Valor + " % " + operadorder.Valor}
}
