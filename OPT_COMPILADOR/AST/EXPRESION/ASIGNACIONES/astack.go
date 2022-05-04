package asignaciones

import (
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"
)

type AsiStack struct {
	Expresion inter.Expresion
}

func Nasistack(exp inter.Expresion) AsiStack {
	asig := AsiStack{Expresion: exp}
	return asig
}

func (asig AsiStack) Opti_Expresion(block *bloque.BloquesOpt) bloque.ObjetoBloque {
	expres := asig.Expresion.Opti_Expresion(block)

	return bloque.ObjetoBloque{
		Tipo:       bloque.TEMPORAL,
		Asignacion: "",
		OpeIz:      expres.Valor,
		OpeDe:      "",
		Ope:        "STACK",
		Valor:      expres.Valor}
}
