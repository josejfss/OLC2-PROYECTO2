package relacional

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
)

type OperacionMenor struct {
	OpIzquierda     interfaz.Expresion
	OperadorDerecha interfaz.Expresion
}

func Noperacionmenor(oi interfaz.Expresion, od interfaz.Expresion) OperacionMenor {
	menor := OperacionMenor{OpIzquierda: oi, OperadorDerecha: od}
	return menor
}

func (menor OperacionMenor) Optimizar_Expresion(block *objeto.Bloque) objeto.ObjetoBloque {
	operadorizq := menor.OpIzquierda.Optimizar_Expresion(block)
	operadorder := menor.OperadorDerecha.Optimizar_Expresion(block)

	return objeto.ObjetoBloque{Operacion: true,
		Opiz:  operadorizq.Valor,
		Opde:  operadorder.Valor,
		Ope:   "<",
		Valor: operadorizq.Valor + " < " + operadorder.Valor}
}
