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
	if fun.Nombre == "main" {
		simbetiq := objeto.ObjetoBloque{
			Valor: "int main() {",
		}
		block.ListaTodo.Add("fnmain")
		nuevo_bloquesito.Guardar_Declaracion("fnmain", simbetiq)
	} else {
		simbetiq := objeto.ObjetoBloque{
			Valor: "void " + fun.Nombre + "() {",
		}
		block.ListaTodo.Add("fn" + fun.Nombre)
		nuevo_bloquesito.Guardar_Declaracion("fn"+fun.Nombre, simbetiq)
	}

	for i := 0; i < fun.Lista.Len(); i++ {
		act := fun.Lista.GetValue(i).(interfaz.Instruccion)
		if interfaz.TYPEOF(act) == "sif.SenteIf" {
			noms := act.Optimizar_Instruccion(nuevo_bloquesito)
			block.Guardar_Bloque(nuevo_bloquesito)
			block.ListaTodo.Add(noms.(string))
			nuevo_bloquesito = objeto.Nuevo_Bloque()
			nuevo_bloquesito.NombreFuncion = fun.Nombre
		} else if interfaz.TYPEOF(act) == "etiquetas.Etiqueta" {

			block.Guardar_Bloque(nuevo_bloquesito)
			nuevo_bloquesito = objeto.Nuevo_Bloque()
			nuevo_bloquesito.NombreFuncion = fun.Nombre
			eti := act.Optimizar_Instruccion(nuevo_bloquesito)

			block.ListaTodo.Add(eti.(string))
		} else if interfaz.TYPEOF(act) == "retorno.Returns" {
			val := act.Optimizar_Instruccion(nuevo_bloquesito)
			block.Guardar_Bloque(nuevo_bloquesito)
			block.ListaTodo.Add(val)
		} else {
			g := act.Optimizar_Instruccion(nuevo_bloquesito)
			block.ListaTodo.Add(g.(string))
		}
	}
	return 0
}
