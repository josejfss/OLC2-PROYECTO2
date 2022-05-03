package etiquetas

import (
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
	"strconv"
)

type Goto struct {
	Nombre string
}

func Ngoto(nom string) Goto {
	got := Goto{Nombre: nom}
	return got
}

func (got Goto) Optimizar_Instruccion(block *objeto.Bloque) interface{} {
	simbgoto := objeto.ObjetoBloque{Operacion: false, Opiz: got.Nombre, Opde: got.Nombre, Valor: "goto " + got.Nombre + ";"}
	nomb := "goto" + strconv.Itoa(objeto.ContadorGoto)
	objeto.ContadorGoto = objeto.ContadorGoto + 1
	block.Guardar_Declaracion(nomb, simbgoto)
	block.Guardar_Declaracion1(nomb, simbgoto)
	block.ListaTemporales.Add(nomb)
	return nomb
}
