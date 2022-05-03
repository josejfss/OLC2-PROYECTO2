package llamada

import (
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"
	"strconv"
)

type Llama struct {
	Nombre string
}

func Nllama(nom string) Llama {
	llamita := Llama{Nombre: nom}
	return llamita
}

func (lla Llama) Optimizar_Instruccion(block *objeto.Bloque) interface{} {
	simbllama := objeto.ObjetoBloque{
		Operacion: false,
		Temporal:  "",
		Opiz:      lla.Nombre,
		Opde:      lla.Nombre,
		Valor:     lla.Nombre + "();"}
	nomstack := "llama" + strconv.Itoa(objeto.ContadorLlamada)
	objeto.ContadorLlamada = objeto.ContadorLlamada + 1
	block.Guardar_Declaracion(lla.Nombre, simbllama)
	block.ListaTemporales.Add(nomstack)
	return lla.Nombre
}
