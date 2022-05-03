package aritmetica

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
)

type Negativo struct {
	Operador interfaz.Expresion
}

func Nnegativo(op interfaz.Expresion) Negativo {
	nega := Negativo{Operador: op}
	return nega
}

func (nega Negativo) Optimizar_Expresion(block *objeto.Bloque) objeto.ObjetoBloque {
	neg := nega.Operador.Optimizar_Expresion(block)
	return objeto.ObjetoBloque{Operacion: false,
		EsTemporal: false,
		Constante:  false,
		Opiz:       "-" + neg.Valor,
		Opde:       "-" + neg.Valor,
		Ope:        "-",
		Valor:      " - " + neg.Valor}
}
