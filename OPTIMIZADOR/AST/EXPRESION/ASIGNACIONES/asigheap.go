package asignaciones

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
)

type AsiHeap struct {
	Expresion interfaz.Expresion
}

func Nasiheap(exp interfaz.Expresion) AsiHeap {
	asig := AsiHeap{Expresion: exp}
	return asig
}

func (asig AsiHeap) Optimizar_Expresion(block *objeto.Bloque) objeto.ObjetoBloque {
	expres := asig.Expresion.Optimizar_Expresion(block)

	return objeto.ObjetoBloque{Operacion: false, Opiz: expres.Valor, Opde: expres.Valor, Valor: "HEAP[(int)" + expres.Valor + "]"}
}
