package aritmetica

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
)

type OperacionSuma struct {
	OpIzquierda     interfaz.Expresion
	OperadorDerecha interfaz.Expresion
}

func Noperacionsuma(oi interfaz.Expresion, od interfaz.Expresion) OperacionSuma {
	suma := OperacionSuma{OpIzquierda: oi, OperadorDerecha: od}
	return suma
}

func (suma OperacionSuma) Optimizar_Expresion(block *objeto.Bloque) objeto.ObjetoBloque {
	operadorizq := suma.OpIzquierda.Optimizar_Expresion(block)
	operadorder := suma.OperadorDerecha.Optimizar_Expresion(block)

	return objeto.ObjetoBloque{Operacion: true,
		EsTemporal: false,
		Constante:  false,
		Opiz:       operadorizq.Valor,
		Opde:       operadorder.Valor,
		Ope:        "+",
		Valor:      operadorizq.Valor + " + " + operadorder.Valor}
}
