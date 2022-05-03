package asignaciones

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
)

type AsiStack struct {
	Expresion interfaz.Expresion
}

func Nasistack(exp interfaz.Expresion) AsiStack {
	asig := AsiStack{Expresion: exp}
	return asig
}

func (asig AsiStack) Optimizar_Expresion(block *objeto.Bloque) objeto.ObjetoBloque {
	expres := asig.Expresion.Optimizar_Expresion(block)

	return objeto.ObjetoBloque{Operacion: false, Opiz: expres.Valor, Opde: expres.Valor, Valor: "STACK[(int)" + expres.Valor + "]"}
}
