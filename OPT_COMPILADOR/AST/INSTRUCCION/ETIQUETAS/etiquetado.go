package etiquetas

import (
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	generando "OLC2-PROYECTO2/OPT_COMPILADOR/GENERANDO"
)

type Etiqueta struct {
	Nombre string
}

func Netiqueta(nom string) Etiqueta {
	eti := Etiqueta{Nombre: nom}
	return eti
}

func (eti Etiqueta) Opti_Instruccion(block *bloque.BloquesOpt, genopt *generando.GenerandoOpti) interface{} {
	guardar := bloque.Nobjetolista("etiqueta", eti.Nombre+":")
	genopt.AgegarCodigoOpt(guardar)
	return 0
}
