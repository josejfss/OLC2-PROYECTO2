package funcion

import (
	bloque "OLC2-PROYECTO2/OPT_COMPILADOR/BLOQUE"
	generando "OLC2-PROYECTO2/OPT_COMPILADOR/GENERANDO"
	inter "OLC2-PROYECTO2/OPT_COMPILADOR/INTER"

	"github.com/colegno/arraylist"
)

type FuncaMain struct {
	Nombre string
	Lista  *arraylist.List
}

func Nfuncamain(nom string, list *arraylist.List) FuncaMain {
	fun := FuncaMain{Nombre: nom, Lista: list}
	return fun
}

func (mains FuncaMain) Opti_Instruccion(block *bloque.BloquesOpt, genopt *generando.GenerandoOpti) interface{} {
	guardar := bloque.Nobjetolista(mains.Nombre, "int "+mains.Nombre+"() {\nSP=0;\nHP=0;")
	genopt.AgegarCodigoOpt(guardar)
	for i := 0; i < mains.Lista.Len(); i++ {
		mains.Lista.GetValue(i).(inter.Instruccion).Opti_Instruccion(block, genopt)
	}
	guardar1 := bloque.Nobjetolista("return", "return 0;\n}\n")
	genopt.AgegarCodigoOpt(guardar1)
	return 0
}
