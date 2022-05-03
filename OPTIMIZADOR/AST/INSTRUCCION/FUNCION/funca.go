package funcion

import (
	interfaz "OLC2-PROYECTO2/OPTIMIZADOR/INTERFAZ"
	objeto "OLC2-PROYECTO2/OPTIMIZADOR/OBJETO"

	"github.com/colegno/arraylist"
)

type Funca struct {
	Nombre string
	Lista  *arraylist.List
}

func Nfunca(nom string, list *arraylist.List) Funca {
	fun := Funca{Nombre: nom, Lista: list}
	return fun
}

func (fun Funca) Optimizar_Instruccion(block *objeto.Bloque) interface{} {
	nuevo_bloquesito := objeto.Nuevo_Bloque()
	nuevo_bloquesito.NombreFuncion = fun.Nombre
	for i := 0; i < fun.Lista.Len(); i++ {
		act := fun.Lista.GetValue(i).(interfaz.Instruccion)
		if interfaz.TYPEOF(act) == "sif.SenteIf" {
			act.Optimizar_Instruccion(nuevo_bloquesito)
			block.Guardar_Bloque(nuevo_bloquesito)
			nuevo_bloquesito = objeto.Nuevo_Bloque()
			nuevo_bloquesito.NombreFuncion = fun.Nombre
		} else if interfaz.TYPEOF(act) == "etiquetas.Etiqueta" {
			eti := act.Optimizar_Instruccion(nuevo_bloquesito)
			block.Guardar_Bloque(nuevo_bloquesito)
			nuevo_bloquesito = objeto.Nuevo_Bloque()
			nuevo_bloquesito.NombreFuncion = fun.Nombre
			simbetiq := objeto.ObjetoBloque{
				Operacion: false,
				Temporal:  eti.(string),
				Opiz:      eti.(string),
				Opde:      eti.(string),
				Valor:     eti.(string) + ":",
			}
			nuevo_bloquesito.Guardar_Declaracion(eti.(string), simbetiq)
		} else {
			act.Optimizar_Instruccion(nuevo_bloquesito)
		}

		if i == fun.Lista.Len()-1 {
			block.Guardar_Bloque(nuevo_bloquesito)
		}
	}
	return 0
}
