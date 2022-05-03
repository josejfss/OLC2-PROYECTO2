package relacional

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
)

type OperacionMayor struct {
	OpIzquierda     interfaz.Expresion
	OperadorDerecha interfaz.Expresion
}

func Noperacionmayor(oi interfaz.Expresion, od interfaz.Expresion) OperacionMayor {
	mayor := OperacionMayor{OpIzquierda: oi, OperadorDerecha: od}
	return mayor
}

func (mayor OperacionMayor) Optimizar_Expresion(block *objeto.Bloque) objeto.ObjetoBloque {
	operadorizq := mayor.OpIzquierda.Optimizar_Expresion(block)
	operadorder := mayor.OperadorDerecha.Optimizar_Expresion(block)

	return objeto.ObjetoBloque{Operacion: true,
		Opiz:  operadorizq.Valor,
		Opde:  operadorder.Valor,
		Ope:   ">",
		Valor: operadorizq.Valor + " > " + operadorder.Valor}
}
