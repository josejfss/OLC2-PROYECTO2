package llamada

import (
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	generando "OLC2-PROYECTO2/OPT_COMPILADOR/GENERANDO"
)

type Llama struct {
	Nombre string
}

func Nllama(nom string) Llama {
	llamita := Llama{Nombre: nom}
	return llamita
}

func (got Llama) Opti_Instruccion(block *bloque.BloquesOpt, genopt *generando.GenerandoOpti) interface{} {
	guardar := bloque.Nobjetolista("llamada", got.Nombre+"();")
	genopt.AgegarCodigoOpt(guardar)
	return 0
}
