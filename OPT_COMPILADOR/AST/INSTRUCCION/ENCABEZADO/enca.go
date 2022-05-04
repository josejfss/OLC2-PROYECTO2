package encabezado

import (
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	generando "OLC2-PROYECTO2/OPT_COMPILADOR/GENERANDO"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"

	"github.com/colegno/arraylist"
)

type Enca struct {
	Libreria     string
	DeclaHeap    string
	DeclaStack   string
	PuntHeap     string
	PuntStack    string
	L_Temporales *arraylist.List
	L_Funciones  *arraylist.List
}

func Nenca(lib string, dheap string, dstack string, pheap string, pstack string, ltemps *arraylist.List, lfun *arraylist.List) Enca {
	enc := Enca{Libreria: lib,
		DeclaHeap:    dheap,
		DeclaStack:   dstack,
		PuntHeap:     pheap,
		PuntStack:    pstack,
		L_Temporales: ltemps,
		L_Funciones:  lfun}
	return enc
}

func (enc Enca) Opti_Instruccion(block *bloque.BloquesOpt, genopt *generando.GenerandoOpti) interface{} {
	block.ListaTemporales = enc.L_Temporales.Clone()
	block.Encabezado = enc.Libreria + "\n" + enc.DeclaHeap + "\n" + enc.DeclaStack + "\n" + enc.PuntHeap + "\n" + enc.PuntStack
	for i := 0; i < enc.L_Funciones.Len(); i++ {
		act := enc.L_Funciones.GetValue(i).(inter.Instruccion)
		act.Opti_Instruccion(block, genopt)
	}
	return 0
}
