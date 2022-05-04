package etiquetas

import (
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	generando "OLC2-PROYECTO2/OPT_COMPILADOR/GENERANDO"
)

type Goto struct {
	Nombre string
}

func Ngoto(nom string) Goto {
	got := Goto{Nombre: nom}
	return got
}

func (got Goto) Opti_Instruccion(block *bloque.BloquesOpt, genopt *generando.GenerandoOpti) interface{} {
	guardar := bloque.Nobjetolista("goto", "goto "+got.Nombre+";")
	genopt.AgegarCodigoOpt(guardar)
	return 0
}
