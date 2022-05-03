package etiquetas

import objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"

type Etiqueta struct {
	Nombre string
}

func Netiqueta(nom string) Etiqueta {
	eti := Etiqueta{Nombre: nom}
	return eti
}

func (eti Etiqueta) Optimizar_Instruccion(block *objeto.Bloque) interface{} {
	return eti.Nombre
}
