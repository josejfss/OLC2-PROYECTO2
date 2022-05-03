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
	//block.ListaTemporales.Add(eti.Nombre)
	simbetiq := objeto.ObjetoBloque{
		Operacion: false,
		Temporal:  eti.Nombre,
		Opiz:      eti.Nombre,
		Opde:      eti.Nombre,
		Valor:     eti.Nombre + ":",
	}
	block.Guardar_Declaracion(eti.Nombre, simbetiq)
	block.Guardar_Declaracion1(eti.Nombre, simbetiq)
	block.ListaTemporales.Add(eti.Nombre)
	return eti.Nombre
}
